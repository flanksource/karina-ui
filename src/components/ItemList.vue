<template>





  

        <v-list shaped>
      <v-list-item-group v-model="item" color="primary">
       <!--  <v-list-item
         v-for="(item, i) in items"
         :key="i"
       >
         <v-list-item-icon>
           <v-icon v-text="item.icon"></v-icon>
         </v-list-item-icon>
         <v-list-item-content>
           <v-list-item-title v-text="item.text"></v-list-item-title>
       
         </v-list-item-content>
       </v-list-item> -->

          <ItemCard/>
      </v-list-item-group>
    </v-list>
     
            
</template>



<script>

import ItemCard from './ItemCard.vue'

export default {

  components:{
    ItemCard

  },
    props: {
      source: String,
    },
    data: () => ({
      drawer: null,
    }),
  }
</script>



<style>
 .v-tabs .v-tab {
    background-color: #ffffff !important;
    border-right: 5px solid red !important
}

</style>