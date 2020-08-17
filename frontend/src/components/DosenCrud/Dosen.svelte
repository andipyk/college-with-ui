<script>
    let dataDosen;
  
    import { onMount } from "svelte";
    import DosenGet from "./DosenGet.svelte";
    import DosenCreate from "./DosenCreate.svelte";
    import DosenEdit from "./DosenEdit.svelte";
  
    onMount(async () => {
      await fetch(`http://localhost:3000/admin/list_dosen`)
        .then((r) => r.json())
        .then((json) => {
          dataDosen = json.data;
        });
    });
  </script>
  
  <style>
  
  </style>
  
  <div>
    <DosenCreate />
    <DosenEdit {dataDosen} />
    <DosenGet {dataDosen} />
  </div>
  