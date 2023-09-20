<script>
    import TailwindCss from '../lib/TailwindCSS.svelte';
  import HeaderSell from '../components/HeaderSell.svelte';
</script>

<TailwindCss/>

<main>
    <HeaderSell/>

    <div class="SellDiv">
        <p class="Heading text-white">SELL</p>
        <div>
            <div class="texthere flex-2">Seller's Name <button class="px-16 text-xl font-semibold text-white "></button></div>
            <div class="texthere flex-2">Room Number <button class="px-16 text-xl font-semibold text-white "></button></div>
            <div class="texthere flex-2">Product Name <button class="px-16 text-xl font-semibold text-white "></button></div>
            <div class="texthere flex-2">Product Description <button class="px-16 text-xl font-semibold text-white "></button></div>
            <div class="texthere flex-2">Price <button class="px-16 text-xl font-semibold text-white "></button></div>
        </div>
    </div>
</main>     

<style>
    @import url('https://fonts.googleapis.com/css2?family=Shadows+Into+Light&display=swap');
    button{
        background-color: white;
        height: 36px;
        width: 50%;
        border-radius: 3px;
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
    .Heading{
        font-family: 'Shadows Into Light', cursive;
        font-size: 100px;
    }
</style>