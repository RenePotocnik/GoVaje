<script>

import HomeIcon from 'vue-feather-icons/icons/HomeIcon'
import FileIcon from 'vue-feather-icons/icons/FileIcon'
import SunIcon from 'vue-feather-icons/icons/SunIcon'
import MoonIcon from 'vue-feather-icons/icons/MoonIcon'
import DiscIcon from 'vue-feather-icons/icons/DiscIcon'
import ChevronDownIcon from 'vue-feather-icons/icons/ChevronDownIcon'
import RotateCwIcon from 'vue-feather-icons/icons/RotateCwIcon'
import XIcon from 'vue-feather-icons/icons/XIcon'
import ArrowUpIcon from 'vue-feather-icons/icons/ArrowUpIcon'
import SettingsIcon from 'vue-feather-icons/icons/SettingsIcon'
import HeartIcon from 'vue-feather-icons/icons/HeartIcon'
import MenuIcon from 'vue-feather-icons/icons/MenuIcon'
import BellIcon from 'vue-feather-icons/icons/BellIcon'
import SearchIcon from 'vue-feather-icons/icons/SearchIcon'
import UserIcon from 'vue-feather-icons/icons/UserIcon'
import LogOutIcon from 'vue-feather-icons/icons/LogOutIcon'
import CheckIcon from 'vue-feather-icons/icons/CheckIcon'

const icons = {
  HomeIcon,
  FileIcon,
  SunIcon,
  DiscIcon,
  MoonIcon,
  ChevronDownIcon,
  RotateCwIcon,
  ArrowUpIcon,
  XIcon,
  SettingsIcon,
  HeartIcon,
  MenuIcon,
  BellIcon,
  SearchIcon,
  UserIcon,
  LogOutIcon,
  CheckIcon,
}

export default {
  name: 'FeatherIcon',
  functional: true,
  props: {
    icon: {
      required: true,
      type: [String, Object],
    },
    size: {
      type: String,
      default: '14',
    },
    badge: {
      type: [String, Object, Number],
      default: null,
    },
    badgeClasses: {
      type: [String, Object, Array],
      default: 'badge-primary',
    },
  },
  render(h, { props, data }) {
    // Create Feather Icon
    const svg = h(icons[props.icon], { props: { size: props.size }, ...data })

    // If no badge is provided => Render just SVG
    if (!props.badge) return svg

    // Badge
    const badgeEl = h('span', { staticClass: 'badge badge-up badge-pill', class: props.badgeClasses }, [props.badge])

    // Return span with children
    return h('span', { staticClass: 'feather-icon position-relative' }, [svg, badgeEl])
  },
}
</script>

<style lang="scss">
svg.feather {
  outline: none
}
</style>
