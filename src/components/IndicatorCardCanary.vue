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

  <table>
    <thead>
      <tr class="">
        <th class="border"><strong>Name</strong></th>
        <th class=""><strong>Type</strong></th>
        <th class=""><strong>Uptime</strong></th>
        <th class=""><strong>Latency</strong></th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>{{selectedCheck.name}}</td>
        <td>{{ selectedCheck.type }}</td>
        <td>{{ selectedCheck.uptime }} </td>
        <td>{{ selectedCheck.latency }}</td>
      </tr>
      <tr>
        <td colspan="6" class="">
          <table class="">
            <tbody>
              <tr>
                <th class="border">Description</th>
                <td>{{ selectedCheck.description }}</td>
              </tr>  

              <tr
                v-for="(status, i) in selectedCheck.checkStatuses"
                :key="i"
              >
                <td>Status {{ i + 1 }}</td>
                <td>
                  <tr>
                    <td>
                      Status: {{ status.Status }}
                    </td>
                    <td>
                      Validity: {{ status.Invalid }}
                    </td>
                    <td>
                      <vue-moments-ago 
                        prefix="Flag Raised"
                        suffix="ago"
                        :date="status.Time"
                      >
                      </vue-moments-ago>
                    </td>
                    <td>
                       Duration: {{ status.Duration }}
                    </td>
                    <td>
                       Message: {{ status.Message }} 
                    </td>
                  </tr>
                </td>
              </tr>
            </tbody>
          </table>
        </td>
      </tr>
    </tbody>
  </table>


            </v-card>
          </v-dialog>
      </v-row>
    </v-col>
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
.border{
  border: 2px solid #f00;
}

table {
  padding-left: 5px;
  padding-right: 5px;
}
</style>
