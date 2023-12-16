<script setup lang="ts">
useHead({
  title: 'Jobs'
})

const route = useRoute('jobs')
const pageFromQuery = Number.parseInt(route.query.p as string)
const jobsStore = useJobs()
jobsStore.currentPage = pageFromQuery || 1
const isFetching = ref(true)
onBeforeMount(async () => {
  try {
    await fetchJobs({ withCount: true })
  }
  catch (error) {

  }
  finally {
    isFetching.value = false
  }
})

onUnmounted(() => { jobsStore.clear(); useListingFilters().clearAll() })
</script>

<template>
  <div class="main-wrapper">
    <ListingFilters />
    <div>
      <ListingAppliedFilters />
      <header>
        <span>Results {{ jobsStore.totalCount !== undefined ? `(${jobsStore.totalCount})` : '' }}</span>
      </header>
      <main class="listing-container">
        <ListingOffers :offers="jobsStore.displayedJobs" :is-fetching="isFetching" />
        <ListingPagination v-if="!jobsStore.error.value && jobsStore.displayedJobs.length > 0 && (jobsStore.totalPages ?? 0) > 0" :page-from-query="pageFromQuery" />
      </main>
    </div>
  </div>
  <AppScrollUp />
</template>
