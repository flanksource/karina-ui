<template>
  <div>
      <v-btn
      class="btn-expand"
      @click="selectMode"
      outlined
      x-small
    >
      <i :class="(listExpand? 'up' : 'down')+' outlined'"> </i>
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
<!--              <v-img-->
<!--                :aspect-ratio="1 / 1"-->
<!--                :src="require(`@/assets/svg/${indicator.indicator_icon_lg}`)"-->
<!--                width="48"-->
<!--                class="icon"-->
<!--              />-->
              <i :class="'icon-'+indicator.indicator_icon_lg+' icon large-icon'"></i>

              <span 
                v-if="getDemoAlerts(indicator.indicator_icons) > 0 && !listExpand"
              >
                <v-badge
                  :content="getDemoAlerts(indicator.indicator_icons)"
                  color="#990000"
                  offset-x=-50
                  offset-y=5
                >
                </v-badge>
              </span>
                
              <span
                v-if="getIndicators(indicator.indicator_icons) > 0 && !listExpand"
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

          <v-col cols="10" xs="10" sm="10" md="10" lg="9" xl="9" class="ml-1">
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
    <v-expansion-panels
      v-model="list"
      multiple
      flat
    >
          <v-expansion-panel>
          <v-row no-gutters align="center" class="pl-3">
             <v-col cols="2" xs="2" sm="2" md="2" lg="2" xl="2" class="">
            <div class="parent">
               <i class="icon-birdie icon large-icon"></i>
<!--              <v-img-->
<!--                :aspect-ratio="1 / 1"-->
<!--                :src="require(`@/assets/svg/birdie.svg`)"-->
<!--                width="48"-->
<!--                class="icon"-->
<!--              />-->

            <span 
                v-if="getCanaryAlerts(canarychecks) > 0 && !listExpand"
              >
                <v-badge
                  :content="getCanaryAlerts(canarychecks)"
                  color="#990000"
                  offset-x=-50
                  offset-y=5
                >
                </v-badge>
              </span>
                
              <span
                v-if="getIndicators(canarychecks) > 0 && !listExpand"
              >
                <v-badge
                  :content="getIndicators(canarychecks)"
                  color="#336600"
                  offset-x=-50
                  offset-y=30
                >
                </v-badge>
              </span>
            </div>
          </v-col>
          <v-col cols="10" xs="10" sm="10" md="10" lg="9" xl="9" class="ml-1">
            <v-expansion-panel-content class="regulate-padding bordered">

               <IndicatorCardCanary
               
                :canarychecks="canarychecks"
              />
            </v-expansion-panel-content>
            
          </v-col>
        </v-row>
      </v-expansion-panel>
    </v-expansion-panels>
  </span>

  </div>
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
    expandMode: Boolean
  },

  data () {
    return {
      list: [],
      listExpand: this.expandMode,
      demo: window.DEMO_MODE,
    }
  },

  methods: {
    selectMode() {
      this.listExpand = !this.listExpand;
      if (this.listExpand == true) {
        if (!this.demo) {
          this.list = [...this.canarychecks.keys()].map((k,i) => i)
        } else {
          this.list = [...this.indicators.keys()].map((k,i) => i)
        }
      } else {
        this.list = []
      }
      this.$emit('expand');
    },

    getIndicators(icons) {
      var count = 0;
      if (icons) {
        count = icons.length;
      }
      return count;
    },


    getDemoAlerts(icons) {
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

    getCanaryAlerts(icons) {
      var sum = 0
      if (icons) {
        for (var i=0; i<icons.length; i++) {
          if (icons[i].checkStatuses[0].Status != true) {
            sum = sum + 1;
          }
        }
      }
      return sum;
    }
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
  left: 25%;
  position: absolute;
  width: 64%;
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
i.outlined {
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

@import '../assets/fonts/karina-ui-icons/icons.css';

.icon-birdie:before {
  color: #ffd838 !important;
  font-size: 48px;
}
.icon-heartline:before {
    color: #2f57d3 !important;
}

.icon-server:before {
    color: #12378d !important;
}



.large-icon {
  font-size: 48px;
}
</style>