export async function saveJob(jobId: string, isSaved: boolean): Promise<{ success: boolean }> {
  const response = await fetchApi('/user/saved-jobs', {
    method: 'POST',
    body: {
      jobId,
      isSaved
    }
  })

  return { success: response === 'OK' }
}
