export async function fetchJobs({ withCount } = { withCount: false }) {
  const store = useJobs()
  if (store.error.net && !store.error.net.restored)
    return

  if (store.error.value) {
    if (store.MAX_ERROR_RETRIES === store.retries)
      return

    else {
      if (store.MAX_ERROR_RETRIES - 1 === store.retries)
        store.error.message = 'Something is wrong on our side. Try again later.'
      store.retries++
    }
  }

  const filters = Object.fromEntries(new URLSearchParams(window.location.search).entries())
  if (withCount) {
    store.cid = undefined
    store.paginatedResults.clear()
    fetchApi<{ count: number; totalPages: number }>('/jobs/count', {
      query: {
        cid: store.cid,
        ...filters,
        p: store.currentPage === 1 ? undefined : store.currentPage
      }
    }).then(({ count, totalPages }) => {
      store.totalCount = count
      store.totalPages = totalPages
    }).catch()
  }
  const timeout = setTimeout(() => {
    store.isLoading = true
  }, 100)

  if (store.error.value)
    store.isLoading = true

  store.error.value = false
  try {
    const response = await fetchApi<Response>('/jobs', {
      query: {
        cid: store.cid,
        ...filters,
        p: store.currentPage === 1 ? undefined : store.currentPage
      }
    })
    clearTimeout(timeout)
    useListingFilters().filtersApplied = true

    if ('error' in response) {
      store.error.value = true
      store.isLoading = false
      return
    }

    if (response.data) {
      store.displayedJobs = response!.data
      store.cid = response!.cid
      store.paginatedResults.set(store.currentPage, response!.data)
    }
    store.error = {
      value: false
    }

    return response
  }
  catch (error) {
    store.error.value = true
    clearTimeout(timeout)

    if (!window?.navigator?.onLine)
      store.error.net = { restored: false }

    window.addEventListener('online', () => {
      store.error.net!.restored = true
    }, { once: true })
  }

  finally {
    store.isLoading = false
  }
}

type Response = { data: Offer[]; cid: number } | { error: string }
