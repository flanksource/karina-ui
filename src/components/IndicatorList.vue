<template>


      <v-row dense>
        <v-col
          v-for="(item, i) in items"
          :key="i"
          cols="12"
        >
          <IndicatorCard/>
        </v-col>
      </v-row>
   

</template>


<script>

import IndicatorCard from './IndicatorCard.vue'
  export default {
    components: {
      IndicatorCard

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




