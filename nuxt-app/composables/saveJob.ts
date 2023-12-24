export async function saveJob(jobId: string): Promise<{ success: boolean }> {
  const response = await fetchApi('/user/saved-jobs', {
    method: 'POST',
    body: {
      jobId
    }
  })

  return { success: response === 'OK' }
}
