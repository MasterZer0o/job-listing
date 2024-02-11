<script setup lang="ts">
const theme = ref('dark')
const load = ref(false)
if (process.client) {
  const currentTheme = localStorage.getItem('theme') as string
  theme.value = currentTheme
  load.value = true
}

const checkbox = ref() as Ref<HTMLInputElement>
function changeTheme(ev: MouseEvent) {
  ev.preventDefault()
  checkbox.value.checked = !checkbox.value.checked
  document.body.setAttribute('class', 'theming')
  setTimeout(() => document.body.removeAttribute('class'), 350)

  theme.value = theme.value === 'dark' ? 'light' : 'dark'
  localStorage.setItem('theme', theme.value)
}
</script>

<template>
  <Html :color-scheme="theme" />

  <div class="theme-switcher" @click="changeTheme">
    <ClientOnly>
      <input id="theme-checkbox" ref="checkbox" type="checkbox" :checked="theme === 'light'" />

      <label for="theme-checkbox" aria-label="label for checkbox">
        <svg xmlns="http://www.w3.org/2000/svg" height="15" width="15" fill="currentColor" viewBox="0 0 512 512">
          <path d="M 256 0 Q 234 2 232 24 L 232 88 Q 234 110 256 112 Q 278 110 280 88 L 280 24 Q 278 2 256 0 L 256 0 Z M 256 400 Q 234 402 232 424 L 232 488 Q 234 510 256 512 Q 278 510 280 488 L 280 424 Q 278 402 256 400 L 256 400 Z M 488 280 Q 510 278 512 256 Q 510 234 488 232 L 424 232 Q 402 234 400 256 Q 402 278 424 280 L 488 280 L 488 280 Z M 112 256 Q 110 234 88 232 L 24 232 Q 2 234 0 256 Q 2 278 24 280 L 88 280 Q 110 278 112 256 L 112 256 Z M 437 109 Q 451 92 437 75 Q 420 61 403 75 L 358 120 Q 344 137 358 154 Q 375 168 392 154 L 437 109 L 437 109 Z M 154 358 Q 137 344 120 358 L 75 403 Q 61 420 75 437 Q 92 451 109 437 L 154 392 Q 168 375 154 358 L 154 358 Z M 403 437 Q 420 451 437 437 Q 451 420 437 403 L 392 358 Q 375 344 358 358 Q 344 375 358 392 L 403 437 L 403 437 Z M 154 154 Q 168 137 154 120 L 109 75 Q 92 61 75 75 Q 61 92 75 109 L 120 154 Q 137 168 154 154 L 154 154 Z M 256 368 Q 286 368 312 353 L 312 353 Q 338 338 353 312 Q 368 286 368 256 Q 368 226 353 200 Q 338 174 312 159 Q 286 144 256 144 Q 226 144 200 159 Q 174 174 159 200 Q 144 226 144 256 Q 144 286 159 312 Q 174 338 200 353 Q 226 368 256 368 L 256 368 Z" />
        </svg>
        <svg xmlns="http://www.w3.org/2000/svg" height="15" width="15" fill="currentColor" viewBox="0 0 387 448">
          <path
            d="M 224 0 Q 161 1 111 31 L 111 31 Q 60 60 30 111 L 30 111 Q 1 161 0 224 Q 1 287 31 337 Q 60 388 111 417 Q 161 447 224 448 Q 317 446 379 385 Q 387 376 382 366 Q 377 356 365 357 Q 351 360 335 360 Q 261 358 211 308 Q 162 259 160 184 Q 160 134 184 94 Q 209 54 249 31 Q 259 25 257 13 Q 254 3 242 1 Q 233 0 223 0 L 224 0 Z"
          />
        </svg>
      </label>
    </ClientOnly>
  </div>
</template>
