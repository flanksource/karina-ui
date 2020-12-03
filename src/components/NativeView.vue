<template>
	<v-container
		fluid
	>
		<v-row
			v-if="clusters==null"
			class="fill-height"
			align-content="center"
			justify="center"
		>
			<v-col
				class="subtitle-1 text-center"
				cols="12"
			>
				No Data Found!
			</v-col>
        </v-row>

        <v-row v-else>


			<v-col
				cols="12"
				xs="12"
				sm="12"
				md="6"
				lg="6"
				xl="6"
				v-for="(cluster, i) in clusters"
				:key="i"
			>

				<v-row no-gutters>
					<v-col
						cols="12"
						xs="12"
						sm="12"
						md="12"
						lg="12"
						xl="12"
					>
						<IndicatorPanel
							:cluster="cluster"
							:itemicons="cluster.itemicons"
						/>
					</v-col>
				</v-row>
			</v-col>
		</v-row>
	</v-container>
</template>

<script>
	import IndicatorPanel from "./IndicatorPanel.vue";
	
	export default {
		components: {
			IndicatorPanel,
		},

		props: {
			clusters: Array,
		},
	};
</script>

<style scoped>
	.fill-height{
		min-height: 89vh;
	}
</style>