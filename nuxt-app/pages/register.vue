<script setup lang="ts">
import { email, minLength, object, safeParse, string } from 'valibot'

useHead({
  title: 'Register'
})

definePageMeta({
  pageTransition: {
    name: 'logreg',
    mode: 'out-in'
  }
})

const emailInput = ref() as Ref<HTMLInputElement>
const passwordInput = ref() as Ref<HTMLInputElement>
const passwordRepeatInput = ref() as Ref<HTMLInputElement>

const emailError = ref<string | null>(null)
const passwordError = ref<string | null>(null)
const passwordRepeatError = ref<string | null>(null)

const badInputError = ref<string | null>(null)
const passwordShown = ref(false)

const isLoading = ref(false)
const isSuccess = ref(false)
const rememberedValues = reactive({
  email: emailInput.value?.value,
  password: passwordInput.value?.value,
  passwordRepeat: passwordRepeatInput.value?.value
})

const schema = object({
  email: string([
    email(), minLength(1, 'This field is required.')]),
  password: string([
    minLength(6, 'Your password should contain at least 6 characters'),
    minLength(1, 'This field is required.')]),
  passwordRepeat: string([
    minLength(6, 'Your password should contain at least 6 characters'),
    minLength(1, 'This field is required.')])
})

async function register() {
  const emailValue = emailInput.value.value
  const passwordValue = passwordInput.value.value
  const passwordRepeatValue = passwordRepeatInput.value.value

  if (rememberedValues.email === emailValue
  && rememberedValues.password === passwordValue
  && rememberedValues.passwordRepeat === passwordRepeatValue) {
    emailInput.value.addEventListener('input', () => rememberedValues.email = emailValue, { once: true })
    passwordInput.value.addEventListener('input', () => rememberedValues.password = passwordValue, { once: true })
    passwordRepeatInput.value.addEventListener('input', () => rememberedValues.passwordRepeat = passwordRepeatValue, { once: true })
    return
  }

  const result = safeParse(schema, {
    email: emailValue,
    password: passwordValue,
    passwordRepeat: passwordRepeatValue
  })

  if (!result.success) {
    result.issues.forEach(((issue) => {
      if (issue.path![0].key === 'email') {
        emailError.value = issue.message

        emailInput.value.addEventListener('input', () => {
          emailError.value = null
        }, { once: true })
      }
      else if (issue.path![0].key === 'password') {
        passwordError.value = issue.message

        passwordInput.value.addEventListener('input', () => {
          passwordError.value = null
        }, { once: true })
      }

      passwordRepeatError.value = issue.message

      passwordRepeatInput.value.addEventListener('input', () => {
        passwordRepeatError.value = null
      }, { once: true })
    }))

    return
  }
  if (passwordRepeatValue !== passwordValue) {
    passwordRepeatError.value = 'Passwords do not match'
    passwordInput.value.addEventListener('input', () => {
      passwordRepeatError.value = null
    }, { once: true })

    passwordRepeatInput.value.addEventListener('input', () => {
      passwordRepeatError.value = null
    }, { once: true })

    return
  }

  isLoading.value = true
  rememberedValues.email = emailValue
  rememberedValues.password = passwordValue
  rememberedValues.passwordRepeat = passwordRepeatValue

  const apiUrl = useRuntimeConfig().public.API_BASE
  const response = await $fetch<{ error?: string }>(`${apiUrl}/user/register`, {
    method: 'POST',
    body: {
      email: emailValue,
      password: passwordValue,
      passwordRepeat: passwordRepeatValue
    }
  })

  if (response.error) {
    badInputError.value = response.error
    Array.from([emailInput.value, passwordInput.value])
      .forEach(field => field.addEventListener('input', () => badInputError.value = null, { once: true }))
    isLoading.value = false
    return
  }

  isLoading.value = false
  isSuccess.value = true
  setTimeout(() => {
    navigateTo('/')
  }, 2500)
}
</script>

