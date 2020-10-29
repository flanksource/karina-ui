<template>

  <v-app>
    <v-app-bar app>
      <v-toolbar-title
        ><b>Karina</b> - The Kubernetes Dashboard</v-toolbar-title
      >
    </v-app-bar>

    <div
      v-if="loading == true"
    >
      <Loader/>
    </div>
    
    <div
      v-else-if="loading == false"
    >
      <v-main>
        <v-container fluid>
          <v-row 
            v-if="clusters==null"
            class="fill-height mt-10"
            align-content="center"
            justify="center"
          >
            <v-col
              class="subtitle-1 text-center"
              cols="12"
            >
             No Data Found!
            </v-col>
          </v-row>

          <v-row v-else>
            <v-col
              cols="12"
              xs="12"
              sm="12"
              md="6"
              lg="6"
              xl="6"
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
    </div>
  </v-app>

</template>

<script>
import IndicatorPanel from "./components/IndicatorPanel.vue";
import Loader from "./components/Loader.vue";
import stats from "./data/sample.json";

export default {
  components: {
    IndicatorPanel,
    Loader,
  },

  data() {
    if (window.DEMO_MODE) {
      return {
        clusters: stats.clusters,
        loading: true,
      };
    } else {
      return {
        clusters: null,
        loading: true,
      };
    }
  },

  beforeCreate () {
    if (!window.DEMO_MODE) {
      /*
      this.$axios
      .get("http://localhost:8080/api")
      .then(response => (this.clusters = response.data))
      .catch(error => console.log(error))
      .finally(() => (this.loading = false))  
       */
      this.$axios
        .get("/api")
        .then(response => (this.clusters = response.data))
        .catch(error => console.log(error))
        .finally(() => (this.loading = false)); 

    } else {
      
      setTimeout(() => (this.loading = false)
        .bind(this), (3 * 1000))
        .catch(error => console.log(error));
    }
  }
};
</script>