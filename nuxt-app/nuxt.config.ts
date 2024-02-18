export default defineNuxtConfig({
  modules: ['@nuxtjs/html-validator', '@unocss/nuxt',
    ['@pinia/nuxt', { autoImports: ['defineStore', 'storeToRefs', 'skipHydrate'] }]
  ],
  imports: {
    dirs: ['stores']
  },
  nitro: {
    preset: 'vercel_edge'
  },
  experimental: {
    typedPages: true
  },
  features: {
    inlineStyles: false,
  },

  css: ['@/assets/css/main.scss'],
  postcss: {
    plugins: {
      autoprefixer: {}
    }
  },

  vite: {
    css: {

    },
    build: {
      cssCodeSplit: false,

      rollupOptions: {
        output: {
          assetFileNames: 'public/assets/[hash][extname]',
          chunkFileNames: 'public/[hash].js',
          entryFileNames: 'public/[hash].js'
        }
      }
    }
  },
  devtools: { enabled: false },
  app: {
    rootId: 'app',
    buildAssetsDir: '/public/',

    head: {
      htmlAttrs: {
        lang: 'en'
      },

      link: [

        // {
        //   rel: 'stylesheet',
        //   href: 'https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css'
        // },
        // * font *
        {
          rel: 'preconnect',
          href: 'https://fonts.googleapis.com'
        },
        {
          rel: 'preconnect',
          href: 'https://fonts.gstatic.com',
          crossorigin: ''
        },
        {
          rel: 'stylesheet',
          href: 'https://fonts.googleapis.com/css2?family=Inter:wght@400;700&display=swap'
        },

        // * font *
        {
          rel: 'icon',
          href: '/favicon.svg'
        }
      ],
      script: [{
        innerHTML: '(()=>{let e=localStorage.getItem("theme");if(["dark","light"].includes(e)){document.documentElement.setAttribute("color-scheme", e);return}let t=window.matchMedia&&window.matchMedia("(prefers-color-scheme: dark)").matches?"dark":"light";localStorage.setItem("theme",t),document.documentElement.setAttribute("color-scheme", t)})();',
        tagPosition: 'bodyOpen'
      }],
    }
  },
  runtimeConfig: {
    public: {
      API_BASE: process.env.NODE_ENV === 'development' ? 'http://localhost:5000' : process.env.NUXT_PUBLIC_API_BASE,
      lastUpdate: `${new Date(Date.now() + 60 * 60 * 1000 * Number(process.env.NODE_ENV !== 'development')).toLocaleDateString('pl')} ${new Date(Date.now() + 60 * 60 * 1000 * Number(process.env.NODE_ENV !== 'development')).toLocaleTimeString('pl', { timeStyle: 'short' })}`
    }

  }
})
