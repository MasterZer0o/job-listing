export interface RelatedOffer {
  id: string
  title: string
  remoteAvailable: boolean
  salary: {
    from: string
    to: string
  }
}

export async function fetchRelated(jobId: string) {
  try {
    const store = jobDetailsStore()
    const { results } = await fetchApi<{ results: RelatedOffer[] }>(`/job/${jobId}/related`)
    store.isLoading = false
    logInfo({ results })
    store.results = results
    logInfo(store.results)
  }
  catch (error) {

  }
}
