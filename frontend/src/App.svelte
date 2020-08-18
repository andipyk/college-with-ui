<script>
  import router from "page";
  import AdminPage from "./pages/admin/AdminPage.svelte";
  import Auth from "./pages/Auth.svelte";

  let page;

  router('/', () => (page = Auth))
  router('/admin', () => (page = AdminPage))

  router.start()
</script>


<svelte:component this={page}/>