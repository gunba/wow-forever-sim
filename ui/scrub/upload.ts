// This fork has no upload collector. Local scrubbing/downloads remain available.
export type UploadKind = 'dbcache' | 'damagemeter';
export type UploadResult = { ok: true; receipt: string } | { ok: false; error: string };

export async function upload(_kind: UploadKind, _bytes: Uint8Array, _note: string): Promise<UploadResult> {
	return { ok: false, error: 'Uploads are not configured for this site. Download the file locally instead.' };
}
