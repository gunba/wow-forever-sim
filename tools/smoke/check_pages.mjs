// Loads selected pages, or every page by default, of a built site in a headless
// browser. Run against a local copy of dist:
//
//   python3 -m http.server 8080 --directory dist &
//   SITE_URL=http://localhost:8080/classic/ node tools/smoke/check_pages.mjs
//   SITE_URL=http://localhost:8080/classic/ node tools/smoke/check_pages.mjs --page warrior --page .
//   node tools/smoke/check_pages.mjs --list
//
// The Go tests and tsc both pass on a build whose pages throw at load, which has
// happened twice, so this is the check that actually opens them.
import { chromium } from 'playwright';
import { readdirSync, statSync } from 'fs';
import { join } from 'path';
import { parseArgs } from 'node:util';

const { values } = parseArgs({
	options: {
		page: { type: 'string', multiple: true },
		list: { type: 'boolean' },
	},
});

const siteUrl = process.env.SITE_URL || 'http://localhost:8080/classic/';
const distDir = process.env.DIST_DIR || 'dist/classic';

// Every directory under dist/classic with an index.html is a page.
const availablePages = [''].concat(
	readdirSync(distDir)
		.filter(name => {
			try {
				return statSync(join(distDir, name, 'index.html')).isFile();
			} catch {
				return false;
			}
		})
		.map(name => name + '/'),
);

if (values.list) {
	console.log(availablePages.map(page => page || '.').join('\n'));
	process.exit(0);
}
const pages = values.page
	? [...new Set(values.page.map(page => page === '.' ? '' : page.replace(/\/+$/, '') + '/'))]
	: availablePages;
const unknown = pages.filter(page => !availablePages.includes(page));
if (unknown.length) {
	console.error(`Unknown pages: ${unknown.join(', ')}. Use --list to see built pages.`);
	process.exit(1);
}

// CHROMIUM_PATH points at a browser already on the machine, for running this outside CI.
const browser = await chromium.launch({ executablePath: process.env.CHROMIUM_PATH || undefined });
let failures = 0;
for (const page of pages) {
	const tab = await browser.newPage();
	const errors = [];
	tab.on('pageerror', error => errors.push(error.message));
	tab.on('response', response => {
		const url = new URL(response.url());
		if (url.origin === new URL(siteUrl).origin && /\.(jpg|png|gif|svg)$/i.test(url.pathname) && response.status() >= 400) {
			errors.push(`Image ${response.status()}: ${url.pathname}`);
		}
	});
	// Third party icon and tooltip hosts are not what is under test and are slow.
	await tab.route(/^https?:\/\/([^/]+\.)?(zamimg|wowhead|googletagmanager)\.com\//, route => route.abort());
	try {
		await tab.goto(siteUrl + page, { waitUntil: 'domcontentloaded' });
		await tab.waitForTimeout(3000);
	} catch (error) {
		errors.push(error.message);
	}
	await tab.close();
	if (errors.length) {
		failures++;
		console.log(`FAIL ${page || '(landing)'}\n  ${errors.join('\n  ')}`);
	} else {
		console.log(`ok   ${page || '(landing)'}`);
	}
}
await browser.close();
if (failures) {
	console.log(`${failures} of ${pages.length} pages threw on load`);
	process.exit(1);
}
console.log(`${pages.length} pages loaded clean`);
