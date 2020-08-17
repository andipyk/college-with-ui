<script>
    export let dataMhs;
</script>

<style>

</style>

<div>
    <h1>Data Mahasiswa</h1>
  
    {#if dataMhs}
      {#each dataMhs as mhs}
        <ul>
          <li>
            <h1>{mhs.nama}</h1>
            <small>
              NIM:
              <b>{mhs.nim}</b>
            </small>
          </li>
        </ul>
      {/each}
    {:else}
      <p class="loading">loading...</p>
    {/if}
  
  </div>