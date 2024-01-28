<script>
import { Doughnut, mixins } from "vue-chartjs";

export default {
  extends: Doughnut,
  mixins: [mixins.reactiveProp],
  props: {
    chartData: {
      type: Object,
      default: () => {
        return {};
      },
      required: true
    },
    options: {
      type: Object,
      default: () => {
        return {
          responsive: true,
          maintainAspectRatio: false,
          legend: {
            display: true,
            position: "right"
          }
        };
      }
    }
  },
  mounted() {
    this.renderChart(this.chartData, this.options);
  }
};
</script>
