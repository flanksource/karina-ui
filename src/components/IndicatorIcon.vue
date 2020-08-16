<template>
  <span class="parent">
    <svg-icon :icon="shape" :style="iconStyles" class="icon"></svg-icon>
    
    <v-icon color="white" class="overlay">{{ overlay }}</v-icon> 
  </span>
</template>

<script>
  import SvgIcon from './SVGIcon.vue';

  export default {
    name: 'IndicatorIcon',
    components: {
      SvgIcon,
    },
    props: {
      id: String,
      colour: {
        type: String,
      },
      shape: String,
      overlay: String,
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
  .parent {
    position: relative;
    top: 0;
    left: 0;
  }
  .icon {
    position: absolute;
    top: 0;
    left: 0;
  }
  .overlay {
    position: relative;
    top: -3px;
    left: 2px;
    font-size: 12px;
  }
</style>