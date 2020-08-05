<template>
  <span class="parent">
    <svg-icon :icon="shape" :hasFill="true" :style="iconStyles" class="icon"></svg-icon>
    
    <v-icon icon="" color="white" size="14" class="overlay">
      {{ overlay }}
    </v-icon> 
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
    font-size: 20px;
  }
  .icon {
    position: absolute;
    top: 0;
    left: 0;
    font-size: 20px;
  }
  .overlay {
    position: absolute;
    top: 0;
    left: 3px;
  }
</style>