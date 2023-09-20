<script>
  import axios from "axios";
  import TailwindCss from "../lib/TailwindCSS.svelte";
  import API_URL from "../api";

  const handleLogin = () => {
    axios.get(API_URL + "/auth/url").then((res) => {
      window.location.href = res.data.url;
    });
  };
</script>

<TailwindCss />
<main>
  <header class="Header flex justify-between fixed w-full">
    <div class="mx-5">
      <button class="funt flex"
        ><img src="../public/LOGO.png" class="h-20" />VITAMART</button
      >
    </div>
    <div class="mx-5 my-auto">
      <button on:click={handleLogin} class="px-7 text-2xl login">Login</button>
    </div>
  </header>
  <div class="h-16" />
  <div class="bluediv flex flex-col items-center">
    <h1 class="text-black text-center py-20 text-8xl">
      Feast. Share. Connect.
    </h1>
    <p class="text-black text-center py-1 text-2xl mx-auto">
      Introducing VitaMart, your one-stop online marketplace tailored
      exclusively for the vibrant community within our college hostel.
    </p>
    <button
      on:click={handleLogin}
      class="started rounded-xl text-white px-5 py-3 mt-2">Get Started</button
    >
  </div>

  <div><h3 class="font-bold text-white">Key Features:</h3></div>
  <div class="flex justify-between mx-5 py-5 gap-5">
    <div class="text-left smoldiv text-white text-2xl">
      <h2>
        <div>
          1. Easy Registration: Signing up for VitaMart is a breeze. Use your
          college email to create an account and join the bustling virtual
          marketplace.
          <br />
          2. Customized Listings: List your items with ease, adding detailed descriptions,
          high-quality photos, and your asking price.
          <br />
          3.Secure Messaging: Connect with buyers or sellers through our built-in
          messaging system.
          <br />
          4.Verified Profiles: Trust and transparency are essential.
          <br />
          <br />
          Whether you're looking to satisfy your hunger pangs ,find affordable textbooks
          or get your DA's completed. VitaMart is the digital bazaar that connects
          buyers and sellers within the cozy confines of your campus accommodation.
        </div>
      </h2>
    </div>
    <div class="text-left smoldiv2 text-white text-2xl">
      <h2>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Culpa modi iure
        consectetur, cumque dicta impedit amet. Necessitatibus laboriosam earum
        ratione ipsa cumque ipsam quibusdam, perspiciatis libero, qui soluta
        molestias consequatur!
      </h2>
    </div>
  </div>
</main>

<style>
  @import url("https://fonts.googleapis.com/css2?family=Skranji&display=swap");
  @import url("https://fonts.googleapis.com/css2?family=Montserrat:wght@300&display=swap");
  main {
    background-color: #2b6969;
  }
  h3,
  h2 {
    font-family: "Montserrat", sans-serif;
  }
  h3 {
    margin-top: 20px;
    font-size: 70px;
  }
  .Header {
    background-color: #2b6969;
  }
  .funt {
    padding-left: 1%;
    font-family: "SKranji", cursive;
    font-size: 50px;
    color: white;
  }
  .started {
    background-color: #5a8181;
  }

  p {
    max-width: 500px;
  }

  .bluediv {
    height: 700px;
    background-color: #edf2f2;
  }

  .login {
    color: #edf2f2;
  }

  .login:hover {
    color: #5a8181;
  }
  .smoldiv2 {
  }
</style>