<template>
  <main>
    <section class="form">
      <div class="wrapper">
        <p>Register</p>
        <form data-registerform novalidate @submit="register">
          <fieldset>
            <legend></legend>
            <input id="email" ref="emailInput" type="email" required placeholder="" />
            <label for="email">Email</label>
            <span v-if="emailError !== null" class="validation-error">{{ emailError }}</span>
          </fieldset>

          <fieldset>
            <legend></legend>
            <input id="password" ref="passwordInput" :type="passwordShown ? 'text' : 'password'" required placeholder="">
            <label for="password">Password</label>
            <button
              :aria-label="`Click to ${passwordShown ? 'hide' : 'reveal'} password`"
              class="input-eye"
              type="button" :title="`Click to ${passwordShown ? 'hide' : 'reveal'} password`"
              @click="passwordShown = !passwordShown">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <template v-if="!passwordShown">
                  <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                  <path d="M10.585 10.587a2 2 0 0 0 2.829 2.828"></path>
                  <path d="M16.681 16.673a8.717 8.717 0 0 1 -4.681 1.327c-3.6 0 -6.6 -2 -9 -6c1.272 -2.12 2.712 -3.678 4.32 -4.674m2.86 -1.146a9.055 9.055 0 0 1 1.82 -.18c3.6 0 6.6 2 9 6c-.666 1.11 -1.379 2.067 -2.138 2.87"></path>
                  <path d="M3 3l18 18"></path>
                </template>
                <template v-else>
                  <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                  <path d="M10 12a2 2 0 1 0 4 0a2 2 0 0 0 -4 0"></path>
                  <path d="M11.102 17.957c-3.204 -.307 -5.904 -2.294 -8.102 -5.957c2.4 -4 5.4 -6 9 -6c3.6 0 6.6 2 9 6a19.5 19.5 0 0 1 -.663 1.032"></path>
                  <path d="M15 19l2 2l4 -4"></path>
                </template>
              </svg>
            </button>
            <span v-if="passwordError !== null" class="validation-error">{{ passwordError }}</span>
          </fieldset>

          <fieldset>
            <legend></legend>
            <input id="password-repeat" ref="passwordRepeatInput" :type="passwordShown ? 'text' : 'password'" required placeholder="">
            <label for="password-repeat">Repeat Password</label>
            <button
              :aria-label="`Click to ${passwordShown ? 'hide' : 'reveal'} password`"
              class="input-eye"
              type="button" :title="`Click to ${passwordShown ? 'hide' : 'reveal'} password`"
              @click="passwordShown = !passwordShown">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <template v-if="!passwordShown">
                  <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                  <path d="M10.585 10.587a2 2 0 0 0 2.829 2.828"></path>
                  <path d="M16.681 16.673a8.717 8.717 0 0 1 -4.681 1.327c-3.6 0 -6.6 -2 -9 -6c1.272 -2.12 2.712 -3.678 4.32 -4.674m2.86 -1.146a9.055 9.055 0 0 1 1.82 -.18c3.6 0 6.6 2 9 6c-.666 1.11 -1.379 2.067 -2.138 2.87"></path>
                  <path d="M3 3l18 18"></path>
                </template>
                <template v-else>
                  <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                  <path d="M10 12a2 2 0 1 0 4 0a2 2 0 0 0 -4 0"></path>
                  <path d="M11.102 17.957c-3.204 -.307 -5.904 -2.294 -8.102 -5.957c2.4 -4 5.4 -6 9 -6c3.6 0 6.6 2 9 6a19.5 19.5 0 0 1 -.663 1.032"></path>
                  <path d="M15 19l2 2l4 -4"></path>
                </template>
              </svg>
            </button>
            <span v-if="passwordRepeatError !== null" class="validation-error">{{ passwordRepeatError }}</span>
          </fieldset>
          <div>
            <p v-if="badInputError" class="validation-error">
              <svg width="20" height="20" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg">
                <path fill="currentColor" d="M16 2C8.3 2 2 8.3 2 16s6.3 14 14 14s14-6.3 14-14S23.7 2 16 2zm-1.1 6h2.2v11h-2.2V8zM16 25c-.8 0-1.5-.7-1.5-1.5S15.2 22 16 22s1.5.7 1.5 1.5S16.8 25 16 25z" />
              </svg>
              {{ badInputError }}
            </p>
            <button type="submit" @click.prevent="register">
              <template v-if="!isLoading">
                <span v-if="!isSuccess">Register</span>
                <svg v-else class="checkmark" width="30" height="30" stroke="#fff" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 52 52">
                  <circle
                    stroke-dasharray="166" stroke="#619c34" stroke-dashoffset="166" stroke-width="2"
                    stroke-miterlimit="10" cx="26" cy="26" r="25" fill="none" />
                  <path class="checkmark__check" stroke-dasharray="48" stroke-dashoffset="48" fill="none" d="M14.1 27.2l7.1 7.2 16.7-16.8" /></svg>
              </template>

              <svg v-else class="loader" width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path fill="#fff" d="M12,23a9.63,9.63,0,0,1-8-9.5,9.51,9.51,0,0,1,6.79-9.1A1.66,1.66,0,0,0,12,2.81h0a1.67,1.67,0,0,0-1.94-1.64A11,11,0,0,0,12,23Z" /></svg>
            </button>
          </div>
          <NuxtLink to="/login" :no-prefetch="true">
            I have an account
          </NuxtLink>
        </form>
        <DevOnly>
          <button
            type="button"
            class="absolute"
            @click="function () {
              emailInput.value = `test${(Math.random() * 100).toFixed(2)}@test.com`
              passwordInput.value = `123123`
              passwordRepeatInput.value = `123123`
            }">
            insert
          </button>
        </DevOnly>
      </div>
    </section>
  </main>
</template>
