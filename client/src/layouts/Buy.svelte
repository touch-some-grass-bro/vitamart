<script>
// @ts-nocheck

    import TailwindCss from "../lib/TailwindCSS.svelte";
    import Header from "../components/Header.svelte";
    let sellername = ["Param", "Milind", "Ishaan"];
    let sellerdetail = [{name:"Param",price:1200},{name:"Milind",price:300},{name:"Ishaan",price:50}];
    let pricelist = [1200,300,23];
    let buy = "../public/buy.jpg";

</script>

<TailwindCss/>
<main>
    <Header img = {buy}/>
    <div class="maincontainer">
    <h1 class="text-white text-3xl text-center">Products being sold</h1>
    <div class="my-20 mx-10 grid grid-cols-3 gap-6 ">
        {#each sellerdetail as seller}
        <button>
            <img class="relative rounded-lg bg-red-500 z-10" src="../public/sample.jpeg">
            <div class="z-0 absolute inset-0 items-center my-32 text-white">
            <h2>Seller:{seller.name}</h2>
            <h2>{seller.price}Rs</h2>
            </div>
        </button>
        {/each}
    </div>
    </div>

    
</main>


<style>
    @import url('https://fonts.googleapis.com/css2?family=Shadows+Into+Light&display=swap');
    h2{
        font-size: 60px;
    }
    h1{
        font-family: 'Shadows Into Light', cursive;
        font-size: 120px;
        text-shadow: 5px 3px 5px rgba(0,0,0,0.57);
    }

    button:hover{
        box-shadow: 7px 9px 15px 0px rgba(0,0,0,0.75);
        transition: all 0.2s ease-in-out;
    }

    button {
    overflow: hidden;
    position: relative;
}

/* Styles for the image */
button img {
    max-width: 100%;
    height: auto;
    display: block;
    opacity: 1;
    transform: translateY(0); 
    transition: transform 1s; 
}

button img:hover {
    transform: translateY(-100%);
    transition: transform 1s;
}

.maincontainer{
    position: inherit;
}


</style>
