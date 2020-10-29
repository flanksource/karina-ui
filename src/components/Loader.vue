<template>
	<v-main>
		<v-container fluid class="loader-page">
			<v-row
				class="fill-height mt-10"
				align-content="center"
				justify="center"
			>
				<v-col
					cols="3"
				>
					<v-progress-linear
						color="grey"
						indeterminate
						rounded
						height="6"
					/>
				</v-col>
			</v-row>
		</v-container>
	</v-main>
</template>

<style>
  .loader-page {
    min-height: 100vh;
    margin-top: -65px;
    padding-top: 40vh;
  }
</style>