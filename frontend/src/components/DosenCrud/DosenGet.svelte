<script>
    export let dataDosen;
</script>

<style>

</style>

<div>
    <h1>Data Dosen</h1>
  
    {#if dataDosen}
      {#each dataDosen as dosen}
        <ul>
          <li>
            <h1>{dosen.nama}</h1>
            <small>
              NIK:
              <b>{dosen.nik}</b>
            </small>
          </li>
        </ul>
      {/each}
    {:else}
      <p class="loading">loading...</p>
    {/if}
  
  </div>