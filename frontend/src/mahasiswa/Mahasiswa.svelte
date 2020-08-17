<script>
  let dataMhs;

  import { onMount } from "svelte";
  import MhsGet from "./MhsGet.svelte";
import MhsCreate from "./MhsCreate.svelte";
import MhsEdit from "./MhsEdit.svelte";

  onMount(async () => {
    await fetch(`http://localhost:3000/admin/list_mahasiswa`)
      .then((r) => r.json())
      .then((json) => {
        dataMhs = json.data;
      });
  });
</script>

<style>

</style>

<div>
  <MhsCreate />
  <MhsEdit {dataMhs} />
  <MhsGet {dataMhs} />
</div>
