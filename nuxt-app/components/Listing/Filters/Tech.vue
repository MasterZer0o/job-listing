<script setup lang="ts">
const { techSkills } = storeToRefs(useListingFilters())

function addSkillFilter(skill: number) {
  if (techSkills.value.includes(skill))
    techSkills.value.splice(techSkills.value.indexOf(skill), 1)
  else
    techSkills.value.push(skill)
}
</script>

<template>
  <span>Tech skills</span>
  <section class="tech-skills">
    <label aria-label="Search tech skill"><input type="text" placeholder="Skill name"></label>
    <ul>
      <li
        v-for="skill in 9" :key="skill"
        :class="{ selected: techSkills.includes(skill) }" @click="addSkillFilter(skill)">
        skill{{ skill }}
      </li>
    </ul>
  </section>
</template>
