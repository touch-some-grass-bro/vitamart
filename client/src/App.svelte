<script>
// @ts-nocheck

  import { Router, Route, Link } from "svelte-routing";
  import Login from "./layouts/Login.svelte";
  import Gender from "./layouts/Gender.svelte";
  import Sell from "./layouts/sell.svelte";

</script>

<Router>

  <main>
    <Route path="login"> <Login/> </Route>
    <Route path="gender"> <Gender/> </Route>
    <Route path="sell"> <Sell/> </Route>
  </main>
</Router>

<style></style>
