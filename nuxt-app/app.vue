<script setup lang="ts">
if (process.server) {
  try {
    const response = await $fetch<{ success: boolean }>(`${useRuntimeConfig().public.API_BASE}/user/auth`, {
      credentials: 'include',
      headers: {
        ...useRequestHeaders()!
      }
    })

    if (response.success) {
      useUser().loggedIn = true
    }
    logInfo('app.vue ', useUser().loggedIn)
  }
  catch (error) {
    logError(error)
  }
}

useHead({
  titleTemplate: title => title ?? 'Home'
})
const theme = ref('dark')
</script>

<template>
  <Html :color-scheme="theme" />
  <DevOnly>
    <button
      type="button"
      style="position: absolute;"
      @click="theme = theme === 'dark' ? 'light' : 'dark'">
      theme
    </button>
  </DevOnly>

  <NuxtLoadingIndicator
    :throttle="0"
    :duration="500"
    :height="2"
    color="repeating-linear-gradient(to right,var(--brand-3) 0%,var(--brand-2) 50%,var(--brand) 100%)"
  />
  <AppTheNavbar />
  <NuxtPage>
  </NuxtPage>
  <AppTheFooter />
</template>
