<script lang="ts">
  interface TrackingStep {
    time: string;
    location: string;
    status: string;
  }

  interface TrackingData {
    status: string;
    lastUpdated: string;
    origin: string;
    destination: string;
    steps: TrackingStep[];
  }

  let trackingNumber = '';
  let trackingResult: TrackingData | null = null;

  function trackShipment() {
    if (trackingNumber.trim() === '') return;

    trackingResult = {
      status: 'In Transit',
      lastUpdated: '2025-08-05 14:33',
      origin: 'Jakarta, Indonesia',
      destination: 'Surabaya, Indonesia',
      steps: [
        { time: '2025-08-03 10:00', location: 'Jakarta Hub', status: 'Picked Up' },
        { time: '2025-08-04 08:30', location: 'Central Warehouse', status: 'Sorting' },
        { time: '2025-08-05 11:45', location: 'Surabaya Gateway', status: 'Out for Delivery' }
      ]
    };
  }
</script>

<div class="w-full max-w-6xl bg-item-gk-item p-12 mx-auto min-h-screen flex flex-col mt-24">
  <header class="mb-12 text-center">
    <h1 class="text-5xl md:text-6xl font-TikTok text-font-color2 mb-4">
      Track Your Shipment
    </h1>
    <p class="text-xl md:text-2xl text-font-color2 font-TikTok">
      Enter your tracking number:
    </p>
  </header>

  <div class="mb-8 flex justify-center">
    <input
      type="text"
      bind:value={trackingNumber}
      placeholder="Enter Tracking Number"
      class="w-full md:w-2/3 lg:w-1/2 px-6 py-3 rounded-xl border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400 font-TikTok text-font-color2"
      on:keydown={(e) => e.key === 'Enter' && trackShipment()}
    />
    <button
      on:click={trackShipment}
      class="ml-4 px-6 py-3 rounded-xl bg-card-bg font-TikTok transition"
    >
      Track
    </button>
  </div>

  {#if trackingResult}
    <div class="bg-card-bg p-8 rounded-2xl shadow-md mx-auto max-w-3xl">
      <h2 class="text-2xl font-TikTok mb-4">Status: {trackingResult.status}</h2>
      <p class="text-font-color2 font-TikTok mb-2">Last Updated: {trackingResult.lastUpdated}</p>
      <p class="text-font-color2 font-TikTok mb-4">
        From {trackingResult.origin} → {trackingResult.destination}
      </p>

      <div class="border-t border-gray-300 mt-4 pt-4">
        {#each trackingResult.steps as step}
          <div class="mb-3">
            <p class="font-TikTok text-font-color2 text-lg">{step.time}</p>
            <p class="font-TikTok text-gray-600 text-md">{step.location} – {step.status}</p>
          </div>
        {/each}
      </div>
    </div>
  {/if}

  <footer class="mt-16 text-font-color2 text-center font-TikTok text-sm">
    This is a demo tracking page. No real data is available.
  </footer>
</div>
