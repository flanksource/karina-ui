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

        <v-dialog v-model="dialog" max-width="75%">
          <v-card class="pa-3">
            <table style="width:100%">
              <col>
              <col>
              <colgroup span="12"></colgroup>
              <thead>
                <tr>
                  <th colspan="3" scope="colgroup">Name</th>
                  <th colspan="3" scope="colgroup">Type</th>
                  <th colspan="3" scope="colgroup">Uptime</th>
                  <th colspan="3" scope="colgroup">Latency</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td colspan="3">{{selectedCheck.name}}</td>
                    <td colspan="3">{{ selectedCheck.type }}</td>
                    <td colspan="3">{{ selectedCheck.uptime }}</td>
                    <td colspan="3">{{ selectedCheck.latency }}</td>
                </tr>
              </tbody>
              <tbody>
                <tr>
                  <th colspan="2" scope="rowgroup">Description</th>
                  <td colspan="10" scope="colgroup">{{ selectedCheck.description }}</td>
                </tr>
              </tbody>
              <tbody>
                <tr>
                  <th colspan="1" scope="rowgroup"></th>
                  <th colspan="1">Status</th>
                  <th colspan="1">Valid</th>
                  <th colspan="1">Duration</th>
                  <th colspan="2">Raised</th>
                  <th colspan="6">Message</th>
                </tr>   
              </tbody>
              <tbody
                v-for="(status, i) in selectedCheck.checkStatuses"
                :key="i"
              >
                <tr>
                  <th colspan="1" scope="rowgroup"> Check {{ i + 1 }}</th>
                  <td colspan="1">{{ status.Status }}</td>
                  <td colspan="1">{{ status.Invalid }}</td>
                  <td colspan="1">{{ status.Duration }}</td>
                  <td colspan="2">
                    <vue-moments-ago 
                      prefix=""
                      suffix="ago"
                      :date="status.Time"
                    />
                  </td>
                  <td colspan="6">{{ status.Message }} </td>
                </tr>
              </tbody>
            </table>
          </v-card>
        </v-dialog>
      </v-row>
</template>
<!-- 
// <script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js">
  

//</script>
 -->
<script>
import IndicatorIcon from "./IndicatorIcon.vue";
import VueMomentsAgo from "vue-moments-ago";




export default {
  name: "IndicatorCard",

  components: {
    IndicatorIcon,
    VueMomentsAgo,
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

table, th, td {
  border: 1px solid black;
  border-collapse: collapse;
}
th, td {
  padding: 5px;
  text-align: left;    
}
</style>
