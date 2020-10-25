<template>
  <span>
    <span class="cluster-title">
      {{ name }}
    </span>

    <span>
      Prometheus Data
    </span>

    <span
      v-for="(prometheusProp, i) in prometheusProps"
      :key="i"
      class="ten-abreast"
    >
      <span v-if="prometheusProp.data.result[0].metric.__name__ == 'node_memory_MemAvailable_bytes'">
        <ItemCard
          :label="prometheusProp.data.result[0].value[1]+' bytes'" 
          icon="ram" 
          alerts=""
          alertscolour=""
          class="clickable-icon"
        />
      </span>

      <span v-if="prometheusProp.data.result[0].metric.__name__ == 'machine_cpu_cores'">
        <ItemCard
          :label="prometheusProp.data.result[0].value[1]" 
          icon="cpu" 
          alerts=""
          alertscolour=""
          class="clickable-icon"
        />
      </span>

      <span v-if="prometheusProp.data.result[0].metric.__name__ == 'node_filesystem_size_bytes'">
        <ItemCard
          :label="prometheusProp.data.result[0].value[1]+' bytes'" 
          icon="storage" 
          alerts=""
          alertscolour=""
          class="clickable-icon"
        />
      </span>
    </span>
  </span>
</template>

<script>
import ItemCard from "./ItemCard.vue";
  export default {
    name: 'ItemList',
    
    components:{
      ItemCard
    },
    props: {
      name: String,
      properties: Object,
      prometheusProps: Array,
    },

  }
</script>

<style scoped>
  .no-events{
    pointer-events: none
  }
  
  .cluster-title {
    text-decoration: underline;
    font-weight: bold;
  }
  .clickable-icon {
    cursor: pointer;
  }
</style>