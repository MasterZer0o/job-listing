export const useJobs = defineStore('jobs', () => {
  const displayedJobs = ref<Offer[]>([])
  const totalCount = ref<number>()
  const totalPages = ref<number>()
  const cid = ref<number>()
  const paginatedResults = ref(new Map<PageNumber, Offer[]>())

  const currentPage = ref<number>(1)

  const isLoading = ref(false)
  const isFetchingInitial = ref(true)

  function clear() {
    displayedJobs.value = []
    totalCount.value = undefined
    totalPages.value = undefined
    currentPage.value = 1
    cid.value = undefined
  }
  const error = ref<Error>({ value: false })
  const MAX_ERROR_RETRIES = 2
  let retries = 0

  return {
    displayedJobs,
    totalCount,
    totalPages,
    currentPage,
    cid,
    clear,
    paginatedResults,
    isLoading,
    error,
    MAX_ERROR_RETRIES,
    retries,
    isFetchingInitial
  }
})
interface Error {
  value: boolean
  message?: string
  net?: {
    restored: boolean
  }
}

type PageNumber = number
