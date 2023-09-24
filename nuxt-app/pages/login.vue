<script setup lang="ts">
useHead({
  title: 'Login'
})
</script>

<template>
  <main>
    <section class="login-form">
      <div class="wrapper">
        <p>Login</p>
        <form>
          <fieldset>
            <legend></legend>
            <input id="email" type="email" required placeholder="" />
            <label for="email">Email</label>
          </fieldset>

          <fieldset>
            <legend></legend>
            <input id="password" type="password" required placeholder="">
            <label for="password">Password</label>
          </fieldset>
          <button type="submit" @click.prevent>
            Login
          </button>
        </form>
      </div>
    </section>
  </main>
</template>
