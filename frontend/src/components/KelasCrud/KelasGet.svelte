<script>
    export let dataKelas;
</script>

<style>

</style>

<div>
    <h1>Data Kelas</h1>
  
    {#if dataKelas}
      {#each dataKelas as kelas}
        <ul>
          <li>
            <h1>{kelas.nama}</h1>
            <small>
              NIK:
              <b>{kelas.kode}</b>
            </small>
          </li>
        </ul>
      {/each}
    {:else}
      <p class="loading">loading...</p>
    {/if}
  
  </div>