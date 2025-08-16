<script>
    const routes = {
      shipment: '/shipment',
      tracking: '/track',
      input: '/input'
    }
</script>

<div class="w-full max-w-6xl bg-item-gk-item p-12">
  <header class="mb-12 text-center">
    <h1 class="text-5xl md:text-6xl font-TikTok text-font-color2 mb-4">
      Digital Logistic System
    </h1>
    <p class="text-xl md:text-2xl text-font-color2 font-TikTok">
      Streamline your supply chain with real-time visibility and smart tools.
    </p>
  </header>
  
  <div class="grid grid-cols-1 md:grid-cols-3 gap-10">

    <a href={routes.input} class="block bg-card-bg p-8 rounded-2xl shadow-md hover:shadow-lg transition no-underline">
      <h2 class="text-2xl font-TikTok mb-2">New Shipment</h2>
      <p class="font-TikTok text-lg">
        Register new shipments with ease. Enter details and track them instantly.
      </p>
    </a>

    <a href={routes.tracking} class="block bg-card-bg p-8 rounded-2xl shadow-md hover:shadow-lg transition no-underline">
      <h2 class="text-2xl font-TikTok mb-2">Shipment Tracking</h2>
      <p class="font-TikTok text-lg">
        Monitor your shipments across the globe with accurate location and delivery time updates.
      </p>
    </a>

    <a href={routes.shipment} class="block bg-card-bg p-8 rounded-2xl shadow-md hover:shadow-lg transition no-underline">
      <h2 class="text-2xl font-TikTok mb-2">Port Management</h2>
      <p class="font-TikTok text-lg">
        Manage terminal schedules, monitor storage capacity, and optimize port operations.
      </p>
    </a>

  </div>

  <footer class="mt-16 text-font-color2 text-center font-TikTok text-sm">
    This is a demo tracking page. No real data is available.
  </footer>
</div>