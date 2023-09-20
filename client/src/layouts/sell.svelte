<script>
    import TailwindCss from '../lib/TailwindCSS.svelte';
    import Header from '../components/Header.svelte';
</script>

<TailwindCss/>

<main>
    <header class="flex h-60px align-middle">
        <div><img class="Main_Image max-w-full my-0 " src="../public/SELL.jpeg"></div>
        <div><img class="top-left1 h-36 my-0" src="../public/Logo.svg"></div>
        <div class="top-left2 my-auto"><h1 class="text-5xl font-bold">VITAMART</h1></div>
        <div class="top-right my-auto ml-auto flex">
            <a>Change Hostel Details</a>
            <a>Sign Out</a>
        </div>
    </header>

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
    h1{
        font-family: 'Shadows Into Light', cursive;
        font-size: 70px;
        color:rgb(193,226,0);
    }
    .texthere{
        color:rgb(193,226,0);
        font-family: 'Shadows Into Light', cursive;
        font-size: 50px;
    }
    .top-left1{
        position: absolute;
        top: 0px;
        left: 16px;
    }
    .top-left2 {
        position: absolute;
        top: 30px;
        left: 130px;
    }
    .top-right {
        position: absolute;
        top: 55px;
        right: 16px;
}
    .Main_Image{
        border-radius: 0px 0px 400px 0px;
    }
    .Heading{
        font-family: 'Shadows Into Light', cursive;
        font-size: 100px;
    }
</style>