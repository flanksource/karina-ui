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
        <Menu :drawer="drawer" :clusters="clusters" @selectState="setState($event)"/>
        <span
          v-if="state == 'iframe'"
        >
          <IframeView class="ml-10"/>
        </span>
   
        <span 
          v-if="state == 'native'"
        > 
          <NativeView :clusters="clusters" class="ml-10"/>
        </span> 
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
    max-width: 8vw;
    height: 92vh;
  }

  .right-pane {
    margin-left: 8vw;
  }

  .v-application--is-ltr .v-toolbar__content > .v-btn.v-btn--icon:first-child + .v-toolbar__title, .v-application--is-ltr .v-toolbar__extension > .v-btn.v-btn--icon:first-child + .v-toolbar__title {
  padding-left: unset !important;
}
</style>