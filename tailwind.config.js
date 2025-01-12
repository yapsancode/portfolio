/** @type {import('tailwindcss').Config} */
import colors, { gray } from 'tailwindcss/colors';
export const mode = 'jit';
export const content = [
  './templates/**/*.html', // Include all your HTML files in the templates folder
  './pages/**/*.html', // Include your page-specific HTML files
  './partials/**/*.html', // Include any partial templates
  './static/js/**/*.js', // Include your JavaScript files
];
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
        100: '#f2f3f5',
      },
    },
    animation: {
      'fade-in-left': 'fadeInLeft 1.5s ease-out',
    },
    keyframes: {
      fadeInLeft: {
        '0%': { opacity: 0, transform: 'translateX(-50px)' },
        '100%': { opacity: 1, transform: 'translateX(0)' },
      },
    },
  },
};
export const variants = {
  extend: {},
};
export const plugins = [];
