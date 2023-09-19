<script>
    import TailwindCss from '../lib/TailwindCSS.svelte';
    let HOSTELS = ["A","B","C","D","D Annex","E","F","G","H","J","K","L","M","N","P","Q","R"];
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

    <div class="maindiv rounded-lg grid grid-cols-4">
        {#each HOSTELS as hostel}
            <div class="divinside rounded-lg mx-auto my-6 h-44 w-60 text-center py-16 text-4xl">{hostel}</div>
        {/each}
    </div>
</main>

<style>
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
        .maindiv{
            box-shadow: 9px 10px 7px 1px rgba(38, 38, 38, 0.94);
            background-color: rgb(46, 79, 79);
        }
        .divinside{
            background-color: rgb(36, 60, 60);
            position: relative;
            overflow: hidden;
            transition: background-position 0.3s ease;
        }
        .divinside::before {
            content: "";
            position: absolute;
            bottom: 0;
            left: 0;
            width: 100%;
            height: 0;
            background-color: rgb(107, 142, 142);
            transition: height 0.3s ease;
        }

        .divinside:hover::before {
            height: 100%;
        }
</style>