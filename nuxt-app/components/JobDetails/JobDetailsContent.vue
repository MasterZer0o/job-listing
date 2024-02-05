<script setup lang="ts">
const props = defineProps<{
  postedAt: string
  content: string
}>()

// Description, tech, responsibilities, requirements, benefits, what we offer, recruitment process
const postedAt = new Date(props.postedAt).toLocaleDateString('pl')
</script>

<template>
  <section class="details-content">
    <span>Posted at {{ postedAt }}</span>
    <div v-html="content"></div>
  </section>
</template>
