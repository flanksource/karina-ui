<template>
  <span>
    <div v-if="count == 1">
      <v-badge bottom :content="count" offset-x="15" offset-y="20" overlap>
        <svg-icon :icon="icon" :style="iconStyles" class="icon" />
      </v-badge>
    </div>

    <span class="text-size"> {{ label }} </span>
  </span>
</template>

<script>
import SvgIcon from "./SVGIcon.vue";

export default {
  name: "ItemIcon",

  components: {
    SvgIcon,
  },

  props: {
    id: String,
    label: String,
    icon: String,
    count: [Number, String],
    color: String,
  },

  computed: {
    iconStyles() {
      return {
        color: this.color,
      };
    },
  },
};
</script>

<style scoped>
.icon {
  font-size: 48px;
}

.text-size {
  font-size: 12px;
}
</style>
