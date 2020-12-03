<template>
    <v-navigation-drawer
      v-model="drawer"
      absolute
      temporary
    >
      <v-list-item
        link
        @click="selectState('native')"
      > 
        <v-list-item-icon class="icon" >
          <v-icon>mdi-home</v-icon>
        </v-list-item-icon>

        <v-list-item-title class="title">
          Home
        </v-list-item-title>
      </v-list-item>

      <v-divider></v-divider>

      <v-list
         
      >
       <span v-for="(menuItem, i) in menuItems" :key="i">
        <v-list-item
          link
          :href="menuItem.url"
          target="menuFrame"
          @click="selectState('iframe')"
          @click.stop="drawer = !drawer"
        >
          <svg-icon :icon="menuItem.icon" class="icon" />
          <v-list-item-title>{{menuItem.title}}</v-list-item-title>
        </v-list-item>
      </span>
      </v-list>
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
    drawer: Boolean,
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
    }
  },
}
</script>

<style scoped>
.icon {
  position: relative;
  top: 0;
  left: 0;
  padding-right: 15px;
  padding-left: 5px;
}

.v-application--is-ltr .v-list-item__action:first-child, .v-application--is-ltr .v-list-item__icon:first-child {
  margin-right: -3px !important;
}

.v-list-item__icon {
  align-self: flex-start;
  margin: 11px 0 !important;
}
</style>