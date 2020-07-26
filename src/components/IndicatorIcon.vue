<template>
  <div class="smile-icon">
 

 
    <svg-icon :icon="shape" :hasFill="false" :color="color" class="smile-icon"></svg-icon>

  </div>
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
  font-size: 32px;
  color: #000;
  fill: #0F0;
 
}
</style>
