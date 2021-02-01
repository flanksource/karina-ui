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
            <i :class="'icon-'+menuItem.icon"></i>
          <v-list-item-title>{{menuItem.title}}</v-list-item-title>
        </v-list-item>
      </span>
      </v-list>
    </v-navigation-drawer>
</template>

<script>
 
export default {
  name: "Menu",

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
@import '../assets/fonts/karina-ui-icons/icons.css';

.icon-grafana:before {
    color: #e5512b;
}

.icon-karma:before {
    color: #3d9c8a;
}

.icon-prometheus:before {
    color: #e5512b;
}

</style>