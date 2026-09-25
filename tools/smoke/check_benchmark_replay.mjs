// Browser imports must preserve the scaled Alliance benchmark, including
// Shaman buffs, Windfury and the selected rank of Blessing of Might.
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { chromium } from 'playwright';

const base = process.env.SITE_URL || 'http://localhost:8080/classic/';
const scenario = JSON.parse(readFileSync('ui/core/forever_ranked_profiles.json', 'utf8')).gearScenario;
const prefix = ['modeled-65-v1', 'modeled-65-v2'].includes(scenario)
	? 'artifacts/modelled_gear/' : 'artifacts/';
const data = JSON.parse(readFileSync(`${prefix}sensitivity/gear_150.json`, 'utf8'));
const row = data.Results.find(r => r.Key === 'fury' && r.Race === 'Gnome');
const req = row.Request, raid = req.raid, party = raid.parties[0], player = party.players[0];
const settings = {
	settings: {
		iterations: req.simOptions.iterations, fixedRngSeed: req.simOptions.randomSeed,
		ruleset: req.simOptions.ruleset, phase: 1, faction: row.Faction, showDamageMetrics: true,
	},
	raidBuffs: raid.buffs, partyBuffs: party.buffs, debuffs: raid.debuffs,
	tanks: raid.tanks || [], player, encounter: req.encounter,
};
const browser = await chromium.launch({ executablePath: process.env.CHROMIUM_PATH || undefined });
try {
	const context = await browser.newContext();
	const page = await context.newPage();
	await page.route(/^https?:\/\/([^/]+\.)?(zamimg|wowhead|googletagmanager)\.com\//, route => route.abort());
	await page.goto(`${base}warrior/`, { waitUntil: 'networkidle' });
	await page.getByRole('button', { name: 'Simulate', exact: true }).waitFor();
	await page.locator('.import-dropdown .import-link').evaluate(element => element.click());
	await page.locator('.import-dropdown').getByRole('button', { name: 'JSON', exact: true }).click();
	const dialog = page.getByRole('dialog');
	await dialog.getByRole('textbox').fill(JSON.stringify(settings));
	await dialog.getByRole('button', { name: /Import/ }).click();
	await dialog.waitFor({ state: 'hidden' });
	const stored = await page.evaluate(() => JSON.parse(localStorage.getItem('__classic_warrior__currentSettings__')));
	assert.deepEqual(stored.raidBuffs, raid.buffs);
	assert.deepEqual(stored.player.buffs, player.buffs);
	assert.deepEqual(stored.player.consumes, player.consumes);
	assert.deepEqual(stored.player.bonusStats.stats, player.bonusStats.stats);
	const pseudoStats = stored.player.bonusStats.pseudoStats || [];
	assert.deepEqual(pseudoStats, player.bonusStats.pseudoStats || Array(pseudoStats.length).fill(0));
	assert.equal(stored.player.equipmentScale, 1.5);
	assert.equal(stored.player.talentsString, player.talentsString);
	await page.getByRole('button', { name: 'Simulate', exact: true }).click();
	await page.getByText('Save as Reference', { exact: true }).first().waitFor({ timeout: 180000 });
	const actual = Number(await page.locator('.results-sim-dps .topline-result-avg').first().innerText());
	assert.ok(Math.abs(actual - row.DPS) < 0.015, `WASM ${actual} vs native ${row.DPS}`);
	console.log(`Gnome Fury +50%: ${actual} DPS; inputs preserved; native match`);

	const invalidShaman = JSON.parse(readFileSync(`${prefix}ui_profiles/enhancement__orc.json`, 'utf8'));
	invalidShaman.player.equipment.items[14] = { id: 279261 };
	invalidShaman.player.equipment.items[15] = { id: 22384 };
	await page.goto(`${base}enhancement_shaman/`, { waitUntil: 'networkidle' });
	await page.getByRole('button', { name: 'Simulate', exact: true }).waitFor();
	await page.locator('.import-dropdown .import-link').evaluate(element => element.click());
	await page.locator('.import-dropdown').getByRole('button', { name: 'JSON', exact: true }).click();
	await dialog.getByRole('textbox').fill(JSON.stringify(invalidShaman));
	await dialog.getByRole('button', { name: /Import/ }).click();
	await page.getByText(/Import error: Cannot equip .* in the off hand for this class/).waitFor();
	assert.equal(await page.getByText(/Perfect Sync|Delayed Offhand|Sync Type/).count(), 0);
	console.log('Shaman dual-wield import rejected; obsolete synchronization controls absent');
	await context.close();
} finally {
	await browser.close();
}
