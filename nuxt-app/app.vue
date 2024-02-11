<script setup lang="ts">
if (process.server) {
  try {
    const response = await fetchApi<{ success: boolean }>('/user/auth', {
      headers: {
        ...useRequestHeaders()
      }
    })

    if (response.success) {
      useUser().loggedIn = true
    }
  }
  catch (error) {
    logError(error)
  }
}
useHead({
  titleTemplate: title => title ?? 'Home'
})
</script>

<template>
  <NuxtLoadingIndicator
    :duration="500"
    :height="2"
    color="repeating-linear-gradient(to right,var(--brand-3) 0%,var(--brand-2) 50%,var(--brand) 100%)"
  />
  <AppTheNavbar />
  <NuxtPage>
  </NuxtPage>
  <AppTheFooter />
</template>
