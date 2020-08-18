<script>
    import CrudMahasiswa from "../../components/MhsCrud/Mahasiswa.svelte";
    import CrudDosen from "../../components/DosenCrud/Dosen.svelte";
    import CrudKelas from "../../components/KelasCrud/Kelas.svelte";
   
  
  </script>
  
  <style>
  </style>
  
  <div>
    <h1>=== Mahasiswa ===</h1>
    <CrudMahasiswa/>
    
    <h1>=== Dosen ===</h1>
    <CrudDosen />

    <h2>=== Kelas ===</h2>
    <CrudKelas />
  </div>
  