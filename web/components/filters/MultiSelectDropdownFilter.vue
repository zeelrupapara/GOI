<script>
import Multiselect from 'vue-multiselect'
export default {
  components: {
    Multiselect
  },
  props: {
    options: {
      type: Array,
      default: []
    },
    placeholder: {
      type: String,
      default: "Select Options"
    },
    values: {
      type: Array,
      default: []
    }
  },
}
</script>
<template>
  <multiselect v-model="values" :options="options" :placeholder="placeholder" :multiple="true" />
</template>

<style src="vue-multiselect/dist/vue-multiselect.min.css">
</style>



