<script>
  import { ClockHub } from "./stores/clock.spindlyhubs";

  let { ClockFace, TimeZone } = ClockHub();

  export let Zone = "";
  if (Zone) {
    $TimeZone = Zone;
  }
</script>

<div class="container">
  <h1>[ {$ClockFace} ]</h1>

  <select bind:value={$TimeZone}>
    <option>Local</option>
    <option>UTC</option>
    <option>America/New_York</option>
    <option>Asia/Colombo</option>
    <option>Asia/Singapore</option>
    <option>Asia/Tokyo</option>
    <option>Australia/Sydney</option>
    <option>Europe/Paris</option>
    <option>Europe/London</option>
    <option>Israel</option>
  </select>
</div>

<style>
  h1 {
    color: rgb(0, 128, 68);
    font-family: "Comic Sans MS";
  }

  .container {
    width: 20rem;
    margin-left: auto;
    margin-right: auto;
    margin-top: 1rem;
    margin-bottom: 1rem;
    font-size: 1.1rem;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    padding: 0.3rem;
  }
</style>
