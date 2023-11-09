export async function fetchJobs() {
  const store = useJobs()
  try {
    const response = await $fetch<{ data: Offer[]; cid: number }>(`${useRuntimeConfig().public.API_BASE}/jobs`, {
      query: {
        p: store.currentPage === 1 ? undefined : store.currentPage,
        cid: store.cid
      }
    })

    return response
  }
  catch (error) {

  }
}
