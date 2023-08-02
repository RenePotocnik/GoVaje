<template>
  <div class="navbar-container d-flex content align-items-center">

    <!-- Nav Menu Toggler -->
    <ul class="nav navbar-nav d-xl-none">
      <li class="nav-item">
        <b-link
          class="nav-link"
          @click="toggleVerticalMenuActive"
        >
          <feather-icon
            icon="MenuIcon"
            size="21"
          />
        </b-link>
      </li>
    </ul>

    <!-- Left Col -->
    <div class="bookmark-wrapper align-items-center flex-grow-1 d-none d-lg-flex">
      <dark-Toggler class="d-none d-lg-block" />
<!--        <h3 style="margin: 0 0 0 1vw;"> {{ $route.meta.pageTitle }} </h3>-->
      <b-col cols="12">
        <h2 class="content-header-title float-left pr-1 mb-0">
          {{ $route.meta.pageTitle }}
        </h2>
        <div class="breadcrumb-wrapper">
          <b-breadcrumb>
            <b-breadcrumb-item to="/">
              <feather-icon
                  icon="HomeIcon"
                  size="16"
                  class="align-text-top"
              />
            </b-breadcrumb-item>
            <b-breadcrumb-item
                v-for="item in $route.meta.breadcrumb"
                :key="item.text"
                :active="item.active"
                :to="item.to"
            >
              {{ item.text }}
            </b-breadcrumb-item>
          </b-breadcrumb>
        </div>
      </b-col>
    </div>

    <b-navbar-nav class="nav align-items-center ml-auto">
      <b-nav-item-dropdown
        right
        toggle-class="d-flex align-items-center dropdown-user-link"
        class="dropdown-user"
      >
        <template #button-content>
          <div class="d-sm-flex d-none user-nav">
            <p class="user-name font-weight-bolder mb-0">
              {{ username }}
            </p>
            <span class="user-status"></span>
          </div>
          <b-avatar
            size="40"
            variant="light-primary"
            badge
            :src="require('@/assets/images/avatars/13-small.png')"
            class="badge-minimal"
            badge-variant="success"
          />
        </template>

        <b-dropdown-item link-class="d-flex align-items-center"
                         @click="$router.push('/')">
          <feather-icon
            size="16"
            icon="HomeIcon"
            class="mr-50"
          />
          <span>Home</span>
        </b-dropdown-item>

        <b-dropdown-item link-class="d-flex align-items-center"
                         @click="$router.push('todo-page')">
          <feather-icon
            size="16"
            icon="CheckIcon"
            class="mr-50"
          />
          <span>Tasks</span>
        </b-dropdown-item>

        <b-dropdown-divider />

        <b-dropdown-item link-class="d-flex align-items-center"
                         @click="logoutUser">
          <feather-icon
            size="16"
            icon="LogOutIcon"
            class="mr-50"
          />
          <span>Logout</span>
        </b-dropdown-item>
      </b-nav-item-dropdown>
    </b-navbar-nav>
  </div>
</template>

<script>
import {
  BLink, BNavbarNav, BNavItemDropdown, BDropdownItem, BDropdownDivider, BAvatar,
} from 'bootstrap-vue'
import DarkToggler from '@core/layouts/components/app-navbar/components/DarkToggler.vue'
import {logoutUser} from "@/main";

export default {
  methods: {
    logoutUser() {
      logoutUser();
    },
  },
  components: {
    BLink,
    BNavbarNav,
    BNavItemDropdown,
    BDropdownItem,
    BDropdownDivider,
    BAvatar,

    // Navbar Components
    DarkToggler,
  },
  data() {
    return {
      username: localStorage.getItem('username') || 'User', // Default to 'User' if no username in local storage
    }
  },
  props: {
    toggleVerticalMenuActive: {
      type: Function,
      default: () => {},
    },
  },
}
</script>
