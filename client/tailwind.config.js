/* eslint-disable @typescript-eslint/no-var-requires */
const { fontFamily } = require('tailwindcss/defaultTheme');

/** @type {import("@types/tailwindcss/tailwind-config").TailwindConfig } */
module.exports = {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        primary: ['Roboto', ...fontFamily.sans],
      },
      colors: {
        primary: {
          // Customize it on globals.css :root
          50: '#BAA6DA',
          75: '#BAA6DA',
          100: '#9D82CA',
          200: '#734CB3',
          300: '#5627A4',
          400: '#3C1B73',
          500: '#341864',
        },
        dark: '#222222',
        secondary: {
          50: '#F2F7FA',
          75: '#CBDFE9',
          100: '#B5D2E0',
          200: '#96BFD2',
          300: '#80B2C9',
          400: '#5A7D8D',
          500: '#4E6D7B',
        },
        success: {
          50: '#ECFBED',
          75: '#B2EFB4',
          100: '#92E895',
          200: '#63DE68',
          300: '#43D749',
          400: '#2F9733',
          500: '#29832D',
        },
        warning: {
          50: '#FDF6ED',
          75: '#F5DBB4',
          100: '#F1CC94',
          200: '#EBB666',
          300: '#E7A747',
          400: '#A27532',
          500: '#8D662B',
        },
        danger: {
          50: '#FCE9E9',
          75: '#F4A3A3',
          100: '#EF7D7D',
          200: '#E94545',
          300: '#E41F1F',
          400: '#A01616',
          500: '#8B1313',
        },
        neutrals: {
          20: '#F5F6F7',
          30: '#EBEDF0',
          40: '#DFE2E6',
          50: '#C2C7D0',
          60: '#B3B9C4',
          70: '#A6AEBB',
          80: '#98A1B0',
          90: '#8993A4',
          100: '#7A8699',
          200: '#6B788E',
          300: '#5D6B82',
          400: '#505F79',
          500: '#42526D',
          600: '#354764',
          700: '#243757',
          800: '#15294B',
          900: '#091E42',
        },
        white: '#FAFBFB',
      },
      keyframes: {
        flicker: {
          '0%, 19.999%, 22%, 62.999%, 64%, 64.999%, 70%, 100%': {
            opacity: 0.99,
            filter:
              'drop-shadow(0 0 1px rgba(252, 211, 77)) drop-shadow(0 0 15px rgba(245, 158, 11)) drop-shadow(0 0 1px rgba(252, 211, 77))',
          },
          '20%, 21.999%, 63%, 63.999%, 65%, 69.999%': {
            opacity: 0.4,
            filter: 'none',
          },
        },
        shimmer: {
          '0%': {
            backgroundPosition: '-700px 0',
          },
          '100%': {
            backgroundPosition: '700px 0',
          },
        },
      },
      dropShadow: {
        lg: '0px 4px 4px rgba(0, 0, 0, 0.25)',
      },
      animation: {
        flicker: 'flicker 3s linear infinite',
        shimmer: 'shimmer 1.3s linear infinite',
      },
    },
  },
  plugins: [require('@tailwindcss/forms')],
};
