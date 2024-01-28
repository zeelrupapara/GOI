<template>
  <div class="card">
    <b-overlay :show="!contentLoaded" opacity="0.3" blur="1px">
      <div v-if="firstTimeLoaded" class="card-body">
        <h4 class="header-title">{{ title }}</h4>
        <div class="widget-chart">
          <div class="d-flex align-items-center justify-content-center">
          </div>
          <pie-chart v-if="Object.keys(orgChartData).length" class="mt-4" :chart-data="orgChartData" :height="300" />
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
      orgChartData: {},
      firstTimeLoaded: true,
      contentLoaded: true
    }
  },
  watch: {
    "$route.query": {
      handler() {
        this.getOrganizationContributionsData()
      }
    }
  },
  async mounted() {
    this.firstTimeLoaded = false;
    await this.getOrganizationContributionsData();
  },
  methods: {
    async getOrganizationContributionsData() {
      this.contentLoaded = !this.firstTimeLoaded;
      const queryParams = this.$route.query;
      await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/organization`, { params: queryParams })
        .then((res) => {
          if (res.data.data) {
            const orgContributions = res.data.data;
            if (orgContributions.length > 0) {
              const pieChartLables = orgContributions.map(item => item.organization_name);
              const pieChartDatasets = [{
                data: orgContributions.map(item => item.total_prs + item.total_issues),
                backgroundColor: orgContributions.map(item => this.$utils.getColor(item.organization_name)),
              }]
              this.orgChartData = {
                labels: pieChartLables,
                datasets: pieChartDatasets
              }
            }
          } else {
            this.orgChartData = {};
          }
        })
        .catch((err) => {
          this.orgChartData = {};
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
