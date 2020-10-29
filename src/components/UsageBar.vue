<template>
	<div class="bar">
		<span class="unit-min">{{min}}</span>
		<span class="unit-max">{{max}}</span>

		<div class="scale">
			<v-progress-linear
				:value="getValue(value, min, max)"
				height="5"
				rounded
				background-color="#4a4a4a"
				:color="getColour(value,optimum)"
				:title="metric + ': ' + getTitle(value, min, max)"
			>
			</v-progress-linear>
		</div>

		<div
			class="marker"
			v-bind:style="{marginLeft: `${getValue(optimum, min, max)}%`}"
		>
			<svg-icon
				icon="triangle-marker"
				:style="markerStyles"
				:title="getTitle(optimum, min, max)"
			/>
		</div>
	</div>
</template>

<script>
	import SvgIcon from "./SVGIcon.vue";
	
	export default{

		components: {
			SvgIcon,
		},

		props: {
			value: Number,
			min: Number,
			max: Number,
			colour: String,
			optimum: Number,
			metric: String,
		},

		methods : {
			getValue(value, min, max) {

				var usageValue;

				if (min < max) {

					usageValue = (value/max) * 100;
					return usageValue;
				}

				if (min > max) {

					usageValue =  (value/min) *100;
					return usageValue;
				}
			},

			getOptimum(optimum, min, max) {

				var optimumValue;

				if (min < max) {

					optimumValue = (optimum/max) * 100;
					return optimumValue;
				}

				if (min > max) {

					optimumValue =  (optimum/min) *100;
					return optimumValue;
				}
			},

			getTitle(value, min, max){

				var title;

				if (min < max) {

					title = value;
					return title;
				}

				if (min > max) {

					title =  value * -1;
					return title;
				}
			},

			getColour(value, optimum){

				var colour;

				if (value < optimum) {

					colour = "#29d1b3";
					return colour;
				}

				if (value > optimum) {

					colour = "#fa3861";
					return colour;
				}
			},

		},

		computed: {
			markerStyles() {

				var colour;

				if (this.value < this.optimum) {
					colour = "#29d1b3";
				}

				else if (this.value > this.optimum) {
					colour = "#fa3861";	
				}

				return {
					color: colour,
				};
			},
		},
	}
</script>


<style scoped>

	.bar {
		padding: 25px 0;
		position: relative;
		margin: 0 10px;
	}

	.scale {
		position: absolute;
		margin-top: 0;
		margin-left: 5px;
		width: 100%;
	}

	.marker {
		position: absolute;
		margin-top: 2px;
		font-size: 10px;

	}

	.unit-min {
		position: relative;
		float: left;
		margin-left: 5px;
		margin-top: -15px;
		z-index: 200;
		font-size: 12px;
	}

	.unit-max {
		position: relative;
		float: right;
		margin-right: -5px;
		margin-top: -15px;	
		z-index: 200;
		font-size: 12px;
	}
	
</style>
