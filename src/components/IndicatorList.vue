<template>

 
      <v-row dense>
    <v-col
       v-for="(item, i) in indicators"
       :key="i"
       cols="12"
     > 

          <v-card flat>
            <v-card-text>
       
              <IndicatorCard 
                :indicatorx="item.indicator_icons"/>
            
                <v-avatar
                class="ma-3"
                size="64"
                tile
              >
                <v-img v-bind:src="item.indicator_icon_lg"></v-img>


              </v-avatar>

             

            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
   

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




