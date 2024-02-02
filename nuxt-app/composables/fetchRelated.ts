export interface RelatedOffer {
  id: string
  title: string
  remoteAvailable: boolean
  companyName: string
  isSaved: boolean
  skills: string[]
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
    store.results = results
  }
  catch (error) {

  }
}
