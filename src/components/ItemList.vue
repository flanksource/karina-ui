<template>
  <span>
    <span class="cluster-title">
      {{ name }}
     </span>
<!--     <span
      v-for="(property, i) in properties"
      :key="i"
      class="no-events"
    >
      <ItemCard
        :label="property.label" 
        :icon="property.type" 
        :alerts="getPropertyAlerts(property.alerts)"
        :alertscolour="property.badgecolour"
      />

    </span> -->


     <ItemCard
        :label="properties.allocatable.memory" 
        icon="memory" 
        :alerts="getPropertyAlerts(properties.null)"
        alertscolour="property.badgecolour"
        class="clickable-icon"
        @click.native="selectProperties(properties.capacity.memory, properties.allocatable.memory)"
      />
        <ItemCard
        :label="properties.capacity.cpu" 
        icon="cpu" 
        :alerts="getPropertyAlerts(properties.null)"
        alertscolour="property.badgecolour"
      />
        <ItemCard
        :label="properties.capacity['ephemeral-storage']" 
        icon="disk" 
        :alerts="getPropertyAlerts(properties.null)"       
        alertscolour="property.badgecolour"
      />
        <ItemCard
        :label="properties.nodeInfo.machineID" 
        icon="node" 
        :alerts="getPropertyAlerts(properties.null)"
        alertscolour="property.badgecolour"
      />
        <ItemCard
        :label="properties.nodeInfo.kernelVersion" 
        icon="linux" 
        :alerts="getPropertyAlerts(properties.null)"
        alertscolour="property.badgecolour"
      />
        <ItemCard
        :label="properties.nodeInfo.kubeProxyVersion" 
        icon="kubernetes" 
        :alerts="getPropertyAlerts(properties.null)"
        alertscolour="property.badgecolour"
      />
       <ItemCard
        :label="properties.nodeInfo.osImage" 
        icon="ubuntu" 
        :alerts="getPropertyAlerts(properties.null)"
        alertscolour="property.badgecolour"
      />
       <ItemCard
        label="" 
        icon="threat" 
        :alerts="getPropertyAlerts(properties.conditions)"
        alertscolour="property.badgecolour"
        class="clickable-icon"
        @click.native="selectProperty(properties.conditions)"
      />
       <ItemCard
        label="" 
        icon="safe" 
        :alerts="getPropertyAlerts(properties.null)"
        alertscolour="property.badgecolour"
      />

    <v-dialog v-model="dialog" max-width="50%">
            <v-card>
              <v-card-title class="headline">{{
              
              }}</v-card-title>

              <v-card-text>
                <v-row no-gutters>
                  {{ selectedProperty }}
                  <!-- <v-col class="ml-0">
                    <b>Type</b>: {{ selectedCheck.type }}
                  </v-col>

                  <v-col> <b>Uptime</b>: {{ selectedCheck.uptime }} </v-col>

                  <v-col> <b>Latency</b>: {{ selectedCheck.latency }} </v-col>
                </v-row>

                <v-row>
                  <b>Description</b>: {{ selectedCheck.description }}
                </v-row>

                <v-row>
                  <span
                    v-for="(status, i) in selectedCheck.checkStatuses"
                    :key="i"
                  >
                    <b>Status {{ i + 1 }}</b
                    >: <br />
                    Status: {{ status.Status }} <br />
                    Validity: {{ status.Invalid }} <br />
                    Time: {{ status.Time }} <br />
                    Duration: {{ status.Duration }} <br />
                    Message: {{ status.Message }} <br />
                  </span> -->
                </v-row>
              </v-card-text>
            </v-card>
          </v-dialog>

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
    },

    data() {
      return {
        dialog: false,
        selectedProperty: {},
        propertyAllocatable: {},
        propertyCapacity: {},
      };
    },

    methods: {
      getPropertyAlerts(alerts){

        var count = 0;
        if (alerts) {
          count = alerts.length;
        }
        return count;      
      },
      selectProperty(property) {
        this.selectedProperty = property
        this.dialog = true;
      },
      selectProperties(propertycapacity, propertyallocatable) {
        this.propertyAllocatable = propertyallocatable
        this.propertyCapacity = propertycapacity
        this.dialog = true;
      },
    }
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
