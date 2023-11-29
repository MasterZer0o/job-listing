export async function fetchJobs({ withCount } = { withCount: false }) {
  const store = useJobs()
  const API_URL = useRuntimeConfig().public.API_BASE
  const filters = Object.fromEntries(new URLSearchParams(window.location.search).entries())

  if (withCount) {
    $fetch<{ count: number;totalPages: number }>(`${API_URL}/jobs/count`).then(({ count, totalPages }) => {
      store.totalCount = count
      store.totalPages = totalPages
    }).catch()
  }

  try {
    const response = await $fetch<{ data: Offer[]; cid: number }>(`${API_URL}/jobs`, {
      query: {
        p: store.currentPage === 1 ? undefined : store.currentPage,
        cid: store.cid,
        ...filters
      }
    })

    if (response.data) {
      store.displayedJobs = response!.data
      store.cid = response!.cid
      store.paginatedResults.set(store.currentPage, response!.data)
    }

    return response
  }
  catch (error) {

  }
}
