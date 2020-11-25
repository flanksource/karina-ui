<template> 
  <div class="pb-12">
    <v-row no-gutters>
      <v-col cols="12" xs="12" sm="2" md="2" lg="3" xl="3" class="px-2">
        <v-card :class="{ 'card-wide': expandUsageBar, 'card-narrow': !expandUsageBar  }">
          <v-btn
            class="btn-usage"
            @click="selectUsageBarMode"
            outlined
            x-small
          >
            <i :class="{ 'up': expandUsageBar, 'down': !expandUsageBar  }"> </i>
          </v-btn> 
          <div class="pt-4">

            <UsageBar metric="metric1" v-bind:value="45" v-bind:min="0" v-bind:max="100" v-bind:optimum="75" colour="#0f0" :expandUsageBar="expandUsageBar"/>
            <UsageBar metric="metric2" v-bind:value="80" v-bind:min="0" v-bind:max="100" v-bind:optimum="75" colour="#0f0" :expandUsageBar="expandUsageBar"/>
            <UsageBar metric="metric3" v-bind:value="25" v-bind:min="100" v-bind:max="0" v-bind:optimum="66" colour="#f00" :expandUsageBar="expandUsageBar"/>
            <UsageBar metric="metric4" v-bind:value="78" v-bind:min="100" v-bind:max="0" v-bind:optimum="66" colour="#f00" :expandUsageBar="expandUsageBar"/>
          </div>
        </v-card>
      </v-col>
      
      <v-col cols="12" xs="12" sm="3" md="3" lg="2" xl="2" class="item-list">
        <v-card flat class="ml-3 py-3 scroll">
          <ItemList :properties="cluster.properties" :name="cluster.name"/>
        </v-card>
      </v-col>

      <v-col
        cols="12"
        xs="12"
        sm="7"
        md="7"
        lg="6"
        xl="6"
        :class="{ 'expand-list': expandMode, 'collapse-list': !expandMode  }"

       
      >
        <v-card flat class="mr-4 py-3">
          <IndicatorList
            :canarychecks="cluster.canary_checks"
            :indicators="cluster.indicators"
            :expandMode="expandMode"
            :list="list"
            @expand="expansion"
          />

        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import IndicatorList from "./IndicatorList.vue";
import ItemList from "./ItemList.vue";
import UsageBar from "./UsageBar";
  
export default {
  name: "IndicatorPanel",

  components: {
    ItemList,
    IndicatorList,
    UsageBar,
  },

  props: {
    cluster: Object,
    bgColor: {
      type: String,
      default: "#0099CC"
    },
    height: {
      type: Number,
      default: 100
    },
    maxWidth: {
      type: Number,
      default: 150,
    },
  },

  data () {
    return {
      list: [],
      expandMode: false,
      expandUsageBar: false,
      demo: window.DEMO_MODE,
    }
  },

  methods: {
    expansion() {
      this.expandMode = !this.expandMode;
    },
    selectUsageBarMode() {
      this.expandUsageBar = !this.expandUsageBar;
    },
  }
};
</script>

<style scoped>
.item-list {
  border-top-left-radius: 20px !important;
  border-bottom-left-radius: 20px !important;
  min-height: 100%;
  border-right: none;
  border: 1px solid #444;
}

.expand-list {
  border-top-right-radius: 20px !important;
  border-bottom-right-radius: 20px !important;
  border-left: 2px solid #222;
  border: 1px solid #444;
}

.collapse-list {
  border-top-right-radius: 20px !important;
  border-bottom-right-radius: 20px !important;
  border-left: 2px solid #222;
  border: 1px solid #444;
  max-width: 150px !important;
}

.btn-usage {
  position: absolute;
  right: 6px;
  top: 5px;
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

.card-narrow {
  border: 1px solid #111;
  width: 50px;
  padding-top: 20px;
  height:100%;
}
.card-wide {
  border: 1px solid #111;
  height: 100%;
  border-bottom-right-radius: 25px;
}
</style>
