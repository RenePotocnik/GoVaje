import Vue from 'vue'
import { BootstrapVue, IconsPlugin, ModalPlugin } from 'bootstrap-vue'
import VueCompositionAPI from '@vue/composition-api'
import dayjs from 'dayjs'
import router from './router'
import store from './store'
import App from './App.vue'

// Global Components
import './global-components'

// 3rd party plugins
import '@/libs/portal-vue'
import '@/libs/toastification'
import '@/libs/axios'
import '@/libs/font-awesome'

// add dayjs
Vue.prototype.$dayjs = dayjs

// BSV Plugin Registration
Vue.use(ModalPlugin)
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

// Composition API
Vue.use(VueCompositionAPI)

// import core styles
require('@core/scss/core.scss')
import '@/main.css'

// Import Bootstrap and BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

// import assets styles
require('@/assets/scss/style.scss')

Vue.config.productionTip = false

import i18n from '@/libs/i18n'

new Vue({
  router,
  store,
  i18n,
  render: h => h(App)
}).$mount('#app')
