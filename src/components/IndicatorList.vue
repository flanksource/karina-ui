<template>

 
      <span>
    <span
       v-for="(item, i) in indicators"
       :key="i"
       cols="12"
     > 

          
               

              <IndicatorCard 
                :indicatorx="item.indicator_icons" :iconlg="item.indicator_icon_lg"/>
        
           
             

   </span>
      </span>
   

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




