<script setup lang="ts">
import { custom, email, minLength, object, safeParse, string } from 'valibot'

useHead({
  title: 'Register'
})

const emailInput = ref() as Ref<HTMLInputElement>
const passwordInput = ref() as Ref<HTMLInputElement>
const passwordRepeatInput = ref() as Ref<HTMLInputElement>

const emailError = ref<string | null>(null)
const passwordError = ref<string | null>(null)
const passwordRepeatError = ref<string | null>(null)
const registerError = ref(false)

const isLoading = ref(false)
const rememberedValues = reactive({
  email: emailInput.value?.value,
  password: passwordInput.value?.value
})

const schema = object({
  email: string([

    email(), minLength(1, 'This field is required.')]),
  password: string([
    minLength(6, 'Your password has minimum 6 characters'),
    minLength(1, 'This field is required.')]),
  passwordRepeat: string([
    minLength(6, 'Your password has minimum 6 characters'),
    minLength(1, 'This field is required.')])
}, [
  custom(({ password, passwordRepeat }) => password === passwordRepeat, 'Passwords do not match')
])

async function register() {
  const emailValue = emailInput.value.value
  const passwordValue = passwordInput.value.value
  const passwordRepeatValue = passwordRepeatInput.value.value

  // if (rememberedValues.email === emailValue && rememberedValues.password === passwordValue) {
  //   emailInput.value.addEventListener('input', () => rememberedValues.email = emailValue, { once: true })
  //   passwordInput.value.addEventListener('input', () => rememberedValues.password = passwordValue, { once: true })
  //   return
  // }

  const result = safeParse(schema, {
    email: emailValue,
    password: passwordValue,
    passwordRepeat: passwordRepeatValue
  })

  if (!result.success) {
    logInfo(result.issues)
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
      else {
        passwordRepeatError.value = issue.message

        passwordRepeatInput.value.addEventListener('input', () => {
          passwordRepeatError.value = null
        }, { once: true })
      }
    }))

    return
  }
  isLoading.value = true
  rememberedValues.email = emailValue
  rememberedValues.password = passwordValue

  await new Promise(resolve => setTimeout(resolve, 1000))

  registerError.value = true

  Array.from([emailInput.value, passwordInput.value])
    .forEach(field => field.addEventListener('input', () => registerError.value = false, { once: true }))

  isLoading.value = false
}
</script>

<template>
  <main>
    <section class="login-form">
      <div class="wrapper">
        <p>Register</p>
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

          <fieldset>
            <legend></legend>
            <input id="password-repeat" ref="passwordRepeatInput" type="password" required placeholder="">
            <label for="password-repeat">Repeat Password</label>
            <span v-if="passwordRepeatError !== null" class="validation-error">{{ passwordRepeatError }}</span>
          </fieldset>
          <div>
            <p v-if="registerError" class="validation-error">
              Email or password is incorrect.
            </p>
            <button type="submit" @click.prevent="register">
              <template v-if="!isLoading">
                Register
              </template>

              <svg v-else class="loader" width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path fill="#fff" d="M12,23a9.63,9.63,0,0,1-8-9.5,9.51,9.51,0,0,1,6.79-9.1A1.66,1.66,0,0,0,12,2.81h0a1.67,1.67,0,0,0-1.94-1.64A11,11,0,0,0,12,23Z" /></svg>
            </button>
          </div>
          <NuxtLink to="/login" :no-prefetch="true">
            I have an account
          </NuxtLink>
        </form>
      </div>
    </section>
  </main>
</template>
