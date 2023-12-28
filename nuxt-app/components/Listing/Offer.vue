<script setup lang="ts">
defineProps<{
  offer: Offer
}>()
const router = useRouter()

async function linkToLocation(e: MouseEvent) {
  e.preventDefault()

  await router.push({
    query: {
      location: 'remote'
    }
  })
}
async function linkToCompany(e: MouseEvent) {
  e.preventDefault()

  await router.push({
    query: {
      company: 'comp'
    }
  })
}
</script>

<template>
  <article>
    <NuxtLink :to="`/job/${offer.id}`" :prefetch="false">
      <section class="company-image">
        <img src="https://placehold.co/400x400" alt="">
      </section>
      <div class="offer-info">
        <div class="row">
          <section class="title">
            {{ offer.title }}
          </section>
          <section class="rmt-lvl">
            <div>
              <svg xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="20" x2="12" y2="10"></line><line x1="18" y1="20" x2="18" y2="4"></line><line x1="6" y1="20" x2="6" y2="16"></line></svg>
              {{ offer.level }}
            </div>
            <div @click="linkToLocation">
              <svg width="15" height="15" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2">
                  <circle cx="12" cy="10" r="3" />
                  <path d="M12 2a8 8 0 0 0-8 8c0 1.892.402 3.13 1.5 4.5L12 22l6.5-7.5c1.098-1.37 1.5-2.608 1.5-4.5a8 8 0 0 0-8-8Z" />
                </g>
              </svg>
              {{ offer.location }}
            </div>
          </section>
          <section class="pay">
            <strong>{{ offer.salaryFrom }} - {{ offer.salaryTo }}{{ offer.currency }}</strong>
          </section>
        </div>
        <div class="row">
          <section class="company" @click="linkToCompany">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 14 14"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5" /><path d="M3.5 9h7L7 3.5L3.5 9z" /></g></svg>
            {{ offer.company.name }}
          </section>
          <section class="skills">
            <ul>
              <li v-for="skill in offer.skills" :key="skill">
                <div>{{ skill }}</div>
              </li>
            </ul>
          </section>
        </div>
      </div>
    </NuxtLink>
  </article>
</template>
