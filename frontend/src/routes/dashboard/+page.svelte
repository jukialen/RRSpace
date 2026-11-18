<script>
    import { onMount } from 'svelte';
    import axios from 'axios';

    let user = $state(null);
    let studios = $state([]);

    onMount(async () => {
        const token = localStorage.getItem('token');
        if (!token) {
            window.location.href = '/login';
            return;
        }
        try {
            const res = await axios.get('/api/users/me', {
                headers: { Authorization: `Bearer ${token}` },
            });
            user = res.data;
            if (user.role === 'owner') {
                const sRes = await axios.get('/api/studios', {
                    headers: { Authorization: `Bearer ${token}` },
                });
                studios = sRes.data;
            }
        } catch {
            window.location.href = '/login';
        }
    });
</script>

{#if user}
    <div class="dashboard">
        <h1>Witaj, {user.name}!</h1>
        {#if user.role === 'owner'}
            <section>
                <h2>Moje studia</h2>
                <a href="/studios/new" class="btn btn-primary">+ Nowe studio</a>
                {#each studios as studio}
                    <a href="/studios/{studio.id}" class="card">{studio.name}</a>
                {/each}
            </section>
        {:else}
            <section>
                <h2>Przeglądaj studia</h2>
                <a href="/studios" class="btn">Zobacz studia</a>
            </section>
        {/if}
    </div>
{/if}