<template>
  <div>
    <b-row>
      <b-col md="8" class="mt-2 mt-md-0">
          <WidgetLineChart
            class="h-100"
            title="Issue Activity"
          />
        </b-col>
        <b-col md="4">
          <WidgetPieChart
            class="h-100"
            title="Organizations Activity"
          />
        </b-col>
    </b-row>
    <b-row class="mt-2 mt-lg-3" height="100">
      <b-col md="12">
        <WidgetLineChart
          class="h-100"
          title="Pull Request Activity"
        />
      </b-col>
    </b-row>
  </div>
</template>

<script>
import WidgetPieChart from '~/components/widgets/WidgetPieChart.vue';
import WidgetLineChart from '~/components/widgets/WidgetLineChart.vue';
export default {
  name: 'ChartSection',

  components: {
    WidgetPieChart,
    WidgetLineChart
  },
  data(){
    return {
      pieChartTitle: this.$constants.CHART.ORG_CHART_TITLE
    }
  }
}
</script>
