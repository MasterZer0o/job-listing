<script setup lang="ts">
const props = defineProps<{
  data: RelatedOffer
}>()
const emit = defineEmits(['saveJob'])
const salaryFrom = props.data?.salary?.from[0]
const salaryTo = props.data?.salary?.to[0]
const isExpanded = ref(false)
</script>

<template>
  <li>
    <div class="top" @click="isExpanded = !isExpanded">
      <div>
        <div>
          <p :title="data.title">
            {{ data.title }}
          </p>
          <div>
            <span class="comp-name" :title="data.companyName">{{ data.companyName }}</span>
            <span v-if="data.remoteAvailable">
              Remote

              <svg width="13" height="13" viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg">
                <path fill="#228be6" d="M504 256c0 136.967-111.033 248-248 248S8 392.967 8 256S119.033 8 256 8s248 111.033 248 248M227.314 387.314l184-184c6.248-6.248 6.248-16.379 0-22.627l-22.627-22.627c-6.248-6.249-16.379-6.249-22.628 0L216 308.118l-70.059-70.059c-6.248-6.248-16.379-6.248-22.628 0l-22.627 22.627c-6.248 6.248-6.248 16.379 0 22.627l104 104c6.249 6.249 16.379 6.249 22.628.001" />
              </svg>
            </span>
          </div>
        </div>

        <div>
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="-0.895 -0.895 25 25" height="20" width="20">
            <g>
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.605 22.38107142857143c5.951541357142857 0 10.776071428571429 -4.824530071428572 10.776071428571429 -10.776071428571429C22.38107142857143 5.653541535714286 17.55654135714286 0.8289285714285715 11.605 0.8289285714285715 5.653541535714286 0.8289285714285715 0.8289285714285715 5.653541535714286 0.8289285714285715 11.605c0 5.951541357142857 4.8246129642857145 10.776071428571429 10.776071428571429 10.776071428571429Z" stroke-width="1.79"></path>
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.42642417857143 8.657959985714287c-0.09703437857142858 -0.2745411428571429 -0.24677203571428574 -0.5242144285714285 -0.43706087857142856 -0.7368180285714286 -0.4047658214285714 -0.4522302714285715 -0.9929569571428573 -0.7368180285714286 -1.6476281642857142 -0.7368180285714286H10.63076025c-1.0897426571428572 0 -1.9731484142857143 0.8834057571428572 -1.9731484142857143 1.9731318357142857 0 0.9272560785714286 0.6456524642857143 1.7294102571428573 1.5514890285714287 1.9275739214285714l2.605007507142857 0.5698386571428572c1.014790935714286 0.2219870714285714 1.7380974285714286 1.1212751 1.7380974285714286 2.1600718071428573 0 1.220812842857143 -0.9896578214285714 2.2112830142857143 -2.2104706642857144 2.2112830142857143H10.868082500000002c-0.9624523857142857 0 -1.7812348714285713 -0.6151147357142858 -2.084689042857143 -1.4736526357142856" stroke-width="1.79"></path>
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.605 7.184042092857143V4.973571428571429" stroke-width="1.79"></path>
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.605 18.236097V16.02567607142857" stroke-width="1.79"></path>
            </g>
          </svg>
          <span>{{ salaryFrom }}K - {{ salaryTo }}K</span>
        </div>
      </div>
      <ul v-show="isExpanded">
        <li v-for="skill in data.skills" :key="skill">
          <span>{{ skill }}</span>
        </li>
      </ul>
    </div>
    <div class="related-bottom">
      <ul>
        <li @click="emit('saveJob')">
          <button type="button">
            <svg width="18" height="18" xmlns="http://www.w3.org/2000/svg" :fill="data.isSaved ? '#ebcb8b' : 'none'" viewBox="0 0 21 20">
              <path :stroke="data.isSaved ? '#ebcb8b' : 'currentColor'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.479 1.712 2.367 4.8a.532.532 0 0 0 .4.292l5.294.769a.534.534 0 0 1 .3.91l-3.83 3.735a.534.534 0 0 0-.154.473l.9 5.272a.535.535 0 0 1-.775.563l-4.734-2.49a.536.536 0 0 0-.5 0l-4.73 2.487a.534.534 0 0 1-.775-.563l.9-5.272a.534.534 0 0 0-.154-.473L2.158 8.48a.534.534 0 0 1 .3-.911l5.294-.77a.532.532 0 0 0 .4-.292l2.367-4.8a.534.534 0 0 1 .96.004Z" />
            </svg>

            <template v-if="data.isSaved">
              <strong>Saved</strong>
            </template>
            <template v-else>
              Save
            </template>
          </button>
        </li>
        <li>
          <a :href="`/job/${data.id}`" target="_blank">
            Open
            <svg width="16" height="16" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path fill="currentColor" d="M19 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h6v2H5v12h12v-6zM13 3v2h4.586l-7.793 7.793l1.414 1.414L19 6.414V11h2V3z" />
            </svg>
          </a>
        </li>
      </ul>
    </div>
  </li>
</template>
