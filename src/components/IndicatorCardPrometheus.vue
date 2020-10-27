<template>
  <v-row align="center" class="pl-5">
    <v-col cols="2" xs="2" sm="2" md="2" lg="2" xl="2" class="pa-0">
      <v-img
        :aspect-ratio="1 / 1"
        :src="require(`@/assets/svg/${indicatoricon}`)"
        width="32"
      />
    </v-col>
    
    <v-col cols="10" xs="10" sm="8" md="8" lg="8" xl="8" class="pa-0 bordered">
      <v-row no-gutters>
        <span
          v-for="(prometheusCheck, i) in prometheusChecks"
          :key="i"
          class="ten-abreast"
        >
         <span v-if="prometheusCheck.labels.severity == 'critical'">
            <IndicatorIcon
              shape="square"
              colour="#990000"
              overlay=""
              class="clickable-icon"
              v-bind:title="prometheusCheck.labels.alertname"
              @click.native="selectCheck(prometheusCheck)"
            />
          </span>

          <span v-else-if="prometheusCheck.labels.severity == 'warning'">
            <IndicatorIcon
              shape="square"
              colour="#f68c1f"
              overlay=""
              class="clickable-icon"
              v-bind:title="prometheusCheck.labels.alertname"
              @click.native="selectCheck(prometheusCheck)"
            />
          </span>

          <span v-else-if="prometheusCheck.labels.severity == 'info'">
            <IndicatorIcon
              shape="square"
              colour="#eae229"
              overlay=""
              class="clickable-icon"
              v-bind:title="prometheusCheck.labels.alertname"
              @click.native="selectCheck(prometheusCheck)"
            />
          </span>

          <span v-else-if="prometheusCheck.labels.severity == 'none'">
            <IndicatorIcon
              shape="square"
              colour="#336600"
              overlay=""
              class="clickable-icon"
              v-bind:title="prometheusCheck.labels.alertname"
              @click.native="selectCheck(prometheusCheck)"
            />
          </span>
        </span> 
         <v-dialog v-model="dialog" max-width="50%">
            <v-card>
              <v-card-title class="headline">
                  {{ selectedCheck.labels.alertname }}
              </v-card-title>

              <v-card-text>
                <p>
                  {{ selectedCheck.annotations.message }}
                </p>
                Since: {{ selectedCheck.activeat }}<br/>
                Alert Level: {{ selectedCheck.labels.severity }}<br/>  
              </v-card-text>
            </v-card>
          </v-dialog>
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
  import IndicatorIcon from "./IndicatorIcon.vue";

  export default {
    name: "IndicatorCard",

    components: {
      IndicatorIcon,
    },

    props: {
      prometheusChecks: Array,
      indicatoricon: String,
    },

    data() {
      return {
        dialog: false,
        selectedCheck: {
          labels: {},
          annotations:{},
        },
      };
    },

    methods: {
      selectCheck(prometheusCheck) {
        this.selectedCheck = prometheusCheck;
        this.dialog = true;
      },
    },
  };
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
