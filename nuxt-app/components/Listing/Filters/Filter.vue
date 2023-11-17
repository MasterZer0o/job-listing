<script setup lang="ts">
defineProps<{
  config: {
    title: string
    options: {
      name: string
    }[] }
}>()
</script>

<template>
  <span>{{ config.title }}</span>
  <ul>
    <li v-for="option in config.options" :key="option.name">
      <label>
        <input type="checkbox">
        <span>
          <svg viewBox="0 0 20 20" fill="none" stroke="#20c2e0" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" stroke-dasharray="71" stroke-dashoffset="71">
            <path d="M3,1 L17,1 L17,1 C18.1045695,1 19,1.8954305 19,3 L19,17 L19,17 C19,18.1045695 18.1045695,19 17,19 L3,19 L3,19 C1.8954305,19 1,18.1045695 1,17 L1,3 L1,3 C1,1.8954305 1.8954305,1 3,1 Z"></path>
            <polyline points="4 11 8 15 16 6" fill="none" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" stroke-dasharray="18" stroke-dashoffset="18"></polyline>
          </svg>
        </span>
        {{ option.name }}
      </label>
    </li>
  </ul>
</template>
