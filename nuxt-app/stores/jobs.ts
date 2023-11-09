export const useJobs = defineStore('jobs', () => {
  const displayedJobs = ref<Offer[]>([])
  const totalCount = ref<number>()
  const totalPages = ref<number>()
  const cid = ref<number>()

  const currentPage = ref<number>(1)

  return {
    displayedJobs,
    totalCount,
    totalPages,
    currentPage,
    cid
  }
})
