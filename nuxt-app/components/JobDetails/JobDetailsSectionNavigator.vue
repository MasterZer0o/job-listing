<script setup lang="ts">
const { elements } = defineProps<{
  elements: JobDetailsResponse['data']['contentSections']
}>()

const textsAndLinks = elements.map((el) => {
  const [text, link] = el.split('#')
  return { text, link }
})
const allElements = ref([]) as Ref<HTMLElement[]>
const selectedIndex = ref<number>(1)
const indicatorWidth = ref<number>(0)
const indicatorPos = ref<number>(0)
watch(allElements, () => selectedIndex.value = 0, { deep: true, once: true })

watch(selectedIndex, () => {
  const selectedEl = allElements.value[selectedIndex.value]
  indicatorWidth.value = selectedEl.clientWidth
  indicatorPos.value = selectedEl.offsetLeft
})
const scrollContainer = ref() as Ref<HTMLElement>
function scrollWheel(evt: WheelEvent) {
  const cont = scrollContainer.value
  evt.preventDefault()

  if (evt.deltaY >= -15 && evt.deltaY <= 15)
    cont.scrollLeft += (evt.deltaY * 40)
  else
    cont.scrollLeft += (evt.deltaY * 5)
}
onMounted(() => scrollContainer.value.addEventListener('wheel', scrollWheel))
onBeforeUnmount(() => scrollContainer.value.removeEventListener('wheel', scrollWheel))

const indicatorCSS = computed(() => {
  return `--indicator-width:${indicatorWidth.value - (indicatorWidth.value * 0.1)}px;--indicator-posX:${indicatorPos.value + (indicatorWidth.value * 0.05)}px;`
})
</script>

<template>
  <section class="section-navigator">
    <ul ref="scrollContainer" :style="indicatorCSS">
      <li v-for="({ link, text }, i) in textsAndLinks" ref="allElements" :key="link" @click="selectedIndex = i">
        <a :href="`#${link}`">{{ text }}</a>
      </li>
    </ul>
  </section>
</template>
