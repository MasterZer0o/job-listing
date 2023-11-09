<script setup lang="ts">
defineProps<{ pageFromQuery: number }>()
const store = useJobs()

type PageNumber = number
const paginatedResults = new Map<PageNumber, Offer[]>()

paginatedResults.set(store.currentPage, store.displayedJobs)

const lastPage = computed(() => store.totalPages ?? 0)

const nextPage = computed(() => {
  return store.currentPage === lastPage.value ? lastPage.value : store.currentPage + 1
})
const previousPage = computed(() => {
  return store.currentPage === 1 ? 1 : store.currentPage - 1
})

const TOTAL_VISIBLE_PAGES = 5
const visiblePages = computed(() => {
  const totalPages = store.totalPages
  const currPage = store.currentPage
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
  if (page === store.currentPage)
    return

  const cachedPage = paginatedResults.get(page)

  if (cachedPage) {
    store.displayedJobs = cachedPage
    store.currentPage = page
    return
  }
  // go back to uncached page OR jump pages
  if (store.currentPage > page || Math.abs(page - store.currentPage) !== 1) {
    store.cid = undefined
  }

  store.currentPage = page
  const newJobs = await fetchJobs()
  store.displayedJobs = newJobs!.data
  store.cid = newJobs!.cid
  paginatedResults.set(store.currentPage, newJobs!.data)

  // document.querySelector('.listing-container')?.scrollIntoView({ behavior: 'smooth' })
}
</script>

<template>
  <section class="pagination">
    <div>
      <NuxtLink
        :class="{ disabled: store.currentPage === 1 }"
        to="/jobs?p=1"
        aria-label="Go to page 1"
        @click.prevent="changePage(1)">
        <svg width="15" height="15" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 1 1 5l4 4m6-8L7 5l4 4" />
        </svg>
      </NuxtLink>
      <NuxtLink
        :to="`/jobs?p=${previousPage}`" :aria-label="`Go to page ${previousPage}`"
        :class="{ disabled: store.currentPage === 1 }"
        @click.prevent="changePage(previousPage)">
        <svg width="15" height="15" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 1 1 5l4 4" />
        </svg>
      </NuxtLink>
      <ul>
        <li v-for="page in visiblePages" :key="page" :class="{ active: store.currentPage === page }" @click.prevent="changePage(page)">
          <NuxtLink :to="`/jobs?p=${page}`" :aria-label="`Go to page ${page}`">
            {{ page }}
          </NuxtLink>
        </li>
        <li class=":uno: px-4">
          ...
        </li>
        <li :class="{ active: store.currentPage === lastPage }" @click.prevent="changePage(lastPage)">
          <NuxtLink :to="`/jobs?p=${lastPage}`">
            {{ lastPage }}
          </NuxtLink>
        </li>
      </ul>
      <NuxtLink
        :to="`/jobs?p=${nextPage}`" :aria-label="`Go to page ${nextPage}`"
        :class="{ disabled: store.currentPage === lastPage }"
        @click.prevent="changePage(nextPage)">
        <svg width="15" height="15" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4" />
        </svg>
      </NuxtLink>
      <NuxtLink
        :to="`/jobs?p=${lastPage}`" :aria-label="`Go to page ${lastPage}`"
        :class="{ disabled: store.currentPage === lastPage }"
        @click.prevent="changePage(lastPage)">
        <svg width="15" height="15" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9 4-4-4-4M1 9l4-4-4-4" />
        </svg>
      </NuxtLink>
    </div>
  </section>
</template>
