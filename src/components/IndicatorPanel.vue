<template>
  <v-card
  class="mx-auto"
    max-width="600">
   <!--  <v-toolbar flat color="primary" dark>
     <v-toolbar-title>User Profile</v-toolbar-title>
   </v-toolbar> -->
    <v-tabs vertical>
      <v-tab>

        <ItemList/>
       
 
      </v-tab>
      <v-tab>
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
         mdi-cpu-64-bit
        </v-icon>
   
          96
        </v-badge>
      </v-tab>
      <v-tab>
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
      </v-tab>

  

      <v-tab-item>
        <v-card flat>
          <v-card-text>
             <IndicatorList/>
          

          </v-card-text>
        </v-card>
      </v-tab-item>
      <v-tab-item>
        <v-card flat>
          <v-card-text>
            <p>
              Morbi nec metus. Suspendisse faucibus, nunc et pellentesque egestas, lacus ante convallis tellus, vitae iaculis lacus elit id tortor. Sed mollis, eros et ultrices tempus, mauris ipsum aliquam libero, non adipiscing dolor urna a orci. Curabitur ligula sapien, tincidunt non, euismod vitae, posuere imperdiet, leo. Nunc sed turpis.
            </p>

            <p>
              Suspendisse feugiat. Suspendisse faucibus, nunc et pellentesque egestas, lacus ante convallis tellus, vitae iaculis lacus elit id tortor. Proin viverra, ligula sit amet ultrices semper, ligula arcu tristique sapien, a accumsan nisi mauris ac eros. In hac habitasse platea dictumst. Fusce ac felis sit amet ligula pharetra condimentum.
            </p>

            <p>
              Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Nam commodo suscipit quam. In consectetuer turpis ut velit. Sed cursus turpis vitae tortor. Aliquam eu nunc.
            </p>

            <p>
              Etiam ut purus mattis mauris sodales aliquam. Ut varius tincidunt libero. Aenean viverra rhoncus pede. Duis leo. Fusce fermentum odio nec arcu.
            </p>

            <p class="mb-0">
              Donec venenatis vulputate lorem. Aenean viverra rhoncus pede. In dui magna, posuere eget, vestibulum et, tempor auctor, justo. Fusce commodo aliquam arcu. Suspendisse enim turpis, dictum sed, iaculis a, condimentum nec, nisi.
            </p>
          </v-card-text>
        </v-card>
      </v-tab-item>
      <v-tab-item>
        <v-card flat>
          <v-card-text>
            <p>
              Fusce a quam. Phasellus nec sem in justo pellentesque facilisis. Nam eget dui. Proin viverra, ligula sit amet ultrices semper, ligula arcu tristique sapien, a accumsan nisi mauris ac eros. In dui magna, posuere eget, vestibulum et, tempor auctor, justo.
            </p>

            <p class="mb-0">
              Cras sagittis. Phasellus nec sem in justo pellentesque facilisis. Proin sapien ipsum, porta a, auctor quis, euismod ut, mi. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nam at tortor in tellus interdum sagittis.
            </p>
          </v-card-text>
        </v-card>
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