<link rel="stylesheet" href="/style.css"/>

<script>
  import Nav from './Nav.svelte';
</script>

<Nav/>

<h1>New Editor</h1>

<form method="post" action={`/editor`}>
  <!-- Add input fields here -->
  <input type="submit" value="Create Editor" />
</form>

<br />

<a href={`/editor`}>Back</a>
