<template>
  <div>
    <b-row>
      <b-col md="4" class="mt-2 mt-md-0">
        <WidgetLineChart class="h-100" title="Closed Issues" :chart-data="closedIssueContributionsData"
          :content-loaded="contentLoaded" :first-time-loaded="firstTimeLoaded" />
      </b-col>
      <b-col md="5" class="mt-2 mt-md-0">
        <WidgetLineChart class="h-100" title="Open Issues" :chart-data="openIssueContributionsData"
          :content-loaded="contentLoaded" :first-time-loaded="firstTimeLoaded" />
      </b-col>
      <b-col md="3">
        <WidgetPieChart class="h-100" title="Organizations Contribution" />
      </b-col>
    </b-row>
    <b-row class="mt-2 mt-lg-3" height="100">
      <b-col md="4">
        <WidgetLineChart class="h-100" title="Open Pull Request" :chart-data="openPRControbutionsData"
          :content-loaded="contentLoaded" :first-time-loaded="firstTimeLoaded" />
      </b-col>
      <b-col md="4">
        <WidgetLineChart class="h-100" title="Merged Pull Request" :chart-data="mergedPRControbutionsData"
          :content-loaded="contentLoaded" :first-time-loaded="firstTimeLoaded" />
      </b-col>
      <b-col md="4">
        <WidgetLineChart class="h-100" title="Closed Pull Request" :chart-data="closedPRControbutionsData"
          :content-loaded="contentLoaded" :first-time-loaded="firstTimeLoaded" />
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
  data() {
    return {
      openPRControbutionsData: {},
      mergedPRControbutionsData: {},
      closedPRControbutionsData: {},
      openIssueContributionsData: {},
      closedIssueContributionsData: {},
      firstTimeLoaded: false,
      contentLoaded: false
    }
  },
  watch: {
    "$route.query": {
      async handler(newValue) {
        this.firstTimeLoaded = false;
        this.contentLoaded = false;
        await this.getOpenPullRequestContributionsdata();
        await this.getMergedPullRequestContributionsdata();
        await this.getClosedPullRequestContributionsdata();
        await this.getOpenIssueContributionsdata();
        await this.getClosedIssueContributionsdata();
        this.contentLoaded = true;
        this.firstTimeLoaded = true;
      }
    }
  },
  async mounted() {
    this.firstTimeLoaded = false;
    this.contentLoaded = false;
    await this.getOpenPullRequestContributionsdata();
    await this.getMergedPullRequestContributionsdata();
    await this.getClosedPullRequestContributionsdata();
    await this.getOpenIssueContributionsdata();
    await this.getClosedIssueContributionsdata();
    this.contentLoaded = true;
    this.firstTimeLoaded = true;
  },
  methods: {
    async getOpenPullRequestContributionsdata() {
      const queryParams = this.$route.query;
      await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/pullrequest/status/open`, { params: queryParams })
        .then((res) => {
          if (res.data.data.length > 0) {
            const prControbutions = res.data.data;
            if (prControbutions.length > 0) {
              const prControbutionsLables = prControbutions.map(item => this.$utils.getFormattedTimeStamp(item.date));
              const prControbutionsDataSet = prControbutions[0].data.map(item => {
                return {
                  label: item.user,
                  data: prControbutions.map(entry => {
                    const userData = entry.data.find(d => {
                      return d.user === item.user
                    });
                    return userData ? userData.count : 0;
                  }),
                  borderColor: this.$utils.getColor(item.user),
                  fill: false
                };
              });
              this.openPRControbutionsData = {
                labels: prControbutionsLables,
                datasets: prControbutionsDataSet
              }
            }
          } else {
            this.openPRControbutionsData = {};
          }
        })
        .catch((err) => {
          this.openPRControbutionsData = {};
          this.$toaster.error(err);
        })
        .finally(() => {
          // After getting data from API
        });
    },
    async getClosedPullRequestContributionsdata() {
      const queryParams = this.$route.query;
      await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/pullrequest/status/closed`, { params: queryParams })
        .then((res) => {
          if (res.data.data.length > 0) {
            const prControbutions = res.data.data;
            if (prControbutions.length > 0) {
              const prControbutionsLables = prControbutions.map(item => this.$utils.getFormattedTimeStamp(item.date));
              const prControbutionsDataSet = prControbutions[0].data.map(item => {
                return {
                  label: item.user,
                  data: prControbutions.map(entry => {
                    const userData = entry.data.find(d => {
                      return d.user === item.user
                    });
                    return userData ? userData.count : 0;
                  }),
                  borderColor: this.$utils.getColor(item.user),
                  fill: false
                };
              });
              this.closedPRControbutionsData = {
                labels: prControbutionsLables,
                datasets: prControbutionsDataSet
              }
            }
          } else {
            this.closedPRControbutionsData = {};
          }
        })
        .catch((err) => {
          this.closedPRControbutionsData = {};
          this.$toaster.error(err);
        })
        .finally(() => {
          // After getting data from API
        });
    },
    async getMergedPullRequestContributionsdata() {
      const queryParams = this.$route.query;
      await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/pullrequest/status/merged`, { params: queryParams })
        .then((res) => {
          if (res.data.data.length > 0) {
            const prControbutions = res.data.data;
            if (prControbutions.length > 0) {
              const prControbutionsLables = prControbutions.map(item => this.$utils.getFormattedTimeStamp(item.date));
              const prControbutionsDataSet = prControbutions[0].data.map(item => {
                return {
                  label: item.user,
                  data: prControbutions.map(entry => {
                    const userData = entry.data.find(d => {
                      return d.user === item.user
                    });
                    return userData ? userData.count : 0;
                  }),
                  borderColor: this.$utils.getColor(item.user),
                  fill: false
                };
              });
              this.mergedPRControbutionsData = {
                labels: prControbutionsLables,
                datasets: prControbutionsDataSet
              }
            }
          } else {
            this.mergedPRControbutionsData = {};
          }
        })
        .catch((err) => {
          this.mergedPRControbutionsData = {};
          this.$toaster.error(err);
        })
        .finally(() => {
          // After getting data from API
        });
    },
    async getOpenIssueContributionsdata() {
      const queryParams = this.$route.query;
      await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/issue/status/open`, { params: queryParams })
        .then((res) => {
          if (res.data.data.length > 0) {
            const issueContributions = res.data.data;
            if (issueContributions.length > 0) {
              const issueContributionsLables = issueContributions.map(item => this.$utils.getFormattedTimeStamp(item.date));
              const issueContributionsDatasets = issueContributions[0].data.map(item => {
                return {
                  label: item.user,
                  data: issueContributions.map(entry => {
                    const userData = entry.data.find(d => {
                      return d.user === item.user
                    });
                    return userData ? userData.count : 0;
                  }),
                  borderColor: this.$utils.getColor(item.user),
                  fill: false
                };
              });
              this.openIssueContributionsData = {
                labels: issueContributionsLables,
                datasets: issueContributionsDatasets
              }
            }
          } else {
            this.openIssueContributionsData = {};
          }
        })
        .catch((err) => {
          this.openIssueContributionsData = {};
          this.$toaster.error(err);
        })
        .finally(() => {
          // After getting data from API
        });
    },
    async getClosedIssueContributionsdata() {
      const queryParams = this.$route.query;
      await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/issue/status/closed`, { params: queryParams })
        .then((res) => {
          if (res.data.data.length > 0) {
            const issueContributions = res.data.data;
            if (issueContributions.length > 0) {
              const issueContributionsLables = issueContributions.map(item => this.$utils.getFormattedTimeStamp(item.date));
              const issueContributionsDatasets = issueContributions[0].data.map(item => {
                return {
                  label: item.user,
                  data: issueContributions.map(entry => {
                    const userData = entry.data.find(d => {
                      return d.user === item.user
                    });
                    return userData ? userData.count : 0;
                  }),
                  borderColor: this.$utils.getColor(item.user),
                  fill: false
                };
              });
              this.closedIssueContributionsData = {
                labels: issueContributionsLables,
                datasets: issueContributionsDatasets
              }
            }
          } else {
            this.closedIssueContributionsData = {};
          }
        })
        .catch((err) => {
          this.closedIssueContributionsData = {};
          this.$toaster.error(err);
        })
        .finally(() => {
          // After getting data from API
        });
    }
  }
}
</script>
