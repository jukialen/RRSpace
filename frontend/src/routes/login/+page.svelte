<script>
    import { enhance } from '$app/forms';
    import { goto, invalidate } from '$app/navigation';
    import type { PageProps } from './$types';
    import { m } from '$lib/paraglide/messages.js';

    import styles from './+page.module.css'
    import ShowPass from '$lib/assets/icons/show-pass.svg?component';
    import HidePass from '$lib/assets/icons/hide-pass.svg?component';

    let { data, form }: PageProps = $props();

    let email = $state('');
    let password = $state('');
    let showPass = $state(true);

    let emailInput: HTMLInputElement;
    let passwordInput: HTMLInputElement;

    $effect(() => {
        if (form?.success) {
            email = '';
            password = '';
            invalidate('backend:auth');

            //time for change to spread
            new Promise((resolve) => setTimeout(resolve, 50));

            goto(data.getFirstLogin ? '/chat' : '/avatar');
        }

        if (form?.errors?.email) {
            emailInput?.focus();
        } else if (form?.errors?.password) {
            passwordInput?.focus();
        }
    });

    const changePassVisibility = () => (showPass = !showPass);

</script>

<div class="card">
    <h2>Logowanie</h2>
    <form method="post" use:enhance novalidate>
        <input
                type="email"
                placeholder='E-mail'
                name="email"
                bind:value={email}
                class={form?.errors.email ? styles.input__error : ''}
                bind:this={emailInput}
        />
        {#if !!form?.errors?.email}
            <p class={styles.error}>{form?.errors.email}</p>
        {/if}
        <div class={styles.passContainer}>
            <input
                    type={showPass ? 'password' : 'text'}
                    placeholder={m.form_password()}
                    name="password"
                    bind:value={password}
                    class={form?.errors.password ? styles.input__error : ''}
                    bind:this={passwordInput}
            />
            <button
                    class={styles.passIcon}
                    onclick={changePassVisibility}
                    aria-label="visibly password button"
            >
                {#if showPass}
                    <ShowPass />
                {:else}
                    <HidePass />
                {/if}
            </button>
        </div>
        {#if !!form?.errors?.password}
            <p class={styles.error}>{form.errors.password}</p>
        {/if}

        {#if !!form?.success || !!form?.errors || !!form?.message}
            <Toast message={form?.message} />
        {/if}

        <button class={styles.loginButton} type="submit">
            {m.form_signin()}
        </button>
    </form>
    <p><a href="/register">Nie masz konta? Zarejestruj się</a></p>
</div>