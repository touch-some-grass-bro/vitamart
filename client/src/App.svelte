<script>
// @ts-nocheck


  import { Router, Route, Link } from "svelte-routing";
  import Login from "./layouts/Login.svelte";
  import HostelPage from "./layouts/HostelPage.svelte";
  import Gender from "./layouts/Gender.svelte";
  import Sell from "./layouts/sell.svelte";
</script>

<Router>

{#if location.pathname !== '/login' && location.pathname !== '/hostel'}
  <Link to="/login">LOGIN</Link>
{/if}

{#if location.pathname !== '/login' && location.pathname !== '/hostel'}
  <Link to="/hostel">HOSTEL</Link>
{/if}
{#if location.pathname !== '/login' && location.pathname !== '/hostel'}
  <Link to="/gender">Gender</Link>
{/if}

  <main>
    <Route path ="sell"> <Sell/></Route>
    <Route path="login"> <Login/> </Route>
    <Route path="hostel"> <HostelPage/> </Route>
    <Route path="gender"> <Gender/> </Route>
  </main>
</Router>

<style></style>
