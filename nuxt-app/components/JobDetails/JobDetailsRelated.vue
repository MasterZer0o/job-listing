<script setup lang="ts">
const emit = defineEmits(['saveJob'])
</script>

<template>
  <section class="related-offers">
    <p>Related offers</p>
    <ul class="related-offers-list">
      <JobDetailsRelatedOffer v-for="i in 5" :key="i" @save-job="emit('saveJob')" />
    </ul>
  </section>
</template>
