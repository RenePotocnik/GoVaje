import Vue from 'vue'
import VueI18n from 'vue-i18n'

const messages = {
    en: require('./locales/en.json'),
    hr: require('./locales/hr.json')
}

Vue.use(VueI18n)

const i18n = new VueI18n({
    locale: 'hr', // set locale
    messages, // set locale messages
})



export default  i18n

