export const useJobs = defineStore('jobs', () => {
  const displayedJobs = ref<Offer[]>([])
  const totalCount = ref<number>()
  const totalPages = ref<number>()

  return {
    displayedJobs,
    totalCount,
    totalPages
  }
})
