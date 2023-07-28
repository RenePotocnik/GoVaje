import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { faAirFreshener } from '@fortawesome/free-solid-svg-icons/faAirFreshener'
import { faAmbulance } from '@fortawesome/free-solid-svg-icons/faAmbulance'
import { faTimes } from '@fortawesome/free-solid-svg-icons/faTimes'
import { faCalendar } from '@fortawesome/free-solid-svg-icons/faCalendar'
import { faInfoCircle } from '@fortawesome/free-solid-svg-icons/faInfoCircle'

/**
 * če vam meče ven error Uncaught TypeError: Cannot read properties of undefined (reading 'prefix')
 * je zelo verjetno krivo to, da nimate ikone prav poimenovane v importu
 * AirFreshnener namesto faAirFreshener (manjka fa)
 * faAirFreshner namesto faAirFreshener (typo)
 */

library.add(
  faAirFreshener,
  faAmbulance,
  faTimes,
  faCalendar,
  faInfoCircle
)

// Register the component globally
Vue.component('fa', FontAwesomeIcon)
