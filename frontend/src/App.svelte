<script>
  import router from "page";
  import Auth from "./pages/Auth.svelte";
  import AdminPage from "./pages/admin/AdminPage.svelte";
  import DosenPage from "./pages/dosen/DosenPage.svelte";
  import MahasiswaPage from "./pages/mahasiswa/MahasiswaPage.svelte";

  let page;

  router("/", () => (page = Auth));
  router("/admin", () => (page = AdminPage));
  router("/dosen", () => (page = DosenPage));
  router("/mahasiswa", () => (page = MahasiswaPage));

  router.start();
</script>

<svelte:component this={page} />
