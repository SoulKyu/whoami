import { createApp } from 'vue'
import './styles.css'
import App from './App.vue'
import { createRouter, createWebHistory } from 'vue-router'
import Home from './components/Home.vue'

const routes = [
  { path: '/', component: Home },
  //{ path: '/about', component: About },
  // Ajoutez toutes vos routes ici
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

createApp(App)
  .use(router)
  .mount('#app')
