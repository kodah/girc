<template>
  <v-footer
      app
      color="transparent"
      height="72"
      inset
  >
    <v-text-field
        background-color="grey darken-3"
        dense
        flat
        hide-details
        rounded
        solo
        append-outer-icon="mdi-send"
        hint="Message text"
        placeholder="Message text"
        color="primary"
        @click:append-outer="sendMessage"
        autofocus
    >
    </v-text-field>
  </v-footer>
</template>

<script>
export default {
  name: "MessageBox",
  data: () => ({ drawer: null }),
  methods: {
    sendMessage () {
      console.log("sendMessage triggered")
    }
  }
}
</script>

<style scoped>

</style>