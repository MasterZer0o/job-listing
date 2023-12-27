<script setup lang="ts">
defineProps<{
  isFetching: boolean
  offers: Offer[]
}>()
const store = useJobs()
</script>

<template>
  <section class="offers">
    <template v-if="!isFetching">
      <template v-if="!store.error.value">
        <transition name="ov-fade">
          <div v-if="store.isLoading" class="overlay">
            <svg class="loader" width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path fill="#fff" d="M12,23a9.63,9.63,0,0,1-8-9.5,9.51,9.51,0,0,1,6.79-9.1A1.66,1.66,0,0,0,12,2.81h0a1.67,1.67,0,0,0-1.94-1.64A11,11,0,0,0,12,23Z" /></svg>
          </div>
        </transition>
        <ListingOffer v-for="offer in offers" :key="offer.id" :offer="offer" />
        <p v-if="offers.length === 0 && !store.isLoading" class=":uno: text-center text-3xl mt-10">
          Sorry, no results.
        </p>
      </template>
      <template v-else>
        <p class=":uno: text-[#BF616A] text-center text-3xl mt-10">
          Failed to load job offers :(
        </p>
        <p v-if="store.error.net" class=":uno: text-center">
          {{ store.error.net?.restored ? "Your connection has been restored." : "It seems you have no internet connection." }}  You are <span :style="{ color: store.error.net?.restored ? '#A3BE8C' : '#BF616A' }">{{ store.error.net.restored ? 'online' : 'offline' }}.</span>
        </p>
        <p v-if="store.error.message" class=":uno: text-center m-auto">
          {{ store.error.message }}
        </p>
        <button :style="store.error.net && !store.error.net.restored && 'opacity:0.5;pointer-events:none' || undefined" class=":uno: w-fit px-4 py-2 rounded bg-[var(--surface-2)] self-center hover:bg-[var(--surface-3)] shadow  shadow-dark-400" @click="fetchJobs({ withCount: true })">
          Try again
        </button>
      </template>
    </template>
    <template v-else>
      <ListingSkeleton v-for="i in 20" :key="i" />
    </template>
  </section>
</template>
