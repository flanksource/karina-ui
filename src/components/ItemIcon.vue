<template>
  <span>
    <div v-if="count == 1">
      <v-badge bottom :content="count" offset-x="15" offset-y="20" overlap>
        <i :class="'icon-'+icon" :style="iconStyles" class="icon"></i>
      </v-badge>
    </div>

    <span class="text-size"> {{ label }} </span>
  </span>
</template>

<script>
export default {
  name: "ItemIcon",

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

@import '../assets/fonts/karina-ui-icons/icons.css';

</style>
