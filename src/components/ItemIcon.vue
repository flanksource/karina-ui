<template>
  <span class="item">
    <div v-if="count === '0'">
      <svg-icon :icon="icon" :hasFill="true" :style="iconStyles" class="icon"></svg-icon>
    </div>

    <div v-else>
      <v-badge
        bordered
        bottom
        color="deep-purple accent-4"
        :content="count"
        offset-x="20"
        offset-y="15"
        overlap="true"
      >
     
           <svg-icon :icon="icon" :hasFill="true" :style="iconStyles" class="icon"></svg-icon>
    
      </v-badge>

    </div>

    <div class="caption">
     {{label}}
    </div>


  </span>
</template>

<script>
  import SvgIcon from './SVGIcon.vue';

  export default {
    name: 'my-component',
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
    components: {
      'svg-icon': SvgIcon,
    },
       data: () => ({
        drawer: null,
      }),
  }
</script>

<style scoped>
  .centre{
    align-content: center;
    padding: 0px 20px;
  }
  .icon {
    font-size: 48px; 
  }
  .item {
    vertical-align: top;
    display: inline-block;
    text-align: center;
    width: 120px;
}
.caption {
    display: block;
}
</style>