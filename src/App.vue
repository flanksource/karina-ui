<template>
  <v-app>
  

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
          <v-col cols="1" :class="{ 'left-pane-collapsed': collapsedLNB, 'left-pane-expanded': !collapsedLNB  }">
            <Menu :clusters="clusters" @selectState="setState($event)" @selectLNBState="setLNBState()" :collapsedLNB="collapsedLNB"/>
          </v-col>

          <v-col cols="10" class="right-pane-content">
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
 // import LeftNavBar from "./components/LeftNavBar.vue";
  import Loader from "./components/Loader.vue";
  import Menu from "./components/Menu.vue";
  import NativeView from "./components/NativeView.vue";
  import stats from "./data/sample.json";

export default {
  components: {
    IframeView,
    //LeftNavBar,
    Loader,
    Menu,
    NativeView,
  },

  data() {
    if (window.DEMO_MODE) {
      return {
        clusters: stats.clusters,
        loading: true,
        state: 'native',
        collapsedLNB: true,
      };
    } else {
      return {
        clusters: null,
        loading: true,
        state: 'native',
        collapsedLNB: true,
      };
    }
  },

  methods: {
    setState(state) {
      this.state = state;
    },
    setLNBState() {
      this.collapsedLNB = !this.collapsedLNB;
    },
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

<style scoped>
  
  .left-pane {
    position: fixed;
    max-width: 13vw;
    height: 92vh;
    z-index: 4; 
  }

  .right-pane-menu {
    position: fixed;
    width: 50px;
    margin-left: 0vw;
    height: 92vh;
    z-index: 3;
  }

  .left-pane-collapsed {
    position: fixed;
    width: 50px;
    margin-left: 0vw;
    height: 192vh;
    z-index: 333;

  }

  .left-pane-expanded{
    position: fixed;
    max-width: 220px;
    margin-left: 0vw;
    height: 192vh;
    z-index: 333;
  }

  .right-pane-content {
    margin-left: 50px;
  }

</style>