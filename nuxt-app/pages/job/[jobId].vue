<script setup lang="ts">
const jobId = useRoute('job-jobId').params.jobId

let num = 10000
num.toLocaleString('en-US').replaceAll(',', ' ')

const { data: response } = await useFetch<JobDetailsResponse>(`${useRuntimeConfig().public.API_BASE}/job/${jobId}`)
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
    </main>
    <JobDetailsSidePane />
  </div>
</template>
