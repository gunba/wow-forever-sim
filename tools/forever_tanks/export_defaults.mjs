import assert from 'node:assert/strict';
import { mkdirSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import { chromium } from 'playwright';

const site = process.env.SITE_URL || 'http://127.0.0.1:8767/wow-forever-sim/classic/';
const output = process.argv[2];
if (!output) throw new Error('Usage: node tools/forever_tanks/export_defaults.mjs OUTPUT_DIR');
mkdirSync(output, { recursive: true });
const browser = await chromium.launch({ executablePath: process.env.CHROMIUM_PATH || undefined });
try {
  for (const route of ['tank_warrior', 'protection_paladin', 'feral_tank_druid']) {
    const context = await browser.newContext();
    const page = await context.newPage();
    await page.route(/^https?:\/\/([^/]+\.)?(zamimg|wowhead|googletagmanager)\.com\//, request => request.abort());
    await page.goto(new URL(`${route}/`, site).toString(), { waitUntil: 'networkidle' });
    await page.getByRole('button', { name: 'Simulate', exact: true }).waitFor();
    const key = `__classic_${route}__currentSettings__`;
    await page.waitForFunction(key => localStorage.getItem(key), key);
    const settings = await page.evaluate(key => JSON.parse(localStorage.getItem(key)), key);
    assert.equal(settings.settings.ruleset, 'RulesetForever');
    assert.equal(settings.player.foreverTier1Bonuses, true);
    assert.equal(settings.player.rotation.type, 'TypeAPL');
    assert.ok(settings.player.rotation.priorityList?.length > 0, 'tank default must load its real APL');
    assert.equal(settings.encounter.duration, 300);
    assert.equal(settings.encounter.targets.length, 1);
    assert.equal(settings.encounter.targets[0].level, 63);
    assert.equal(settings.encounter.targets[0].stats[26], 3731);
    assert.equal(settings.encounter.targets[0].minBaseDamage, 3000);
    assert.ok(settings.player.equipment.items.every(item => !item.id || item.id >= 920000001));
    assert.ok(settings.player.equipment.items[14].id);
    writeFileSync(join(output, `${route}.settings.json`), JSON.stringify(settings, null, 2) + '\n');
    console.log(`${route}: ${settings.player.race}, ${settings.player.talentsString}, ${settings.player.equipment.items.filter(item => item.id).length} items, Tier 1 on`);
    await context.close();
  }
} finally {
  await browser.close();
}
