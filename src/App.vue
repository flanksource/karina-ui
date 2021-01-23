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
        <v-row no-gutters>
          <v-col cols="1" class="left-pane">
            <Menu :clusters="clusters" @selectState="setState($event)"/>
          </v-col>

          <v-col cols="11" class="right-pane">

            <span
               v-if="state == 'iframe'"
            >
              <IframeView/>
            </span>
       
            <span 
              v-if="state == 'native'"
            > 
              <NativeView :clusters="clusters"/>
            </span> 
          </v-col>
        </v-row>
      </v-main>
    </div>
  </v-app>
</template>

<script>
  import IframeView from "./components/IframeView.vue";
  import Loader from "./components/Loader.vue";
  import Menu from "./components/Menu.vue";
  import NativeView from "./components/NativeView.vue";
  import stats from "./data/sample.json";

export default {
  components: {
    IframeView,
    Loader,
    Menu,
    NativeView,
  },

  data() {
    if (window.DEMO_MODE) {
      return {
        clusters: stats.clusters,
        loading: true,
        state: 'native'
      };
    } else {
      return {
        clusters: null,
        loading: true,
        state: 'native',
      };
    }
  },

  methods: {
    setState(state) {
      this.state = state;
      console.log(this.state)
    }
  },

  beforeCreate () {
    if (!window.DEMO_MODE) {
      /*
      this.$axios
      .get("http://localhost:8080/api/ui")
      .then(response => (this.clusters = response.data))
      .catch(error => console.log(error))
      .finally(() => (this.loading = false))  
       */
      this.$axios
        .get("/api/ui")
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


<style scoped>

  @import './assets/fonts/karina-ui-icons/icons.css';


  .left-pane {
    position: fixed;
    max-width: 8vw;
    height: 92vh;
  }

  .right-pane {
    margin-left: 8vw;
  }
</style>