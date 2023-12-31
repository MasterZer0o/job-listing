<script setup lang="ts">
export type TechCategory = 'frontend' | 'backend' | 'mobile'

const brandColor = ref<string>('')
onMounted(() =>
  brandColor.value = getComputedStyle(document.documentElement).getPropertyValue('--brand')
)

interface Skill {
  id: number
  name: string
  icon: ReturnType<typeof resolveComponent>
}
const router = useRouter()
const route = useRoute('jobs')

const store = useListingFilters()
const { techSkills } = storeToRefs(store)

const allSelectedToggle: Record<TechCategory, boolean> = {
  frontend: false,
  backend: false,
  mobile: false
}
function clearAll() {
  techSkills.value.clear()
  Object.keys(allSelectedToggle).forEach(e => allSelectedToggle[e as TechCategory] = false)
}
const skills: Record<TechCategory, Skill[]> = {
  frontend: [{
    id: 1,
    name: 'JavaScript',
    icon: resolveComponent('IconsJS')
  },
  {
    id: 2,
    name: 'TypeScript',
    icon: resolveComponent('IconsTS')
  },
  {
    id: 3,
    name: 'Vue',
    icon: resolveComponent('IconsVue')
  },
  {
    id: 4,
    name: 'Angular',
    icon: resolveComponent('IconsAngular')
  },
  {
    id: 5,
    name: 'React',
    icon: resolveComponent('IconsReact')
  },
  {
    id: 6,
    name: 'HTML/CSS',
    icon: resolveComponent('IconsHTML')
  }
  ],
  backend: [{
    id: 7,
    name: 'Node.js',
    icon: resolveComponent('IconsNode')
  },
  {
    id: 8,
    name: 'PHP',
    icon: resolveComponent('IconsPHP')
  },
  {
    id: 9,
    name: 'Go',
    icon: resolveComponent('IconsGO')
  },
  {
    id: 10,
    name: 'Python',
    icon: resolveComponent('IconsPython')
  },
  {
    id: 11,
    name: 'SQL',
    icon: resolveComponent('IconsSQL')
  },
  {
    id: 12,
    name: 'Rust',
    icon: resolveComponent('IconsRust')
  },
  {
    id: 13,
    name: 'C',
    icon: resolveComponent('IconsC')
  },
  {
    id: 14,
    name: 'C#',
    icon: resolveComponent('IconsCS')
  },
  {
    id: 15,
    name: 'C++',
    icon: resolveComponent('IconsCPP')
  },
  {
    id: 16,
    name: 'Java',
    icon: resolveComponent('IconsJava')
  },
  ],
  mobile: [{
    id: 17,
    name: 'Swift',
    icon: resolveComponent('IconsSwift')
  },
  {
    id: 18,
    name: 'Kotlin',
    icon: resolveComponent('IconsKotlin')
  },
  {
    id: 19,
    name: 'Flutter',
    icon: resolveComponent('IconsFlutter')
  }
  ]
}

const query = (route.query?.tech ?? '') as string
const ids = query.split(',').map(e => Number.parseInt(e))
const techSkillsFromUrl = [...skills.backend, ...skills.frontend, ...skills.mobile].filter((skill) => {
  return ids.includes(skill.id)
})

const selectedSkillsIds = ref(new Set<number>())
techSkillsFromUrl.forEach((skill) => {
  techSkills.value.set(skill.id, { name: skill.name })
  selectedSkillsIds.value.add(skill.id)
})
const jobsStore = useJobs()
function addSkillFilter(skill: Skill | Skill[], category?: TechCategory) {
  if (Array.isArray(skill)) {
    if (allSelectedToggle[category!])
      skill.forEach((e) => {
        techSkills.value.delete(e.id)
        selectedSkillsIds.value.delete(e.id)
      })
    else {
      skill.forEach((s) => {
        selectedSkillsIds.value.add(s.id)
        techSkills.value.set(s.id, { name: s.name })
      })
    }

    allSelectedToggle[category!] = !allSelectedToggle[category!]
    return
  }

  if (techSkills.value.has(skill.id)) {
    techSkills.value.delete(skill.id)
    selectedSkillsIds.value.delete(skill.id)
  }
  else {
    techSkills.value.set(skill.id, { name: skill.name })
    selectedSkillsIds.value.add(skill.id)
  }

  store.filtersApplied = false
}
watch(techSkills.value, async () => {
  jobsStore.currentPage = 1
  const ids: number[] = []
  techSkills.value.forEach((_, k) => ids.push(k))
  await router.push({
    query: {
      ...route.query,
      tech: ids.length !== 0 ? ids.join(',') : undefined
    }
  })
  useLoadingIndicator().finish()
})
</script>

