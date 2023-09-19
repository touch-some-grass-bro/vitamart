<script>
  import TailwindCss from './lib/TailwindCSS.svelte';
</script>
<TailwindCss/>

<main>

  <header class="flex h-60px align-middle">
    <div><img class="h-36 my-0" src="../public/VitaMartlogo.svg"></div>
    <div class="font-mono my-auto"><h1 class="text-5xl font-bold">VITAMART</h1></div>
    <div class="my-auto ml-auto flex">
        <a class="text-white ">Hello _Name_</a>
    </div>
</header>

  <div class="DIV1 py-20 px-56 rounded-lg ">
    <p class="text-4xl font-semibold text-white">Choose your hostel</p><br>
    <div class = "Div2 flex">
      <button class="px-16 text-xl font-semibold text-white "><img class="px-18 w-64" src="../public/male.svg">BOYS</button>
      <button class="px-16 ml-3.5 font-semibold text-xl text-white"><img class="w-64" src="../public/female.svg">GIRLS</button>

    </div>
  </div>
</main>

<style>
  .DIV1{
    box-shadow: 9px 10px 7px 1px rgba(38,38,38,0.94);
    background-color: rgb(46,79,79);
  }

    button:hover{
      box-shadow: 9px 10px 7px 1px rgba(38,38,38,0.54);
      transition: box-shadow 0.25s ease-in-out;
    }
  button{
    background-color: rgb(193,226,0);
  }
  h1{
    font-size: 50px;
    color:rgb(193,226,0);
  }
</style>