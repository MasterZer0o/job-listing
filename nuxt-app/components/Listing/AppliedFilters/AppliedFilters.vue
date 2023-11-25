<script setup lang="ts">
const filters = useListingFilters()
const router = useRouter()
async function clearAllFilters() {
  await router.push({
    query: undefined
  })

  filters.minSalary = 0
  filters.experienceLevels = []
  filters.employmentTypes = []
  filters.techSkills.clear()
  filters.workModes = []
  filters.typesOfWork = []
}
</script>

<template>
  <section v-show="filters.shouldShowApplied" class="applied-filters">
    <div>
      <ListingAppliedFiltersAppliedFilter
        v-if="filters.minSalary > 0"
        filter-query-name="ms"
        store="minSalary"
        :option="{ name: `Min. ${filters.minSalary} PLN`, id: 0 }" />

      <ListingAppliedFiltersAppliedFilter
        v-for="workMode in filters.workModes"
        filter-query-name="mode"
        :option="workMode"
        store="workModes" />

      <ListingAppliedFiltersAppliedFilter
        v-for="employmentType in filters.employmentTypes"
        filter-query-name="emp"
        :option="employmentType"
        store="employmentTypes" />

      <ListingAppliedFiltersAppliedFilter
        v-for="typeOfWork in filters.typesOfWork"
        filter-query-name="tow"
        :option="typeOfWork"
        store="typesOfWork" />

      <ListingAppliedFiltersAppliedFilter
        v-for="expLevel in filters.experienceLevels"
        filter-query-name="exp"
        store="experienceLevels"
        :option="expLevel" />
      <ListingAppliedFiltersAppliedFilter
        v-for="techSkill in filters.techSkills"
        filter-query-name="tech"
        store="techSkills"
        :option="{ name: techSkill[1].name, id: techSkill[0] }" />
    </div>
    <button type="button" @click="clearAllFilters">
      clear all
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="-0.5 -0.5 14 14" height="14" width="14"><g><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 12.535714285714286c3.3334785714285715 0 6.035714285714286 -2.702235714285714 6.035714285714286 -6.035714285714286C12.535714285714286 3.1665678571428573 9.833478571428572 0.4642857142857143 6.5 0.4642857142857143 3.1665678571428573 0.4642857142857143 0.4642857142857143 3.1665678571428573 0.4642857142857143 6.5c0 3.3334785714285715 2.702282142857143 6.035714285714286 6.035714285714286 6.035714285714286Z" stroke-width="1.5"></path><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.7142857142857144 6.5h5.571428571428571" stroke-width="2"></path></g></svg>
    </button>
  </section>
</template>
