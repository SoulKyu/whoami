<template>
    <div class="w-64 bg-gray-200 border-r border-gray-300 min-h-screen sticky top-0 px-4 py-2">
      <div class="flex items-center space-x-4 p-2">
        <img src="logo.svg" alt="logo-svg" class="mx-auto w-16 h-16">
        <span class="text-2xl font-bold text-gray-700">WhoAmI</span>
      </div>
      <div class="h-0.5 bg-gray-300 my-2"></div>
      <ul>
        <li><router-link to="/" class="block py-2 px-4 text-gray-800 hover:bg-blue-200 transition duration-200 rounded-lg">Accueil</router-link></li>
        <li><router-link to="/tutoriels" class="block py-2 px-4 text-gray-800 hover:bg-blue-200 transition duration-200 rounded-lg">Mes Tutoriels</router-link></li>
        <li><router-link to="/whoami" class="block py-2 px-4 text-gray-800 hover:bg-blue-200 transition duration-200 rounded-lg">Whoami</router-link></li>
        <!-- Temporarily commented out for authentication
        <li v-if="authenticated"><router-link to="/newPage" class="block py-2 px-4 text-gray-800 hover:bg-blue-200 transition duration-200 rounded-lg">Create Tutoriels</router-link></li>
        <li v-if="authenticated"><router-link to="/kubernetes" class="block py-2 px-4 text-gray-800 hover:bg-blue-200 transition duration-200 rounded-lg">Kubernetes</router-link></li>
        -->
      </ul>
      <div class="h-0.5 bg-gray-300 my-2"></div>
    </div>
</template>
  

<script>
export default {
  data() {
    return {
      // authenticated: false, // à décommenter quand l'authentification sera nécessaire
    }
  },
}
</script>