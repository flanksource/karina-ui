<template>

  <span>
    <span
      v-for="(indicator, i) in indicators"
      :key="i"
      class="no-events"
    > 
      <IndicatorCard
        :iconlg="indicator.indicator_icon_lg"
        :indicatoricons="indicator.indicator_icons"
      />
    </span>

    <span 
      v-for="(cluster, i) in apidata"
      :key="i"
    >  
      <IndicatorCardCanary
        indicatoricon="birdie.svg"
        :checktypes="cluster.canary_checks"
      /> 

    </span>

       
  </span>

</template>

<script>

  import IndicatorCard from './IndicatorCard.vue'
  import IndicatorCardCanary from './IndicatorCardCanary.vue'

  export default {
    name: 'IndicatorList',

    components: {
      IndicatorCard,
      IndicatorCardCanary
    },

    props:{
      indicators: Array,
    },

    data () {
      return {
        apidata: null,
      }
    },

    mounted() {
      this.$axios
        .get('http://localhost:8080/api')
        .then(response => (this.apidata = response.data))
        .catch(error => console.log(error))  
    }

  }

</script>

<style scoped>

  .no-events {
    pointer-events: none
  }

</style>