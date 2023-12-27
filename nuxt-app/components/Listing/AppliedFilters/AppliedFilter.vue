<script setup lang="ts">
import { debounce } from 'perfect-debounce'

const props = defineProps<{
  filterQueryName: FilterQueryName
  option: {
    name: string
    id: number
  }
  store: FilterStoreNames
}>()
function removeFilter(id: number) {
  const store = storeToRefs(useListingFilters())[props.store!]

  if (Array.isArray(store.value))
    store.value.splice(store.value.findIndex(e => e.id === id), 1)
  else if (props.store === 'minSalary')
    store.value = 0
  else
    (store.value as Map<number, { name: string }>).delete(id)
}

const refetch = debounce(async () => {
  if (useListingFilters().filtersApplied)
    await fetchJobs({ withCount: true })
}, 1000)
</script>

<template>
  <div>
    <span>{{ option.name }}</span>
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="-0.5 -0.5 14 14" height="14" width="14" @click="[removeFilter(option.id), refetch()]"><g><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 12.535714285714286c3.3334785714285715 0 6.035714285714286 -2.702235714285714 6.035714285714286 -6.035714285714286C12.535714285714286 3.1665678571428573 9.833478571428572 0.4642857142857143 6.5 0.4642857142857143 3.1665678571428573 0.4642857142857143 0.4642857142857143 3.1665678571428573 0.4642857142857143 6.5c0 3.3334785714285715 2.702282142857143 6.035714285714286 6.035714285714286 6.035714285714286Z" stroke-width="1.5"></path><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.7142857142857144 6.5h5.571428571428571" stroke-width="2"></path></g></svg>
  </div>
</template>
