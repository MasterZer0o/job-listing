<script setup lang="ts">
const { minSalary } = storeToRefs(useListingFilters())

const input = ref() as Ref<HTMLInputElement>
watch(minSalary, () => {
  if (minSalary.value === 0)
    input.value.value = ''
})

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
      <input ref="input" type="number" min="0" placeholder="0" @change="addSalaryFilter">
      <span>pln</span>
    </label>
  </section>
</template>
