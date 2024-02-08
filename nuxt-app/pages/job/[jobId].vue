<script setup lang="ts">
const jobId = useRoute('job-jobId').params.jobId
provide('job_details_jobId', jobId)

const { data: response } = await useAsyncData(() => fetchApi<JobDetailsResponse>(`/job/${jobId}`, {
  headers: {
    cookie: useCookie('session').value ? `session=${useCookie('session').value}` : undefined!
  }
}),
)

provide('is_job_saved', response.value?.data.isSaved)
onMounted(() => {
  fetchRelated(jobId)
})
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
      <JobDetailsSectionNavigator v-if="response.data.contentSections.length > 0" :elements="response.data.contentSections" />
      <JobDetailsContent :posted-at="response.data.createdAt" :content="response.data.content" />
    </main>
    <JobDetailsSidePane />
  </div>
</template>
