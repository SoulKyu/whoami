<template>
  <div class="w-64 bg-blue-200 border-r border-gray-300 min-h-screen sticky top-0 p-4">
    <!-- liens de navigation -->
    <ul>
      <li>
        <hr class="mt-16 border-t border-black mx-full mb-full">
        <router-link to="/" class="block py-2 px-4 text-xl text-center text-gray-900 hover:bg-blue-100 transition duration-200 rounded-lg">Accueil</router-link>
      </li>
      <hr class="border-t border-black mx-full my-2">
      <li>
        <router-link to="/" class="block py-2 px-4 text-xl text-center text-gray-900 hover:bg-blue-100 transition duration-200 rounded-lg">Tutoriels</router-link>
      </li>
      <hr class="border-t border-black mx-full my-2">
      <li>
        <router-link to="/" class="block py-2 px-4 text-xl text-center text-gray-900 hover:bg-blue-100 transition duration-200 rounded-lg">A Propos</router-link>
      </li>
      <hr class="border-t border-black mx-full my-2">
      <!-- Ajoutez vos liens ici -->
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Sidebar'
}
</script>