<template>
  <div>
    <span>Tech skills</span>
    <button type="button" @click="clearAll">
      clear
    </button>
  </div>
  <section class="tech-skills">
    <div>
      <p title="Click to select all" @click="addSkillFilter(skills.frontend, 'frontend')">
        Frontend
      </p>
      <ul>
        <li
          v-for="frontendSkill in skills.frontend"
          :key="frontendSkill.name" class="tech-skill"
          :class="{ selected: techSkills.has(frontendSkill.id) }" @click="addSkillFilter(frontendSkill)">
          <component :is="frontendSkill.icon" />
          {{ frontendSkill.name }}
          <svg v-if="techSkills.has(frontendSkill.id)" width="20" height="20" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" class="check">
            <g fill="none" fill-rule="evenodd">
              <path d="M24 0v24H0V0h24ZM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018Zm.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022Zm-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01l-.184-.092Z" />
              <path :fill="brandColor" d="M21.546 5.111a1.5 1.5 0 0 1 0 2.121L10.303 18.475a1.6 1.6 0 0 1-2.263 0L2.454 12.89a1.5 1.5 0 1 1 2.121-2.121l4.596 4.596L19.424 5.111a1.5 1.5 0 0 1 2.122 0Z" />
            </g>
          </svg>
        </li>
      </ul>
      <p title="Click to select all" @click="addSkillFilter(skills.backend, 'backend')">
        Backend
      </p>
      <ul>
        <li
          v-for="backendSkill in skills.backend"
          :key="backendSkill.name" class="tech-skill"
          :class="{ selected: techSkills.has(backendSkill.id) }" @click="addSkillFilter(backendSkill)">
          <component :is="backendSkill.icon" />

          {{ backendSkill.name }}
          <svg v-if="techSkills.has(backendSkill.id)" width="20" height="20" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" class="check">
            <g fill="none" fill-rule="evenodd">
              <path d="M24 0v24H0V0h24ZM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018Zm.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022Zm-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01l-.184-.092Z" />
              <path :fill="brandColor" d="M21.546 5.111a1.5 1.5 0 0 1 0 2.121L10.303 18.475a1.6 1.6 0 0 1-2.263 0L2.454 12.89a1.5 1.5 0 1 1 2.121-2.121l4.596 4.596L19.424 5.111a1.5 1.5 0 0 1 2.122 0Z" />
            </g>
          </svg>
        </li>
      </ul>

      <p title="Click to select all" @click="addSkillFilter(skills.mobile, 'mobile')">
        Mobile
      </p>
      <ul>
        <li
          v-for="mobileSkill in skills.mobile"
          :key="mobileSkill.name" class="tech-skill"
          :class="{ selected: techSkills.has(mobileSkill.id) }" @click="addSkillFilter(mobileSkill)">
          <component :is="mobileSkill.icon" />

          {{ mobileSkill.name }}
          <svg v-if="techSkills.has(mobileSkill.id)" width="20" height="20" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" class="check">
            <g fill="none" fill-rule="evenodd">
              <path d="M24 0v24H0V0h24ZM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018Zm.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022Zm-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01l-.184-.092Z" />
              <path :fill="brandColor" d="M21.546 5.111a1.5 1.5 0 0 1 0 2.121L10.303 18.475a1.6 1.6 0 0 1-2.263 0L2.454 12.89a1.5 1.5 0 1 1 2.121-2.121l4.596 4.596L19.424 5.111a1.5 1.5 0 0 1 2.122 0Z" />
            </g>
          </svg>
        </li>
      </ul>
    </div>
  </section>
</template>
