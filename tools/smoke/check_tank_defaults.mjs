import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { chromium } from 'playwright';

const site = process.env.SITE_URL || 'http://127.0.0.1:8767/wow-forever-sim/classic/';
const engine = process.env.TANK_BENCH || '/tmp/forever-tank-bench';
const browser = await chromium.launch({ executablePath: process.env.CHROMIUM_PATH || undefined });
const temp = mkdtempSync(join(tmpdir(), 'forever-tanks-'));
try {
  for (const route of ['tank_warrior', 'protection_paladin', 'feral_tank_druid']) {
    const context = await browser.newContext();
    const page = await context.newPage();
    await page.route(/^https?:\/\/([^/]+\.)?(zamimg|wowhead|googletagmanager)\.com\//, request => request.abort());
    await page.goto(new URL(`${route}/`, site).toString(), { waitUntil: 'networkidle' });
    await page.getByRole('button', { name: 'Simulate', exact: true }).waitFor();
    const key = `__classic_${route}__currentSettings__`;
    await page.waitForFunction(key => localStorage.getItem(key), key);
    await page.evaluate(key => {
      const state = JSON.parse(localStorage.getItem(key));
      state.settings.iterations = 300;
      state.settings.fixedRngSeed = 20261993;
      localStorage.setItem(key, JSON.stringify(state));
    }, key);
    await page.reload({ waitUntil: 'networkidle' });
    const state = await page.evaluate(key => JSON.parse(localStorage.getItem(key)), key);
    assert.equal(state.player.rotation.type, 'TypeAPL');
    assert.ok(state.player.rotation.priorityList.length > 0);
    assert.equal(state.player.foreverTier1Bonuses, true);
    assert.equal(state.player.healingModel.hps, 1500);
    assert.equal(state.encounter.duration, 300);
    assert.equal(state.encounter.targets[0].stats[17], 805);
    assert.equal(state.encounter.targets[0].stats[26], 3731);
    assert.ok(state.player.equipment.items.every(item => !item.id || item.id >= 920000001));
    const request = {
      raid: {
        parties: [{ players: [state.player], buffs: state.partyBuffs }],
        buffs: state.raidBuffs,
        debuffs: state.debuffs,
        tanks: state.tanks,
      },
      encounter: state.encounter,
      simOptions: { iterations: 300, randomSeed: 20261993, ruleset: state.settings.ruleset },
    };
    const input = join(temp, `${route}.request.json`);
    const output = join(temp, `${route}.result.json`);
    writeFileSync(input, JSON.stringify(request));
    execFileSync(engine, ['-request', input, '-output', output], { stdio: 'pipe' });
    const native = JSON.parse(readFileSync(output, 'utf8'));
    assert.equal(native.warnings?.length || 0, 0, `${route}: APL warning`);
    const metrics = native.result.raidMetrics.parties[0].players[0];
    await page.getByRole('button', { name: 'Simulate', exact: true }).click();
    await page.getByText('Save as Reference', { exact: true }).first().waitFor({ timeout: 180000 });
    await page.getByRole('tab', { name: 'Results', exact: true }).click();
    const row = page.locator('main table').first().locator('tbody tr').first();
    const cells = await row.locator('td').allTextContents();
    for (const [index, expected, label] of [[0, metrics.dps.avg, 'DPS'], [1, metrics.threat.avg, 'TPS'], [2, metrics.dtps.avg, 'DTPS']]) {
      const web = Number(cells[index].trim().replaceAll(',', '').split(/\s/)[0]);
      assert.ok(Math.abs(web - expected) < 0.03, `${route}: ${label} web ${web}, native ${expected}`);
    }
    console.log(`${route}: tank APL, Tier 1, modeled gear and native/WASM DPS/TPS/DTPS match`);
    if (route === 'tank_warrior') {
      const invalid = structuredClone(state);
      const strike = invalid.player.rotation.priorityList.find(
        item => item.action?.castSpell?.spellId?.spellId === 25286,
      );
      assert.equal(strike.action.castSpell.spellId.tag, 1);
      delete strike.action.castSpell.spellId.tag;
      await page.evaluate(({ key, invalid }) => localStorage.setItem(key, JSON.stringify(invalid)), { key, invalid });
      await page.reload({ waitUntil: 'networkidle' });
      await page.getByRole('tab', { name: 'Rotation', exact: true }).click();
      const warning = page.locator('.apl-warnings:visible').first();
      await warning.waitFor({ state: 'visible' });
      await warning.hover();
      await page.getByText(/replaces the next melee swing; use its queue action/).waitFor();
      console.log('tank_warrior: saved direct-cast Heroic Strike is rejected with a queue warning');
    }
    await context.close();
  }
} finally {
  await browser.close();
  rmSync(temp, { recursive: true, force: true });
}
