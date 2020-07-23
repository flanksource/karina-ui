<!--Demo Page

  -->

<template>
  <v-app id="inspire">
    <v-app-bar
      app
      color="indigo"
      dark
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Sysmonitor</v-toolbar-title>
    </v-app-bar>

    <v-main class="grey lighten-5">
      <v-container class="grey lighten-5" fluid>
        <v-row align="left" justify="center">
          <v-col cols="12" sm="9">

            <v-card class="pa-2" outlined tile>
             <v-banner two-line>
              <v-avatar
                slot="icon"
                color="indigo"
                dark
                size="40"
              >
                <v-icon
                  icon="mdi-lock"
                  color="white"
                >
                  mdi-information
                </v-icon>
              </v-avatar>

              <h3>Real-time CPU and Memory Usage</h3>

              <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
          
             </v-banner>
            </v-card>
          </v-col>
          <v-col cols="12"  sm="3" class="text-center">
            <MemoryCard/>
            <br>
            <CPUCard/>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <v-footer
      color="indigo"
      app
    >
      <span class="white--text">&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>

//Calls two components for memory usage and for CPU usage
import MemoryCard from './components/MemoryCard.vue'
import CPUCard from './components/CPUCard.vue'

export default {

  components:{
  MemoryCard, CPUCard
  },
    props: {
      source: String,
    },
    data: () => ({
      drawer: null,
    }),
  }
</script>