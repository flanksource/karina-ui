<template>
  <v-card flat class="fit-parent px-3">
    <v-row align="center" no-gutters>
      <v-col cols="12" md="2">
        <img v-bind:src="iconlg" alt="" width=48/>
      </v-col>

      <v-col cols="12" md="10" class="card-bottom">
        <v-row>
          <v-col
            v-for="(indicatoricon, i) in indicatoricons"
            :key="i"
            cols="12" md="1" class="ten-cols"
          >
            <IndicatorIcon :shape="indicatoricon.shape" :colour="indicatoricon.color" :overlay="indicatoricon.overlay"/> 
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
  import IndicatorIcon from './IndicatorIcon.vue'

  export default {
    name: 'IndicatorCard',
    components: {
      IndicatorIcon
    },
    props: {
      indicatoricons: Object,
      iconlg: String,
    },
  }
</script>

<style scoped>
  .card-bottom  {
    border-bottom: 1px dashed black !important;
  }
  .ten-cols {
    min-width: 10%;
  }
  .fit-parent{
    width: 100%;
    padding-top: 10px;
  }
</style>