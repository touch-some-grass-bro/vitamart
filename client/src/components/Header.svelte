<script></script>

<main>
    <header class="flex h-60px align-middle">
        <div><img class="h-36 my-0" src="../public/Logo.svg"></div>
        <div class="my-auto"><h1 class="text-5xl font-bold">VITAMART</h1></div>
        <div class="my-auto ml-auto flex">
            <a>Change Hostel Details</a>
            <a>Sign Out</a>
        </div>
    </header>
</main>

<style>
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