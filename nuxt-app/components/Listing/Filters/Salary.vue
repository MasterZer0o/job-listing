<script setup lang="ts">
const { minSalary } = storeToRefs(useListingFilters())
const jobsStore = useJobs()
const input = ref() as Ref<HTMLInputElement>
const route = useRoute()
const router = useRouter()
watch(minSalary, async () => {
  jobsStore.currentPage = 1
  await router.push({
    query: {
      ...route.query,
      ms: minSalary.value !== 0 ? minSalary.value : undefined
    }
  })
  useListingFilters().filtersApplied = false

  useLoadingIndicator().finish()
  if (minSalary.value === 0)
    input.value.value = ''
})
const valueFromUrl = Number.parseInt(route.query.ms as string) || 0
minSalary.value = valueFromUrl

function addSalaryFilter() {
  const value = input.value.valueAsNumber

  if (Number.isNaN(value) || value === 0 || value < 0) {
    input.value.value = ''
    minSalary.value = 0
    return
  }

  if (input.value.value.startsWith('0')) {
    input.value.value = value.toString()
  }

  minSalary.value = value
}
</script>

<template>
  <section class="min-salary">
    <p>Minimum salary</p>
    <label>
      <input ref="input" type="number" min="0" placeholder="0" :value="minSalary === 0 ? undefined : minSalary" @change="addSalaryFilter">
      <span>PLN</span>
    </label>
  </section>
</template>
