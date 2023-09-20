<script>
// @ts-nocheck

  import { Router, Route, Link } from "svelte-routing";
  import Login from "./layouts/Login.svelte";
  import HostelPage from "./layouts/HostelPageog.svelte";
  import Gender from "./layouts/Gender.svelte";
  import Buy from "./layouts/Buy.svelte";

</script>

<Router>

  <main>
    <Route path ="sell"> <Sell/></Route>
    <Route path="login"> <Login/> </Route>
    <Route path="gender"> <Gender/> </Route>
    <Route path="buy"> <Buy/> </Route>
  </main>
</Router>

<style></style>
