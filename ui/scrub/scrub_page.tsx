import { SITE_BASE, SITE_REPO_URL } from '../core/constants/other';
import { ActorRecord, plausible, readRecords, scrub } from './scrub';

const ISSUE_URL = `${SITE_REPO_URL}/issues/new?template=hotfix_cache.md`;

const totals = (records: ActorRecord[]) => {
	const byActor = new Map<string, { className: string; damage: number; hits: number }>();
	const bySpell = new Map<number, { damage: number; hits: number }>();
	for (const record of records.filter(plausible)) {
		const actor = byActor.get(record.name) ?? { className: record.className, damage: 0, hits: 0 };
		actor.damage += record.damage;
		actor.hits += record.hits;
		byActor.set(record.name, actor);

		const spell = bySpell.get(record.spellId) ?? { damage: 0, hits: 0 };
		spell.damage += record.damage;
		spell.hits += record.hits;
		bySpell.set(record.spellId, spell);
	}
	return { byActor, bySpell };
};

const summaryRow = (label: string, cells: string[]) => (
	<li className="scrub-row">
		<span className="scrub-row-label">{label}</span>
		{cells.map(cell => (
			<span className="scrub-row-cell">{cell}</span>
		))}
	</li>
);

const dropZone = (id: string, title: string, hint: string, onFile: (file: File | undefined) => void) => {
	const label = (
		<label className="scrub-drop" id={id}>
			<input type="file" onchange={(event: Event) => onFile((event.target as HTMLInputElement).files?.[0])} />
			<i className="fas fa-upload" />
			<span className="scrub-drop-title">{title}</span>
			<span className="scrub-drop-hint">{hint}</span>
		</label>
	) as HTMLElement;

	label.addEventListener('dragover', event => {
		event.preventDefault();
		label.classList.add('scrub-drop-over');
	});
	label.addEventListener('dragleave', () => label.classList.remove('scrub-drop-over'));
	label.addEventListener('drop', event => {
		event.preventDefault();
		label.classList.remove('scrub-drop-over');
		onFile((event as DragEvent).dataTransfer?.files?.[0]);
	});
	return label;
};

export class ScrubPage {
	private readonly hotfixResult: HTMLElement;
	private readonly meterResult: HTMLElement;

	constructor(parent: HTMLElement) {
		parent.appendChild(
			<div id="scrub-page">
				<header className="scrub-header">
					<div className="container scrub-header-container">
						<a href={SITE_BASE} className="scrub-home-link">
							<img className="forever-logo" src={`${SITE_BASE}assets/img/forever_logo.png`} alt="World of Warcraft: Forever" />
						</a>
						<div className="scrub-title-block">
							<h1 className="scrub-title">Inspect beta cache files</h1>
							<p className="scrub-subtitle">
								This sim reads Blizzard&apos;s data tables, which say what an ability is <em>meant</em> to do. Two files on your machine say
								things the tables cannot. Files stay in this browser; this site has no upload service.
							</p>
						</div>
					</div>
				</header>

				<main className="container scrub-content">
					<div className="scrub-files">
						<section className="scrub-file">
							<h2 className="scrub-file-title">
								<code>DBCache.bin</code>
							</h2>
							<p className="scrub-file-sub">What Blizzard changed after the build shipped</p>
							<p>
								A client cache can contain hotfixes newer than the static tables used by the simulator.
							</p>
							<p className="scrub-path">
								<code>World of Warcraft\_classic_beta_\Cache\ADB\enUS\</code>
							</p>
							<p className="scrub-file-note">
								This checks the file header locally. It does not upload the cache or apply it to the simulator.
							</p>
							{dropZone('hotfix-drop', 'Choose DBCache.bin', 'or drag it here', file => this.inspectHotfix(file))}
							<div className="scrub-result scrub-result-hotfix" />
						</section>

						<section className="scrub-file">
							<h2 className="scrub-file-title">
								<code>DamageMeter.bin</code>
							</h2>
							<p className="scrub-file-sub">What the server actually paid out</p>
							<p>
								The client&apos;s meter records observed totals, which can help check behavior that static spell tables do not establish.
							</p>
							<p className="scrub-path">
								<code>World of Warcraft\_classic_beta_\Cache\</code>
							</p>
							<p className="scrub-file-note scrub-hint">
								<strong>Not there? That is normal.</strong> The client writes this file while you play and clears it between sessions, so an
								empty folder means the meter has nothing in it yet rather than that you are in the wrong place. Log in, fight something, then
								copy it out &mdash; ideally without closing the game first.
							</p>
							<p className="scrub-file-note scrub-warn">
								This file holds character names. Recognized class-tagged names are replaced <strong>in your browser</strong>.
								Pet names, unrecognized records and other identifying data may remain. The download is not guaranteed anonymous.
							</p>
							{dropZone('meter-drop', 'Choose DamageMeter.bin', 'or drag it here', file => this.scrubMeter(file))}
							<div className="scrub-result scrub-result-meter" />
						</section>
					</div>

					<section className="scrub-drop-section">
						<p className="scrub-file-note">
							Review any file before sharing it. You can{' '}
							<a href={ISSUE_URL} target="_blank" rel="noreferrer">
								open an issue
							</a>{' '}
							to discuss a finding without attaching a file.
						</p>
					</section>
				</main>
			</div>,
		);
		this.hotfixResult = parent.querySelector('.scrub-result-hotfix') as HTMLElement;
		this.meterResult = parent.querySelector('.scrub-result-meter') as HTMLElement;
	}

