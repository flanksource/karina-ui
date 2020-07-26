<template>
<!--   <v-badge
 color="deep-purple accent-4"
 content="4"
 inline="true"
>
  <v-icon
    :color="indigo"
    class="mr-3"
    size="24"
  >
    {{icon}}
  </v-icon>
   
  {{title}}
</v-badge>   
           -->
 <v-layout row wrap>
       <v-badge
 color="deep-purple accent-4"
 :content="content"
 inline="true"
>
<v-list-item-icon>
           <v-icon v-text="icon"></v-icon>
         </v-list-item-icon>
   
  <v-list-item-content>
           <v-list-item-title v-text="label"></v-list-item-title>
               
         </v-list-item-content>
</v-badge>  
         <!--  <v-list-item-icon>
           <v-icon v-text="icon"></v-icon>
         </v-list-item-icon>
             
         
         <v-list-item-content>
           <v-list-item-title v-text="title"></v-list-item-title>
               
         </v-list-item-content> -->
</v-layout>

</template>



<script>

export default {

  components:{

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

