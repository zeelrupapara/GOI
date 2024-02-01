<script>
import { Line, mixins } from "vue-chartjs";
// import utils from "@/utils/utils.js";
export default {
  extends: Line,
  mixins: [mixins.reactiveProp],
  props: {
    chartData: {
      type: Object,
      default: () => {
        return {};
      }
    },
    options: {
      type: Object,
      default: () => {
        return {
          tooltips: {
            intersect: false,
            caretSize: 0,
            enabled: true,
            mode: "index",
            // callbacks: {
            //   label(tooltipItem, data) {
            //     const sum = data.datasets.reduce((sum, dataset) => {
            //       return sum + dataset.data[tooltipItem.index];
            //     }, 0);
            //     let percent =
            //       (data.datasets[tooltipItem.datasetIndex].data[
            //         tooltipItem.index
            //       ] /
            //         sum) *
            //       100;
            //     percent = percent.toFixed(2);
            //     return `${data.datasets[tooltipItem.datasetIndex].label}:  ${
            //       data.datasets[tooltipItem.datasetIndex].data[
            //         tooltipItem.index
            //       ]
            //     } Out of ${sum}  (${percent}%)`;
            //   }
            // }
          },
          responsive: true,
          maintainAspectRatio: false,
          legend: {
            display: true,
            position: "bottom"
          },
          elements: {
            point: {
              backgroundColor: "rgba(0, 0, 0, 0)",
              borderColor: "rgba(0, 0, 0, 0)",
              radius: 0,
              borderWidth: 0,
              hitRadius: 4
            },
          },
          gridLines: {
            reverse: true
          },
        };
      }
    }
  },
  mounted() {
    this.renderChart(this.chartData, this.options);
  }
};
</script>
