<script>
    import TailwindCss from '../lib/TailwindCSS.svelte';
  import HeaderSell from '../components/HeaderSell.svelte';

  
  let prodname = '';
  let proddes = '';
  let price = '';
  let propic = '';

  function handleSubmit() {
    // Handle the form submission here
    // console.log('First Name:', sellname);
    
    console.log('Product name', prodname);
    console.log('Product description', proddes);
    console.log('Price', price);
    console.log('Product Picture', propic);
  }
</script>

<TailwindCss/>

<main>
    <HeaderSell/>

    <p class="text">Hungry,Grab a snickers.</p>

    <div class="SellDiv">
        <p class="Heading text-white">SELL</p>
        <div>
            <form on:submit={handleSubmit}>
                <!-- <label class="texthere" for="sellname">Seller's Name</label>
                <input type="text" id="sellname" bind:value={sellname} /><br>

                <label class="texthere" for="roomname">Room Number</label>
                <input type="text" id="roomname" bind:value={roomname} /><br> -->

                <label class="texthere" for="prodname">Product Name</label>
                <input type="text" id="prodname" bind:value={prodname} /><br>

                <label class="texthere" for="proddes">Product Description</label>
                <input type="text" id="proddes" bind:value={proddes} /><br>

                <label class="texthere" for="price">Price</label>
                <input type="text" id="price" bind:value={price} /><br>
                
                <label class="texthere" for="propic"></label>
                <input type="file" id="propic" bind:value={propic} /><br>

                <button class="BuTTon" type="submit">Submit</button>
              </form>

              
              
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
    .BuTTon{
        font-family: 'Shadows Into Light', cursive;
        font-size: 100px;
        border-radius: 30px;
        border-color: rgb(193,226,0);
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

        p.text {
             writing-mode: vertical-rl;
        }
</style>

  
  
  