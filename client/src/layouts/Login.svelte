<script>
  import TailwindCss from '../lib/TailwindCSS.svelte';
  import axios from 'axios'
  import API_URL from '../api.js'

  const signInWithGoogle = () => {
    axios
      .get(API_URL + '/auth/url')
      .then((res) => {
        // console.log
        window.location.href = res.data.url
      })
      .catch((err) => {
        console.log(err)
      })
  }
</script>

<TailwindCss/>

<main class="py-20">
  <div class="logindiv py-10 px-16 rounded-lg">
      <h1 class="">VITAMART</h1>
      <p class="text-white">Feast. Share. Connect</p>
      <img class="h-80 m-auto" src="../public/Logo.svg">
      <button class="px-36 py-4 m-auto text-white" on:click={signInWithGoogle}>SignIn with Google</button>
  </div>
</main>

<style>

  .logindiv {
      max-width: 1280px;
      margin: 0 auto;
      padding: 2rem;
      text-align: center;
      box-shadow: 9px 10px 7px 1px rgba(38, 38, 38, 0.94);
      background-color: rgb(46, 79, 79);
  }

  button:hover {
      box-shadow: 9px 10px 7px 1px rgba(38, 38, 38, 0.54);
      transition: box-shadow 0.25s ease-in-out;
  }

  button {
      background-color: rgb(193, 226, 0);
  }

  h1 {
      font-size: 50px;
      color: rgb(193, 226, 0);
  }
</style>
