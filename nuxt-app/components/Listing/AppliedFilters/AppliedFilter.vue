<script setup lang="ts">
const props = defineProps<{
  name: string
  store: FilterStoreNames
}>()
function removeFilter(name: string) {
  const store = storeToRefs(useListingFilters())[props.store!]

  if (Array.isArray(store.value))
    store.value.splice(store.value.findIndex(e => e.name === name), 1)
  else if (props.store === 'minSalary')
    store.value = 0
  else
    (store.value as Set<any>).delete(name)
}
// TODO: initialize filters from pupulated URL
</script>

<template>
  <div>
    <span>{{ name }}</span>
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="-0.5 -0.5 14 14" height="14" width="14" @click="removeFilter(name)"><g id="subtract-circle--delete-add-circle-subtract-button-buttons-remove-mathematics-math-minus"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 12.535714285714286c3.3334785714285715 0 6.035714285714286 -2.702235714285714 6.035714285714286 -6.035714285714286C12.535714285714286 3.1665678571428573 9.833478571428572 0.4642857142857143 6.5 0.4642857142857143 3.1665678571428573 0.4642857142857143 0.4642857142857143 3.1665678571428573 0.4642857142857143 6.5c0 3.3334785714285715 2.702282142857143 6.035714285714286 6.035714285714286 6.035714285714286Z" stroke-width="1.5"></path><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.7142857142857144 6.5h5.571428571428571" stroke-width="2"></path></g></svg>
  </div>
</template>
