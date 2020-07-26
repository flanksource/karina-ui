<template>


 
          <v-card
            color="#fff"
            dark
          >
            <div class="d-flex flex-no-wrap justify-space-between">
                <v-avatar
                class="ma-3"
                size="64"
                tile
              >
                <v-img src="https://cdn.vuetifyjs.com/images/cards/foster.jpg"></v-img>
              </v-avatar>
            
              <v-layout row>
                      <v-item
          v-for="n in 42"
          :key="n"
    
        >
       
           <IndicatorIcon/>
        
        </v-item>
      </v-layout>
             

            
            </div>
          </v-card>
        

</template>


<script>

import IndicatorIcon from './IndicatorIcon.vue'
  export default {
    components: {
      IndicatorIcon

    },
    data: () => ({
       items: [
        {
          color: '#fff',
          icon: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          title: 'Supermodel',
          artist: 'Foster the People',
          indicator_icon: 'smile',
        },
        {
          color: '#fff',
          icon: 'https://cdn.vuetifyjs.com/images/cards/halcyon.png',
          title: 'Halcyon Days',
          artist: 'Ellie Goulding',
          indicator_icon: 'test'
        },
      ],
    }),
  }
</script>




