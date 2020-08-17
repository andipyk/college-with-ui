<script>
    let dataKelas;
  
    import { onMount } from "svelte";
    import KelasGet from "./KelasGet.svelte";
    import KelasCreate from "./KelasCreate.svelte";
  
    onMount(async () => {
      await fetch(`http://localhost:3000/admin/list_kelas`)
        .then((r) => r.json())
        .then((json) => {
          dataKelas = json.data;
        });
    });
  </script>
  
  <style>
  
  </style>
  
  <div>
    <KelasCreate />
    <KelasGet {dataKelas} />
  </div>
  