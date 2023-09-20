<script>
    import TailwindCss from '../lib/TailwindCSS.svelte';
    import Header from '../components/Header.svelte';
</script>

<TailwindCss/>

<main>
    <!-- <header class="flex h-60px align-middle">
        <div><img class="Main_Image max-w-full my-0 " src="../public/SELL.jpeg"></div>
        <div><img class="top-left1 h-36 my-0" src="../public/Logo.svg"></div>
        <div class="top-left2 my-auto"><h1 class="text-5xl font-bold">VITAMART</h1></div>
        <div >
            <button type="submit" class="top-right my-auto ml-auto flex"><img class=" h-14 w-14" src="../public/Settings.svg"></button>
                
        </div>
    </header> -->

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
    .button{
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
    h1{
        font-family: 'Shadows Into Light', cursive;
        font-size: 70px;
        color:rgb(193,226,0);
    }
    .texthere{
        color:rgb(193,226,0);
        font-size: 50px;
    }
    .top-left1{
        position: fixed;
        top: 0px;
        left: 16px;
    }
    .top-left2 {
        position: fixed;
        top: 30px;
        left: 130px;
    }
    .top-right {
        position: fixed;
        top: 24px;
        right: 16px;
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
</style>