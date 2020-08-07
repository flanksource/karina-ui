<template>
  <v-list class="indcator-list pa-0" width="73%">
    <v-list-item-group>
      <v-list-item
        v-for="(indicator, i) in indicators"
        :key="i"
        class="px-0 no-events"
      > 
        <IndicatorCard :iconlg="indicator.indicator_icon_lg" :indicatoricons="indicator.indicator_icons"/>
      </v-list-item>
    </v-list-item-group>
  </v-list> 
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