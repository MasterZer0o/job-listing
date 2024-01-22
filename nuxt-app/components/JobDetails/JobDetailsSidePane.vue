<script setup lang="ts">
const jobId = inject<string>('job_details_jobId', '')
const isSaved = ref(inject('is_job_saved', false))
const shouldShowLoginModal = ref(false)
const shouldImportLoginModal = ref(false)
const copyButtonText = ref('Copy')
function copyJobLink() {
  window.navigator.clipboard.writeText(window.location.href)
  copyButtonText.value = 'Copied!'
  setTimeout(() => copyButtonText.value = 'Copy', 10000)
}
function onLoginSuccess(_isSaved?: boolean) {
  shouldShowLoginModal.value = false
  if (_isSaved)
    isSaved.value = true
  else _saveJob()
}
async function _saveJob() {
  if (!useUser().loggedIn) {
    shouldShowLoginModal.value = true
    if (!shouldShowLoginModal.value)
      shouldImportLoginModal.value = true

    return
  }
  try {
    const response = await saveJob(jobId, isSaved.value)

    if (response.success) {
      isSaved.value = !isSaved.value
    }
  }
  catch (error) {

  }
}
</script>

<template>
  <section class="side-pane">
    <div>
      <button type="button">
        Apply
      </button>
      <ul>
        <li @click="_saveJob">
          <button type="button">
            <svg width="18" height="18" xmlns="http://www.w3.org/2000/svg" :fill="isSaved ? '#ebcb8b' : 'none'" viewBox="0 0 21 20">
              <path :stroke="isSaved ? '#ebcb8b' : 'currentColor'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.479 1.712 2.367 4.8a.532.532 0 0 0 .4.292l5.294.769a.534.534 0 0 1 .3.91l-3.83 3.735a.534.534 0 0 0-.154.473l.9 5.272a.535.535 0 0 1-.775.563l-4.734-2.49a.536.536 0 0 0-.5 0l-4.73 2.487a.534.534 0 0 1-.775-.563l.9-5.272a.534.534 0 0 0-.154-.473L2.158 8.48a.534.534 0 0 1 .3-.911l5.294-.77a.532.532 0 0 0 .4-.292l2.367-4.8a.534.534 0 0 1 .96.004Z" />
            </svg>

            {{ isSaved ? 'Saved' : 'Save' }}
          </button>
        </li>
        <li>
          <button type="button">
            <svg width="18" height="18" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 18 18">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.953 7.467 6.094-2.612m.096 8.114L5.857 9.676m.305-1.192a2.581 2.581 0 1 1-5.162 0 2.581 2.581 0 0 1 5.162 0ZM17 3.84a2.581 2.581 0 1 1-5.162 0 2.581 2.581 0 0 1 5.162 0Zm0 10.322a2.581 2.581 0 1 1-5.162 0 2.581 2.581 0 0 1 5.162 0Z" />
            </svg>
            Share
          </button>
        </li>
        <li @click="copyJobLink">
          <button type="button">
            <svg width="18" height="18" viewBox="0 0 408 480" xmlns="http://www.w3.org/2000/svg">
              <path fill="currentColor" d="M299 5v43H43v299H0V48q0-18 12.5-30.5T43 5zm64 86q17 0 29.5 12.5T405 133v299q0 18-12.5 30.5T363 475H128q-18 0-30.5-12.5T85 432V133q0-17 12.5-29.5T128 91zm0 341V133H128v299z" />
            </svg>
            {{ copyButtonText }}
          </button>
        </li>
      </ul>
    </div>
    <JobDetailsRelated @save-job="_saveJob" />
  </section>

  <Teleport to="body">
    <ClientOnly>
      <Transition v-show="shouldShowLoginModal" name="logreg" mode="out-in">
        <LazyAppLoginModal v-if="shouldShowLoginModal" :callback="onLoginSuccess" @close="shouldShowLoginModal = false" />
      </Transition>
    </ClientOnly>
  </Teleport>
</template>
