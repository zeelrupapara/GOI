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
  },
  data() {
    return {
      firstTimeLoaded: false,
      contentLoaded: true,
      chartData: {
        labels: ['Jan', 'Fab', 'March', 'April', 'May'],
        datasets: [{
          label: 'red',
          data: [300],
          backgroundColor: [
            'rgba(255, 99, 132, 0.2)', 'rgba(255, 159, 64, 0.2)', 'rgba(255, 205, 86, 0.2)', 'rgba(75, 192, 192, 0.2)', 'rgba(54, 162, 235, 0.2)'
          ]
        },{
          label: 'blue',
          data: [50],
          backgroundColor: [
            'rgba(255, 99, 132, 0.2)', 'rgba(255, 159, 64, 0.2)', 'rgba(255, 205, 86, 0.2)', 'rgba(75, 192, 192, 0.2)', 'rgba(54, 162, 235, 0.2)'
          ]
        },{
          label: 'green',
          data: [100],
          backgroundColor: [
            'rgba(255, 99, 132, 0.2)', 'rgba(255, 159, 64, 0.2)', 'rgba(255, 205, 86, 0.2)', 'rgba(75, 192, 192, 0.2)', 'rgba(54, 162, 235, 0.2)'
          ]
        },{
          label: 'white',
          data: [300, 50, 100, 40, 120],

          backgroundColor: [
            'rgba(255, 99, 132, 0.2)', 'rgba(255, 159, 64, 0.2)', 'rgba(255, 205, 86, 0.2)', 'rgba(75, 192, 192, 0.2)', 'rgba(54, 162, 235, 0.2)'
          ]
        },{
          label: 'black',
          data: [300, 50, 100, 40, 120],

          backgroundColor: [
            'rgba(255, 99, 132, 0.2)', 'rgba(255, 159, 64, 0.2)', 'rgba(255, 205, 86, 0.2)', 'rgba(75, 192, 192, 0.2)', 'rgba(54, 162, 235, 0.2)'
          ]
        }]
      },
      lineChartData: {}
    };
  },
  computed: {
    // records() {
    //   return this.lineChartData.records || [];
    // },
  },
  // watch: {
  //   query: {
  //     handler(newVal, oldVal) {
  //       if (
  //         newVal.from !== oldVal.from ||
  //         newVal.to !== oldVal.to ||
  //         newVal.job_id !== oldVal.job_id ||
  //         newVal.status !== oldVal.status ||
  //         newVal.site_key !== oldVal.site_key
  //       ) {
  //         this.getLineChartData();
  //       }
  //     }
  //   }
  // },
  mounted() {
    this.firstTimeLoaded = true;
    // this.getLineChartData();
  },
  methods: {
    getLineChartData() {
      if (!this.query.from || !this.query.to) {
        return;
      }
      this.contentLoaded = !this.firstTimeLoaded;
      this.$axios
        .get(this.getUrl)
        .then((res) => {
          res = res.data.data;
          this.lineChartData = res;
          if (res.records && res.records.length > 0) {
            this.refreshChart();
          }
        })
        .catch((err) => {
          this.$toaster.error(err);
          this.lineChartData = [];
        })
        .finally(() => {
          this.contentLoaded = true;
          this.firstTimeLoaded = true;
        });
    },

    refreshChart() {
      const labels = this.records.map((record) => {
        return this.$utils.getFormattedTimestamp(
          record.timestamp,
          this.lineChartData.aggregate_duration
        );
      });

      // Visible only status codes that is supplied in query params
      const includeStatusCode =
        (this.query &&
          this.query.status &&
          Object.keys(this.query.status).length &&
          this.query.status.split(",")) ||
        [];

      // Get single record, exclude timestamp line
      const firstRecord = this.records[0];
      const keys = Object.keys(firstRecord);

      // Make lines of only those in query params, else include all keys exluding timestamp
      const lines = keys.filter(
        (key) =>
          key !== "timestamp" &&
          key !== "total" &&
          (includeStatusCode.length ? includeStatusCode.includes(key) : key)
      );

      // Iterate each line, and take record for line
      const ds = [];
      lines.forEach((line) => {
        const obj = {
          label: line,
          data: [],
          borderColor: this.$constants.HTTP_COLORS_CODE[line],
          backgroundColor: this.$constants.HTTP_COLORS_CODE[line],
          fill: false
        };

        this.records.forEach((record) => {
          obj.data.push(record[line]);
        });

        ds.push(obj);
      });

      this.chartData = {
        labels,
        datasets: ds
      };
    }
  }
};
</script>
