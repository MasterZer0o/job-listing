<script setup lang="ts">

</script>

<template>
  <section class="related-offers">
    <p>Related offers</p>
    <ul>
      <JobDetailsRelatedOffer v-for="i in 5" :key="i" />
    </ul>
  </section>
</template>
