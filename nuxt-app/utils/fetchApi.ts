export function fetchApi<T>(path: Parameters<typeof $fetch>[0], opts?: Parameters<typeof $fetch>[1]) {
  return $fetch<T>(`${useRuntimeConfig().public.API_BASE}${path}`, {
    credentials: 'include',
    ...opts,
    headers: {
      cookie: useCookie('session').value ? `session=${useCookie('session').value}` : undefined!
    },
  })
}
