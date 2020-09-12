<template>

  <v-row 
    align="center"
    class="pl-5"
  >
    <v-col
      cols="2" xs="2" sm="2" md="2" lg="2" xl="2"
      class="pa-0"
    >
      <v-img
        :aspect-ratio="1/1"
        :src="require(`@/assets/svg/${iconlg}`)"
        width=32
      />   
    </v-col>

    <v-col
      cols="10" xs="10" sm="8" md="8" lg="8" xl="8"
      class="pa-0 bordered"
    >
      <v-row no-gutters >
        <span
          v-for="(indicatoricon, i) in indicatoricons"
          :key="i"
          class="ten-abreast" 
        >
          <IndicatorIcon
            :shape="indicatoricon.shape"
            :colour="indicatoricon.color"
            :overlay="indicatoricon.overlay"
          /> 
        </span>
      </v-row>
    </v-col>
  </v-row>

</template>

<script>

  import IndicatorIcon from './IndicatorIcon.vue'

  export default {
    name: 'IndicatorCard',

    components: {
      IndicatorIcon
    },

    props: {
      indicatoricons: Array,
      iconlg: String,
    },
  }

</script>

<style scoped>

  .ten-abreast {
    min-width: 10%;
  }

  .bordered {
    border-bottom: 1px dashed #111;
  }

</style>