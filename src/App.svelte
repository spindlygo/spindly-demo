<script>
  import { Global } from "./stores/global.spindlyhubs.js";
  import { ExampleHub } from "./stores/example.spindlyhubs.js";
  import Clock from "./Clock.svelte";

  import { IsHostAlive } from "spindly-hubs";

  let NameJS = "";
  let GreatingJS = "";

  // This is how svelte reactivity works
  $: if (NameJS.length > 0) {
    let num = (NameJS.length * 7) % 20;
    GreatingJS = "Hello " + NameJS + ", your lucky number is " + num;
  } else {
    GreatingJS = "Hello from JavaScript!";
  }

  let { Name, Greating } = ExampleHub("example1");

  let { AppName, HelloMessage, Events, SaidHello } = Global;

  if (!$Events) {
    $Events = ["Hi"];
  }

  let deadHostTimout = null;
  IsHostAlive.subscribe((alive) => {
    if (alive) {
      if (deadHostTimout) {
        clearTimeout(deadHostTimout);
        deadHostTimout = null;
      }
    } else {
      if (!deadHostTimout) {
        deadHostTimout = setTimeout(() => {
          clearTimeout(deadHostTimout);
          deadHostTimout = null;
          window.stop();
          window.close();
        }, 6500);
      }
    }
  });
</script>

<div>
  <div style="margin-left: auto; margin-right: auto; text-align: center;">
    <h1>{$AppName}</h1>
    <h3>Write once, run anywhere app framework for desktop and mobile</h3>
  </div>

  {#if $IsHostAlive}
    <div style="display: flex; flex-wrap: wrap; margin: 5vw; ">
      <div class="half box">
        <h3>This logic is written in JavaScript</h3>

        <input
          type="text"
          bind:value={NameJS}
          placeholder="Type your name here"
        />
        <p class="store">{GreatingJS}</p>
      </div>

      <div class="half box">
        <h3>This logic is written in Go</h3>

        <input
          type="text"
          bind:value={$Name}
          placeholder="Type your name here"
        />
        <p class="store">{$Greating}</p>
      </div>

      <div class="half box">
        <h3>This is a pre-defined instance store</h3>
        <p>Go can access these directly by the variable name</p>
        <p class="store">{$HelloMessage}</p>
        <input type="button" value="Say hello" on:click={SaidHello} />

        <p>
          A store can hold any type of value that a JSON can represent. Here we
          are storing an array
        </p>
        <div>
          {#each $Events as event}
            <p class="store">{event}</p>
          {/each}
        </div>
      </div>

      <div class="half box">
        <h3>These clocks are running in Go host</h3>
        <p>This demonstraits how simple it is to update the UI and respond</p>

        <div style="display: flex; flex-wrap: wrap; margin: 1vw; ">
          <Clock />
          <Clock Zone="UTC" />
          <Clock Zone="Europe/Paris" />
          <Clock Zone="Asia/Singapore" />
        </div>

        <p>
          Each clock has its own instance of a 'ClockHub'. Each holds its state
          and respond independently.
        </p>
      </div>

      <div style="margin-left: auto; margin-right: auto; text-align: center;">
        <p>
          It took less than 30 JavaScript code lines and less than 50 executable
          Go code lines to create this sample project
        </p>
      </div>
    </div>
  {:else}
    <div style="display: flex; flex-wrap: wrap; margin: 5vw; ">
      <div class="half box">
        <h3>Host is not responding. Try restarting the app.</h3>
      </div>
    </div>
  {/if}
</div>

<style>
  .half {
    width: min(50rem, 80%);
  }

  .box {
    margin-left: auto;
    margin-right: auto;
    text-align: center;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    padding: 1rem;
    background-color: #f5f5f5;
    border-radius: 0.3rem;
    margin-top: 1rem;
    margin-bottom: 1rem;
  }

  .store {
    width: 80%;
    margin-left: auto;
    margin-right: auto;
    font-size: 1.1rem;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    padding: 1rem;
  }
</style>
