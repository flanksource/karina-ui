<template>
  <span>
    <span class="cluster-title">
      {{ id }}
    </span>

    <span v-for="(item, i) in items" :key="i" class="no-events">
      <ItemCard
        :label="item.label"
        :icon="item.icon"
        :content="item.badge"
        :color="item.badgecolor"
      />
    </span>

    <span v-if="metricsPresent">
      <UsageBarItemCard
        icon="cpu"
        content="2"
        badgecolor="#336600"
        v-bind:value="metrics.cpu_resources.current"
        v-bind:min="0"
        v-bind:max="metrics.cpu_resources.total"
        v-bind:optimum="metrics.cpu_resources.min"
        v-bind:marker="metrics.cpu_resources.max"
        orientation="horizontal"
      />

      <!--<UsageBarItemCard
        icon="cpu"
        content="2"
        badgecolor="#336600"
        v-bind:value="45"
        v-bind:min="0"
        v-bind:max="100"
        v-bind:optimum="75"
        orientation="vertical"
      /> -->

      <UsageBarItemCard
        icon="ram"
        content="0"
        badgecolor="#f00"
        v-bind:value="getGb(metrics.memory_resources.current)"
        v-bind:min="0"
        v-bind:max="getGb(metrics.memory_resources.total)"
        v-bind:optimum="getGb(metrics.memory_resources.min)"
        v-bind:marker="getGb(metrics.memory_resources.max)"
        orientation="horizontal"
      />

      <!-- <UsageBarItemCard
        icon="ram"
        content="0"
        badgecolor="#f00"
        v-bind:value="45"
        v-bind:min="0"
        v-bind:max="100"
        v-bind:optimum="75"
        orientation="vertical"
      /> -->
    </span>
  </span>
</template>

<script>
import ItemCard from "./ItemCard.vue";
import UsageBarItemCard from "./UsageBarItemCard.vue";

export default {
  name: "ItemList",

  components: {
    ItemCard,
    UsageBarItemCard,
  },

  props: {
    id: String,
    items: Array,
    metrics: Object,
  },

  methods: {
    getGb(value) {
      return Math.round(value / (1024 * 1024 * 1024) * 100) / 100
    }
  },

  computed: {
    metricsPresent() {
      return Object.keys(this.metrics).length > 0
    }
  }
};
</script>

<style scoped>
.no-events {
  pointer-events: none;
}

.cluster-title {
  text-decoration: underline;
  font-weight: bold;
}
</style>
