<script setup lang="ts">
const jobId = useRoute('job-jobId').params.jobId
provide('job_details_jobId', jobId)

const { data: response } = await useAsyncData(() => fetchApi<JobDetailsResponse>(`/job/${jobId}`, {
  headers: {
    cookie: useRequestHeaders().cookie
  }
}),
)
provide('is_job_saved', response.value?.data.isSaved)

useHead({
  titleTemplate() {
    return response.value?.data.title ?? 'Job offer'
  },
})
</script>

<template>
  <div v-if="response !== null" class="job-details-wrapper">
    <main>
      <JobDetailsHeader
        :company="response.data.company" :title="response.data.title"
        :salary="{ from: response.data.salaryFrom, to: response.data.salaryTo }" />
      <JobDetailsContent :posted-at="response.data.createdAt" />
    </main>
    <JobDetailsSidePane />
  </div>
</template>
