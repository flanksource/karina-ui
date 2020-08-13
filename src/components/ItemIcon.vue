<template>
  <span>
    <div v-if="count === ' '">
      <v-badge color="transparent" :content="count">
        <svg-icon :icon="icon" :style="iconStyles" class="icon"></svg-icon>
      </v-badge>
    </div>

    <div v-else>
      <v-badge bottom :content="count"
        offset-x="15"
        offset-y="20"
        overlap="true"
      >
        <svg-icon :icon="icon" :style="iconStyles" class="icon"></svg-icon>
      </v-badge>
    </div>

    <v-text class="text-size">{{label}}</v-text>
  </span>
</template>

<script>
  import SvgIcon from './SVGIcon.vue';

  export default {
    name: 'ItemIcon',
    components: {
      SvgIcon,
    },
    props: {
      id: String,
      label: String,
      icon: String,
      count: Number,
      colour: String,
    },
    computed: {
      iconStyles() {
        return {
          "color": this.colour,
        };
      }
    },
  }
</script>

<style scoped>
  .icon {
    font-size: 48px; 
  }
  .text-size{
    font-size: 12px;
  }
</style>