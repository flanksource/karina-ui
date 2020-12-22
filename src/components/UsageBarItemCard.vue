<template>
	<span>
		<div v-if="content > 0">
			<v-badge
				v-if="orientation == 'horizontal'"
				:color="badgecolor"
				:content="content"
				inline
			>
				<svg-icon :icon="icon" :hasFill="false" :color="badgecolor" />
				<div class="bar">
					<span class="unit-min">{{min}}</span>
					<span class="unit-max">{{max}}</span>

					<div class="scale">
						<v-progress-linear
							:value="getValue(value, min, max)"
							height="5"
							rounded
							background-color="#4a4a4a"
							:color="getColor(value,optimum)"
							:title="getTitle(value, min, max)"
						/>
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
			</v-badge>

			<v-badge
				v-if="orientation == 'vertical'"
				:color="badgecolor"
				:content="content"
				inline
			>
				<svg-icon :icon="icon" :hasFill="false" :color="badgecolor" />
				<div class="vertical-bar">
					<span class="vertical-unit-min">{{min}}</span>
					<span class="vertical-unit-max">{{max}}</span>

					<div class="vertical-scale">
						<v-progress-linear
							:value="getValue(value, min, max)"
							height="10"
							rounded
							background-color="#4a4a4a"
							:color="getColor(value,optimum)"
							:title="getCollapsedTitle(value, optimum, min, max)"
						/>
					</div>

					<div
						class="vertical-marker"
						v-bind:style="{marginLeft: `20%`}"
					>
						<svg-icon
							icon="triangle-marker"
							title="title"
						/>
					</div>
				</div>
			</v-badge>
		</div>

		<div v-else inline>
			<v-badge
				v-if="orientation == 'horizontal'"
				value="false"
				color="white"
				inline
			>
				<svg-icon :icon="icon" :hasFill="false" :color="badgecolor" />
				<div class="bar">
					<span class="unit-min">{{min}}</span>
					<span class="unit-max">{{max}}</span>

					<div class="scale">
						<v-progress-linear
							:value="getValue(value, min, max)"
							height="5"
							rounded
							background-color="#4a4a4a"
							:color="getColor(value,optimum)"
							:title="getTitle(value, min, max)"
						/>
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
			</v-badge>


			<v-badge
				v-if="orientation == 'vertical'"
				value="false"
				color="white"
				inline
			>
				<svg-icon :icon="icon" :hasFill="false" :color="badgecolor" />
				<div class="vertical-bar">
					<span class="vertical-unit-min">{{min}}</span>
					<span class="vertical-unit-max">{{max}}</span>

					<div class="vertical-scale">
						<v-progress-linear
							:value="getValue(value, min, max)"
							height="10"
							rounded
							background-color="#4a4a4a"
							:color="getColor(value,optimum)"
							:title="getCollapsedTitle(value, optimum, min, max)"
						/>
					</div>

					<div
						class="vertical-marker"
						v-bind:style="{marginLeft: `20%`}"
					>
						<svg-icon
							icon="triangle-marker"
							title="title"
						/>
					</div>
				</div>
			</v-badge>
		</div>
	</span>
</template>

<script>
	import SvgIcon from "./SVGIcon.vue";

	export default{
		name: "UsageBarItemCard",

		components: {
			"svg-icon": SvgIcon,
		},

		props: {
			icon: String,
			content: String,
			badgecolor: String,
			value: Number,
			min: Number,
			max: Number,
			optimum: Number,
			orientation: String,
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

			getColor(value, optimum){
				var color;

				if (value < optimum) {
					color = "#29d1b3";
					return color;
				}

				if (value > optimum) {
					color = "#fa3861";
					return color;
				}
			},

			getCollapsedTitle(value, optimum, min, max) {

				var title;

				if (min < max) {

					title = ' value(' + value + ')' + ' optimum(' + optimum + ')';
					return title;
				}

				if (min > max) {

					title = ' value(' + value * -1 + ')' + ' optimum(' + optimum * -1 + ')';
					return title;
				}
			}		
		},

		computed: {
			markerStyles() {

				var color;

				if (this.value < this.optimum) {
					color = "#29d1b3";
				}

				else if (this.value > this.optimum) {
					color = "#fa3861";	
				}

				return {
					color: color,
				};
			},
		},
	}
</script>


<style scoped>

	.bar {
		padding: 15px 0;
		position: relative;
		margin: 0 2px 0 0;
		width: 60px;
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
		z-index: 3;
		font-size: 12px;
	}

	.unit-max {
		position: relative;
		float: right;
		margin-right: -5px;
		margin-top: -15px;	
		z-index: 3;
		font-size: 12px;
	}

	.vertical-bar {
		transform: rotate(-90deg);
		padding: 15px 0;
		position: relative;
		width: 35px;
		height: 50px;
	}

	.vertical-scale {
		position: absolute;
		margin-top: 0;
		width: 100%;
	}

	.vertical-marker {
		display: none;
	}

	.vertical-unit-min {
		display: none;
	}

	.vertical-unit-max {
		display: none;
	}
</style>
