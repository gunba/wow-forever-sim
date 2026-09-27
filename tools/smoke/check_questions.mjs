// Regression for the former unstructured/stale checklist above the matrix.
import assert from 'node:assert/strict';
import { chromium } from 'playwright';

const base = process.env.SITE_URL || 'http://localhost:8080/classic/';
const browser = await chromium.launch({ executablePath: process.env.CHROMIUM_PATH || undefined });
try {
	const page = await browser.newPage({ viewport: { width: 1440, height: 1000 } });
	const errors = [];
	page.on('pageerror', e => errors.push(e.message));
	await page.goto(new URL('review/', base).href, { waitUntil: 'domcontentloaded' });
	const register = await (await page.request.get(new URL('review/uncertainties.json', base).href)).json();
	assert.equal(await page.locator('.question-card').count(), register.items.length);
	assert(await page.evaluate(() =>
		document.querySelector('#questions').compareDocumentPosition(document.querySelector('#matrix'))
		& Node.DOCUMENT_POSITION_FOLLOWING));
	assert.equal(await page.locator('.question-group:not([hidden])[open]').count(), 0);
	assert.equal(await page.locator('.question-card[data-status="resolved"]:not([hidden])').count(), 0);
	assert.equal(await page.locator('.question-card[data-status="provisional"]:not([hidden])').count(), 0);
	await page.locator('#question-search').fill('Maelstrom');
	assert(await page.locator('#SHA-001').isVisible());
	assert(await page.locator('.question-group[open]').count() > 0);
	await page.locator('#SHA-001 > summary').click();
	assert(await page.locator('#SHA-001').getAttribute('open') !== null);
	await page.locator('#question-search').fill('this-is-not-a-real-spell');
	assert.equal(await page.locator('.question-card:not([hidden])').count(), 0);
	await page.goto(new URL('review/#DONE-007', base).href, { waitUntil: 'domcontentloaded' });
	assert(await page.locator('#DONE-007').isVisible());
	assert(await page.locator('#DONE-007').getAttribute('open') !== null);
	assert.equal(await page.locator('#question-status').inputValue(), 'all');
	await page.locator('#question-search').fill('');
	await page.locator('#question-status').selectOption('needs-code');
	await page.locator('#question-priority').selectOption('high');
	await page.locator('#question-access').selectOption('offline');
	assert(await page.locator('.question-card:not([hidden])').count() > 0);
	assert.equal(await page.locator('.question-card:not([hidden]):not([data-status="needs-code"])').count(), 0);
	assert.equal(await page.locator('.question-card:not([hidden]):not([data-priority="high"])').count(), 0);
	for (const stage of ['current-beta', 'level30', 'launch', 'offline']) {
		await page.locator(`button[data-stage="${stage}"]`).click();
		const expected = register.items.filter(item =>
			['needs-evidence', 'needs-code', 'model-choice'].includes(item.status)
			&& item.review.tests.some(test => test.stage === stage)).map(item => item.id).sort();
		const visible = await page.locator('.question-card:not([hidden])').evaluateAll(cards => cards.map(c => c.id).sort());
		assert.deepEqual(visible, expected, `${stage}: multi-stage filter mismatch`);
	}
	await page.locator('button[data-stage="current-beta"]').click();
	await page.setViewportSize({ width: 390, height: 844 });
	assert(await page.evaluate(() => document.querySelector('#questions').scrollWidth <= window.innerWidth));
	assert.deepEqual(errors, []);
	console.log(`${register.items.length} register entries: categories, filters, empty state, deep links and mobile layout pass`);
} finally {
	await browser.close();
}
