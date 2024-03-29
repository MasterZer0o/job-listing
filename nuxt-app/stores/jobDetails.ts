import type { RelatedOffer } from '#imports'

export const jobDetailsStore = defineStore('jobDetailsStore', () => {
  const results = ref<RelatedOffer[]>([])
  const isLoading = ref(true)

  return {
    results, isLoading
  }
})
