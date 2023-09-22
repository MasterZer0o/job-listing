import { defineConfig, presetWind } from 'unocss'
import transformerCompileClass from '@unocss/transformer-compile-class'

export default defineConfig({
  transformers: [
    transformerCompileClass({ classPrefix: 'u-', trigger: ':uno:', keepUnknown: true })
  ],
  presets: [
    presetWind()
  ]
})
