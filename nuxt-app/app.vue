<script setup lang="ts">
useHead({
  titleTemplate: title => title ?? 'Home'
})
</script>

<template>
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
