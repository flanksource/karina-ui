<template>
  

     <span class="centre">

<div>
    <v-badge
        bordered
        bottom
        color="deep-purple accent-4"
        :content="count"
        offset-x="25"
        offset-y="15"
        overlap="true"
      >
        <v-avatar :color="colour" size="64">
           <v-icon dark>{{icon}}</v-icon>

        </v-avatar>
      </v-badge>
</div>
<div>

     

    <span class="pb-0">{{label}}</span>
  </div>
  </span>


</template>

<script>
  export default{
    props:{
      label: String,
      icon: String,
      count: Number,
      colour: String,
    }
  }

</script>


<style>
  
  .centre{
    align-content: center;
    padding: 0px 20px;
  }
</style>