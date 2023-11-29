<script setup lang="ts">
const props = defineProps<{
  config: {
    store: DefaultFilterStoreNames
    title: string
    queryParamName: string
    options: DefaultFilterField[]
  }
}>()

const selectedFields = storeToRefs(useListingFilters())[props.config.store]
const jobsStore = useJobs()
let toggledAll = false
function clearAll() {
  selectedFields.value = []
  toggledAll = false
}
const selectedFieldsIds = computed(() => {
  return selectedFields.value.map(e => e.id).join(',')
})
const route = useRoute('jobs')
const router = useRouter()

function selectField(field: DefaultFilterField | DefaultFilterField[]) {
  if (Array.isArray(field)) {
    if (toggledAll)
      selectedFields.value = []
    else
      selectedFields.value = [...field]

    toggledAll = !toggledAll
    return
  }
  if (selectedFields.value.some(selectedField => selectedField.name === field.name)) {
    selectedFields.value.splice(selectedFields.value.findIndex(e => e.name === field.name), 1)
  }
  else
    selectedFields.value.push(field)
}
watch(selectedFields, () => {
  jobsStore.currentPage = 1
  router.push({
    query: {
      ...route.query,
      [props.config.queryParamName]: selectedFieldsIds.value.length !== 0 ? selectedFieldsIds.value : undefined
    }
  })
}, { deep: true })
</script>

<template>
  <div>
    <span @click="selectField(config.options)">{{ config.title }}</span>
    <button type="button" @click="clearAll">
      clear
    </button>
  </div>
  <ul>
    <li v-for="option in config.options" :key="option.name">
      <label>
        <input type="checkbox" :checked="selectedFields.some(e => e.name === option.name)" @input="selectField(option)">
        <span>
          <svg viewBox="0 0 20 20" fill="none" stroke="#20c2e0" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" stroke-dasharray="71" stroke-dashoffset="71">
            <path d="M3,1 L17,1 L17,1 C18.1045695,1 19,1.8954305 19,3 L19,17 L19,17 C19,18.1045695 18.1045695,19 17,19 L3,19 L3,19 C1.8954305,19 1,18.1045695 1,17 L1,3 L1,3 C1,1.8954305 1.8954305,1 3,1 Z"></path>
            <polyline points="4 11 8 15 16 6" fill="none" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" stroke-dasharray="18" stroke-dashoffset="18"></polyline>
          </svg>
        </span>
        {{ option.name }}
      </label>
    </li>
  </ul>
</template>
