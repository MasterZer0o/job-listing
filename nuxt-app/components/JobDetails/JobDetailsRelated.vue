<script setup lang="ts">
const emit = defineEmits(['saveJob'])
const store = jobDetailsStore()
</script>

<template>
  <section class="related-offers">
    <p>Related offers</p>
    <ul v-if="store.isLoading" class=":uno: flex flex-col gap-6 py-4">
      <JobDetailsRelatedOfferSkeleton v-for="i in 8" :key="i" />
    </ul>
    <ul v-else class="related-offers-list">
      <JobDetailsRelatedOffer v-for="item in store.results" :key="item.id" :data="item" @save-job="emit('saveJob')" />
    </ul>
  </section>
</template>
