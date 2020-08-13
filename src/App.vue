<template>
  <v-app>
    <v-app-bar app>
      <v-toolbar-title><b>Karina</b> - The Kubernetes Dashboard</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <v-row>
          <v-col cols="12" xs="12" sm="12" md="6" lg="6" xl="6">
            <v-row no-gutters>
              <v-col cols="12" xs="12" sm="12" md="12" lg="12" xl="12">
                <IndicatorPanel/>
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

  export default {
    components:{
     IndicatorPanel
    },
  }
</script>
