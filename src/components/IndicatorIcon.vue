<template>

  <span
    v-if="shape == 'none'"
    class="parent"
  />
    
  <span
    v-else
    class="parent"
  >
    <svg-icon
      :icon="shape"
      :style="iconStyles"
      class="icon"
    />
    
    <v-icon
      color="white"
      class="overlay"
    >
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
      colour: String,
      shape: String,
      overlay: String,
    },

    computed: {
      iconStyles() {
        return {
          "fill": this.colour,
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