<template>
  <p class="mb-0" style="display: flex; justify-content: space-between;">
    <span class="float-md-left d-block d-md-inline-block mt-25">
      COPYRIGHT  © {{ new Date().getFullYear() }}
      <b-link
          class="ml-25"
          href="https://1.envato.market/pixinvent_portfolio"
          target="_blank"
      >Pixinvent</b-link>
      <span class="d-none d-sm-inline-block">, All rights Reserved</span>
    </span>

    <span class="float-md-left d-block d-md-inline-block mt-25">
      I am also here, Hello! <em>-Rene Potocnik</em>
    </span>

    <span class="">Hand-crafted &amp; Made with
      <feather-icon
          icon="HeartIcon"
          size="21"
          class="text-danger stroke-current"
      />
    </span>
  </p>
</template>

<script>
import { BLink } from 'bootstrap-vue'

export default {
  components: {
    BLink,
  },
}
</script>
