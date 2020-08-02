<template>
  <v-list class="indcator-list pa-0" width="73%">
    <v-list-item-group>
      <v-list-item
        v-for="(item, i) in indicators"
        :key="i"
        class="px-0"
      > 
        <IndicatorCard :iconlg="item.indicator_icon_lg" :indicatoricons="item.indicator_icons"/>
      </v-list-item>
    </v-list-item-group>
  </v-list> 
</template>

<script>
  import IndicatorCard from './IndicatorCard.vue'

  export default {
    components: {
      IndicatorCard
    },

    props:{
      indicators: Object,
    }
  }
</script>

<style>
  .indicator-list  {
    width: 73%;
  }
</style>