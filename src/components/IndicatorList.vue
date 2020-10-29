<template>
  <span>
    <v-btn
      class="btn-expand"
      @click="expansion"
      outlined
      x-small
    >
      <i :class="{ 'up': expand, 'down': !expand  }"> </i>
    </v-btn>  
        
    <v-expansion-panels
      v-model="list"
      multiple
      flat
    >
      <v-expansion-panel
        v-for="(indicator,i) in indicators"
        :key="i"
      >
        <v-row no-gutters align="center" class="pl-3">
          <v-col cols="2" xs="2" sm="2" md="2" lg="2" xl="2" class="">
            <div class="parent">
              <v-img
                :aspect-ratio="1 / 1"
                :src="require(`@/assets/svg/${indicator.indicator_icon_lg}`)"
                width="48"
                class="icon"
              />
            
              <span 
                v-if="getAlerts(indicator.indicator_icons) > 0"
              >
                <v-badge
                  :content="getAlerts(indicator.indicator_icons)"
                  color="#990000"
                  offset-x=-50
                  offset-y=5
                >
                </v-badge>
              </span>
                
              <span
                v-if="getIndicators(indicator.indicator_icons) > 0"
              >
                <v-badge
                  :content="getIndicators(indicator.indicator_icons)"
                  color="#336600"
                  offset-x=-50
                  offset-y=30
                >
                </v-badge>
              </span>
            </div>
          </v-col>

          <v-col cols="10" xs="10" sm="10" md="10" lg="9" xl="9" class="ml-4">
            <v-expansion-panel-content class="regulate-padding bordered">
              <IndicatorCard
                :indicatoricons="indicator.indicator_icons"
              />
            </v-expansion-panel-content>
          </v-col> 
        </v-row>
      </v-expansion-panel>
    </v-expansion-panels>
    <span
      v-if="!demo"
    >
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
      expand: false,
      demo: window.DEMO_MODE,
    }
  },

  methods: {
    expansion() {
      this.expand = !this.expand;
      if (this.expand == true) {
        this.list = [...this.indicators.keys()].map((k,i) => i)
      } else {
        this.list = []
      }
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
          if (icons[i].color == '#990000' || icons[i].color == '#994c00'){
            sum = sum + 1;
          }
        }
      }
       return sum;
    },
  },
}
</script>

<style lang="css" scoped>
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
  height: 1px;
  left: 30%;
  position: absolute;
  width: 60%;
}
.text-size {
  font-size: 12px;
}
.parent {
  position: relative;
  top: 0;
  left: 0;
  width: 48px;
  height: 48px;
}
.icon {
  position: absolute;
  top: 0;
  left: 0;
}
.alert-overlay {
  position: absolute;
  top: 2px;
  right: -18px;
  font-size: 12px;
  color: #777;
}
.indicator-overlay {
  position: absolute;
  bottom: 2px;
  right: -18px;
  font-size: 12px;
  color: #777;
}
.regulate-padding {
  padding-top: 16px;
  padding-left: 0px !important;
}
.btn-expand {
  position: absolute;
  right: 3px;
  z-index: 3;
}
i {
  border: solid black;
  border-width: 0 3px 3px 0;
  display: inline-block;
  padding: 3px;
}
.down {
  transform: rotate(45deg);
  margin-top: -3px;
}
.up {
  transform: rotate(-135deg);
  margin-top: 2px;
}
.border{
  border: 2px solid #0f0;
  margin-left: 15px;
}
</style>