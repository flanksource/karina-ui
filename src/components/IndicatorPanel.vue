<template>
  <v-card
  class="mx-auto"
    max-width="600">
  <v-tabs vertical>

    
      <ItemList :items="items"/>
  

      <v-tab-item

      v-for="(item, i) in items"
      :key="i"
      >
      {{ item.service }}
             <IndicatorList :indicators="item.indicators"/>
          

      </v-tab-item>

    </v-tabs>
  </v-card>
</template>



<script>

//Calls two components for memory usage and for CPU usage

import ItemList from './ItemList.vue'
import IndicatorList from './IndicatorList.vue'

export default {

  components:{
  ItemList, IndicatorList
  },
    props: {
    
    },
        data: () => ({
       items: [
        {
          id: 1,
          service: 'memory',
          icon: 'mdi-memory',
          label: '512GB',
          badge: '1',
          indicators: [
         { 
          id: 1,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},
             { 
          id: 2,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

 { 
          id: 3,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

        ],
        },
            {
          id: 2,
          service: 'cpu',
          icon: 'mdi-cpu-64-bit',
          label: '96',
          badge: '2',
          indicators: [
         { 
          id: 1,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},
             { 
          id: 2,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

 { 
          id: 3,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

        ],
        },
            {
          id: 3,
          service: 'storage',
          icon: 'mdi-server',
          label: '3TB',
          badge: '3',
          indicators: [
         { 
          id: 1,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},
             { 
          id: 2,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

 { 
          id: 3,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

        ],
        },
            {
          id: 4,
          service: 'memory',
          icon: 'mdi-memory',
          label: '512GB',
          badge: '1',
     indicators: [
         { 
          id: 1,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},
             { 
          id: 2,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

 { 
          id: 3,
          indicator_icon_lg: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

        ],
        },
      ],
    }),
  }
</script>
