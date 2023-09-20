<script>
  // @ts-nocheck

  import TailwindCss from "../lib/TailwindCSS.svelte";
  import HeaderBuy from "../components/HeaderBuy.svelte";
  let sellerdetail = [
    { name: "Param", price: 1200 },
    { name: "Milind", price: 300 },
    { name: "Ishaan", price: 50 },
  ];
  let buy = "../public/buy.jpg";
</script>

<TailwindCss />
<main>
  <HeaderBuy />
  <div>
    <h1 class="text-white text-3xl text-center">Products being sold</h1>
    <div class="my-20 mx-10 grid grid-cols-3 gap-6">
      {#each sellerdetail as seller}
        <div
          class="w-full max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700"
        >
          <a href="#">
            <img
              class="p-8 rounded-t-lg"
              src="https://flowbite.com/docs/images/products/apple-watch.png"
              alt="product image"
            />
          </a>
          <div class="px-5 pb-5">
            <a href="#">
              <h5
                class="text-xl font-semibold tracking-tight text-gray-900 dark:text-white"
              >
                Apple Watch Series 7 GPS, Aluminium Case, Starlight Sport
              </h5>
            </a>
            <div class="flex items-center mt-2.5 mb-5">
              
            </div>
            <div class="flex items-center justify-between">
              <span class="text-3xl font-bold text-gray-900 dark:text-white"
                >$599</span
              >
              <a
                href="#"
                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >Add to cart</a
              >
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</main>

<style>
  @import url("https://fonts.googleapis.com/css2?family=Shadows+Into+Light&display=swap");
  h2 {
    font-size: 60px;
  }
  h1 {
    font-family: "Shadows Into Light", cursive;
    font-size: 120px;
    text-shadow: 5px 3px 5px rgba(0, 0, 0, 0.57);
  }

  button:hover {
    box-shadow: 7px 9px 15px 0px rgba(0, 0, 0, 0.75);
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
</style>
