<template>
  <v-app id="inspire">
    <v-app-bar app color="indigo" dark>
      <v-toolbar-title><b>Karina</b> - The Kubernetes Dashboard</v-toolbar-title>
    </v-app-bar>

    <v-main class="grey lighten-5">
      <v-container class="grey lighten-5" fluid>
        <v-row align="left" justify="center">
          <v-col cols="12" sm="9">
            <v-row>
              <v-col cols="12" sm="8">
                   <IndicatorPanel/>
              </v-col>
            </v-row>
          </v-col>

          <v-col cols="12"  sm="3" class="text-center">
            <MemoryCard/>
            <br>
            <CPUCard/>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <v-footer color="indigo" app>
      <span class="white--text">&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>

  import MemoryCard from './components/MemoryCard.vue'
  import CPUCard from './components/CPUCard.vue'
  import IndicatorPanel from './components/IndicatorPanel.vue'

  export default {
    components:{
      MemoryCard, CPUCard, IndicatorPanel
    },
  }
</script>