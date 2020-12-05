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
        :colour="item.badgecolour"
      />
    </span>

    <UsageBarItemCard
      icon="cpu"
      content="2"
      badgecolour="#336600"
      v-bind:value="45"
      v-bind:min="0"
      v-bind:max="100"
      v-bind:optimum="75"
      orientation="horizontal"
    />

    <UsageBarItemCard
      icon="cpu"
      content="2"
      badgecolour="#336600"
      v-bind:value="45"
      v-bind:min="0"
      v-bind:max="100"
      v-bind:optimum="75"
      orientation="vertical"
    />

    <UsageBarItemCard
      icon="ram"
      content="0"
      badgecolour="#f00"
      v-bind:value="45"
      v-bind:min="0"
      v-bind:max="100"
      v-bind:optimum="75"
      orientation="horizontal"
    />

    <UsageBarItemCard
      icon="ram"
      content="0"
      badgecolour="#f00"
      v-bind:value="45"
      v-bind:min="0"
      v-bind:max="100"
      v-bind:optimum="75"
      orientation="vertical"
    />
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
  },
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
