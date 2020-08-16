<template>
  <span>
    <v-span
      v-for="(indicator, i) in indicators"
      :key="i"
      class="no-events"
    > 
      <IndicatorCard :iconlg="indicator.indicator_icon_lg" :indicatoricons="indicator.indicator_icons"/>
    </v-span>
  </span>
</template>

<script>
  import IndicatorCard from './IndicatorCard.vue'

  export default {
    name: 'IndicatorList',
    components: {
      IndicatorCard
    },
    props:{
      indicators: Object,
    }
  }
</script>

<style scoped>
  .no-events{
    pointer-events: none
  }
</style>