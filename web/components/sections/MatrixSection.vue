<template>
  <div class="mt-3">
      <VueDraggable v-if="matricsData.length > 0" class="row">
        <b-col v-for="(matrix, index) in matricsData" :key="index" class="col-md-6 col-xl-3">
          <WidgetCountCard :count="matrix.count" :title="matrix.title" />
        </b-col>
      </VueDraggable>
  </div>
</template>
<script>
import WidgetCountCard from "~/components/widgets/WidgetCountCard.vue"
import { VueDraggable } from 'vue-draggable-plus';
export default {
  components: {
    WidgetCountCard,
    VueDraggable
  },
  data() {
    return {
      matricsData: []
    }
  },
  watch:{
    "$route.query":{
      handler(){
        this.getMetrixData()
      }
    }
  },
  async mounted() {
    await this.getMetrixData()
  },
  methods:{
    getMetrixData(){
      const queryParams = this.$route.query;
      // call the API
      this.$axios.get(`${this.$constants.API_URL_PREFIX}/matrics`, { params: queryParams }).then((res) => {
        this.matricsData = res.data.data
      }).catch((err) => {
        this.$toaster.error(err)
      }).finally(() => {
        // After getting data from API
      })
    }
  }
}
</script>
