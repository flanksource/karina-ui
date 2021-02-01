<template>
  <span>
    <div v-if="content > 0">
      <v-badge :color="color" :content="content" inline>
        <i :class="'icon-'+icon"></i>
        <span class="pl-1 text-size">{{ label }} </span>
      </v-badge>
    </div>

    <div v-else inline>
      <v-badge value="false" color="white" inline>
        <i :class="'icon-'+icon"></i>
        <span class="pl-1 text-size">{{ label }} </span>
      </v-badge>
    </div>
  </span>
</template>

<script>
export default {
  name: "ItemCard",


  props: {
    source: String,
    icon: String,
    label: String,
    content: Number,
    color: String,
  },
};
</script>

<style scoped>
.text-size {
  font-size: 12px;
}

@import '../assets/fonts/karina-ui-icons/icons.css';



.icon-cpu:before {
    color: #9a2812 !important;
}
.icon-memory:before {
    color: #336610 !important;
}
.icon-cargo-ship-50:before {
    color: #98ccfc !important;

}


</style>
