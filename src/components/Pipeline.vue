<template> 
  <div class="pb-12">
  <v-row no-gutters justify="" class="">   
    <v-col cols="12" xs="12" sm="2" md="2" lg="3" xl="3" class="px-2">
       <v-text class="ml-5" justify=""><b><u>Panel {{ cluster.id }}</u></b></v-text> 

       <!-- <button @click='changeName()'>Change ID</button> -->
      <br/>
      <MemoryCard/>
      <br/>
      <CPUCard/>
      <br/>
    </v-col>
    
    <v-col cols="12" xs="12" sm="3" md="3" lg="2" xl="2" class="item-list">
      <v-card flat class="ml-3 py-3">
        <ItemList :items="cluster.items"/>
      </v-card>
    </v-col>
    
    <v-col cols="12" xs="12" sm="7" md="7" lg="6" xl="6" class="indicator-list">
      <v-card flat class="mr-4 py-3">
        <IndicatorList :indicators="cluster.indicators"/>
      </v-card>
    </v-col>
    </v-row>
</div>
</template>
     
<script>
  import CPUCard from './CPUCard.vue'
  import IndicatorList from './IndicatorList.vue'
  import ItemList from './ItemList.vue'
  import MemoryCard from './MemoryCard.vue'

  export default {
    name: 'IndicatorPanel',
    components:{
      ItemList, IndicatorList, MemoryCard, CPUCard
    },
    props: {
      cluster: Object,
    },
 /*   methods: {
      changeName() {
        this.$emit('changeName', 5)
      }
    }*/
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