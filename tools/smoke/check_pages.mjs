// Loads every page of a built site in a headless browser and fails on the first
// uncaught error. Run against a local copy of dist:
//
//   python3 -m http.server 8080 --directory dist &
//   SITE_URL=http://localhost:8080/classic/ node tools/smoke/check_pages.mjs
//
// The Go tests and tsc both pass on a build whose pages throw at load, which has
// happened twice, so this is the check that actually opens them.
import { chromium } from 'playwright';
import { readdirSync, statSync } from 'fs';
import { join } from 'path';

const siteUrl = process.env.SITE_URL || 'http://localhost:8080/classic/';
const distDir = process.env.DIST_DIR || 'dist/classic';

// Every directory under dist/classic with an index.html is a page.
const pages = [''].concat(
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
