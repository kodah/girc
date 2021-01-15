<template>
  <div>
    <div class="text-lg-subtitle-2 d-inline red--text lighten-1 font-weight-black pt-3 pl-3">{{ user }}</div>
    <div class="text-sm-caption text--disabled d-inline pt-3 pl-2">{{ time }}</div>
    <div class="text-body-2 text--secondary pl-3">{{ message }}</div>
  </div>
</template>

<script>
export default {
  name: "Message",
  props: [
      'user',
      'time',
      'message'
  ]
}
</script>

<style scoped>

</style>