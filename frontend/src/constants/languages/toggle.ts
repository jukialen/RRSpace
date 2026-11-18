import { get, readonly, writable } from 'svelte/store';

export const langActive = writable<boolean>(false);
export const langActiveReadOnly = readonly(langActive);
export const langActiveWriteOnlyCurrentValue = get(langActiveReadOnly);
