<div class="w-full max-w-6xl bg-item-gk-item p-12 mt-140">
  <header class="mb-12 text-center">
    <h1 class="text-5xl md:text-6xl font-TikTok text-font-color2 mb-4">
      New Shipment Entry
    </h1>
    <p class="text-xl md:text-2xl text-font-color2 font-TikTok">
      Register new shipments with ease and start tracking them instantly.
    </p>
  </header>

  <form class="bg-card-bg p-8 rounded-2xl shadow-md space-y-6">
    
    <div>
      <label class="block mb-2 font-TikTok">Cargo Type</label>
      <input type="text" placeholder="e.g., Electronics" class="w-full p-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400" required>
    </div>

    <div>
      <label class="block mb-2 font-TikTok">Weight (kg)</label>
      <input type="number" step="0.01" placeholder="e.g., 1250" class="w-full p-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400" required>
    </div>

    <div>
      <label class="block mb-2 font-TikTok">Dimensions (L×W×H in cm)</label>
      <input type="text" placeholder="e.g., 120×80×100" class="w-full p-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400">
    </div>

    <div>
      <label class="block mb-2 font-TikTok">Ship Name</label>
      <input type="text" placeholder="e.g., Ocean Trader" class="w-full p-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400">
    </div>

    <div>
      <label class="block mb-2 font-TikTok">Origin Port</label>
      <input type="text" placeholder="e.g., Jakarta Port" class="w-full p-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400">
    </div>

    <div>
      <label class="block mb-2 font-TikTok">Destination Port</label>
      <input type="text" placeholder="e.g., Singapore Port" class="w-full p-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400">
    </div>

    <div>
      <label class="block mb-2 font-TikTok">Expected Arrival Date</label>
      <input type="date" class="w-full p-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400">
    </div>

    <button type="submit" class="w-full bg-blue-500 text-white p-3 rounded-lg font-TikTok text-lg hover:bg-blue-600 transition">
      Submit Shipment
    </button>

  </form>

  <footer class="mt-16 text-font-color2 text-center font-TikTok text-sm">
    This is a demo input page. Data entered here is not stored.
  </footer>
</div>
