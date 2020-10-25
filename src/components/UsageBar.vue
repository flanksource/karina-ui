<template>
	<div class="region">
		<span class="unit-min">{{min}}</span>
		<span class="unit-max">{{max}}</span>

		<v-progress-linear
			v-model="value"
			height="25"
			rounded
			background-color="#333"
			:color="colour"
		>
			
		</v-progress-linear>
	</div>
</template>



<script>
	
	export default{

		props: {
			value: Number,
			min: Number,
			max: Number,
			colour: String,
		},

		data() {
			
		},
	}
</script>


<style scoped>
	.unit-min{
		position: relative;
		margin-top: -1px;
		margin-left: 0;
		float: left;
		z-index: 200;

	}

	.unit-max{
		position: relative;
		margin-top: -1px;
		margin-right: 0;
		float: right;
		z-index: 200;

	}

	.region{
		border: 2px solid #000;
	}


	
</style>

<!-- <UsageBar
min=35
max=50
green=45
red=0
greenArrow=50
redArrow=-1
/> -->