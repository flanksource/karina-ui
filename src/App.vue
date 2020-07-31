<!--Demo Page

  -->

<template>
  <v-app id="inspire">
    <v-app-bar app color="indigo" dark>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title><b>Karina</b> - Kubernetes Dashboard</v-toolbar-title>
    </v-app-bar>

    <v-main class="grey lighten-5">
      <v-container class="grey lighten-5" fluid>
        <v-row align="left" justify="center">
          <v-col cols="12" sm="9">
            <div>
              <v-card class="pa-2" outlined tile>
                <v-row>
                  <ItemIcon icon="git" colour="#f8cecc" label="Change Config" count="5"/>
                  <ItemIcon icon="git" colour="#666666" label="Downgrade Version" count="2"/>
                  <ItemIcon icon="git" colour="#d5e8d4" label="Bump Replicas" count="0"/>
                </v-row>
              </v-card>
            </div>

            <v-row >
              <v-col cols="12" sm="4">
                <v-card class="" max-width="" height="410">
                  <v-card-text>
                    <div> Lorem ipsum dolor</div>
                    
                    <p class="display-1 text--primary">
                       Lorem ipsum
                    </p>

                    <p>Lorem ipsum dolor</p>

                    <div class="text--primary">
                      Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non...
                    </div>
                  </v-card-text>

                  <v-card-actions>
                    <v-btn text color="deep-purple accent-4">
                      Learn More
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-col>

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
  import ItemIcon from './components/ItemIcon.vue'

  export default {
    components:{
      MemoryCard, CPUCard, IndicatorPanel, ItemIcon
    },
  }
</script>