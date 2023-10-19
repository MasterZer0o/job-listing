export const useUser = defineStore('user', () => {
  const loggedIn = ref<boolean>(false)

  function logout() {
    loggedIn.value = false
  }

  return {
    loggedIn,
    logout
  }
})
