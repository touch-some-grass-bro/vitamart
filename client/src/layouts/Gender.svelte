<script>
    import TailwindCss from '../lib/TailwindCSS.svelte';
  </script>
  
  <TailwindCss/>
  
  <main>
  
    <header class="flex h-60px align-middle">
        <div><img class="h-36 my-0" src="../public/Logo.svg"></div>
        <div class="my-auto"><h1 class="text-5xl font-bold">VITAMART</h1></div>
        <div class="my-auto ml-auto flex">
            <a>Change Hostel Details</a>
            <a>Sign Out</a>
        </div>
    </header>
  
    <div class="DIV1 py-20 px-56 rounded-lg ">
      <p class="text-4xl font-semibold text-white text-center">Choose Your Hostel</p><br>
      <div class = "Div2 flex">
        <button class="px-16 text-xl font-semibold text-white "><img class="px-18 w-64" src="../public/male.svg">BOYS</button>
        <button class="px-16 ml-auto font-semibold text-xl text-white"><img class="w-64" src="../public/female.svg">GIRLS</button>
      </div>
    </div>
  </main>
  
  <style>
    .DIV1{
      box-shadow: 9px 10px 7px 1px rgba(38,38,38,0.94);
      background-color: rgb(46,79,79);
    }
      button:hover{
        box-shadow: 9px 10px 7px 1px rgba(38,38,38,0.54);
        transition: box-shadow 0.25s ease-in-out;
      }
    button{
      background-color: rgb(193,226,0);
    }
    h1{
      font-size: 50px;
      color:rgb(193,226,0);
    }
    a {
            font-size: 30px;
            text-decoration: none;
            color: black;
            position: relative; 
            margin: auto 15px;
        }

        a::before {
            content: "";
            position: absolute;
            width: 100%;
            height: 2px; /* Set the desired underline thickness */
            bottom: 0;
            left: 0;
            background-color: black; /* Set the desired underline color */
            transform: scaleX(0); /* Initial state: no underline */
            transform-origin: bottom right;
            transition: transform 0.3s ease; /* Add a smooth transition */
        }

        /* Animate the underline on hover */
        a:hover::before {
            transform: scaleX(1); /* Fully underline on hover */
            transform-origin: bottom left;
        }

        a:hover{
            cursor: pointer;
        }
  </style>