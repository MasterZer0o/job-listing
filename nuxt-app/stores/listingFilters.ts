export const useListingFilters = defineStore('listingFilters', () => {
  const techSkills = ref<Set<string>>(new Set<string>())

  return { techSkills }
})
