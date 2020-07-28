<template>
  <span class="parent">
    <svg-icon :icon="shape" :hasFill="true" :style="iconStyles" class="icon "></svg-icon>
    
    <v-icon icon="" color="white" size="14" class="overlay">
      {{ overlay }}
    </v-icon> 
  </span>
</template>

<script>
  import SvgIcon from './SVGIcon.vue';

  export default {
    name: 'my-component',
    props: {
      id: String,
      colour: {
        type: String,
        /*default: "#24ea1f"*/
      },
      shape: String,
      overlay: String,
      indy: Object,

    },
     computed: {
      iconStyles() {
        return {
          "color": this.colour,
        };
      }
    },
    components: {
      'svg-icon': SvgIcon,
    },
       data: () => ({
        drawer: null,
      }),
  }
</script>

<style >
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
    top: -3px;
    left: 3px;
   
  }
</style>