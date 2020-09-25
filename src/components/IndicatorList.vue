<template>
  <span>
    <span v-for="(indicator, i) in indicators" :key="i">
      <IndicatorCard
        :iconlg="indicator.indicator_icon_lg"
        :indicatoricons="indicator.indicator_icons"
      />
    </span>

    <span>
      <IndicatorCardCanary
        indicatoricon="birdie.svg"
        :canarychecks="canarychecks"
      />
    </span>
  </span>
</template>


<script>
import IndicatorCard from "./IndicatorCard.vue";
import IndicatorCardCanary from "./IndicatorCardCanary.vue";

export default {
  name: "IndicatorList",

  components: {
    IndicatorCard,
    IndicatorCardCanary,
  },

  props: {
    indicators: Array,
    canarychecks: Array,
  },
};
</script>


<style scoped>
.no-events {
  pointer-events: none;
}
</style>
