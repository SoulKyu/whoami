<template>
    <nav class="bg-blue-200 border-gray-300 p-4 flex text-black">
        <div class="col-span-1 border-gray-300"></div>
        <span class="text-2xl font-bold text-gray-700 ml-14">WhoAmI</span>
        <div v-if="authenticated" class="flex items-center justify-start ml-auto">
            <p class="text-sm">Bienvenue, {{username}} !</p>
            <form action="/logout" method="GET" class="flex space-x-3">
                <button type="submit" class="h-8 w-100 px-2 py-1 rounded-md bg-white text-blue-500 text-xs">Se déconnecter</button>
            </form>
        </div>
        <div v-else class="flex items-center justify-start ml-auto">
        <form action="/login" method="POST" class="flex space-x-3">
            <input type="text" name="username" placeholder="Nom d'utilisateur" required class="h-8 w-100 px-2 py-1 rounded-md">
            <input type="password" name="password" placeholder="Mot de passe" required class="h-8 w-100 px-2 py-1 rounded-md">
            <button type="submit" class="h-8 w-100 px-20 py-1 rounded-md bg-green-500 text-white whitespace-nowrap">Se connecter</button>
        </form>
        </div>
    </nav>
</template>
  

<script>
export default {
  data() {
    return {
      authenticated: false, // Vous devez remplacer cette valeur par la valeur d'authentification correcte
      username: 'Utilisateur', // Vous devez remplacer cette valeur par le nom d'utilisateur correct
    }
  },
}
</script>
