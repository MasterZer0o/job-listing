export const useListingFilters = defineStore('listingFilters', () => {
  const techSkills = ref<Map<number, { name: string }>>(new Map())
  const employmentTypes = ref<DefaultFilterField[]>([])
  const experienceLevels = ref<DefaultFilterField[]>([])
  const typesOfWork = ref<DefaultFilterField[]>([])
  const workModes = ref<DefaultFilterField[]>([])
  const minSalary = ref<number>(0)
  const filtersApplied = ref(true)

  const shouldShowApplied = computed(() => {
    return techSkills.value.size !== 0
      || filtersAreNotEmpty(employmentTypes, experienceLevels, typesOfWork, workModes)
      || minSalary.value > 0
  })

  function clearAll() {
    minSalary.value = 0
    experienceLevels.value = []
    employmentTypes.value = []
    techSkills.value.clear()
    workModes.value = []
    typesOfWork.value = []
  }

  return {
    techSkills,
    employmentTypes,
    experienceLevels,
    typesOfWork,
    workModes,
    minSalary,
    shouldShowApplied,
    clearAll,
    filtersApplied
  }
})

function filtersAreNotEmpty(...arr: Ref<any[]>[]) {
  return arr.some(filter => filter.value.length !== 0)
}

declare global {
  export type FilterStoreNames = 'employmentTypes' | 'experienceLevels' | 'typesOfWork' | 'workModes' | 'techSkills' | 'minSalary'
  export type DefaultFilterStoreNames = Exclude<FilterStoreNames, 'techSkills' | 'minSalary'>
  export interface DefaultFilterField {
    id: number
    name: string
  }
  export type FilterQueryName = 'emp' | 'tech' | 'tow' | 'mode' | 'exp' | 'ms' | 'p'
}
