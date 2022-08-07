<link rel="stylesheet" href="/style.css"/>

<script>
    import Nav from './Nav.svelte';

    export let editors = []
</script>

<Nav/>

<main>
    <header>
        <h2>Onboarding</h2>
    </header>
    <section>
        {#each editors as editor}
            <aside>
                <h3>{editor.name}</h3>
                <p>{editor.description}</p>
                <a href={`/editor/${editor.id || 0}/edit`}>
                    <p>Edit</p>
                </a>
            </aside>
        {/each}
        <aside>
            <h3><a href={`/editor/new`}>Add a new onboarding</a></h3>
        </aside>

    </section>
</main>

<br/>

