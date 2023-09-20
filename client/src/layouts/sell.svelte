<script>
    import TailwindCss from '../lib/TailwindCSS.svelte';
</script>

<TailwindCss/>

<main>
    <header class="flex h-60px align-middle">
        <div><img class="max-w-full my-0" src="../public/Sell_Image.jpg"></div>
        <div class="top-left my-auto"><h1 class="text-5xl font-bold">VITAMART</h1></div>
    </header>
    <div class="SellDiv">

        <img class="h-36 my-0" src="../public/Logo.svg">
        <p class="text-white">SELL</p>
        <div>
            <div class="texthere flex">Seller's Name <button class="px-16 text-xl font-semibold text-white "></button></div>
            <div class="texthere flex">Room Number <button class="px-16 text-xl font-semibold text-white "></button></div>
            <div class="texthere flex">Product Name <button class="px-16 text-xl font-semibold text-white "></button></div>
            <div class="texthere flex">Product Description <button class="px-16 text-xl font-semibold text-white "></button></div>
            <div class="texthere flex">Price<button class="px-16 text-xl font-semibold text-white "></button></div>
        </div>
    </div>
</main>     

<style>
    button{
      background-color: white;
    }
    button:hover{
        box-shadow: 9px 10px 7px 1px rgba(38,38,38,0.54);
        transition: box-shadow 0.25s ease-in-out;
    }
    .SellDiv{
        max-width: 1280px;
        margin: 0 auto;
        padding: 2rem;
        text-align: center;
    }
    h1{
      font-size: 50px;
      color:rgb(193,226,0);
    }
    .texthere{
        color:rgb(193,226,0);
    }
    .top-left {
        position: absolute;
        top: 20px;
        left: 16px;
    }
    
</style>