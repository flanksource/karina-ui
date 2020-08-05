<template>      
  <v-card>
    <v-layout row class="fill-size">
      <v-card flat class="pa-2">
        <v-row>
          <ItemIcon icon="git" colour="#f8cecc" label="Change Config" count="5"/>
          <ItemIcon icon="git" colour="#666666" label="Downgrade Version" count="2"/>
          <ItemIcon icon="git" colour="#d5e8d4" label="Bump Replicas" count="0"/>
          <ItemIcon icon="git" colour="#d5e8d4" label="Sample" count="0"/>
          <ItemIcon icon="git" colour="#d5e8d4" label="Sample" count="0"/>
        </v-row>
      </v-card>
    </v-layout>
          
    <v-layout row class="fill-size">
      <ItemList :items="stats.items"/>
      <IndicatorList :indicators="stats.indicators"/>
    </v-layout>

    <v-layout row class="fill-size">
      <v-card flat class="pa-2">
        <v-row>
          <ItemIcon icon="git" colour="#d5e8d4" label="Sample" count="0"/>
          <ItemIcon icon="git" colour="#d5e8d4" label="Sample" count="0"/>
          <ItemIcon icon="git" colour="#d5e8d4" label="Sample" count="0"/>              
        </v-row>
      </v-card>
    </v-layout>
  </v-card>
</template>

<script>
  import ItemList from './ItemList.vue'
  import IndicatorList from './IndicatorList.vue'
  import ItemIcon from './ItemIcon.vue'
  import stats from '../data/sample.json'

  export default {
    components:{
      ItemList, IndicatorList, ItemIcon
    },
    data () {
      return {
        stats: stats
      }
    }
  }
</script>

<style scoped>
  .fill-size{
    width:100%;
    margin: 10px;
  }
</style>