<template>
  <v-card class="mx-auto" max-width="600" min-height="400">
  <v-tabs vertical>

      <ItemList :items="items"/>

      <vr/>
  
      <v-tab-item
        v-for="(item, i) in items"
        :key="i"
      >
      
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
          icon: 'ram',
          label: '512GB',
          badge: '',
          indicators: [
         { 
          id: 1,
          indicator_icon_lg: require('@/assets/svg/heartline.svg'),
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:'mdi-information'},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},
             { 
          id: 2,
          indicator_icon_lg: require('@/assets/svg/cluster.svg'),
          indicator_icons: [
            {id:1,color:'#f00', shape:'smile', overlay:'mdi-information'},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

 { 
          id: 3,
          indicator_icon_lg: require('@/assets/svg/birdie.svg'),
          indicator_icons: [
            {id:1,color:'#f00', shape:'heart', overlay:''},
            {id:2,color:'#0f0', shape:'heart', overlay:'mdi-information'},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
            {id:4,color:'#0f0', shape:'heart', overlay:''},
            {id:5,color:'#0f0', shape:'heart', overlay:''},
            {id:6,color:'#0f0', shape:'heart', overlay:''},
            {id:7,color:'#0f0', shape:'heart', overlay:''},
            {id:8,color:'#0f0', shape:'heart', overlay:''},
            {id:9,color:'#0f0', shape:'heart', overlay:''},
            {id:10,color:'#0f0', shape:'heart', overlay:''},
            {id:11,color:'#f00', shape:'heart', overlay:''},
            {id:12,color:'#0f0', shape:'heart', overlay:''},
            {id:13,color:'#0f0', shape:'heart', overlay:''},
            {id:14,color:'#0f0', shape:'heart', overlay:''},
            {id:15,color:'#0f0', shape:'heart', overlay:''},
            {id:16,color:'#0f0', shape:'heart', overlay:''},
            {id:17,color:'#0f0', shape:'heart', overlay:''},
            {id:18,color:'#0f0', shape:'heart', overlay:''},
            {id:19,color:'#0f0', shape:'heart', overlay:''},
            {id:20,color:'#0f0', shape:'heart', overlay:''},
        ],},
         { 
          id: 4,
          indicator_icon_lg: require('@/assets/svg/idea.svg'),
          indicator_icons: [
            {id:1,color:'#f00', shape:'square', overlay:''},
            {id:2,color:'#0f0', shape:'square', overlay:''},
            {id:3,color:'#0f0', shape:'square', overlay:''},
            {id:4,color:'#0f0', shape:'square', overlay:''},
        ],},
         { 
          id: 4,
          indicator_icon_lg: require('@/assets/svg/server.svg'),
          indicator_icons: [
            {id:1,color:'#f00', shape:'square', overlay:''},
            {id:2,color:'#0f0', shape:'square', overlay:''},
            {id:3,color:'#0f0', shape:'square', overlay:''},
            {id:4,color:'#0f0', shape:'square', overlay:''},
            {id:5,color:'#0f0', shape:'square', overlay:''},
            {id:6,color:'#0f0', shape:'square', overlay:''},
            {id:7,color:'#0f0', shape:'square', overlay:''},
            {id:8,color:'#0f0', shape:'square', overlay:''},
            {id:9,color:'#0f0', shape:'square', overlay:''},
            {id:10,color:'#0f0', shape:'square', overlay:''},
            {id:11,color:'#f00', shape:'square', overlay:''},
            {id:12,color:'#0f0', shape:'square', overlay:''},
            {id:13,color:'#0f0', shape:'square', overlay:''},
            {id:14,color:'#0f0', shape:'square', overlay:''},
            {id:15,color:'#0f0', shape:'square', overlay:''},
            {id:16,color:'#0f0', shape:'square', overlay:''},
            {id:17,color:'#0f0', shape:'square', overlay:''},
            {id:18,color:'#0f0', shape:'square', overlay:''},
            {id:19,color:'#0f0', shape:'square', overlay:''},
            {id:20,color:'#0f0', shape:'square', overlay:''},
            {id:21,color:'#f00', shape:'square', overlay:''},
            {id:22,color:'#0f0', shape:'square', overlay:''},
            {id:23,color:'#0f0', shape:'square', overlay:''},
            {id:24,color:'#0f0', shape:'square', overlay:''},
            {id:25,color:'#0f0', shape:'square', overlay:''},
            {id:26,color:'#0f0', shape:'square', overlay:''},
            {id:27,color:'#0f0', shape:'square', overlay:''},
            {id:28,color:'#0f0', shape:'square', overlay:''},
            {id:29,color:'#0f0', shape:'square', overlay:''},
            {id:30,color:'#0f0', shape:'square', overlay:''},
            {id:31,color:'#f00', shape:'square', overlay:''},
            {id:32,color:'#0f0', shape:'square', overlay:''},
            {id:33,color:'#0f0', shape:'square', overlay:''},
            {id:34,color:'#0f0', shape:'square', overlay:''},
            {id:35,color:'#0f0', shape:'square', overlay:''},
            {id:36,color:'#0f0', shape:'square', overlay:''},
            {id:37,color:'#0f0', shape:'square', overlay:''},
            {id:38,color:'#0f0', shape:'square', overlay:''},
            {id:39,color:'#0f0', shape:'square', overlay:''},
            {id:40,color:'#0f0', shape:'square', overlay:''},
        ],},

        ],
        },
            {
          id: 2,
          service: 'cpu',
          icon: 'core',
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
            {id:1,color:'#f00', shape:'smile', overlay:'mdi-information'},
            {id:2,color:'#0f0', shape:'heart', overlay:''},
            {id:3,color:'#0f0', shape:'heart', overlay:''},
        ],},

        { 
          id: 4,
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
          icon: 'storage',
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
          service: 'node',
          icon: 'node',
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
