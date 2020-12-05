<template>
    <v-navigation-drawer
      permanent
      expand-on-hover
    >

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
     <!--      <svg-icon :icon="menuItem.icon" class="icon" /> -->
          <v-list-item-icon class="icon" >
          <v-icon>{{ menuItem.icon }}</v-icon>
        </v-list-item-icon>
          <v-list-item-title>{{menuItem.title}}</v-list-item-title>
        </v-list-item>
      </span>
      </v-list>
    </v-navigation-drawer>
</template>

<script>
 // import SvgIcon from "./SVGIcon.vue";
 
export default {
  name: "Menu",

  components: {
  //  SvgIcon,
  },

  props: {
    state: String,
    clusters: Array,
  },

  data: () => ({
    menuItems: [
      {
        "title": "Menu Item",
        "icon": "mdi-menu-open",
        "url": "",
      },
      {
        "title": "Menu Item",
        "icon": "mdi-menu-open",
        "url": "",
      },
      {
        "title": "Menu Item",
        "icon": "mdi-menu-open",
        "url": "",
      },
      {
         "title": "Menu Item",
        "icon": "mdi-menu-open",
        "url": "",
      },
      {
         "title": "Menu Item",
        "icon": "mdi-menu-open",
        "url": "",
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
  padding-right: 15px !important;
  padding-left: 5px;
}

.v-application--is-ltr .v-list-item__action:first-child, .v-application--is-ltr .v-list-item__icon:first-child {
   margin-right: unset !important;
}

.v-list-item__icon {
  align-self: flex-start;
  margin: 11px 0 !important;
}
</style>