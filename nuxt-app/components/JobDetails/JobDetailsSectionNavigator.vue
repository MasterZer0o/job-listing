<script setup lang="ts">
defineProps<{
  elements: string[]
}>()
const allElements = ref([]) as Ref<HTMLElement[]>
const selectedIndex = ref<number>(1)
const indicatorWidth = ref<number>(0)
const indicatorDistance = ref<number>(0)
watch(allElements, () => selectedIndex.value = 0, { deep: true, once: true })

watch(selectedIndex, () => {
  const selectedEl = allElements.value[selectedIndex.value]
  indicatorWidth.value = selectedEl.clientWidth
  indicatorDistance.value = selectedEl.offsetLeft
})

const style = computed(() => {
  return `--indicator-width:${indicatorWidth.value - (indicatorWidth.value * 0.1)}px;--indicator-distance:${indicatorDistance.value + (indicatorWidth.value * 0.05)}px;`
})
</script>

<template>
  <section class="section-navigator">
    <ul :style="style">
      <li v-for="(option, i) in elements" ref="allElements" :key="option" @click="selectedIndex = i">
        {{ option }}
      </li>
    </ul>
  </section>
</template>
