<template>
  <div class="mt-3">
      <VueDraggable v-model="list" animation="150" ghostClass="ghost" class="row">
        <b-col v-for="matrix in matricsData" :key="matrix.id" class="col-md-6 col-xl-3">
          <CountWidget :count="matrix.count" :title="matrix.title" />
        </b-col>
      </VueDraggable>
  </div>
</template>
<script>
import CountWidget from "~/components/widgets/CountWidget.vue"
import { VueDraggable } from 'vue-draggable-plus';
export default {
  components: {
    CountWidget,
    VueDraggable
  },
  data() {
    return {
      matricsData: [
        {
          id: 1,
          title: "PRs",
          count: 100,
        },
        {
          id: 2,
          title: "Issues",
          count: 20
        },
        {
          id: 3,
          title: "Repositories",
          count: 2
        },
        {
          id: 4,
          title: "Organizations",
          count: 1
        }
      ]
    }
  }
}
</script>
