<script setup lang="ts">
function scrollUp() {
  window.scroll(0, 0)
}

const showButton = ref<boolean>(false)
const controller = new AbortController()

process.client && window.addEventListener('scroll',
  () => showButton.value = window.scrollY > 150,
  { signal: controller.signal }
)

onUnmounted(() => controller.abort())
</script>

<template>
  <button v-show="showButton" class="scroll-up" title="scroll to top" type="button" aria-label="scroll to top" @click="scrollUp">
    <svg
      color="currentColor"
      width="35"
      height="35"
      stroke-width="1.5"
      viewBox="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M12 16V8M12 8L15.5 11.5M12 8L8.5 11.5" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" />
      <path
        d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z"
        stroke="currentColor"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  </button>
</template>
