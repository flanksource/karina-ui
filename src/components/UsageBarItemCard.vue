<template>
	<span>
		<div v-if="content > 0">
			<v-badge
				v-if="orientation == 'horizontal'"
				:color="badgecolor"
				:content="content"
				inline
			>
				<i :class="'icon icon-'+icon" :color="badgecolor" ></i>
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
						<i class="icon icon-triangle-marker" :style="markerStyles"></i>

					</div>

					<div
						class="marker"
						v-bind:style="{marginLeft: `${getValue(marker, min, max)}%`}"
						v-if="markerAvailable()"
					>
						<i class="icon icon-triangle-marker" :style="markerStyles"></i>
					</div>
				</div>
			</v-badge>

			<v-badge
				v-if="orientation == 'vertical'"
				:color="badgecolor"
				:content="content"
				inline
			>
				<i :class="'icon icon-'+icon" :color="badgecolor" ></i>
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
						<i class="icon icon-triangle-marker" ></i>
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
				<i :class="'icon icon-'+icon" :color="badgecolor" ></i>
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
						<i class="icon icon-triangle-marker" :style="markerStyles"></i>
					</div>

					<div
						class="marker"
						v-bind:style="{marginLeft: `${getValue(marker, min, max)}%`}"
						v-if="markerAvailable()"
					>
						<i class="icon icon-triangle-marker" :style="markerStyles"></i>
					</div>
				</div>
			</v-badge>


			<v-badge
				v-if="orientation == 'vertical'"
				value="false"
				color="white"
				inline
			>
				<i :class="'icon icon-'+icon" :color="badgecolor" ></i>
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
						<i class="icon icon-triangle-marker" :style="markerStyles"></i>
					</div>
				</div>
			</v-badge>
		</div>
	</span>
</template>

<script>


	export default{
		name: "UsageBarItemCard",


		props: {
			icon: String,
			content: String,
			badgecolor: String,
			value: Number,
			min: Number,
			max: Number,
			optimum: Number,
			marker: Number,
			orientation: String,
		},

		methods: {
			getValue(value, min, max) {
				if (min <= max) {
					return (value/max) * 100;
				}

				return (value/min) *100;
			},

			getOptimum(optimum, min, max) {
				if (min <= max) {
					return (optimum/max) * 100;
				}

				return  (optimum/min) *100;
			},

			getTitle(value, min, max){
				if (min <= max) {
					return value
				}

				return value * -1
			},

			getColor(value, optimum){
				if (value <= optimum) {
					return "#29d1b3"
				}

				return "#fa3861";
			},

			getCollapsedTitle(value, optimum, min, max) {
				if (min <= max) {
					return ' value(' + value + ')' + ' optimum(' + optimum + ')';
				}

				return' value(' + value * -1 + ')' + ' optimum(' + optimum * -1 + ')';
			},

			markerAvailable() {
				return this.marker && this.marker > 0
			},
		},

		computed: {
			markerStyles() {
				if (this.value <= this.optimum) {
					return {
						color: "#29d1b3"
					}
				}

				return {
					color: "#fa3861"
				}
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

	@import '../assets/fonts/karina-ui-icons/icons.css';

</style>
