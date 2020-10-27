<template>
  <span>
    <v-expansion-panels
      v-model="list"
      multiple
      flat
    >
      <v-expansion-panel
        v-for="(indicator,i) in indicators"
        :key="i"
      >
        <v-row no-gutters align="">
          <v-col cols="10" xs="10" sm="10" md="10" lg="10" xl="10" class="pa-0">
            <v-expansion-panels flat class="bordered">
              <v-expansion-panel class="">
                <v-expansion-panel-header class="pl-2">

                  <v-col cols="2" xs="2" sm="2" md="2" lg="2" xl="2" class="pa-0">
                    <v-img
                      :aspect-ratio="1 / 1"
                      :src="require(`@/assets/svg/${indicator.indicator_icon_lg}`)"
                      width="32"
                    />
                  </v-col>
                </v-expansion-panel-header>

                <v-expansion-panel-content class="pl-9 pb-0">
                  <IndicatorCard
                    :indicatoricons="indicator.indicator_icons"
                  />
                </v-expansion-panel-content>
              </v-expansion-panel>
            </v-expansion-panels>
          </v-col> 
          <v-col cols="2" xs="2" sm="2" md="2" lg="1" xl="1" class="px-0">
            <span 
              class="d-block text-size alerts"
              v-if="getAlerts(indicator.indicator_icons) > 0"
            >
              {{ getAlerts(indicator.indicator_icons) }}</span>
              
            <span class="d-block text-size indicators"
            v-if="getIndicators(indicator.indicator_icons) > 0"
            >{{ getIndicators(indicator.indicator_icons) }}</span>
          </v-col>
        </v-row>
      </v-expansion-panel>
    </v-expansion-panels>

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
    indicator: Array,
  },

  data () {
    return {
      list: [],
    }
  },

  methods: {
    expand () {
      this.list = [...this.indicators.keys()].map((k,i) => i)
    },

    collapse () {
      this.list = []
    },
    getIndicators(icons) {
      var count = 0;
      if (icons) {
        count = icons.length;
      }
      return count;
    },
    getAlerts(icons) {
      
      var sum = 0
      if (icons) {
        for (var i=0; i<icons.length; i++) {
         // console.log(icons[i].color);
          
          if (icons[i].color == '#990000' || icons[i].color == '#994c00'){
            sum = sum + 1;
            console.log(sum)
          }
        }
      }
       return sum;
    },
  },
}

</script>


<style scoped>
  .no-events {
    pointer-events: none;
  }
  .v-expansion-panels .v-expansion-panel {
    margin-top: 0 !important;
  }
  .v-expansion-panel-header {
     padding: 0 24px 0 24px;
     min-height: 30px;
  }
  .v-expansion-panel-header:active {
     padding: 0 24px 0 24px;
     min-height: 30px;
  }
  .v-expansion-panel:active > .v-expansion-panel-header {
    min-height: 40px;
  }
 .v-expansion-panel > .v-expansion-panel-header {
    min-height: 40px;
  }
  .bordered::after {
    border: 1px dashed #111 ; 
    bottom: 0;
    content: '';
    display: block;
    height: 2px;
    left: 20%;
    position: absolute;
    width: 80%;
  }
  .text-size {
    font-size: 12px;
  }
  .alerts {
    margin-top: -1px;
    margin-left: -20px;
    position: absolute;
    z-index: 222;
  }
  .indicators {
    margin-top: 20px;
    margin-left: -20px;
    position: absolute;
    z-index: 222;
  }
</style>
