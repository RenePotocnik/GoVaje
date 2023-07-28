import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

/* webpackPrefetch: true */
/*
 * webpack magic commet, which will prefetch next page.
 * if we know that user is definitely going to that one page (for example If users is going to products page while in categories, we add comment to products)
 * this will add < link rel="prefetch" href="/products" />
 */

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  scrollBehavior() {
    return { x: 0, y: 0 }
  },
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import(/* webpackPrefetch: true */ '@/views/Home.vue'),
      meta: {
        pageTitle: 'Home',
        breadcrumb: [
          {
            text: 'Home',
            active: true
          }
        ]
      }
    },
    {
      path: '/second-page',
      name: 'second-page',
      component: () => import('@/views/SecondPage.vue'),
      meta: {
        pageTitle: 'Second Page',
        breadcrumb: [
          {
            text: 'Second Page',
            active: true
          }
        ]
      }
    },
    {
      path: '/login',
      name: 'login',
      component: () => import(/* webpackPrefetch: true */ '@/views/Login.vue'),
      meta: {
        layout: 'full'
      }
    },
    {
      path: '/error-404',
      name: 'error-404',
      component: () => import('@/views/error/Error404.vue'),
      meta: {
        layout: 'full'
      }
    },
    {
      path: '*',
      redirect: 'error-404'
    }
  ]
})

/*
 * LOGIKA ZA PREVERJANJE AVTORIZACIJE UPORANIKA
 * LOGIKA ZA PRIDOBIVANJE PODATKOV IZ LOCAL STORAGE, KI SE SHRANIJO V VUEX
*/
// import store from '@/store'
// router.beforeEach((to, from, next) => {

//   // pridobi loggedIn spremenljivko iz localStorage, ki jo ob loginu nastavimo na true
//   const loggedIn = localStorage.getItem('loggedIn')
//   if (loggedIn) {
//     // ob reloadu strani (Ctrl + R) se vrednosti iz store (Vuex Store) pobrišejo
//     // zato jih pridobimo iz localStorage in jih shranimo
//     store.state.user.loggedIn = true
//     const userData = localStorage.getItem('userData')
//     if (userData) {
//       store.state.user.userData = JSON.parse(userData)
//     }
//   }

//   // če stran, na katero želimo iti vsebuje meta podatek requiresAuth
//   // in iz authService zvemo, da user ni prijavljen, ga pošljemo na login
//   // const isLoggedIn = false; // TODO - add is logged in logic
//   if (to.meta.requiresAuth && !isLoggedIn) {
//     return next('/login')
//   }

//   // če meta vsebuje array roles (npr ['user', 'superuser'] ali ['admin'])
//   // pogledamo v vuex store, če ima prijavljen uporabnik ta role
//   if (to.meta.roles) {
//     if (to.meta.roles.includes(store.state.user.userData.type)) {
//       return next() // z next() spustimo dalje
//     }
//     return next('/login')
//   }
//   next()
// })

// ? For splash screen
// Remove afterEach hook if you are not using splash screen
router.afterEach(() => {
  // Remove initial loading
  const appLoading = document.getElementById('loading-bg')
  if (appLoading) {
    appLoading.style.display = 'none'
  }
})

export default router
