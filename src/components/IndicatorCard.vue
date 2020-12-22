<template>

      <v-row no-gutters>
        <span
          v-for="(indicatoricon, i) in indicatoricons"
          :key="i"
          class="ten-abreast"
        >
          <IndicatorIcon
            :shape="indicatoricon.shape"
            :color="indicatoricon.color"
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

