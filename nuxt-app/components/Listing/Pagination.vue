<script setup lang="ts">
const props = defineProps<{ pageFromQuery: number }>()

// TODO: paginated results Map<pagenumber, results>
const store = useJobs()
const currentPage = ref<number>(Number.isNaN(props.pageFromQuery) ? 1 : props.pageFromQuery)
const lastPage = computed(() => store.totalPages ?? 0)

const nextPage = computed(() => {
  return currentPage.value === lastPage.value ? lastPage.value : currentPage.value + 1
})
const previousPage = computed(() => {
  return currentPage.value === 1 ? 1 : currentPage.value - 1
})
const TOTAL_VISIBLE_PAGES = 5
const visiblePages = computed(() => {
  const totalPages = store.totalPages
  const currPage = currentPage.value
  if (totalPages === 1)
    return [1]

  if (totalPages && totalPages - currPage < TOTAL_VISIBLE_PAGES) {
    return [totalPages - 5, totalPages - 4, totalPages - 3, totalPages - 2, totalPages - 1]
  }

  if (totalPages && totalPages > TOTAL_VISIBLE_PAGES)
    return [currPage, currPage + 1, currPage + 2, currPage + 3, currPage + 4]

  return []
})

async function changePage(page: number) {
  if (page === currentPage.value)
    return

  currentPage.value = page
  const newJobs = await fetchJobs()
  // TODO: 

  document.querySelector('.listing-container')?.scrollIntoView({ behavior: 'smooth' })
}
</script>

<template>
  <section v-if="store.displayedJobs.length > 0 && lastPage > 0" class="pagination">
    <div>
      <NuxtLink
        :to="`/jobs?p=${previousPage}`" :aria-label="`Page ${previousPage}`"
        :class="{ disabled: currentPage === 1 }"
        @click.prevent="changePage(previousPage)">
        <svg width="15" height="15" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 1 1 5l4 4" />
        </svg>
      </NuxtLink>
      <ul>
        <li v-for="page in visiblePages" :key="page" :class="{ active: currentPage === page }" @click.prevent="changePage(page)">
          <NuxtLink :to="`/jobs?p=${page}`" :aria-label="`Page ${page}`">
            {{ page }}
          </NuxtLink>
        </li>
        <li class=":uno: px-4">
          ...
        </li>
        <li :class="{ active: currentPage === lastPage }" @click.prevent="changePage(lastPage)">
          <NuxtLink :to="`/jobs?p=${lastPage}`">
            {{ lastPage }}
          </NuxtLink>
        </li>
      </ul>
      <NuxtLink
        :to="`/jobs?p=${nextPage}`" :aria-label="`Page ${nextPage}`"
        :class="{ disabled: currentPage === lastPage }"
        @click.prevent="changePage(nextPage)">
        <svg width="15" height="15" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4" />
        </svg>
      </NuxtLink>
    </div>
  </section>
</template>
