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

                <v-expansion-panel-content class="pl-9">
         

                  <IndicatorCard
                    :indicatoricons="indicator.indicator_icons"
                  
                  />
               
                </v-expansion-panel-content>
              </v-expansion-panel>
            </v-expansion-panels>
          </v-col> 
 
            <v-col cols="2" xs="2" sm="2" md="2" lg="1" xl="1" class="px-0">
          
             <span class="d-block">{{ getIndicators(indicator.indicator_icons) }}</span>

             <span class="d-block">{{ getIndicators(indicator.indicator_icons) }}</span>

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
      var count = 0;
      if (icons) {
        count = icons.length;
      }
      return count;
    },
  },
}

</script>


<style scoped>
  .no-events {
    pointer-events: none;
  }
  .alerts {
    color:#4a4a4a;
    float: right;
    margin-top: 0px;
    margin-right: 10px;
  }

  .indicators {
    color: #4a4a4a;
    float: right;
 
    margin-top: 25px;
    margin-right: 0px;

  }

  .right-text{
    border: 2px solid #00f;
    padding: 0;
    align-items: unset;
    display: block;
    position: unset;
    min-height: 50px;
    justify-content: flex-end;
    
  }

.v-expansion-panels .v-expansion-panel {
    margin-top: 0 !important;


  }

  .bordered {
    
  }

  .bordered::after {

    border: 1px dashed #111 ; 
    bottom: 0;
    content: '';
    display: block;
    height: 2px;
    left: 20%;
    position: absolute;

    width: 73%;

  }
/*for testing*/
.border{
  border: 2px solid #00f;
}

</style>
