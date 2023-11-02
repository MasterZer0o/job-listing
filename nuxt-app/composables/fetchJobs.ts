export async function fetchJobs() {
  const store = useJobs()
  try {
    const response = await $fetch<Offer[]>(`${useRuntimeConfig().public.API_BASE}/jobs`, {
      query: {
        p: store.currentPage === 1 ? undefined : store.currentPage
      }
    })

    return response
  }
  catch (error) {

  }
}
