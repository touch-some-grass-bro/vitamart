<script>
  import axios from "axios";
  import { link, navigate } from "svelte-routing";

  import API_URL from "../api";

  const logoutOutHandler = () => {
    window.location.href = `${API_URL}/auth/logout`;
  };
</script>

<main>
  <div class="header">
    <div class="header-content align-text-top">
      <div><img class="h-36 my-0" src="../public/LOGO.png" alt="Logo" /></div>
      <div class="my-auto"><h1 class="text-5xl font-bold">VITAMART</h1></div>
      <div class="my-auto ml-auto flex">
        <a href="/gender" use:link>Change Hostel Details</a>
        <button on:click={logoutOutHandler}>Sign Out</button>
      </div>
    </div>
  </div>
</main>

<style>
  @import url("https://fonts.googleapis.com/css2?family=Skranji&display=swap");

  .header {
    border-radius: 0 0 920px 0;
    background-image: url("/public/sell.jpg");
    background-size: cover;
    height: 650px;
    display: flex;
    align-items: center;
    padding-bottom: 27%;
    margin-bottom: 50px;
  }

  .header-content {
    display: flex;
    flex-grow: 1;
    justify-content: space-between;
  }

  h1 {
    font-family: "SKranji", cursive;
    font-size: 50px;
    color: #2b6969;
  }

  a {
    font-size: 30px;
    text-decoration: none;
    color: #2b6969;
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
    background-color: #2b6969; /* Set the desired underline color */
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
    color: #2b6969;
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
    background-color: #2b6969; /* Set the desired underline color */
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

