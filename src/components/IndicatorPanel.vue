<template> 
  <v-card>
    <v-row no-gutters align="center" justify="space-around" class="py-2 ml-4">
      <v-col cols="3" sm="3" md="3" lg="3" xl="3">
        <ItemIcon icon="git" colour="#f8cecc" label="Change Config" count="5"/>
      </v-col>
      
      <v-col cols="3" sm="3" md="3" lg="3" xl="3">
        <ItemIcon icon="git" colour="#666666" label="Downgrade Version" count="2"/>
      </v-col>

      <v-col cols="3" sm="3" md="3" lg="3" xl="3">
        <ItemIcon icon="git" colour="#d5e8d4" label="Bump Replicas" count=" "/>
      </v-col>

      <v-col cols="3" sm="3" md="3" lg="3" xl="3">
        <ItemIcon icon="git" colour="#d5e8d4" label="Sample" count=" "/>
      </v-col>
    </v-row>
        
    <v-row no-gutters justify="">
      <v-col cols="12" xs="12" sm="2" md="2" lg="3" xl="3" class="px-2">
        <br/>
        <MemoryCard/>
        <br/>
        <CPUCard/>
        <br/>
      </v-col>

      <v-col cols="12" xs="12" sm="3" md="3" lg="2" xl="2" class="item-list">
        <v-card flat class="ml-3 py-3">
          <ItemList :items="stats.items"/>
        </v-card>
      </v-col>

      <v-col cols="12" xs="12" sm="7" md="7" lg="6" xl="6" class="indicator-list">
        <v-card flat class="mr-4 py-3">
          <IndicatorList :indicators="stats.indicators"/>
        </v-card>
      </v-col>
    </v-row>

    <v-row no-gutters align="center" justify="space-around" class="py-2">
      <v-col cols="3" xs="12" sm="3" md="3" lg="3" xl="3">
       <ItemIcon icon="git" colour="#d5e8d4" label="Sample" count=" "/>
      </v-col>

      <v-col cols="3" xs="12" sm="3" md="3" lg="3" xl="3">
        <ItemIcon icon="git" colour="#d5e8d4" label="Sample" count=" "/>
      </v-col>

      <v-col cols="3" xs="12" sm="3" md="3" lg="3" xl="3">
        <ItemIcon icon="git" colour="#d5e8d4" label="Sample" count=" "/>
      </v-col>
    </v-row>
  </v-card>
</template>
     
<script>
  import CPUCard from './CPUCard.vue'
  import IndicatorList from './IndicatorList.vue'
  import ItemIcon from './ItemIcon.vue'
  import ItemList from './ItemList.vue'
  import MemoryCard from './MemoryCard.vue'
  import stats from '../data/sample.json'

  export default {
    name: 'IndicatorPanel',
    components:{
      ItemList, IndicatorList, ItemIcon, MemoryCard, CPUCard
    },
    data () {
      return {
        stats: stats
      }
    }
  }
</script>

<style scoped>
  .item-list{
    border-top-left-radius: 20px !important;
    border-bottom-left-radius: 20px !important;
    min-height: 100%;
    border-right: none;
    border: 1px solid #444;
  }
 .indicator-list{
    border-top-right-radius: 20px !important;
    border-bottom-right-radius: 20px !important;
    border-left: 2px solid #222;
    border: 1px solid #444;
  }
</style>