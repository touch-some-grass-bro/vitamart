<script>
  import TailwindCss from "../lib/TailwindCSS.svelte";
</script>
  <TailwindCss/>
<main>
  <header class="flex bg-white justify-between fixed w-full">
      <div class="mx-5"><button class="funt">VITAMART</button></div>
      <div class="mx-5 my-auto">
          <a href="#" class="px-7 text-2xl">Login</a>
          <button class="rounded-xl started text-white px-5 py-3">Get Started</button>
      </div>
  </header>
  <div class="h-16"></div>
  <div class="bluediv bg-blue-300 flex flex-col  items-center">
      <h1 class="text-white text-center py-20 text-8xl">Feast. Share. Connect</h1>
      <p class="text-white text-center py-1 text-2xl mx-auto">Introducing VitaMart, your one-stop online marketplace tailored exclusively for the vibrant community within our college hostel.</p>
      <button class="rounded-xl started text-white px-5 py-3 mt-2">Get Started</button>
  </div>

  <div class="flex justify-between mx-5 py-5 gap-5">
  <div class="text-left smoldiv text-white text-3xl">  
    <h3>
    Key Features:
    <br>
    <br>
    1. Easy Registration: Signing up for VitaMart is a breeze. Use your college email to create an account and join the bustling virtual marketplace. 
    <br>
    2. Customized Listings: List your items with ease, adding detailed descriptions, high-quality photos, and your asking price.
    <br>
    3.Secure Messaging: Connect with buyers or sellers through our built-in messaging system.
    <br>
    4.Verified Profiles: Trust and transparency are essential.
    <br>
    <br>
    Whether you're looking to satisfy your hunger pangs ,find affordable textbooks or get your DA's completed. VitaMart is the digital bazaar that connects buyers and sellers within the cozy confines of your campus accommodation.
  </h3>
  </div>
  <div class="text-center smoldiv text-white text-3xl">
    <h3>Lorem ipsum dolor sit amet consectetur adipisicing elit. Culpa modi iure consectetur, cumque dicta impedit amet. Necessitatibus laboriosam earum ratione ipsa cumque ipsam quibusdam, perspiciatis libero, qui soluta molestias consequatur!</h3>
  </div>
</div>

</main>

<style>
 @import url('https://fonts.googleapis.com/css2?family=Skranji&display=swap');
 .smoldiv{

 }
  .funt{
      padding-left: 1%;
      font-family: "SKranji",cursive;
    font-size: 50px;
    color:rgb(193,226,0);
  }
  .started{
      background-color: rgb(193,226,0);
  }

  p{
      max-width: 500px;
  }

  .bluediv{
      height: 700px;
  }
</style>