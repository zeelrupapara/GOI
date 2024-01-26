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
    query: {
      type: String,
      default: null
    }
  },
  data() {
    return {
      values: null
    }
  },
  methods: {
    applyFilter() {
      // get currunt query params
      const queryParams = { ...this.$route.query };
      const valuesKey = this.values.map(value => value.key)

      // set filters wise query params
      switch (this.query) {
        case this.$constants.FILTERS.ORG_QP:
          delete queryParams.orgs;
          if (valuesKey.length > 0) {
            queryParams.orgs = JSON.stringify(valuesKey);
          }
          break;

        case this.$constants.FILTERS.REPO_QP:
          delete queryParams.repos;
          if (valuesKey.length > 0) {
            queryParams.repos = JSON.stringify(valuesKey);
          }
          break;

        case this.$constants.FILTERS.MEMBER_QP:
          delete queryParams.membs;
          if (valuesKey.length > 0) {
            queryParams.membs = JSON.stringify(valuesKey);
          }
          break;
        case null:
          break;
      }
      this.$router.push({
        query: queryParams
      })
    }
  }
}
</script>
<template>
  <multiselect @input="applyFilter" style="font-size: 0.8rem;" v-model="values" label="name" track-by="name"
    :options="options" :placeholder="placeholder" :multiple="true" />
</template>

<style src="vue-multiselect/dist/vue-multiselect.min.css">
</style>
