<script setup lang="ts">

</script>

<template>
  <nav class="app-nav">
    <div class=":uno: text-4xl mr-10">
      LOGO
    </div>
    <ul>
      <li>
        <NuxtLink to="/">
          Link 1
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

    <AppNavUser />
  </nav>
</template>
