<script>
  export let dataDosen;

  let nama = "";
  let nik = "";

  async function doUpdate() {
    const res = await fetch("http://localhost:3000/admin/update/dosen", {
      method: "PUT",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        nama,
        nik,
      }),
    });

    const json = await res.json();
    result = JSON.stringify(json);
  }
</script>

<style>
  .content {
    display: grid;
    grid-template-columns: 10% 50%;
    grid-column-gap: 10px;
  }
</style>


<div>
  <h1>Form Edit</h1>
  <form class="content">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>Nama</label>
    <input type="text" bind:value={nama} />
    
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label>nik</label>
    <select name="nik" bind:value={nik}>
      {#if dataDosen}
        {#each dataDosen as dosen}
          <option value={dosen.nik}>{dosen.nik}</option>
        {/each}
      {:else}<option disabled>loading....</option>{/if}
    </select>
    <button type="button" on:click={doUpdate}>Simpan</button>
  </form>
</div>
