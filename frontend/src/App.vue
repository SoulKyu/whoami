<script setup>
import { defineProps } from 'vue'
import Sidebar from './components/Sidebar.vue'

const props = defineProps({
  msg: String
})
</script>

<template>
  <div class="m-0 p-0">
    <Sidebar />
  </div>
</template>