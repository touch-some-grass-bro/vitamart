<script>
    import TailwindCss from '../lib/TailwindCSS.svelte';
  import HeaderSell from '../components/HeaderSell.svelte';
  let userInput = ''; // Initialize a variable to store the user's input
  
    function handleSubmit() {
      // Handle the form submission here, you can access userInput to get the input value
      console.log('User input:', userInput);
      }
</script>

<TailwindCss/>

<main>
    <HeaderSell/>

    <div class="SellDiv">
        <p class="Heading text-white">SELL</p>
        <div>
            <form on:submit={handleSubmit}>
                <label class="texthere" for="inputField">Seller's Name</label>
                <input type="text" id="inputField" bind:value={userInput} /><br>
                <label class="texthere" for="inputField">Room Number</label>
                <input type="text" id="inputField" bind:value={userInput} /><br>
                <label class="texthere" for="inputField">Product Name</label>
                <input type="text" id="inputField" bind:value={userInput} /><br>
                <label class="texthere" for="inputField">Product Description</label>
                <input type="text" id="inputField" bind:value={userInput} /><br>
                <label class="texthere" for="inputField">Price</label>
                <input type="text" id="inputField" bind:value={userInput} /><br>
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