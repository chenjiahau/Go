/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./*.html",
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        darkPrimary: '#054673',
        primary: '#5bbcff',
        lightPrimary: '#cbe0ff',
        secondary: '#fffab7',
        third: '#ffd1e3',
        fourth: '#7ea1ff',
        success: '#28a745',
        info: '#17a2b8',
        warning: '#ffc107',
        danger: '#dc3545',
        white: '#ffffff',
        darkGray: '#c0c0c0',
        gray: '#f8f9fa',
        lightGray: '#f0f0f0',
        black: '#000000',
        active: '#0AF8F4',
        hover: '#0AF8F4',
        modal: '#FFFFFF',
        reverseModal: '#5bbcff',
        hoverModal: '#0AF8F4',
      },
      boxShadow: {
        darkPrimary: '4px 4px 2px rgba(5, 70, 155, 0.4)',
        lightPrimary: '4px 4px 2px rgba(203, 224, 255, 0.4)',
        save: '4px 4px 2px rgba(91, 188, 255, 0.4)',
        cancel: '4px 4px 2px rgba(220, 53, 69, 0.4)',
      },
      fontSize: {
        'small': '0.9rem',
        'general': '1rem',
        'title': '1.5rem',
        'button': '1.1rem',
        'large': '2rem',
      },
      fontFamily: {
        sans: ['Poppins', 'sans-serif'],
      },
    },
    plugins: [],
  }
}