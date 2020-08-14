<template>
  <span>
    <div v-if="content === ' '">
      <v-badge :color="colour" :content="content" inline="true">
        <svg-icon :icon="icon" :hasFill="false" :color="color"></svg-icon>

        <v-text class="pl-1 text-size">{{label}}</v-text>
      </v-badge>  
    </div>

    <div v-else >
      <v-badge :color="colour" :content="content" inline="true">
        <svg-icon :icon="icon" :hasFill="false" :color="color"></svg-icon>

        <v-text class="pl-1 text-size">{{label}}</v-text>
      </v-badge>  
    </div>
  </span>
</template>

<script>
  import SvgIcon from './SVGIcon.vue';

  export default {
    name: 'ItemCard',
    components: {
      'svg-icon': SvgIcon,
    },
    props: {
      source: String,
      icon: String,
      label: String,
      content: String,
      colour: String,
    },
  }
</script>

<style scoped>
  .text-size{
    font-size: 12px;
  }
</style>