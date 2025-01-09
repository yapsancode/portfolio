/** @type {import('tailwindcss').Config} */
import colors, { gray } from 'tailwindcss/colors'
export const mode = 'jit'
export const purge = [
  './src/**/*.{js,jsx,ts,tsx}',
  './public/index.html',
  './src/**/*.{js,jsx,ts,tsx}',
]
export const darkMode = false
export const content = []
export const theme = {
  extend: {
    colors: {
      primary: '#202225',
      secondary: '#5865F2',
      gray: colors.gray,
      gray: {
        900: '#202225',
        800: '#2f3136',
        700: '#36393f',
        600: '#4f545c',
        500: '#72767d',
        400: '#a3a6aa',
        300: '#dcddde',
        200: '#ebedef',
        100: '#f2f3f5'
      }
    },
  },
}
export const variants = {
  extend: {},
}
export const plugins = []

