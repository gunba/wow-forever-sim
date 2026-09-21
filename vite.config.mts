/** @type {import('vite').UserConfig} */

import { execFileSync } from 'child_process';
import fs from 'fs';
import { IncomingMessage, ServerResponse } from 'http';
import path from 'path';
import { ConfigEnv, defineConfig, PluginOption, UserConfigExport } from 'vite';
import { checker } from 'vite-plugin-checker';

import { specPages } from './tools/vite/spec_pages.mjs';

export const BASE_PATH = path.resolve(__dirname, 'ui');
export const OUT_DIR = path.join(__dirname, 'dist', 'classic');

// Where the site will be served from, with a trailing slash. Defaults to serving from the
// root; override with SITE_BASE to hang it off a path, e.g. a Github project
// page at /<repo>/classic/. Everything that needs the prefix reads it from here: the
// TypeScript through import.meta.env.BASE_URL, the stylesheets through $site-base, and
// the page template through the specPages plugin.
export const SITE_BASE = process.env.SITE_BASE || '/classic/';

// The Github repository the UI points at for source, issues, crash reports and releases.
// Defaults to this fork; the deploy workflow passes the repository it is running in, and a
// local build or another fork overrides it
// with SITE_REPO=<owner>/<repo>; the deploy workflow passes the repository it is running
// in. Read it through SITE_REPO in core/constants/other.ts for TypeScript, and through the
// @@REPO@@ placeholder for the hand-written homepage.
export const SITE_REPO = process.env.SITE_REPO || 'gunba/wow-forever-sim';

// The version the UI shows, so a user can say which build they are looking at. Taken from
// git rather than hand-maintained: the newest tag plus the commits since it, or the bare
// commit when the fork has no tags yet, with '-dirty' appended for an uncommitted build.
// Builds from outside a git checkout - a source tarball, a Docker image that only copies
// the tree - have nothing to read, so they fall back to 'unknown' instead of failing.
export const SITE_VERSION = (() => {
	try {
		return execFileSync('git', ['describe', '--tags', '--always', '--dirty'], {
			cwd: __dirname,
			encoding: 'utf-8',
			stdio: ['ignore', 'pipe', 'ignore'],
		}).trim();
	} catch {
		// Not a git checkout, or no git on PATH.
		return 'unknown';
	}
})();

function serveExternalAssets() {
	const workerMappings = {
		'/classic/sim_worker.js': '/classic/local_worker.js',
		'/classic/net_worker.js': '/classic/net_worker.js',
		'/classic/lib.wasm': '/classic/lib.wasm',
	};

	return {
		name: 'serve-external-assets',
		configureServer(server) {
			server.middlewares.use((req, res, next) => {
				const url = req.url!;

				if (Object.keys(workerMappings).includes(url)) {
					const targetPath = workerMappings[url as keyof typeof workerMappings];
					const assetsPath = path.resolve(__dirname, './dist/classic');
					const requestedPath = path.join(assetsPath, targetPath.replace('/classic/', ''));
					serveFile(res, requestedPath);
					return;
				}

				if (url.includes('/classic/assets')) {
					const assetsPath = path.resolve(__dirname, './assets');
					const assetRelativePath = url.split('/classic/assets')[1];
					const requestedPath = path.join(assetsPath, assetRelativePath);

					serveFile(res, requestedPath);
					return;
				} else {
					next();
				}
			});
		},
	} satisfies PluginOption;
}

function serveFile(res: ServerResponse<IncomingMessage>, filePath: string) {
	if (fs.existsSync(filePath)) {
		const contentType = determineContentType(filePath);
		res.writeHead(200, { 'Content-Type': contentType });
		fs.createReadStream(filePath).pipe(res);
	} else {
		console.log('Not found on filesystem: ', filePath);
		res.writeHead(404, { 'Content-Type': 'text/plain' });
		res.end('Not Found');
	}
}

function determineContentType(filePath: string) {
	const extension = path.extname(filePath).toLowerCase();
	switch (extension) {
		case '.jpg':
		case '.jpeg':
			return 'image/jpeg';
		case '.png':
			return 'image/png';
		case '.gif':
			return 'image/gif';
		case '.css':
			return 'text/css';
		case '.js':
			return 'text/javascript';
		case '.woff':
		case '.woff2':
			return 'font/woff2';
		case '.json':
			return 'application/json';
		case '.wasm':
			return 'application/wasm'; // Adding MIME type for WebAssembly files
		// Add more cases as needed
		default:
			return 'application/octet-stream';
	}
}

export const getBaseConfig = ({ command, mode }: ConfigEnv) =>
	({
		base: SITE_BASE,
		root: path.join(__dirname, 'ui'),
		define: {
			__SITE_VERSION__: JSON.stringify(SITE_VERSION),
			__SITE_REPO__: JSON.stringify(SITE_REPO),
		},
		build: {
			outDir: OUT_DIR,
			minify: mode === 'development' ? false : 'terser',
			sourcemap: command === 'serve' ? 'inline' : false,
			target: ['es2020'],
		},
	}) satisfies Partial<UserConfigExport>;

export default defineConfig(({ command, mode }) => {
	const baseConfig = getBaseConfig({ command, mode });
	return {
		...baseConfig,
		plugins: [
			// The homepage is hand-written rather than generated from ui/index_template.html,
			// so substitute its repository placeholder here.
			{
				name: 'site-repo-html',
				transformIndexHtml: (html: string) => html.replaceAll('@@REPO@@', SITE_REPO),
			},
			specPages(BASE_PATH, SITE_BASE),
			serveExternalAssets(),
			checker({
				root: path.resolve(__dirname, 'ui'),
				typescript: true,
				enableBuild: true,
			}),
		],
		esbuild: {
			jsxInject: "import { element, fragment } from 'tsx-vanilla';",
		},
		css: {
			preprocessorOptions: {
				scss: {
					additionalData: `$site-base: '${SITE_BASE}';`,
				},
			},
		},
		build: {
			...baseConfig.build,
			rollupOptions: {
				// The per-page entries are added by the specPages plugin.
				input: {
					'ui/index.html': path.resolve(BASE_PATH, 'index.html'),
				},
				output: {
					assetFileNames: () => 'bundle/[name]-[hash].style.css',
					entryFileNames: () => 'bundle/[name]-[hash].entry.js',
					chunkFileNames: () => 'bundle/[name]-[hash].chunk.js',
				},
			},
			server: {
				origin: 'http://localhost:3000',
				// Adding custom middleware to serve 'dist' directory in development
			},
		},
	};
});
