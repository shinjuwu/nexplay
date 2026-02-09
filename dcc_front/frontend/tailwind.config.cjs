/* eslint-env node */
module.exports = {
  content: ['./src/**/*.{html,js,vue}'],
  theme: {
    screens: {
      sm: '576px',
      md: '768px',
      lg: '992px',
      xl: '1200px',
      '2xl': '1400px',
    },
    extend: {
      backgroundImage: {
        login: 'url(@/base/assets/images/bg-login.svg)',
      },
    },
  },
  plugins: [],
}
