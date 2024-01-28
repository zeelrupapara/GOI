<template>
  <div class="card">
    <b-overlay :show="!contentLoaded" opacity="0.3" blur="1px">
      <div class="card-body">
        <h4 class="header-title">{{ title }}</h4>
        <div v-if="firstTimeLoaded">
          <LineChart
            class="mt-4 pt-1"
            :chart-data="chartData"
            :height="340"
          />
          <!-- <div v-else data-test-id="no-line-chart-data" class="text-center">
            <img
              src="~/assets/images/emptyGraph.svg"
              alt="no overtime request(s)"
              class="d-block ml-auto mr-auto py-4"
            />
            <span class="d-block">{{ $constants.NO_CHART_DATA_MESSAGE }}</span>
          </div> -->
        </div>
        <div v-else>
          <div class="d-flex justify-content-center">
            <div class="spinner-border" role="status">
              <span class="sr-only">Loading...</span>
            </div>
          </div>
        </div>
      </div>
    </b-overlay>
  </div>
</template>

<script>
import LineChart from "@/components/charts/LineChart.vue";
export default {
  components: {
    LineChart
  },
  props: {
    title: {
      type: String,
      default: ""
    },
    chartData: {
      type: Object,
      default: () => {}
    },
    contentLoaded:{
      type: Boolean,
      default: false
    },
    firstTimeLoaded:{
      type: Boolean,
      default: false
    }
  },
};
</script>
