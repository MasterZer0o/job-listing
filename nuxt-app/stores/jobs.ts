export const useJobs = defineStore('jobs', () => {
  const displayedJobs = ref<Offer[]>([])
  const totalCount = ref<number>()
  const totalPages = ref<number>()
  const cid = ref<number>()
  const paginatedResults = ref(new Map<PageNumber, Offer[]>())

  const currentPage = ref<number>(1)

  const isLoading = ref(false)

  function clear() {
    displayedJobs.value = []
    totalCount.value = undefined
    totalPages.value = undefined
    currentPage.value = 1
    cid.value = undefined
  }

  return {
    displayedJobs,
    totalCount,
    totalPages,
    currentPage,
    cid,
    clear,
    paginatedResults,
    isLoading
  }
})

type PageNumber = number
