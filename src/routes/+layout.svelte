<script>
  let navLinks = [
    { name: "Home", href: "/" },
    { name: "Input", href: "/input" },
    { name: "Tracking", href: "/track" },
    { name: "Manage", href: "/manage" },
  ];
</script>

<div class="h-screen bg-navbar-color flex flex-col">
  <nav class="w-full bg-navbar shadow-md z-50">
    <div class="max-w-6xl mx-auto px-4 py-4 flex justify-between items-center">
      <div class="text-2xl text-font-color font-TikTok">
        Digital Logistic
      </div>
      <ul class="flex space-x-6">
        {#each navLinks as link}
          <li>
            <a href={link.href} class="text-font-color font-TikTok hover:text-card-bg transition">
              {link.name}
            </a>
          </li>
        {/each}
      </ul>
    </div>
  </nav>

  <main class="flex-1 flex flex-col items-center justify-center bg-item-gk-item p-4 sm:p-6 md:p-12 overflow-auto">
    <slot />
  </main>
</div>