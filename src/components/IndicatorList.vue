<template>
  <span>

       <div class="text-center d-flex pb-4">
      <v-btn @click="expand">
        expand all
      </v-btn>
      <v-btn @click="collapse">
        collapse all
      </v-btn>
    </div>

   <!--  <span v-for="(indicator, i) in indicators" :key="i"> -->

    <v-expansion-panels
      v-model="list"
      multiple
    >
      <v-expansion-panel
        v-for="(indicator,i) in indicators"
        :key="i"
      >
        <v-expansion-panel-header>
          <v-col cols="2" xs="2" sm="2" md="2" lg="2" xl="2" class="pa-0">
            <v-img
              :aspect-ratio="1 / 1"
              :src="require(`@/assets/svg/${indicator.indicator_icon_lg}`)"
              width="32"
            />
          </v-col>
          <v-col cols="10" xs="10" sm="8" md="8" lg="8" xl="8" class="pa-0 bordered">
            total
            total alerts

          </v-col>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <IndicatorCard
          
            :indicatoricons="indicator.indicator_icons"
          />
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>

     
<!--     </span>
 -->












    <span>
      <IndicatorCardCanary
        indicatoricon="birdie.svg"
        :canarychecks="canarychecks"
      />
    </span>






  </span>
</template>


<script>
import IndicatorCard from "./IndicatorCard.vue";
import IndicatorCardCanary from "./IndicatorCardCanary.vue";

export default {
  name: "IndicatorList",

  components: {
    IndicatorCard,
    IndicatorCardCanary,
  },

  props: {
    indicators: Array,
    canarychecks: Array,
  },

  data () {
    return {
      list: []
    }
  },

  methods: {
    expand () {
      this.list = [...this.indicators.keys()].map((k,i) => i)
    },

    collapse () {
      this.list = []
    },
  },
}

</script>


<style scoped>
.no-events {
  pointer-events: none;
}
</style>
