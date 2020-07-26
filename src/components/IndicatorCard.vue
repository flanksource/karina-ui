<template>


 
      <div class="d-flex flex-no-wrap justify-space-between indie">



{{ indicatorx}}
        <div  v-for="(item, i) in indicatorx"
       :key="i"
       cols="12">


        
            
         <!--    {{ indicatoricons }} -->
              <v-layout row>
                      <v-item
        
          
        >

           <IndicatorIcon :indy="indicatorx"/> 
        
        </v-item>
      </v-layout>
             

            </div>
            </div>
        
        

</template>


<script>

//import IndicatorIcon from './IndicatorIcon.vue'
  export default {
    components: {
      //IndicatorIcon

    },
    props: {

      indicatorx: Object,
   /*   shape: String,
      color: String,
      overlay: String,*/
    },
    data: () => ({
      drawer: null,
    }),
  }
</script>



<style>
 .indie  {
    border-bottom: 1px dashed black !important
}

</style>


