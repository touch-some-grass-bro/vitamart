<script>
  import TailwindCss from "../lib/TailwindCSS.svelte";
  import HeaderSell from "../components/HeaderSell.svelte";
  import axios from "axios";
  import API_URL from "../api";
    import { navigate } from "svelte-routing";

  //   let sellname = '';

  let prodname = "";
  let proddes = "";
  let price = "";
  let propic = "";

  // E is an event
  function handleSubmit(e) {
    // Get the image selected
    // @ts-ignore
    const file = document.getElementById("propic").files[0];

    // Create a new form data object
    const formData = new FormData();
    formData.append("name", prodname);
    formData.append("description", proddes);
    formData.append("price", price);
    formData.append("image", file);

    // Send the form data to the server using axios withCredtials: true
    axios
      .post(`${API_URL}/items`, formData, {
        withCredentials: true,
      })
      .then((res) => {
        console.log(res);
        navigate("/buy")
      })
      .catch((err) => {
        console.log(err);
      });
  }
</script>

<TailwindCss />

<main>
  <HeaderSell />


  <div class="SellDiv">
    <p class="Heading">SELL</p>
    <div>
      <form
        on:submit|preventDefault={(e) => {
          handleSubmit(e);
        }}
      >

        <label class="texthere" for="prodname">Product Name</label>
        <input type="text" id="prodname" bind:value={prodname} /><br />

        <label class="texthere" for="proddes">Product Description</label>
        <input type="text" id="proddes" bind:value={proddes} /><br />

        <label class="texthere" for="price">Price</label>
        <input type="text" id="price" bind:value={price} /><br />

        <label class="texthere" for="propic" />
        <!-- Make an input that takes an image and gets the file data -->
        <input class="onlyimput" type="file" id="propic" bind:value={propic} /><br />

        <button class="BuTTon" type="submit">Submit</button>
      </form>
    </div>
  </div>
</main>

<style>
  @import url("https://fonts.googleapis.com/css2?family=Shadows+Into+Light&display=swap");

  button:hover {
    box-shadow: 9px 10px 7px 1px rgba(38, 38, 38, 0.54);
    transition: box-shadow 0.25s ease-in-out;
  }
  .SellDiv {
    max-width: 1280px;
    margin: 0 auto;
    padding: 2rem;
    text-align: center;
  }
  .BuTTon {
    font-family: "Shadows Into Light", cursive;
    font-size: 100px;
    border-radius: 30px;
    border-color: #5a8181;
  }
  input{ 
    margin-left: 10px;
    height: 40px;
    border-radius: 20px;
    background-color: #5a8181;
  }

  .onlyimput{
    background-color: white;
    border-radius: 0px;
  }
 
  .Heading {
    font-family: "Shadows Into Light", cursive;
    font-size: 100px;
  }

  .SUB {
    background-color: rgb(200, 0, 0);
    color: white;
    border: none;
    height: 36px;
    width: 105px;
    border-radius: 3px;
    cursor: pointer;
    margin-right: 8px;
    transition: opacity 0.15s;
  }
  .SUB:hover {
    opacity: 0.8;
  }
  .SUB:active {
    opacity: 0.4;
  }
  .texthere {
    font-size: 50px;
  }

</style>
