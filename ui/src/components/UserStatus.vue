<template>
  <v-sheet
      color="grey darken-3"
      height="128"
      width="100%"
  >
    <v-card
        class="mx-auto pl-16"
        max-width="344"
        outlined
    >
      <div class="mb-1">{{ user }}</div>
      <div class="text-sm-caption text--disabled text-uppercase mb-4">{{ network }}</div>

      <v-card-actions>
        <v-btn text>Settings</v-btn>
        <v-btn text>Connect</v-btn>
      </v-card-actions>
    </v-card>
  </v-sheet>
</template>

<script>
export default {
  name: "UserStatus",
  props: [
      'user',
      'network'
  ]
}
</script>

<style scoped>

</style>