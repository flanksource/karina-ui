<template>
  <span class="smile-icon">
    <svg-icon :icon="shape" :hasFill="false" :color="color" class="smile-icon"></svg-icon>
   <!--  <v-icon icon="" color="black" size="14">
     {{ overlay }}
   </v-icon> -->
 
  </span>
</template>

<script>
import SvgIcon from './SVGIcon.vue';

export default {
  name: 'my-component',
  props: {
    id: String,
    color: String,
    shape: String,
    overlay: String,
    indy: Object,

  },
  computed: {
    style () {
       return {
       fill: this.colorx
    }
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
.smile-icon {
  font-size: 20px;
  color: #FC0;
  fill: #0F0;
 
}
</style>
