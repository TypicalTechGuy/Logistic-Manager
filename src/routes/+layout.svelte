<script>
  let navLinks = [
    { name: "Home", href: "/" },
    { name: "Tracking", href: "/track" },
    { name: "Inventory", href: "/inventory" },
    { name: "Reports", href: "/reports" }
  ];
</script>

<div class="h-screen bg-blue-fun flex flex-col">
  <nav class="w-full bg-white shadow-md z-50">
    <div class="max-w-6xl mx-auto px-4 py-4 flex justify-between items-center">
      <div class="text-2xl text-blue-light font-TikTok">
        Digital Logistic
      </div>
      <ul class="flex space-x-6">
        {#each navLinks as link}
          <li>
            <a href={link.href} class="text-blue-light font-TikTok hover:text-blue-800 transition">
              {link.name}
            </a>
          </li>
        {/each}
      </ul>
    </div>
  </nav>

  <main class="flex-1 flex flex-col items-center justify-center p-4 sm:p-6 md:p-12 overflow-auto">
    <slot />
  </main>
</div>