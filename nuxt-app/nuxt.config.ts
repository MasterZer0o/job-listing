import openProps from 'open-props'

export default defineNuxtConfig({
  modules: ['@nuxtjs/html-validator', '@unocss/nuxt',
  ['@pinia/nuxt', { autoImports: ['defineStore', 'storeToRefs', 'skipHydrate'] }]
  ],
  imports: {
    dirs: ['stores']
  },
  experimental: {
    inlineSSRStyles: false,
    typedPages: true
  },
  css: ['@/assets/css/main.scss'],
  postcss: {
    plugins: {
      'postcss-jit-props': openProps,
      'autoprefixer': {}
    }
  },
  vite: {
    build: {
      rollupOptions: {
        output: {
          assetFileNames: 'public/assets/[hash][extname]',
          chunkFileNames: 'public/[hash].js',
          entryFileNames: 'public/[hash].js'
        }
      }
    }
  },
    
  app: {
    rootId: 'app',
    buildAssetsDir: '/public/',

    head: {
      htmlAttrs: {
        lang: 'en'
      },
      link: [
        // * font *
        // {
        //   rel: 'stylesheet',
        //   href: 'https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css'
        // },
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
      ]
    }
  },
  runtimeConfig: {
    public: {}
  }
})
