<script setup lang="ts">
const props = defineProps<{
  postedAt: string
}>()
const postedAt = new Date(props.postedAt).toLocaleDateString('pl')
</script>

<template>
  <section class="details-content">
    <span>Posted: {{ postedAt }}</span>
    <h2>Job Offer: Senior Software Engineer</h2>

    <p>Tech Company Inc. is seeking a highly skilled Senior Software Engineer to join our dynamic team. If you are passionate about cutting-edge technologies and solving complex problems, we want you on board!</p>

    <h3>Key Responsibilities:</h3>
    <ul>
      <li>Lead the development of innovative software solutions</li>
      <li>Collaborate with cross-functional teams to define, design, and ship new features</li>
      <li>Debug and optimize existing software</li>
      <li>Conduct code reviews and provide constructive feedback</li>
    </ul>

    <h3>Qualifications:</h3>
    <ul>
      <li>Bachelor's degree in Computer Science or related field</li>
      <li>Minimum of 5 years of experience in software development</li>
      <li>Proficient in multiple programming languages (e.g., Java, Python, C++)</li>
      <li>Strong problem-solving and analytical skills</li>
      <li>Excellent communication and teamwork abilities</li>
    </ul>

    <h3>Perks:</h3>
    <ul>
      <li>Competitive salary</li>
      <li>Flexible work hours</li>
      <li>Health and dental insurance</li>
      <li>401(k) retirement plan</li>
      <li>On-site gym and recreation area</li>
    </ul>

    <p>If you're ready to take on exciting challenges and be part of a forward-thinking tech company, apply now by sending your resume and cover letter to careers@techcompanyinc.com.</p>
  </section>
</template>
