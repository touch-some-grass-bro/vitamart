<script>
  import axios from "axios";
  import { link, navigate } from "svelte-routing";
  import API_URL from "../api";

  let img;

  const logoutOutHandler = () => {
    window.location.href = `${API_URL}/auth/logout`
  };
</script>

<main>
  <header class="header">
    <div><img class="h-36 my-0" src="../public/LOGO.png" alt="Logo" /></div>
    <div class="my-auto"><h1 class="text-5xl font-bold">VITAMART</h1></div>
    <div class="my-auto ml-auto flex">
      <a href="/gender" use:link>Change Hostel Details</a>
      <button on:click={logoutOutHandler}>Sign Out</button>
    </div>
  </header>
</main>

<style>
  @import url("https://fonts.googleapis.com/css2?family=Skranji&display=swap");

  .header {
    display: flex;
    align-items: center;
    background-color: transparent;
    position: relative;
    background-color: #2B6969;
  }

  .bgimg {
    width: 100%;
    height: auto;
    position: absolute;
    z-index: -1; /* Place the image behind the header content */
    border-radius: 0 0 50px 0;
  }

  .header-content {
    display: flex;
    flex-grow: 1; /* Allow the header content to grow and fill the available space */
    justify-content: space-between;
    /* padding: 20px; Adjust the padding as needed */
    /* background-color: rgba(0, 0, 0, 0.7); Add a background color for the header content */
  }

  h1 {
    font-family: "SKranji", cursive;
    font-size: 50px;
    color: #edf2f2;
  }

  a {
    font-size: 30px;
    text-decoration: none;
    color: white;
    position: relative;
    margin: auto 15px;
  }

  a::before {
    content: "";
    position: absolute;
    width: 100%;
    height: 2px; /* Set the desired underline thickness */
    bottom: 0;
    left: 0;
    background-color: white; /* Set the desired underline color */
    transform: scaleX(0); /* Initial state: no underline */
    transform-origin: bottom right;
    transition: transform 0.3s ease; /* Add a smooth transition */
  }

  /* Animate the underline on hover */
  a:hover::before {
    transform: scaleX(1); /* Fully underline on hover */
    transform-origin: bottom left;
  }

  a:hover {
    cursor: pointer;
  }

  button {
    font-size: 30px;
    text-decoration: none;
    color: white;
    position: relative;
    margin: auto 15px;
  }

  button::before {
    content: "";
    position: absolute;
    width: 100%;
    height: 2px; /* Set the desired underline thickness */
    bottom: 0;
    left: 0;
    background-color: white; /* Set the desired underline color */
    transform: scaleX(0); /* Initial state: no underline */
    transform-origin: bottom right;
    transition: transform 0.3s ease; /* Add a smooth transition */
  }

  /* Animate the underline on hover */
  button:hover::before {
    transform: scaleX(1); /* Fully underline on hover */
    transform-origin: bottom left;
  }

  button:hover {
    cursor: pointer;
  }
</style>

