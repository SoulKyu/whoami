import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import HelloWorld from './components/HelloWorld.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: HelloWorld },
    //{ path: '/about', component: About },
    // Ajoutez toutes vos routes ici
  ]
})

const app = createApp(App)
app.use(router)
app.mount('#app')