<template>

  <span>
    <span>
      <IndicatorCardCanary
        indicatoricon="birdie.svg"
        :canarychecks="canarychecks"
      /> 

    </span>

    
  </span>

</template>


<script>

  import IndicatorCardCanary from './IndicatorCardCanary.vue'

  export default {
    name: 'IndicatorList',

    components: {
      IndicatorCardCanary
    },

    props:{
      indicators: Array,
      canarychecks: Array
    },
  }

</script>


<style scoped>

  .no-events {
    pointer-events: none
  }

</style>