<template>
 <!--  <v-row align="center" class="pl-5"> -->
   <!--  <v-col cols="2" xs="2" sm="2" md="2" lg="2" xl="2" class="pa-0">
      <v-img
        :aspect-ratio="1 / 1"
        :src="require(`@/assets/svg/${iconlg}`)"
        width="32"
      />
    </v-col>

    <v-col cols="10" xs="10" sm="8" md="8" lg="8" xl="8" class="pa-0 bordered"> -->
      <v-row no-gutters>
        <span
          v-for="(indicatoricon, i) in indicatoricons"
          :key="i"
          class="ten-abreast"
        >
          <IndicatorIcon
            :shape="indicatoricon.shape"
            :colour="indicatoricon.color"
            :overlay="indicatoricon.overlay"
            class="clickable-icon"
            @click.native="selectCheck()"
          />
        </span>

        <v-dialog v-model="dialog" max-width="50%">
          <v-card>
            <v-card-title class="headline"> Check Name</v-card-title>

            <v-card-text>
              <v-row no-gutters>
                <v-col class="ml-0"> <b>Type</b>: </v-col>

                <v-col> <b>Uptime</b>: </v-col>

                <v-col> <b>Latency</b>: </v-col>
              </v-row>

              <v-row> <b>Description</b>: </v-row>

              <v-row>
                <span>
                  <b>Status</b>:<br />
                  Status:<br />
                  Validity:<br />
                  Time:<br />
                  Duration:<br />
                  Message:<br />
                </span>
              </v-row>
            </v-card-text>
          </v-card>
        </v-dialog>
     <!--  </v-row> -->
<!--     </v-col> -->
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
    indicatoricons: Array,
    iconlg: String,
  },

  data() {
    return {
      dialog: false,
    };
  },

  methods: {
    selectCheck() {
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

