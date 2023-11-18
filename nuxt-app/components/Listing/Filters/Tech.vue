<script setup lang="ts">
const { techSkills } = storeToRefs(useListingFilters())
type TechCategory = 'frontend' | 'backend' | 'mobile'

interface Skill {
  name: string
  icon: ReturnType<typeof resolveComponent>
}

const allSelectedToggle: Record<TechCategory, boolean> = {
  frontend: false,
  backend: false,
  mobile: false
}

function addSkillFilter(skill: string | Skill[], category?: TechCategory) {
  if (Array.isArray(skill)) {
    if (allSelectedToggle[category!])
      skill.forEach((e) => {
        techSkills.value.delete(e.name)
      })
    else
      skill.forEach(s => techSkills.value.add(s.name))

    allSelectedToggle[category!] = !allSelectedToggle[category!]
    return
  }

  if (techSkills.value.has(skill))
    techSkills.value.delete(skill)
  else
    techSkills.value.add(skill)
}

const skills: Record<TechCategory, Skill[]> = {
  frontend: [{
    name: 'JavaScript',
    icon: resolveComponent('IconsJS')
  },
  {
    name: 'TypeScript',
    icon: resolveComponent('IconsTS')
  },
  {
    name: 'Vue.js',
    icon: resolveComponent('IconsVue')
  },
  {
    name: 'Angular.js',
    icon: resolveComponent('IconsAngular')
  },
  {
    name: 'React.js',
    icon: resolveComponent('IconsReact')
  },
  {
    name: 'HTML/CSS',
    icon: resolveComponent('IconsHTML')
  }
  ],
  backend: [{
    name: 'Node.js',
    icon: resolveComponent('IconsNode')
  },
  {
    name: 'PHP',
    icon: resolveComponent('IconsPHP')
  },
  {
    name: 'Go',
    icon: resolveComponent('IconsGO')
  },
  {
    name: 'Python',
    icon: resolveComponent('IconsPython')
  },
  {
    name: 'SQL',
    icon: resolveComponent('IconsSQL')
  },
  {
    name: 'Rust',
    icon: resolveComponent('IconsRust')
  },
  {
    name: 'C',
    icon: resolveComponent('IconsC')
  },
  {
    name: 'C#',
    icon: resolveComponent('IconsCS')
  },
  {
    name: 'C++',
    icon: resolveComponent('IconsCPP')
  },
  {
    name: 'Java',
    icon: resolveComponent('IconsJava')
  },
  ],
  mobile: [{
    name: 'Swift',
    icon: resolveComponent('IconsSwift')
  },
  {
    name: 'Kotlin',
    icon: resolveComponent('IconsKotlin')
  },
  {
    name: 'Flutter',
    icon: resolveComponent('IconsFlutter')
  }
  ]
}
const brandColor = ref<string>()
if (process.client) {
  brandColor.value = getComputedStyle(document.documentElement).getPropertyValue('--brand')
}
</script>

<template>
  <span>Tech skills</span>
  <section class="tech-skills">
    <label aria-label="Search tech skill"><input type="text" placeholder="Skill name"></label>
    <div>
      <p @click="addSkillFilter(skills.frontend, 'frontend')">
        Frontend
      </p>
      <ul>
        <li
          v-for="frontendSkill in skills.frontend"
          :key="frontendSkill.name" class="tech-skill"
          :class="{ selected: techSkills.has(frontendSkill.name) }" @click="addSkillFilter(frontendSkill.name)">
          <component :is="frontendSkill.icon" />
          {{ frontendSkill.name }}
          <svg v-if="techSkills.has(frontendSkill.name)" width="20" height="20" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" class="check">
            <g fill="none" fill-rule="evenodd">
              <path d="M24 0v24H0V0h24ZM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018Zm.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022Zm-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01l-.184-.092Z" />
              <path :fill="brandColor" d="M21.546 5.111a1.5 1.5 0 0 1 0 2.121L10.303 18.475a1.6 1.6 0 0 1-2.263 0L2.454 12.89a1.5 1.5 0 1 1 2.121-2.121l4.596 4.596L19.424 5.111a1.5 1.5 0 0 1 2.122 0Z" />
            </g>
          </svg>
        </li>
      </ul>
      <p @click="addSkillFilter(skills.backend, 'backend')">
        Backend
      </p>
      <ul>
        <li
          v-for="backendSkill in skills.backend"
          :key="backendSkill.name" class="tech-skill"
          :class="{ selected: techSkills.has(backendSkill.name) }" @click="addSkillFilter(backendSkill.name)">
          <component :is="backendSkill.icon" />

          {{ backendSkill.name }}
          <svg v-if="techSkills.has(backendSkill.name)" width="20" height="20" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" class="check">
            <g fill="none" fill-rule="evenodd">
              <path d="M24 0v24H0V0h24ZM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018Zm.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022Zm-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01l-.184-.092Z" />
              <path :fill="brandColor" d="M21.546 5.111a1.5 1.5 0 0 1 0 2.121L10.303 18.475a1.6 1.6 0 0 1-2.263 0L2.454 12.89a1.5 1.5 0 1 1 2.121-2.121l4.596 4.596L19.424 5.111a1.5 1.5 0 0 1 2.122 0Z" />
            </g>
          </svg>
        </li>
      </ul>

      <p @click="addSkillFilter(skills.mobile, 'mobile')">
        Mobile
      </p>
      <ul>
        <li
          v-for="mobileSkill in skills.mobile"
          :key="mobileSkill.name" class="tech-skill"
          :class="{ selected: techSkills.has(mobileSkill.name) }" @click="addSkillFilter(mobileSkill.name)">
          <component :is="mobileSkill.icon" />

          {{ mobileSkill.name }}
          <svg v-if="techSkills.has(mobileSkill.name)" width="20" height="20" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" class="check">
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
