import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { chromium } from 'playwright';

const base = process.env.SITE_URL || 'http://localhost:8080/classic/';
const profiles = JSON.parse(readFileSync('ui/core/forever_ranked_profiles.json', 'utf8')).profiles;
const representatives = [...new Map(profiles.map(profile => [profile.key, profile])).values()];
const routeFor = profile => {
	const cls = profile.settings.player.class.replace(/^Class/, '').toLowerCase();
	if (cls === 'druid') return profile.key === 'balance' ? 'balance_druid' : 'feral_druid';
	if (cls === 'priest') return profile.key === 'smite' ? 'smite_priest' : 'shadow_priest';
	if (cls === 'paladin') return 'retribution_paladin';
	if (cls === 'shaman') return profile.key === 'enhancement' ? 'enhancement_shaman' : 'elemental_shaman';
	return cls;
};
const browser = await chromium.launch({ executablePath: process.env.CHROMIUM_PATH || undefined });
try {
	const page = await browser.newPage();
	// Display metadata and images must work without third-party requests.
	await page.route(/^https?:\/\/([^/]+\.)?(zamimg|wowhead|googletagmanager)\.com\//, route => route.abort());
	for (const profile of representatives) {
		const route = routeFor(profile);
		await page.goto(`${base}${route}/`, { waitUntil: 'networkidle' });
		await page.locator(`option[value="${profile.id}"]`).waitFor({ state: 'attached' });
		await page.getByRole('combobox', { name: 'Ranked build', exact: true }).selectOption(profile.id);
		await page.getByRole('button', { name: 'Load build', exact: true }).click();
		for (const tab of ['Gear', 'Settings', 'Talents', 'Rotation']) {
			await page.getByRole('tab', { name: tab, exact: true }).click();
			const problems = await page.evaluate(async () => {
				const urls = new Set();
				for (const element of document.querySelectorAll('img, a, div, button')) {
					if (element.tagName === 'IMG' && element.getAttribute('src')) urls.add(element.src);
					// Include closed dropdown options, not only visible controls.
					for (const match of getComputedStyle(element).backgroundImage.matchAll(/url\(["']?(.*?)["']?\)/g)) {
						urls.add(match[1]);
					}
				}
				const failures = (await Promise.all([...urls].map(async url => {
					const image = new Image();
					image.src = url;
					try { await image.decode(); return null; } catch { return url; }
				}))).filter(Boolean);
				for (const button of document.querySelectorAll('a.icon-picker-button.active')) {
					if (!button.getClientRects().length) continue;
					const style = getComputedStyle(button);
					if (style.backgroundImage === 'none' && !button.textContent.trim()) {
						failures.push(`Blank active icon: ${button.outerHTML.slice(0, 300)}`);
					}
					if (button.textContent.trim() === 'Invalid') failures.push('Unsupported active selection');
				}
				return failures;
			});
			assert.deepEqual(problems, [], `${profile.key}/${tab}: icon coverage`);
		}
		console.log(`${profile.key}: gear, settings, talents, rotation and dropdown images decode`);
	}
	await page.goto(`${base}review/`, { waitUntil: 'networkidle' });
	const broken = await page.locator('img').evaluateAll(async images =>
		(await Promise.all(images.map(async image => {
			try { await image.decode(); return null; } catch { return image.src; }
		}))).filter(Boolean));
	assert.deepEqual(broken, [], 'review chart images');
} finally {
	await browser.close();
}
