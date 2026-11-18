declare module '*.svg?component' {
	import type { Component } from 'svelte';
	const content: typeof Component;
	export default content;
}

declare module '*.svg' {
	const content: string;
	export default content;
}
