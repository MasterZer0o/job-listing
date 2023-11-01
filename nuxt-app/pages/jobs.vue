<script setup lang="ts">
useHead({
  title: 'Jobs'
})
const store = useJobs()
async function fetchJobs() {
  const apiUrl = useRuntimeConfig().public.API_BASE
  try {
    $fetch<{ count: number;totalPages: number }>(`${apiUrl}/jobs/count`).then(({ count, totalPages }) => {
      store.totalCount = count
      store.totalPages = totalPages
    })
    const jobs = await $fetch<Offer[]>(`${apiUrl}/jobs`)

    store.displayedJobs = jobs
  }
  catch (error) {

  }
}
onBeforeMount(() => {
  fetchJobs()
})
</script>

<template>
  <main class="listing-container">
    <header>
      <span>Results {{ store.totalCount !== undefined ? `(${store.totalCount})` : '' }}</span>
    </header>
    <ListingWrapper :offers="store.displayedJobs" />
    <ListingPagination />
  </main>
</template>
