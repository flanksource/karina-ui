<template>
  <v-card class="card-height">
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
          :color="itemicon.color"
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