<template>
      <v-row no-gutters>
        <span
          v-for="(canarycheck, i) in canarychecks"
          :key="i"
          class="ten-abreast"
        >
          <span v-if="canarycheck.checkStatuses">
            <span v-if="canarycheck.checkStatuses[0].Status == true">
              <IndicatorIcon
                shape="heart"
                colour="#336600"
                overlay=""
                class="clickable-icon"
                v-bind:title="canarycheck.name"
                @click.native="selectCheck(canarycheck)"
              />
            </span>

            <span v-else>
              <IndicatorIcon
                shape="heart"
                colour="#990000"
                overlay=""
                v-bind:title="canarycheck.name"
                class="clickable-icon"
                @click.native="selectCheck(canarycheck)"
              />
            </span>
          </span>

          <span v-else>
            <IndicatorIcon
              shape="heart"
              colour="#994c00"
              v-bind:title="canarycheck.name"
              overlay=""
              class="clickable-icon"
              @click.native="selectCheck(canarycheck)"
            />
          </span>
        </span>

          <v-dialog v-model="dialog" max-width="50%">
            <v-card>
              <v-card-title class="headline">{{
                selectedCheck.name
              }}</v-card-title>

              <v-card-text>
                <v-row no-gutters>
                  <v-col class="ml-0">
                    <b>Type</b>: {{ selectedCheck.type }}
                  </v-col>

                  <v-col> <b>Uptime</b>: {{ selectedCheck.uptime }} </v-col>

                  <v-col> <b>Latency</b>: {{ selectedCheck.latency }} </v-col>
                </v-row>

                <v-row>
                  <b>Description</b>: {{ selectedCheck.description }}
                </v-row>

                <v-row>
                  <span
                    v-for="(status, i) in selectedCheck.checkStatuses"
                    :key="i"
                  >
                    <b>Status {{ i + 1 }}</b
                    >: <br />
                    Status: {{ status.Status }} <br />
                    Validity: {{ status.Invalid }} <br />
                    Time: {{ status.Time }} <br />
                    Duration: {{ status.Duration }} <br />
                    Message: {{ status.Message }} <br />
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
    canarychecks: Array,
    indicatoricon: String,
  },

  data() {
    return {
      dialog: false,
      selectedCheck: {},
    };
  },

  methods: {
    selectCheck(canarycheck) {
      this.selectedCheck = canarycheck;
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
