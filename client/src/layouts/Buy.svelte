<script>
  import TailwindCss from "../lib/TailwindCSS.svelte";
  import HeaderBuy from "../components/HeaderBuy.svelte";
  let sellerdetail = [
    { name: "Param", price: 1200 },
    { name: "Milind", price: 300 },
    { name: "Ishaan", price: 50 },
    { name: "Param", price: 1200 },
    { name: "Milind", price: 300 },
    { name: "Ishaan", price: 50 },
    { name: "Param", price: 1200 },
    { name: "Milind", price: 300 },
    { name: "Ishaan", price: 50 },
    { name: "Param", price: 1200 },
    { name: "Milind", price: 300 },
    { name: "Ishaan", price: 50 },
    { name: "Param", price: 1200 },
    { name: "Milind", price: 300 },
    { name: "Ishaan", price: 50 },
    { name: "Param", price: 1200 },
    { name: "Milind", price: 300 },
    { name: "Ishaan", price: 50 },
  ];
  let buy = "../public/buy.jpg";
  let scrollAmount = 0;

  function scrollNext() {
    scrollAmount += 250; // Adjust the scroll amount as needed
    scrollTo(scrollAmount);
  }

  function scrollPrev() {
    scrollAmount -= 250; // Adjust the scroll amount as needed
    scrollTo(scrollAmount);
  }

  function scrollTo(amount) {
    const productList = document.getElementById("productList");
    productList.scrollTo({
      left: amount,
      behavior: "smooth",
    });
  }
</script>

<TailwindCss />

<main>
  <HeaderBuy />
  <h1 class="text-white text-3xl text-center">Products being sold</h1>
  <div class="product-carousel">
    <button class="mx-5" on:click={scrollPrev}>Previous</button>
    <div class="product-list" id="productList">
      {#each sellerdetail as seller}
        <div class="product-item">
          <img src="../public/sample.jpg" />
          <h4>{seller.name}</h4>
          <h5>{seller.price}</h5>
        </div>
      {/each}
      <!-- Add more products as needed -->
    </div>
    <button class="mx-5" on:click={scrollNext}>Next</button>
  </div>
</main>

<style>
  .product-carousel {
    display: flex;
    align-items: center;
    justify-content: space-between;
    max-width: 1000px;
    margin: 0 auto;
  }

  .product-list {
    display: flex;
    overflow: hidden;
    width: 100%;
  }

  .product-item {
    flex: 0 0 200px; /* Adjust the width of each product item */
    background-color: #f0f0f0;
    padding: 20px;
    margin: 10px;
    border-radius: 5px;
    text-align: center;
  }

  button {
    padding: 10px 20px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }

  button:hover {
    background-color: #0056b3;
  }
  h1 {
    font-family: "Shadows Into Light", cursive;
    font-size: 120px;
    text-shadow: 5px 3px 5px rgba(0, 0, 0, 0.57);
    margin-bottom: 100px;
  }
</style>
