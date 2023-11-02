<script setup lang="ts">
useHead({
  title: 'Jobs'
})

const route = useRoute('jobs')
const pageFromQuery = Number.parseInt(route.query.p as string)
const store = useJobs()
store.currentPage = pageFromQuery || 1

async function initialFetchJobs() {
  const apiUrl = useRuntimeConfig().public.API_BASE
  try {
    $fetch<{ count: number;totalPages: number }>(`${apiUrl}/jobs/count`).then(({ count, totalPages }) => {
      store.totalCount = count
      store.totalPages = totalPages
    }).catch()

    const jobs = await fetchJobs()
    // TODO: handle errors
    store.displayedJobs = jobs!
  }
  catch (error) {

  }
}
onBeforeMount(() => {
  initialFetchJobs()
})
</script>

<template>
  <main class="listing-container">
    <header>
      <span>Results {{ store.totalCount !== undefined ? `(${store.totalCount})` : '' }}</span>
    </header>
    <ListingWrapper :offers="store.displayedJobs" />
    <ListingPagination :page-from-query="pageFromQuery" />
  </main>
  <AppScrollUp />
</template>
