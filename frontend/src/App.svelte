<script>
  import Mahasiswa from "./Mahasiswa.svelte";

  import { onMount } from "svelte";

  let dataMhs;

  onMount(async () => {
    await fetch(`http://localhost:3000/mahasiswa`)
      .then((r) => r.json())
      .then((json) => {
        dataMhs = json.data;
      });
  });
</script>

<style>
  button {
    background-color: #8bd3dd;
    padding: 0.75em;
    border-radius: 0.25em;
    border: 2px solid black;
    box-shadow: 0.4rem 0.4rem 0 #222222;
  }

  button:hover {
    box-shadow: 0.25rem 0.25rem 0 #222222;
    transition: all 0.4s ease 0s;
    background-color: #8bd;
  }
</style>

<div>
  <h1>Data Mahasiswa</h1>

  {#if dataMhs}
    {#each dataMhs as mhs}
      <ul>
        <li>
          <Mahasiswa {mhs} />
        </li>
      </ul>
    {/each}
  {:else}
    <p class="loading">loading...</p>
  {/if}

</div>
