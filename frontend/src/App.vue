<template>
  <Header />
  <div class="flex flex-col h-screen">
    <div class="flex flex-grow">
      <Sidebar />
      <Home />
    </div>
    <Footer />
  </div>
</template>

<script>
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
import Sidebar from './components/Sidebar.vue'
import Home from './components/Home.vue'

export default {
  components: {
    Header,
    Footer,
    Sidebar,
    Home
  }
}
</script>
