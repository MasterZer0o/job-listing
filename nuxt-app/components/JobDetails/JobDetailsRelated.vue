<script setup lang="ts">
const emit = defineEmits(['saveJob'])
const store = jobDetailsStore()
</script>

<template>
  <section class="related-offers">
    <p>Related offers</p>
    <ul class="related-offers-list">
      <JobDetailsRelatedOffer v-for="item in store.results" :key="item.id" :data="item" :is-saved="true" @save-job="emit('saveJob')" />
    </ul>
  </section>
</template>
