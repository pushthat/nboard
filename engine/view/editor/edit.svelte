<link rel="stylesheet" href="/style.css"/>

<script>
    import Nav from './Nav.svelte';
    import Arrow from './Arrow.svelte';

    export let editor = {}

    function switchTextToField(toto) {
        alert('clicked ' + toto)
    }
</script>

<Nav/>


<main>
    <header>
        <sup>Click to edit</sup> ✍
        <h2 id="name">{editor.name}</h2>
        <h3 id="description" on:click={switchTextToField("lolito")}>{editor.description}</h3>
    </header>
    <section>
        <ul>
            {#each editor.tasks as task, i}
                {#if i === 0}
                    <aside id="{task.id}">
                        <h3>{task.name}</h3>
                        <p>{task.description}</p>
                    </aside>
                {:else}
                    <Arrow/>
                    <aside id="{task.id}">
                        <h3>{task.name}</h3>
                        <p>{task.description}</p>
                    </aside>
                {/if}
            {/each}
            <Arrow/>
            <aside id="new">
                <h3>Add a new task</h3>
                <p>click to add a new</p>
            </aside>
        </ul>
    </section>
</main>

<br/>


<style>
    sup {
        font-size: medium;
    }
</style>
