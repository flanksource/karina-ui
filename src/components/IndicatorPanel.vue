<template>
  <div>
    <v-layout row >
      <ItemList :items="stats.items"/>

      <IndicatorList :indicators="stats.indicators"/>
    </v-layout>
  </div>
</template>

<script>
  import ItemList from './ItemList.vue'
  import IndicatorList from './IndicatorList.vue'
  import stats from '../data/sample.json'

  export default {
    components:{
      ItemList, IndicatorList
    },
    data () {
      return {
        stats: stats
      }
    }
  }
</script>