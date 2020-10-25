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

    <span>
      <IndicatorCardPrometheus
        indicatoricon="idea.svg"
        :prometheusChecks="prometheusChecks"
      />
    </span>
  </span>
</template>

<script>
  import IndicatorCard from "./IndicatorCard.vue";
  import IndicatorCardCanary from "./IndicatorCardCanary.vue";
  import IndicatorCardPrometheus from "./IndicatorCardPrometheus.vue"

  export default {
    name: "IndicatorList",

    components: {
      IndicatorCard,
      IndicatorCardCanary,
      IndicatorCardPrometheus
    },

    props: {
      indicators: Array,
      canarychecks: Array,
      prometheusChecks: Array
    },
  };
</script>


<style scoped>
.no-events {
  pointer-events: none;
}
</style>
