import type { Config } from 'tailwindcss'

const config: Config = {
  theme: {
    container: {
      center: false,
      padding: '0',
      screens: {
        '2xl': '100%',
      },
    },
  },
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  plugins: [],
}

export default config
