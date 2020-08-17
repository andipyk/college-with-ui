<script>
import Button from "../components/Button.svelte";

  let nama = "";
  let nim = "";

  async function doPost() {
    const res = await fetch("http://localhost:3000/admin/input/mahasiswa", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        nama,
        nim,
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
  <h1>Form</h1>
  <form class="content">
    <label>Nama</label>
    <input type="text" bind:value={nama} />
    <label>NIM</label>
    <input type="text" bind:value={nim} />
    <Button on:click={doPost} btnText={"Tambahkan"}/>
  </form>
</div>
