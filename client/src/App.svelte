<script>
// @ts-nocheck

  import { Router, Route, Link } from "svelte-routing";
  import Login from "./layouts/Login.svelte";
  import Gender from "./layouts/Gender.svelte";
  import Buy from "./layouts/Buy.svelte";
  import Sell from "./layouts/sell.svelte";
  import Navbar from "./components/Navbar.svelte";

</script>

<Router>

  <main>
    <Route path="login"> <Login/> </Route>
    <Route path="navbar"> <Navbar/> </Route>
    <Route path="gender"> <Gender/> </Route>
    <Route path="sell"> <Sell/> </Route>
    <Route path="buy"> <Buy/> </Route>
  </main>
</Router>

<style></style>
