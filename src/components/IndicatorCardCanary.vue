<template>

  <v-row 
    align="center"
    class="pl-5"
  >
    <v-col
      cols="2" xs="2" sm="2" md="2" lg="2" xl="2"
      class="pa-0"
    >

      <v-img
        :aspect-ratio="1/1"
        :src="require(`@/assets/svg/${indicatoricon}`)"
        width=32
      />   
    </v-col>

    <v-col
      cols="10" xs="10" sm="8" md="8" lg="8" xl="8"
      class="pa-0 bordered"
    >

      <v-row no-gutters >
        <span
          v-for="(canarycheck, i) in canarychecks"
          :key="i"
          class="ten-abreast" 
        >

     

         <IndicatorIcon
                shape="heart"
                colour="#336600"
                overlay=""
                class="clickable-icon"
                @click.native="selectCheck(canarycheck)"
              /> 
  <br>
     
 
          <span
            v-for="(check, i) in canarycheck.checkStatuses"
            :key="i"
            class="ten-abreast" 
          >       

          {{ check }} 

          </span>  

          <hr> <br> 
    
         <!--   <span

                v-if="check.pass == true"
            > -->
            <!-- <span>
              <IndicatorIcon
                shape="heart"
                colour="#336600"
                overlay=""
                class="clickable-icon"
                @click.native="selectCheck(check)"
              /> 

            </span>
           -->

        <!--     <span
                v-else
            >
              <IndicatorIcon
                shape="heart"
                colour="#990000"
                overlay=""
                class="clickable-icon"
                @click.native="selectCheck(canarycheck)"
              /> 

            </span>

        </span> -->
         
            <v-dialog
              v-model="dialog"
              max-width="500"
            >
              <v-card>

       


                 <v-card-title class="headline">{{selectedCheck.name}}</v-card-title>
                <v-card-text>
                  Type: {{selectedCheck.type}}<br/>
                  Description: {{selectedCheck.description}}<br/>
                  Uptime: {{selectedCheck.uptime}}<br/>
                  Latency: {{selectedCheck.latency}}<br/>
                  Statuses: {{selectedCheck.checkStatuses}}<br/>
                 
                </v-card-text>


              </v-card>
            </v-dialog>


        </span>
      </v-row>
    </v-col>
  </v-row>

</template>

<script>

  import IndicatorIcon from './IndicatorIcon.vue'

  export default {
    name: 'IndicatorCard',

    components: {
      IndicatorIcon
    },

    props: {
      canarychecks : Array,
      canarycheck: Object,
      indicatoricon: String,
      item: String,
     // selectedCheck: Object,
    },

    data () {
      return {
        dialog: false,
        selectedCheck: {},

      }
    },

    methods: {
      selectCheck(canarycheck) {
        this.selectedCheck = canarycheck
        this.dialog = true
      },
    }


  }

</script>

<style scoped>

  .ten-abreast {
    min-width: 10%;
  }

  .bordered {
    border-bottom: 1px dashed #111;
  }

.clickable-icon {
  cursor: pointer;
}

</style>




