<template>
  <v-card>
    <v-row no-gutters align="center" justify="space-around" class="py-2 ml-4">
      <v-col
        cols="3"
        sm="3"
        md="3"
        lg="3"
        xl="3"
        v-for="(itemicon, i) in itemicons"
        :key="i"
      >
        <ItemIcon
          :icon="itemicon.icon"
          :colour="itemicon.colour"
          :label="itemicon.label"
          :count="itemicon.count"
        />
      </v-col>
    </v-row>

    <Pipeline :cluster="cluster" />
  </v-card>
</template>

<script>
import ItemIcon from "./ItemIcon.vue";
import Pipeline from "./Pipeline.vue";

export default {
  name: "IndicatorPanel",

  components: {
    ItemIcon,
    Pipeline,
  },

  props: {
    cluster: Object,
    itemicons: Array,
  },
};
</script>

<style scoped>
.item-list {
  border-top-left-radius: 20px !important;
  border-bottom-left-radius: 20px !important;
  min-height: 100%;
  border-right: none;
  border: 1px solid #444;
}

.indicator-list {
  border-top-right-radius: 20px !important;
  border-bottom-right-radius: 20px !important;
  border-left: 2px solid #222;
  border: 1px solid #444;
}
</style>
