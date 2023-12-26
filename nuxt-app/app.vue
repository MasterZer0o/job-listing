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
const theme = ref('dark')
function changeTheme() {
  document.body.setAttribute('class', 'theming')
  setTimeout(() => {
    document.body.removeAttribute('class')
  }, 350)
  theme.value = theme.value === 'dark' ? 'light' : 'dark'
}
useHead({
  titleTemplate: title => title ?? 'Home'
})
</script>

<template>
  <Html :color-scheme="theme" />
  <DevOnly>
    <button
      type="button"
      style="position: absolute; z-index: 999;"
      @click="changeTheme">
      theme
    </button>
  </DevOnly>

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
