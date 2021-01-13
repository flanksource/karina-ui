<template>
  <span>
    <div v-if="content > 0">
      <v-badge :color="color" :content="content" inline>
<!--        <svg-icon :icon="icon" :hasFill="false" :color="color" />-->
        <i :class="'icon-'+icon"></i>
        <span class="pl-1 text-size">{{ label }} </span>
      </v-badge>
    </div>

    <div v-else inline>
      <v-badge value="false" color="white" inline>
<!--        <svg-icon :icon="icon" :hasFill="false" :color="color" />-->
        <i :class="'icon-'+icon"></i>
        <span class="pl-1 text-size">{{ label }} </span>
      </v-badge>
    </div>
  </span>
</template>

<script>
// import SvgIcon from "./SVGIcon.vue";

export default {
  name: "ItemCard",

  // components: {
  //   "svg-icon": SvgIcon,
  // },

  props: {
    source: String,
    icon: String,
    label: String,
    content: Number,
    color: String,
  },
};
</script>

<style scoped>
.text-size {
  font-size: 12px;
}

@font-face {
  font-family: "karina-ui-icons";
  src: local("icons.eot"), url("../assets/fonts/karina-ui-icons/icons.eot") format("embedded-opentype"),
  local("icons.eot"), url("../assets/fonts/karina-ui-icons/icons.woff2") format("woff2"),
  local("icons.eot"), url("../assets/fonts/karina-ui-icons/icons.woff") format("woff");
}

i[class^="icon-"]:before, i[class*=" icon-"]:before {
  font-family: karina-ui-icons !important;
  font-style: normal;
  font-weight: normal !important;
  font-variant: normal;
  text-transform: none;
  line-height: 1;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}




.icon-cpu:before {
    content: "\f101";
    color: #9a2812 !important;
}
.icon-memory:before {
    content: "\f102";
    color: #336610 !important;
}

</style>
