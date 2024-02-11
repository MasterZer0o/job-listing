<script setup lang="ts">

</script>

<template>
  <nav class="app-nav">
    <span class=":uno: absolute top-0 left-0 text-xs">last update: {{ useRuntimeConfig().public.lastUpdate }}</span>

    <div class=":uno: text-4xl mr-10">
      LOGO
    </div>
    <ul>
      <li>
        <NuxtLink to="/jobs" :prefetch="false">
          Jobs
        </NuxtLink>
      </li>
      <li>
        <NuxtLink to="/">
          Link 2
        </NuxtLink>
      </li>
      <li>
        <NuxtLink to="/">
          Link 3
        </NuxtLink>
      </li>
      <li>
        <NuxtLink to="/">
          Link 4
        </NuxtLink>
      </li>
    </ul>
    <AppThemeSwitcher />
    <AppNavUser />
  </nav>
</template>
