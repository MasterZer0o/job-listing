<script setup lang="ts">
const filters = useListingFilters()
</script>

<template>
  <section v-show="filters.shouldShowApplied" class="applied-filters">
    <ListingAppliedFiltersAppliedFilter v-if="filters.minSalary > 0" store="minSalary" :name="`Min. ${filters.minSalary} PLN`" />
    <ListingAppliedFiltersAppliedFilter v-for="workMode in filters.workModes" :name="workMode.name" store="workModes" />
    <ListingAppliedFiltersAppliedFilter v-for="employmentType in filters.employmentTypes" :name="employmentType.name" store="employmentTypes" />
    <ListingAppliedFiltersAppliedFilter v-for="typeOfWork in filters.typesOfWork" store="typesOfWork" :name="typeOfWork.name" />
    <ListingAppliedFiltersAppliedFilter v-for="expLevel in filters.experienceLevels" store="experienceLevels" :name="expLevel.name" />
    <ListingAppliedFiltersAppliedFilter v-for="techSkill in filters.techSkills" store="techSkills" :name="techSkill" />
  </section>
</template>
