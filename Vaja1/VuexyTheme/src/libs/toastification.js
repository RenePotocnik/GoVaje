import Vue from 'vue'
import Toast from 'vue-toastification'
import ToastificationContent from '@core/components/toastification/ToastificationContent.vue'

// Toast Notification Component Styles
import '@core/scss/vue/libs/toastification.scss'

/**
 * NOTE: If you are using other transition them make sure to import it in `src/@core/scss/vue/libs/notification.scss` from it's source
 */
Vue.use(Toast, {
  hideProgressBar: true,
  closeOnClick: false,
  closeButton: false,
  icon: false,
  timeout: 3000,
  transition: 'Vue-Toastification__fade'
})

Vue.prototype.$printError = function(error) {
  this.$toast({
    component: ToastificationContent,
    props: {
      title: 'Error',
      icon: 'DangerOctagonIcon',
      text: error,
      variant: 'danger'
    }
  })
}

Vue.prototype.$printWarning = function(message) {
  this.$toast({
    component: ToastificationContent,
    props: {
      title: 'Warning',
      icon: 'AlertTriangleIcon',
      text: message,
      variant: 'warning'
    }
  })
}

Vue.prototype.$printSuccess = function(message) {
  this.$toast({
    component: ToastificationContent,
    props: {
      title: 'Success',
      icon: 'CheckCircleIcon',
      text: message,
      variant: 'success'
    }
  })
}
