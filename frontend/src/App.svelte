<script>
    let creators = []

    const res = fetch('http://localhost:8080/api/creators')
        .then(res => res.json())
        .then(_creators => {
            creators = _creators
        })
</script>

<main>
    <h1>Creators</h1>
    <ul>
        {#each creators as { patreonId }}
            <li>
                {patreonId}
            </li>
        {/each}
    </ul>
</main>

<style>
</style>
