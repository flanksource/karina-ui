<template>

  <v-app>
    <v-app-bar app>
      <v-toolbar-title><b>Karina</b> - The Kubernetes Dashboard</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <v-container fluid>


          <v-row>
          <v-col
            cols="12" xs="12" sm="12" md="6" lg="6" xl="6"
            v-for="(apicluster, i) in apiclusters"
            :key="i"
          >
            <v-row no-gutters>
              <v-col cols="12" xs="12" sm="12" md="12" lg="12" xl="12">
                <IndicatorPanel
                  :cluster="apicluster"
                  :itemicons="apicluster.itemicons"
                />
              </v-col>
            </v-row>
          </v-col>         
        </v-row>

<hr>
<span> Old Interface </span>
<hr>

        <v-row>
          <v-col
            cols="12" xs="12" sm="12" md="6" lg="6" xl="6"
            v-for="(cluster, i) in clusters"
            :key="i"
          >
            <v-row no-gutters>
              <v-col cols="12" xs="12" sm="12" md="12" lg="12" xl="12">
                <IndicatorPanel
                  :cluster="cluster"
                  :itemicons="cluster.itemicons"
                />
              </v-col>
            </v-row>
          </v-col>         
        </v-row>

      </v-container>
    </v-main>
  </v-app>

</template>


<script>

  import IndicatorPanel from './components/IndicatorPanel.vue'

  import stats from './data/sample.json'

  export default {

    components:{
      IndicatorPanel,
    },

    data () {
      return {
        apiclusters: null,
        clusters: stats.clusters,
        canary: { 
          "id": 3,
          "indicator_icon_lg": "birdie.svg",
          "indicator_icons": [
            {"id":1,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":2,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":3,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":4,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":5,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":6,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":7,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":8,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":9,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":10,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":11,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":12,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":13,"color":"#994c00", "shape":"heart", "overlay":""},
            {"id":14,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":15,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":16,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":17,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":18,"color":"#990000", "shape":"heart", "overlay":""},
            {"id":19,"color":"#336600", "shape":"heart", "overlay":""},
            {"id":20,"color":"#336600", "shape":"heart", "overlay":""}
          ]
        },

      }
    },

    mounted() {
      this.$axios
        .get('http://localhost:8080/api')
        .then(response => (this.apiclusters = response.data))
        .catch(error => console.log(error))  
    }
  }

</script>
