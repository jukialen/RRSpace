/// <reference types="@sveltejs/kit" />
declare let self: ServiceWorkerGlobalScope;

import { build, files, version } from '$service-worker';
import { isStaticPath } from '$lib/helpers/paths';

const CACHE = `cache-${version}`;
const ASSETS = [...build, ...files];

const cacheWhitelist = [
	'https://www.googletagmanager.com',
	'https://cdn-cookieyes.com',
	'https://fonts.gstatic.com',
	import.meta.env.VITE_SUPABASE_URL
];

if (!import.meta.env.DEV) {
	self.addEventListener('install', (event: ExtendableEvent) => {
		self.skipWaiting();
		event.waitUntil(
			caches.open(CACHE).then(async (cache) => {
				for (const url of ASSETS) {
					try {
						const response = await fetch(url);
						if (response.ok) await cache.put(url, response);
					} catch (err) {
						console.warn(`⚠️ Skipped asset: ${url}`, err);
					}
				}
			})
		);
	});

	self.addEventListener('activate', (event: ExtendableEvent) => {
		event.waitUntil(
			(async () => {
				await self.clients.claim();
				const keys = await caches.keys();
				await Promise.all(keys.map((key) => key !== CACHE && caches.delete(key)));
			})()
		);
	});

	self.addEventListener('fetch', (event: FetchEvent) => {
		const url = new URL(event.request.url);

		// 1. Exclusions (Auth, API, non-GET)
		if (
			url.pathname === '/chat' ||
			url.pathname.startsWith('/auth') ||
			url.pathname.startsWith('/api') ||
			event.request.method !== 'GET' ||
			url.pathname.startsWith('/.svelte-kit')
		) {
			return;
		}

		// 2. Assets (Cache-first)
		if (ASSETS.includes(url.pathname) || isStaticPath(url.pathname)) {
			event.respondWith(
				caches.match(event.request).then((cached) => {
					return (
						cached ||
						fetch(event.request).then(async (res) => {
							if (res.ok) {
								const cache = await caches.open(CACHE);
								await cache.put(event.request, res.clone());
							}
							return res;
						})
					);
				})
			);
			return;
		}

		// 3. HTML & Navigation Data (Network-first with bypass support)
		if (
			event.request.mode === 'navigate' ||
			url.pathname.endsWith('__data.json') ||
			url.search.includes('__data.json')
		) {
			event.respondWith(
				(async () => {
					const cache = await caches.open(CACHE);
					const cachedResponse = await cache.match(event.request);

					try {
						const res = await fetch(event.request);

						// Bypass for redirection statuses and authorisation errors
						if (
							[0, 301, 302, 303, 307, 308, 401, 403].includes(res.status) ||
							res.type === 'opaqueredirect'
						)
							return res;

						if (res.ok && res.status === 200) {
							if (res.headers.get('Vary') !== '*') {
								try {
									await cache.put(event.request, res.clone());
								} catch (e) {
									console.warn('Caching skipped (put error)');
								}
							}
						}
						return res;
					} catch (err) {
						if (cachedResponse) return cachedResponse;
						return new Response('Brak połączenia z internetem', {
							status: 503,
							headers: { 'Content-Type': 'text/html' }
						});
					}
				})()
			);
			return;
		}

		// 4. External CDNs
		if (cacheWhitelist.some((domain) => domain && url.origin.includes(domain))) {
			event.respondWith(
				(async () => {
					const cache = await caches.open(CACHE);
					try {
						const response = await fetch(event.request);
						if (response.ok) await cache.put(event.request, response.clone());
						return response;
					} catch (err) {
						const cached = await cache.match(event.request);
						return cached || new Response('Offline resource', { status: 408 });
					}
				})()
			);
		}
	});
}
