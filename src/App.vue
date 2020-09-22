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

 // var DEMO_MODE = true;

  export default {

    components:{
      IndicatorPanel,
    },

    data () {
      if ( window.DEMO_MODE ){
        return {
          clusters: stats.clusters,
         
        }
      }
      else{
        return {
         clusters: null 
        }
      }
    },

    mounted() {

      if ( !window.DEMO_MODE ) {
       /* this.$axios
          .get('http://localhost:8080/api')
          .then(response => (this.clusters = response.data))
          .catch(error => console.log(error))      */

        this.$axios
          .get('/api')
          .then(response => (this.clusters = response.data))
          .catch(error => console.log(error)) 
      }

    }
  }

</script>
