<script setup lang="ts">
interface Filter {
  store: DefaultFilterStoreNames
  title: string
  queryParamName: FilterQueryName
  options: {
    id: number
    name: string
  }[]
}

const employmentType: Filter = {
  store: 'employmentTypes',
  title: 'Employment type',
  queryParamName: 'emp',
  options: [
    { name: 'B2B', id: 1 },
    { name: 'Permanent', id: 2 },
    { name: 'Internship', id: 3 },
    { name: 'Mandate Contract', id: 4 },
    { name: 'Specific-task Contract', id: 5 }
  ]
}
const experience: Filter = {
  store: 'experienceLevels',
  title: 'Experience',
  queryParamName: 'exp',
  options: [
    { name: 'Junior', id: 1 },
    { name: 'Mid', id: 2 },
    { name: 'Senior', id: 3 }
  ]
}
const typeOfWork: Filter = {
  store: 'typesOfWork',
  title: 'Type of work',
  queryParamName: 'tow',
  options: [
    { name: 'Full-time', id: 1 },
    { name: 'Part-time', id: 2 },
    { name: 'Temporary', id: 3 },
  ]
}
const workMode: Filter = {
  store: 'workModes',
  title: 'Work mode',
  queryParamName: 'mode',
  options: [
    { name: 'Stationary', id: 1 },
    { name: 'Hybrid', id: 2 },
    { name: 'Remote', id: 3 },
  ]
}
const filters: Filter[] = [employmentType, experience, typeOfWork, workMode]
const route = useRoute()
const query = route.query as Partial<Record<Filter['queryParamName'], string>>
const store = useListingFilters()
const keys = Object.keys(query)
const filtersFromUrl = filters.filter((filter) => {
  return keys.includes(filter.queryParamName)
})

filtersFromUrl.forEach((filter) => {
  const ids = query[filter.queryParamName]!.split(',').map(e => Number.parseInt(e))
  const filtered = filter.options.filter(e => ids.includes(e.id))
  store[filter.store] = filtered
})

const pinned = ref(true)
const wrapper = ref() as Ref<HTMLElement>
function updatePin() {
  pinned.value = !pinned.value
}

let savedFilters = { ...route.query as Partial<Record<Filter['queryParamName'], string>> }

const jobsStore = useJobs()
const queryParamNames = ['tech', 'ms', ...filters.map(f => f.queryParamName)]
async function applyFilters() {
  for (const k of queryParamNames) {
    if (savedFilters[k as FilterQueryName] !== route.query[k]) {
      savedFilters = route.query
      const results = await fetchJobs({ withCount: true })

      jobsStore.paginatedResults.clear()
      if (results!.data) {
        jobsStore.displayedJobs = results!.data
      }
      document.querySelector('.main-wrapper > div header')?.scrollIntoView({ behavior: 'smooth' })

      return
    }
  }
}
</script>

<template>
  <section ref="wrapper" class="listing-filters" :class="{ 'not-pinned': !pinned }">
    <svg class="pin-icon" width="20" height="20" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" @click="updatePin">
      <title>Click to {{ pinned ? 'unpin' : 'pin' }} filters panel</title>
      <path :fill="pinned ? 'currentColor' : 'none'" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 20l5-5m0 0l3.956 3.956a1 1 0 0 0 1.626-.314l2.255-5.261a1 1 0 0 1 .548-.535l3.207-1.283a1 1 0 0 0 .336-1.635l-6.856-6.856a1 1 0 0 0-1.635.336l-1.283 3.207a1 1 0 0 1-.535.548L5.358 9.418a1 1 0 0 0-.314 1.626L9 15z" />
    </svg>
    <ul>
      <li>
        <ListingFiltersSalary />
      </li>
      <li v-for="filter in filters.slice(0, filters.length / 2)" :key="filter.title">
        <ListingFiltersFilter :config="filter" />
      </li>
      <li>
        <ListingFiltersTech />
      </li>
      <li v-for="filter in filters.slice(filters.length / 2)" :key="filter.title">
        <ListingFiltersFilter :config="filter" />
      </li>
    </ul>
    <button type="button" @click="applyFilters">
      Apply
    </button>
  </section>
</template>
