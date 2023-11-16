<script setup lang="ts">
interface Filter {
  title: string
  options: {
    name: string
  }[]
}
const employmentType = {
  title: 'Employment type',
  options: [
    {
      name: 'B2B'
    }, { name: 'Permanent' }, { name: 'Internship' }, { name: 'Mandate Contract' }, { name: 'Specific-task Contract' }
  ]
}
const experience: Filter = {
  title: 'Experience',
  options: [
    { name: 'Junior' },
    { name: 'Mid' },
    { name: 'Senior' }
  ]
}
const typeOfWork: Filter = {
  title: 'Type of work',
  options: [
    { name: 'Full time' },
    { name: 'Part time' },
    { name: 'Temporary' },
  ]
}
const workMode: Filter = {
  title: 'Work mode',
  options: [
    { name: 'Stationary' },
    { name: 'Hybrid' },
    { name: 'Remote' },
  ]
}
const filters: Filter[] = [employmentType, experience, typeOfWork, workMode]
const pinned = ref(true)
const wrapper = ref() as Ref<HTMLElement>
function updatePin() {
  pinned.value = !pinned.value
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
  </section>
</template>
