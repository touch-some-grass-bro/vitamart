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
    
    <p class="text"> Hungry,Grab a snickers</p>


</main>

<style>
   @import url('https://fonts.googleapis.com/css2?family=Skranji&display=swap');
   p.text {
             writing-mode: vertical-rl;
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