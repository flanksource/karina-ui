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
      <ItemList :items="items"/>
      <IndicatorList :indicators="indicators"/>
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


  export default {
    name: 'IndicatorPanel',
    components:{
      ItemList, IndicatorList, ItemIcon
    },
    
    data: () => ({
      items: [
        {id:1, service:'memory', icon:'ram', label:'512GB', badge:'0'},
        {id:2, service: 'cpu', icon: 'core', label: '96',   badge: '2'},
        {id:3, service: 'storage',  icon: 'storage',label: '3TB', badge: '3'},
        {id:4, service: 'node', icon: 'node', label: '512GB', badge: '1'},
      ],

      indicators: [
        { 
          id: 1,
          indicator_icon_lg: require('@/assets/svg/heartline.svg'),
          indicator_icons: [
            {id:1,color:'', shape:'', overlay:''},
            {id:2,color:'', shape:'', overlay:''},
            {id:3,color:'', shape:'', overlay:''},
          ],
        },
        { 
          id: 2,
          indicator_icon_lg: require('@/assets/svg/cluster.svg'),
          indicator_icons: [
            {id:1,color:'', shape:'', overlay:''},
            {id:2,color:'', shape:'', overlay:''},
            {id:3,color:'', shape:'', overlay:''},
          ],
        },
        { 
          id: 3,
          indicator_icon_lg: require('@/assets/svg/birdie.svg'),
          indicator_icons: [
            {id:1,color:'#24ea1f', shape:'heart', overlay:''},
            {id:2,color:'#24ea1f', shape:'heart', overlay:''},
            {id:3,color:'#24ea1f', shape:'heart', overlay:''},
            {id:4,color:'#24ea1f', shape:'heart', overlay:''},
            {id:5,color:'#24ea1f', shape:'heart', overlay:''},
            {id:6,color:'#24ea1f', shape:'heart', overlay:''},
            {id:7,color:'#24ea1f', shape:'heart', overlay:''},
            {id:8,color:'#24ea1f', shape:'heart', overlay:''},
            {id:9,color:'#24ea1f', shape:'heart', overlay:''},
            {id:10,color:'#24ea1f', shape:'heart', overlay:''},
            {id:11,color:'#24ea1f', shape:'heart', overlay:''},
            {id:12,color:'#24ea1f', shape:'heart', overlay:''},
            {id:13,color:'#b36807', shape:'heart', overlay:''},
            {id:14,color:'#24ea1f', shape:'heart', overlay:''},
            {id:15,color:'#24ea1f', shape:'heart', overlay:''},
            {id:16,color:'#24ea1f', shape:'heart', overlay:''},
            {id:17,color:'#24ea1f', shape:'heart', overlay:''},
            {id:18,color:'#f00', shape:'heart', overlay:''},
            {id:19,color:'#24ea1f', shape:'heart', overlay:''},
            {id:20,color:'#24ea1f', shape:'heart', overlay:''},
          ],
        },
        { 
          id: 4,
          indicator_icon_lg: require('@/assets/svg/idea.svg'),
          indicator_icons: [
            {id:1,color:'#24ea1f', shape:'square', overlay:'`'},
            {id:2,color:'#24ea1f', shape:'square', overlay:''},
            {id:3,color:'#24ea1f', shape:'square', overlay:''},
            {id:4,color:'#24ea1f', shape:'square', overlay:''},
          ],
        },
        { 
          id: 5,
          indicator_icon_lg: require('@/assets/svg/server.svg'),
          indicator_icons: [
            {id:1,color:'#24ea1f', shape:'square', overlay:''},
            {id:2,color:'#24ea1f', shape:'square', overlay:''},
            {id:3,color:'#24ea1f', shape:'square', overlay:''},
            {id:4,color:'#24ea1f', shape:'square', overlay:''},
            {id:5,color:'#24ea1f', shape:'square', overlay:''},
            {id:6,color:'#24ea1f', shape:'square', overlay:''},
            {id:7,color:'#24ea1f', shape:'square', overlay:''},
            {id:8,color:'#24ea1f', shape:'square', overlay:''},
            {id:9,color:'#24ea1f', shape:'square', overlay:''},
            {id:10,color:'#24ea1f', shape:'square', overlay:''},
            {id:11,color:'#24ea1f', shape:'square', overlay:''},
            {id:12,color:'#24ea1f', shape:'square', overlay:''},
            {id:13,color:'#24ea1f', shape:'square', overlay:'mdi-chart-bell-curve-cumulative'},
            {id:14,color:'#24ea1f', shape:'square', overlay:''},
            {id:15,color:'#24ea1f', shape:'square', overlay:''},
            {id:16,color:'#24ea1f', shape:'square', overlay:''},
            {id:17,color:'#24ea1f', shape:'square', overlay:''},
            {id:18,color:'#24ea1f', shape:'square', overlay:''},
            {id:19,color:'#24ea1f', shape:'square', overlay:''},
            {id:20,color:'#24ea1f', shape:'square', overlay:''},
            {id:21,color:'#24ea1f', shape:'square', overlay:''},
            {id:22,color:'#24ea1f', shape:'square', overlay:''},
            {id:23,color:'#24ea1f', shape:'square', overlay:''},
            {id:24,color:'#24ea1f', shape:'square', overlay:''},
            {id:25,color:'#24ea1f', shape:'square', overlay:''},
            {id:26,color:'#f00', shape:'square', overlay:'mdi-bell'},
            {id:27,color:'#24ea1f', shape:'square', overlay:''},
            {id:28,color:'#24ea1f', shape:'square', overlay:''},
            {id:29,color:'#24ea1f', shape:'square', overlay:''},
            {id:30,color:'#24ea1f', shape:'square', overlay:''},
            {id:31,color:'#24ea1f', shape:'square', overlay:''},
            {id:32,color:'#24ea1f', shape:'square', overlay:''},
            {id:33,color:'#b36807', shape:'square', overlay:'mdi-bell'},
            {id:34,color:'#24ea1f', shape:'square', overlay:''},
            {id:35,color:'#24ea1f', shape:'square', overlay:''},
            {id:36,color:'#24ea1f', shape:'square', overlay:''},
            {id:37,color:'#24ea1f', shape:'square', overlay:''},
            {id:38,color:'#24ea1f', shape:'square', overlay:''},
            {id:39,color:'#24ea1f', shape:'square', overlay:''},
            {id:40,color:'#24ea1f', shape:'square', overlay:''},
          ],
        },
      ],
    }),
  }
</script>

<style scoped>
  .fill-size{
    width:100%;
    margin: 10px;
  }
</style>