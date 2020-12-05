<template>
  <v-app>
    <v-app-bar app>
      <v-app-bar-nav-icon  @click.stop="drawer = !drawer"></v-app-bar-nav-icon>  

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

           <div class="left-pane">
            <left-nav-bar  :clusters="clusters" @selectState="setState($event)"/> 
          </div>

          <v-col cols="1" class="right-pane-menu">
            <Menu :clusters="clusters" @selectState="setState($event)"/>
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
  import LeftNavBar from "./components/LeftNavBar.vue";
  import Loader from "./components/Loader.vue";
  import Menu from "./components/Menu.vue";
  import NativeView from "./components/NativeView.vue";
  import stats from "./data/sample.json";

export default {
  components: {
    IframeView,
    LeftNavBar,
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
        drawer: null,
      };
    } else {
      return {
        clusters: null,
        loading: true,
        state: 'native',
        drawer: null,
      };
    }
  },

  methods: {
    setState(state) {
      this.state = state;
      this.drawer = !this.drawer;
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


<style scoped>
  
  .left-pane {
    position: fixed;
    max-width: 13vw;
    height: 92vh;
    z-index: 4; 
  }

  .right-pane-menu {
    position: fixed;
    max-width: 10vw;
    margin-left: 3vw;
    height: 92vh;
    z-index: 3;
  }

   .right-pane-content {
    margin-left: 13vw;
  }

  .v-application--is-ltr .v-toolbar__content > .v-btn.v-btn--icon:first-child + .v-toolbar__title, .v-application--is-ltr .v-toolbar__extension > .v-btn.v-btn--icon:first-child + .v-toolbar__title {
  padding-left: unset !important;
}
</style>