<template>


       <v-badge
 color="deep-purple accent-4"
 :content="content"
 inline="true"
>
<v-list-item-icon>

    <svg-icon :icon="icon" :hasFill="false" :color="color" class="smile-icon">
          </svg-icon>

        
         </v-list-item-icon>
   
  <v-list-item-content>
           <v-list-item-title v-text="label"></v-list-item-title>
               
         </v-list-item-content>
</v-badge>  


</template>



<script>

  import SvgIcon from './SVGIcon.vue';


export default {

  components: {
    'svg-icon': SvgIcon,
  },
    props: {
      source: String,
      icon: String,
      label: String,
      content: String,
    },
    data: () => ({
      drawer: null,
    }),
  }
</script> 


<!-- <script>
export default {
  props: ['name']
}
</script>
 -->

