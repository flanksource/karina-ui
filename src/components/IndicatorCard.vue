<template>
  <v-card flat class="fit-parent px-3">
    <v-row align="center" no-gutters>
      <v-col cols="12" md="2">
     

    <img :src="require(`@/assets/svg/${iconlg}`)" width=48/>
      

      </v-col>

      <v-col cols="12" md="10" class="card-bottom">
        <v-row>
          <v-col
            v-for="(item, i) in indicatoricons"
            :key="i"
            cols="12" md="1" class="ten-cols"
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
      indicatoricons: Object,
      iconlg: String,
    },
  }
</script>

<style>
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