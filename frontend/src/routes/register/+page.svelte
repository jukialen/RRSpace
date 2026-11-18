<script>
    import axios from 'axios';

    let name = $state('');
    let email = $state('');
    let password = $state('');
    let role = $state('client'); // client | owner
    let error = $state('');

    async function handleRegister() {
        try {
            await axios.post('/api/auth/register', { name, email, password, role });
            window.location.href = '/login';
        } catch (e) {
            error = 'Błąd rejestracji. Sprawdź dane.';
        }
    }
</script>

<div class="card">
    <h2>Rejestracja</h2>
    {#if error}<p class="error">{error}</p>{/if}
    <form on:submit|preventDefault={handleRegister}>
        <label>
            Imię
            <input type="text" bind:value={name} required />
        </label>
        <label>
            Email
            <input type="email" bind:value={email} required />
        </label>
        <label>
            Hasło
            <input type="password" bind:value={password} required />
        </label>
        <label>
            Rola
            <select bind:value={role}>
                <option value="client">Klient</option>
                <option value="owner">Właściciel studia</option>
            </select>
        </label>
        <button type="submit" class="btn btn-primary">Zarejestruj</button>
    </form>
    <p><a href="/login">Masz już konto? Zaloguj</a></p>
</div>