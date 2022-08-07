<link rel="stylesheet" href="/style.css"/>

<script>
    import Nav from './Nav.svelte';

    export let editor = {}
</script>

<Nav/>

<main>
    <header>
        <h2>{editor.name}</h2>
        <h3>{editor.description}</h3>
    </header>
    <section>
        <aside>
        </aside>
    </section>
</main>

<br/>
