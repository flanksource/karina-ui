<template>
  <v-card flat class="righty tops">
    <v-row align="center" no-gutters>
      <v-col
        cols="12" md="2"
      >
        <img v-bind:src="iconlg" alt="" width=48/>
      </v-col>
      <v-col
        cols="12" md="10" class="indie"
      >
        <v-row align="center" no-gutters>

        
        <v-col
          v-for="(item, i) in indicatorx"
          :key="i"
          cols="12" md="1" class="icona"
        >


          <IndicatorIcon :shape="item.shape" :colour="item.color" :overlay="item.overlay"/> 

      </v-col>
      

  </v-row>

  </v-col>
    </v-row>

  </v-card>
</template>


<script>

import IndicatorIcon from './IndicatorIcon.vue'
  export default {
    components: {
      IndicatorIcon

    },
    props: {

      indicatorx: Object,
      iconlg: String,
    },
    computed: {
      getIcon: function (icon) {
  return require(icon)
}

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

 .icona {
    min-width: 10%;
}

.righty{
  margin: 10px 10px 10px 10px;
}

</style>


