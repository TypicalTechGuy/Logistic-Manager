<script>
    const routes = {
      inventory: '/inventory',
      tracking: '/track',
      reports: '/reports'
    }
</script>

<div class="w-full max-w-6xl bg-white shadow-md rounded-3xl p-12">
  <header class="mb-12 text-center">
    <h1 class="text-5xl md:text-6xl font-TikTok text-blue-light mb-4">
      Digital Logistic System
    </h1>
    <p class="text-xl md:text-2xl text-gray-600 font-TikTok">
      Streamline your supply chain with real-time visibility and smart tools.
    </p>
  </header>
  
  <div class="grid grid-cols-1 md:grid-cols-3 gap-10">
    <a href={routes.tracking} class="block bg-blue-50 p-8 rounded-2xl shadow-md hover:shadow-lg transition no-underline">
      <h2 class="text-2xl font-TikTok text-blue-light mb-2">Shipment Tracking</h2>
      <p class="text-gray-600 font-TikTok text-lg">
        Monitor your shipments across the globe with accurate location and delivery time updates.
      </p>
    </a>


    <a href={routes.inventory} class="block bg-blue-50 p-8 rounded-2xl shadow-md hover:shadow-lg transition no-underline">
      <h2 class="text-2xl font-TikTok text-blue-light mb-2">Inventory</h2>
      <p class="text-gray-600 font-TikTok text-lg">
        Track stock levels, manage warehouses, and prevent over/understocking efficiently.
      </p>
    </a>

    <a href={routes.reports} class="block bg-blue-50 p-8 rounded-2xl shadow-md hover:shadow-lg transition no-underline">
      <h2 class="text-2xl font-TikTok text-blue-light mb-2">Reports</h2>
      <p class="text-gray-600 font-TikTok text-lg">
        Generate business insights with customizable reports on deliveries, costs, and KPIs.
      </p>
    </a>
  </div>

  <footer class="mt-16 text-center font-TikTok text-gray-500 text-sm">
    &copy; 2025 Digital Logistic Co. All rights reserved.
  </footer>
</div>