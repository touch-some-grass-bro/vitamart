<script>
    import TailwindCss from '../lib/TailwindCSS.svelte';
    import Dropdown from '../components/Dropdown.svelte';
    import HeaderCopy from '../components/HeaderCopy.svelte';
    let MENHOSTELS = ["A","B","D","D Annex","E","F","G","H","J","K","L","M","N","P","Q","R"];
    let FEMHOSTELS = ["A","B","C","D","E","F","G","H"];
  </script>
  
  <TailwindCss/>
  
  <HeaderCopy/>

  <main>
    <div class="DIV1 py-20 px-56 rounded-lg ">
      <p class="text-4xl font-semibold text-white text-center">Choose Your Hostel</p><br>
      <div class = "Div2 flex">
        <Dropdown img = "../public/male.svg" name = "Boys" hostels = {MENHOSTELS}/>
        <Dropdown img = "../public/female.svg" name = "Female" hostels = {FEMHOSTELS}/>
      </div>
    </div>
  </main>
  
  <style>
    .Div2{
        justify-content: space-between;
    }
    .DIV1{
      box-shadow: 9px 10px 7px 1px rgba(38,38,38,0.94);
      background-color: rgb(46,79,79);
    }
  </style>
