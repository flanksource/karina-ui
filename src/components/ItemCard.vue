<template>
 <v-badge
          color="deep-purple accent-4"
         content="6"
         inline="true"
        >
       


 <v-icon
          :color="indigo"
          class="mr-3"
          size="24"
        >
          mdi-memory
        </v-icon>
   
          512GB
        </v-badge>


  

      
     
            
</template>



<script>

//Calls two components for memory usage and for CPU usage

export default {

  components:{

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