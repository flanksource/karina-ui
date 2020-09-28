<template>
  <span>
    <div v-if="alerts === 0">
      <v-badge color="" :content="alerts" inline>
        <svg-icon :icon="icon" :hasFill="false" color=""/>
        <span class="pl-1 text-size">{{label}}</span>
      </v-badge>  
    </div>

    <div v-else inline>
      <v-badge color="#b3b3b3" :content="alerts" inline>
        <svg-icon :icon="icon" :hasFill="false" color=""/>
        <span class="pl-1 text-size">{{label}}</span>
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
      icon: String,
      label: String,
      alerts: String,
      alertscolour: String,
    },
  }
</script>

<style scoped>
  .text-size{
    font-size: 12px;
  }
</style>