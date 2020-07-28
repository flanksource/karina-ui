<template>




<div class="vr">

 <v-list>
           <v-list-item-group v-model="item" color="primary">
            <v-list-item
      v-for="(item, i) in items"
      :key="i"
      >
            <v-tab>

           <ItemCard :label="item.label" :icon="item.icon" :content="item.badge"/>
    </v-tab>
      
            </v-list-item>
    
           </v-list-item-group>
         </v-list> 




    
            </div>

</template>



<script>

import ItemCard from './ItemCard.vue'

export default {

  components:{
    ItemCard

  },
    props: {
      items: Object,
      /*icon: String,
      label: String,
      badge: String,*/
    },
     
  }
</script>



<style>
 .vr  {
    
    border-right: 2px solid black !important;

}

</style>