	private async inspectHotfix(file: File | undefined) {
		if (!file) return;
		const header = new Uint8Array(await file.slice(0, 4).arrayBuffer());
		const recognized = new TextDecoder().decode(header) === 'XFTH';
		this.hotfixResult.replaceChildren(
			<p className="scrub-message">
				{file.name}: {recognized ? 'recognized XFTH header' : 'not a recognized hotfix-cache header'}.
				Nothing was uploaded.
			</p>,
		);
	}

	/** Show and download the locally scrubbed copy. */
	private async scrubMeter(file: File | undefined) {
		if (!file) return;
		this.meterResult.replaceChildren(<p className="scrub-message">Reading {file.name}...</p>);
		try {
			const bytes = new Uint8Array(await file.arrayBuffer());
			const { scrubbed, records, namesRemoved } = scrub(bytes);
			if (!records.length) {
				this.meterResult.replaceChildren(
					<p className="scrub-warn scrub-message">
						No damage meter records in that file. If the meter had nothing in it the client writes an empty one, so fight something and copy it out
						again.
					</p>,
				);
				return;
			}
			// Summarised from the scrubbed copy, not the original. Showing the real names back
			// would say "24 names removed" above a list of those 24 names, which reads as though
			// the scrub had not worked.
			this.meterResult.replaceChildren(this.report(file.name, scrubbed, readRecords(scrubbed), namesRemoved));
		} catch (error) {
			this.meterResult.replaceChildren(<p className="scrub-warn scrub-message">Could not read that file: {String(error)}</p>);
		}
	}

	private report(name: string, scrubbed: Uint8Array, records: ActorRecord[], namesRemoved: number): Element {
		const { byActor, bySpell } = totals(records);
		const shown = records.filter(plausible).length;
		const url = URL.createObjectURL(new Blob([scrubbed as unknown as BlobPart], { type: 'application/octet-stream' }));

		return (
			<div className="scrub-report">
				<p className="scrub-report-head">
					<strong>
						{String(namesRemoved)} name{namesRemoved === 1 ? '' : 's'} removed
					</strong>{' '}
					from {String(records.length)} records. Same size, every damage number untouched.
				</p>
				<p className="scrub-actions">
					<a className="scrub-button scrub-button-quiet" href={url} download={`scrubbed-${name}`}>
						<i className="fas fa-download" />
						<span>Download locally</span>
					</a>
				</p>
				<p className="scrub-report-note">The {String(shown)} records whose numbers read cleanly are shown below.</p>
				<ul className="scrub-summary">
					{[...byActor.entries()].map(([actor, row]) =>
						summaryRow(actor, [row.className, `${row.damage.toLocaleString()} damage`, `${row.hits.toLocaleString()} hits`]),
					)}
				</ul>
				<ul className="scrub-summary">
					{[...bySpell.entries()]
						.sort((a, b) => b[1].damage - a[1].damage)
						.map(([spellId, row]) =>
							summaryRow(`spell ${spellId}`, [`${row.damage.toLocaleString()} damage`, `${row.hits.toLocaleString()} hits`]),
						)}
				</ul>
			</div>
		);
	}
}
