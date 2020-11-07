<template>
  <v-row align="center" class="pl-5">
    <v-col cols="2" xs="2" sm="2" md="2" lg="2" xl="2" class="pa-0">
      <v-img
        :aspect-ratio="1 / 1"
        :src="require(`@/assets/svg/${iconlg}`)"
        width="32"
      />
    </v-col>

    <v-col cols="10" xs="10" sm="8" md="8" lg="8" xl="8" class="pa-0 bordered">
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
                  <th colspan="3" scope="colgroup">Data</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td colspan="3">{{selectedCheck.color}}</td>
                    <td colspan="3">{{ selectedCheck.shape }}</td>
                    <td colspan="3">{{ selectedCheck.overlay }}</td>
                    <td colspan="3">{{ selectedCheck }}</td>
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
                  <th colspan="1" scope="rowgroup"></th>
                  <th colspan="1">Data</th>
                  <th colspan="1">Data</th>
                  <th colspan="1">Data</th>
                  <th colspan="2">Data</th>
                  <th colspan="6">Data</th>
                </tr>   
              </tbody>
              <tbody
                v-for="(status, i) in selectedCheck"
                :key="i"
              >
                <tr>
                  <th colspan="1" scope="rowgroup"> Check {{ i + 1 }}</th>
                  <td colspan="1">{{ status }}</td>
                  <td colspan="1">{{ status }}</td>
                  <td colspan="1">{{ status }}</td>
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

