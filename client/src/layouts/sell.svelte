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
            <div class="texthere flex-2">Seller's Name <input type="text" class="px-16 text-xl font-semibold text-black "></div>
            <div class="texthere flex-2">Room Number <input class="px-16 text-xl font-semibold text-black "></div>
            <div class="texthere flex-2">Product Name <input class="px-16 text-xl font-semibold text-black "></div>
            <div class="texthere flex-2">Product Description <input class="px-16 text-xl font-semibold text-black "></div>
            <div class="texthere flex-2">Price <input class="px-16 text-xl font-semibold text-black "></div>
            <div><input type="file"></div>
            <div><button class="SUB">Submit</button> </div>
            
        </div>
    </div>
</main>     

<style>
    @import url('https://fonts.googleapis.com/css2?family=Shadows+Into+Light&display=swap');

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

    .SUB {
        background-color: rgb(200, 0, 0);
        color: white;
        border: none ;
        height: 36px;
        width: 105px;
        border-radius: 3px;
        cursor: pointer;
        margin-right: 8px;
        transition: opacity 0.15s;
    }
    .SUB:hover{
        opacity: 0.8;
    }
    .SUB:active{
        opacity: 0.4;
        }
    .texthere{
        color:rgb(193,226,0);
        font-size:50px;
        }
</style>