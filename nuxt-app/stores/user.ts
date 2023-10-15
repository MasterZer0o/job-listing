export const useUser = defineStore('user', () => {
  const loggedIn = ref<boolean>(false)

  return {
    loggedIn
  }
})
