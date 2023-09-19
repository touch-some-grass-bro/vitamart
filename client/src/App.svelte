<script>
// @ts-nocheck


  import { Router, Route, Link } from "svelte-routing";
  import Login from "./layouts/Login.svelte";
  import HostelPage from "./layouts/HostelPage.svelte";
</script>

<Router>

{#if location.pathname !== '/login' && location.pathname !== '/hostel'}
  <Link to="/login">LOGIN</Link>
{/if}

{#if location.pathname !== '/login' && location.pathname !== '/hostel'}
  <Link to="/hostel">HOSTEL</Link>
{/if}

  <main>
    <Route path="login"> <Login/> </Route>
    <Route path="hostel"> <HostelPage/> </Route>
  </main>
</Router>

<style></style>
