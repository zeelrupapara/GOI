<template>
  <div class="card">
    <b-overlay :show="!contentLoaded" opacity="0.3" blur="1px">
      <div v-if="firstTimeLoaded" class="card-body">
        <h4 class="header-title">{{ title }}</h4>
        <div class="widget-chart">
          <div
            class="d-flex align-items-center justify-content-center"
          >
          </div>
          <pie-chart
            v-if="Object.keys(chartData).length"
            class="mt-4"
            :chart-data="chartData"
            :height="300"
          />
          <!-- <div v-else class="text-center">
            <img
              src="~/assets/images/emptyGraph.svg"
              alt="no overall request(s)"
              class="d-block ml-auto mr-auto py-4"
            />
            <span class="d-block" data-test-id="fallback-mesasge">{{
              $constants.NO_CHART_DATA_MESSAGE
            }}</span>
          </div> -->
        </div>
      </div>
      <div v-else class="card-body">
        <h4 class="header-title">{{ title }}</h4>
        <div class="d-flex justify-content-center">
          <div class="spinner-border" role="status">
            <span class="sr-only">Loading...</span>
          </div>
        </div>
      </div>
    </b-overlay>
  </div>
</template>

<script>
import PieChart from "@/components/charts/PieChart.vue";
export default {
  components: {
    PieChart
  },
  props: {
    title: {
      type: String,
      default: ""
    },
  },
  data() {
    return {
      chartData: {
        labels: ['Red', 'Orange'],
        datasets: [
            {
              data: [300, 20, 40],
              backgroundColor: [
                '#FF6384', '#36A2EB', '#FFCE56', '#FF6384', '#36A2EB'
              ]
            }
          ]
      },
      firstTimeLoaded: true,
      totalRequests: 0,
      contentLoaded: true
    };
  },
  computed: {
    // getUrl() {
    //   const jobId = this.query.job_id || "";
    //   const from = this.query.from;
    //   const to = this.query.to;
    //   const status = this.query.status || common.defaultStatus;
    //   const siteKey = this.query.site_key || "";
    //   if (this.params.key) {
    //     return `${this.$constants.API_URL_PREFIX}/accounts/${this.params.slug}/sites/${this.params.key}/widgets/status-code-overall?from=${from}&to=${to}&status=${status}`;
    //   }
    //   return `${this.$constants.API_URL_PREFIX}/accounts/${this.params.slug}/sites/widgets/status-code-overall?from=${from}&to=${to}&job_id=${jobId}&status=${status}&site_key=${siteKey}`;
    // }
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
  //         this.getBotActivityData();
  //       }
  //     }
  //   }
  // },
  mounted() {
    this.firstTimeLoaded = true;
    // this.getBotActivityData();
  },
  methods: {
    getBotActivityData() {
      if (!this.query.from || !this.query.to) {
        return;
      }
      this.contentLoaded = !this.firstTimeLoaded;
      this.$axios
        .get(this.getUrl)
        .then((res) => {
          if (res.data.data) {
            const statusCodedata = res.data.data;
            if (statusCodedata && Object.keys(statusCodedata).length > 0) {
              // Visible only status codes that is supplied in query params
              // If not supplied any status code It includes all
              let statusIncludingTotal =
                (this.query &&
                  this.query.status &&
                  Object.keys(this.query.status).length &&
                  this.query.status.split(",")) ||
                Object.keys(statusCodedata);
              statusIncludingTotal = statusIncludingTotal.sort();
              const statusCodeName = statusIncludingTotal.filter(
                (s) => s !== "total"
              );
              const statusCodeValue = [];
              for (const prop of statusCodeName) {
                statusCodeValue.push(statusCodedata[prop].Total);
              }
              const backgroundColor = statusCodeName.map((statusCode) => {
                return this.$constants.HTTP_COLORS_CODE[statusCode];
              });
              this.totalRequests = statusCodedata.total.Total;
              this.chartData = {
                labels: statusCodeName,
                datasets: [
                  {
                    data: statusCodeValue,
                    backgroundColor,
                    borderColor: "transparent"
                  }
                ]
              };
            }
          } else {
            this.chartData = {};
          }
        })
        .catch((err) => {
          this.chartData = {};
          this.$toaster.error(err);
        })
        .finally(() => {
          this.contentLoaded = true;
          this.firstTimeLoaded = true;
        });
    }
  }
};
</script>
