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

  const API_URL = useRuntimeConfig().public.API_BASE
  const filters = Object.fromEntries(new URLSearchParams(window.location.search).entries())

  if (withCount) {
    store.cid = undefined
    store.paginatedResults.clear()
    $fetch<{ count: number; totalPages: number }>(`${API_URL}/jobs/count`, {
      query: {
        p: store.currentPage === 1 ? undefined : store.currentPage,
        cid: store.cid,
        ...filters
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
    const response = await $fetch<Response>(`${API_URL}/jobs`, {
      query: {
        p: store.currentPage === 1 ? undefined : store.currentPage,
        cid: store.cid,
        ...filters
      }
    })
    clearTimeout(timeout)
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
