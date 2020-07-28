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
      <v-toolbar-title><b>Karina</b> - Kubernetes Dashboard</v-toolbar-title>
    </v-app-bar>

    <v-main class="grey lighten-5">
      <v-container class="grey lighten-5" fluid>
        <v-row align="left" justify="center">
          <v-col cols="12" sm="9">
            <v-card class="pa-2" outlined tile>
          
                 <v-layout row>
              <ItemIcon icon="mdi-github" colour="#f8cecc" label="Change Config" count="5"/>
               <ItemIcon icon="mdi-github" colour="#666666" label="Downgrade Version" count="2"/>
             <ItemIcon icon="mdi-github" colour="#d5e8d4" label="Bump Replicas" count="2"/>
           </v-layout>
            </v-card>
            <br/>

            <IndicatorPanel/>
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
import IndicatorPanel from './components/IndicatorPanel.vue'
import ItemIcon from './components/ItemIcon.vue'

export default {

  components:{
  MemoryCard, CPUCard, IndicatorPanel, ItemIcon
  },
    props: {
      source: String,
    },
    data: () => ({
      drawer: null,
    }),
  }
</script>