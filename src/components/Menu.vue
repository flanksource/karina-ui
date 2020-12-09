<template>
    <v-navigation-drawer permanent>
      <v-list-item
      link
      @click="selectState('native')"
      >
        <v-list-item-title class="title">
          Home
        </v-list-item-title>
      </v-list-item>

      <v-divider></v-divider>

      <v-list
        dense 
      >
       <span v-for="(menuItem, i) in menuItems" :key="i">
        <v-list-item
          link
          :href="menuItem.url"
          target="menuFrame"
          @click="selectState('iframe')"
        >
          <svg-icon :icon="menuItem.icon" class="icon" />
          <v-list-item-title>{{menuItem.title}}</v-list-item-title>
        </v-list-item>
      </span>
      </v-list>

      <v-list-item
        @click="toggleDarkMode"
      >
        <v-icon>mdi-theme-light-dark</v-icon>
         <v-list-item-title class="button-text">
          Switch Theme
        </v-list-item-title>
      </v-list-item>

    </v-navigation-drawer>
</template>

<script>
  import SvgIcon from "./SVGIcon.vue";
 
export default {
  name: "Menu",

  components: {
    SvgIcon,
  },

  props: {
    state: String,
    clusters: Array,
  },

  data: () => ({
    menuItems: [
      {
        "title": "Canary",
        "icon": "canary-checker",
        "url": "https://canaries.hetzner.lab.flanksource.com/",
      },
      {
        "title": "Grafana",
        "icon": "grafana",
        "url": "https://httpstat.us/200",
      },
      {
        "title": "Karma",
        "icon": "karma",
        "url": "https://httpstat.us/200",
      },
      {
        "title": "Kibana",
        "icon": "kibana",
        "url": "https://httpstat.us/200",
      },
      {
        "title": "Prometheus",
        "icon": "prometheus",
        "url": "https://prometheus.hetzner.lab.flanksource.com/graph",
      }
    ]  
  }),

  methods: {
    selectState(state) {
      this.$emit('selectState',state);
    },

    toggleDarkMode() {
      this.$vuetify.theme.dark = !this.$vuetify.theme.dark;
      localStorage.setItem("dark_theme", this.$vuetify.theme.dark.toString());
    }
  },
}
</script>

<style>
.icon {
  position: relative;
  top: 0;
  left: 0;
  padding-right: 5px;
}

.button-text {
  
}
</style>