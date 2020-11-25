<template>

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
            v-bind:title="indicatoricon.shape"
            @click.native="selectCheck(indicatoricon)"
          />
        </span>

        <v-dialog v-model="dialog" max-width="50%">
          <v-card class="pa-3">
            <table style="width:100%">
              <col>
              <col>
              <colgroup span="12"></colgroup>
              <thead>
                <tr>
                  <th colspan="3" scope="colgroup">Colour</th>
                  <th colspan="3" scope="colgroup">Shape</th>
                  <th colspan="3" scope="colgroup">Overlay</th>
                  <th colspan="3" scope="colgroup">Time</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td colspan="3">{{selectedCheck.color}}</td>
                  <td colspan="3">{{ selectedCheck.shape }}</td>
                  <td colspan="3">{{ selectedCheck.overlay }}</td>
                  <td colspan="3">
                    <vue-moments-ago 
                      prefix=""
                      suffix="ago"
                      :date="selectedCheck.time"
                    />
                  </td>
                </tr>
              </tbody>
              <tbody>
                <tr>
                  <th colspan="2" scope="rowgroup">Description</th>
                  <td colspan="10" scope="colgroup">{{ selectedCheck }}</td>
                </tr>
              </tbody>
              <tbody>
                <tr>
                  <th colspan="1" scope="rowgroup">Data Header Col</th>
                  <th colspan="1">Data Col 1</th>
                  <th colspan="1">Data Col 2</th>
                  <th colspan="1">Data Col 3</th>
                  <th colspan="2">Data Col 4</th>
                  <th colspan="6">Data Col 5</th>
                </tr>   
              </tbody>
              <tbody
                v-for="(status, i) in selectedCheck"
                :key="i"
              >
                <tr>
                  <th colspan="1" scope="rowgroup"> </th>
                  <td colspan="1"> </td>
                  <td colspan="1"></td>
                  <td colspan="1">
                  </td>
                  <td colspan="2">
                   
                  </td>
                  <td colspan="6"> </td>
                </tr>
              </tbody>
            </table>
        
          </v-card>
        </v-dialog>
  </v-row>
</template>

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
    indicatoricons: Array,
    iconlg: String,
  },

  data() {
    return {
      dialog: false,
      selectedCheck: {},
    };
  },

  methods: {
    selectCheck(indicatoricon) {
      this.selectedCheck = indicatoricon;
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

