<script setup lang="ts">
defineProps<{
  offers: Offer[]
}>()
</script>

<template>
  <section class="offers">
    <ListingOffer v-for="offer in offers" :key="offer.id" :offer="offer" />
  </section>
</template>
