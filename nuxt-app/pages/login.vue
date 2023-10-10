<script setup lang="ts">
import { email, minLength, object, safeParse, string } from 'valibot'

useHead({
  title: 'Login'
})
const emailInput = ref() as Ref<HTMLInputElement>
const passwordInput = ref() as Ref<HTMLInputElement>

const emailError = ref<string | null>(null)
const passwordError = ref<string | null>(null)
const credentialsError = ref(false)

const isLoading = ref(false)
const rememberedValues = reactive({
  email: emailInput.value?.value,
  password: passwordInput.value?.value
})

const schema = object({
  email: string([email()]),
  password: string([
    minLength(6, 'Your password has minimum 6 characters'),
    minLength(1, 'This field is required.')])
})

async function login() {
  const emailValue = emailInput.value.value
  const passwordValue = passwordInput.value.value

  if (rememberedValues.email === emailValue && rememberedValues.password === passwordValue) {
    emailInput.value.addEventListener('input', () => rememberedValues.email = emailValue, { once: true })
    passwordInput.value.addEventListener('input', () => rememberedValues.password = passwordValue, { once: true })
    return
  }

  const result = safeParse(schema, {
    email: emailValue,
    password: passwordValue
  })

  if (!result.success) {
    result.issues.forEach(((issue) => {
      if (issue.path![0].key === 'email') {
        emailError.value = issue.message

        emailInput.value.addEventListener('input', () => {
          emailError.value = null
        }, { once: true })
      }
      else {
        passwordError.value = issue.message

        passwordInput.value.addEventListener('input', () => {
          passwordError.value = null
        }, { once: true })
      }
    }))

    return
  }
  isLoading.value = true
  rememberedValues.email = emailValue
  rememberedValues.password = passwordValue

  await new Promise(resolve => setTimeout(resolve, 1000))

  credentialsError.value = true

  Array.from([emailInput.value, passwordInput.value])
    .forEach(field => field.addEventListener('input', () => credentialsError.value = false, { once: true }))

  isLoading.value = false
}
</script>

<template>
  <main>
    <section class="form">
      <div class="wrapper">
        <p>Login</p>
        <form novalidate>
          <fieldset>
            <legend></legend>
            <input id="email" ref="emailInput" type="email" required placeholder="" />
            <label for="email">Email</label>
            <span v-if="emailError !== null" class="validation-error">{{ emailError }}</span>
          </fieldset>

          <fieldset>
            <legend></legend>
            <input id="password" ref="passwordInput" type="password" required placeholder="">
            <label for="password">Password</label>
            <span v-if="passwordError !== null" class="validation-error">{{ passwordError }}</span>
          </fieldset>

          <div class="row">
            <fieldset>
              <legend></legend>
              <label>
                <input type="checkbox" />
                <span>
                  <svg
                    fill="none" width="12px" height="11px" viewBox="0 0 12 11" stroke-linecap="rounded" stroke-linejoin="rounded" stroke-width="2" stroke-dasharray="17"
                    stroke-dashoffset="17" stroke="#fff">
                    <polyline points="1 6.29411765 4.5 10 11 1"></polyline>
                  </svg>
                </span>
                <span>Remember me</span>
              </label>
            </fieldset>
            <NuxtLink to="/">
              Forgot password?
            </NuxtLink>
          </div>
          <div>
            <p v-if="credentialsError" class="validation-error">
              Email or password is incorrect.
            </p>
            <button type="submit" @click.prevent="login">
              <template v-if="!isLoading">
                Login
              </template>

              <svg v-else class="loader" width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path fill="#fff" d="M12,23a9.63,9.63,0,0,1-8-9.5,9.51,9.51,0,0,1,6.79-9.1A1.66,1.66,0,0,0,12,2.81h0a1.67,1.67,0,0,0-1.94-1.64A11,11,0,0,0,12,23Z" /></svg>
            </button>
          </div>
          <NuxtLink to="/register" :no-prefetch="true">
            I don't have an account
          </NuxtLink>
        </form>
      </div>
    </section>
  </main>
</template>
