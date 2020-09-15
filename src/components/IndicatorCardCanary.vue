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
          v-for="(checktype, i) in checktypes"
          :key="i"
          class="ten-abreast" 
        >

          <span
            v-for="(check, i) in checktype.status"
            :key="i"
            class="ten-abreast" 
          >           

            <span
                v-if="check.pass == true"
            >
              <IndicatorIcon
                shape="heart"
                colour="#336600"
                overlay=""
                class="clickable-icon"
                @click.native="selectCheck(checktype)"
              /> 

            </span>

            <span
                v-else
            >
              <IndicatorIcon
                shape="heart"
                colour="#990000"
                overlay=""
                class="clickable-icon"
                @click.native="selectCheck(checktype)"
              /> 

            </span>

        </span>
         
            <v-dialog
              v-model="dialog"
              max-width="290"
            >
              <v-card>

       


                 <v-card-title class="headline"></v-card-title>
                <v-card-text>
                   {{selectedChecktype}}

      <!--              {{ checktype.name }}

                  Namespace: {{ checktype.namespace }}<br/>
                  Type: {{ checktype.type }}<br/>
                  Endpoint: {{ checktype.endpoint }}<br/>
                  Latency1H: {{ checktype.latency_1_h }}<br/>
                  Uptime1H: {{ checktype.uptime_1_h }}<br/>
                  Message: {{ checktype.status.message }}<br/>  -->
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
      checktypes: Array,
      indicatoricon: String,
      item: String,
      selectedChecktype: Object,
    },

    data () {
      return {
        dialog: false,
      }
    },

    methods: {
      selectCheck(checktype) {
        this.selectedChecktype = checktype
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