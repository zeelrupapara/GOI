<template>
  <div>
    <b-row>
      <b-col md="8" class="mt-2 mt-md-0">
          <WidgetLineChart
            class="h-100"
            title="Issue Activity"
            :chart-data="issueContributionsData"
            :content-loaded="contentLoaded"
            :first-time-loaded="firstTimeLoaded"
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
          :chart-data="prControbutionsData"
          :content-loaded="contentLoaded"
          :first-time-loaded="firstTimeLoaded"
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
      prControbutionsData: {},
      issueContributionsData: {},
      firstTimeLoaded: false,
      contentLoaded: false
    }
  },
  watch:{
    "$route.query":{
      async handler(newValue){
        this.firstTimeLoaded = false;
        this.contentLoaded = false;
        await this.getPullRequestContributionsdata();
        await this.getIssueContributionsdata();
        this.contentLoaded = true;
        this.firstTimeLoaded = true;
      }
    }
  },
  async mounted(){
    this.firstTimeLoaded = false;
    this.contentLoaded = false;
    await this.getPullRequestContributionsdata();
    await this.getIssueContributionsdata();
    this.contentLoaded = true;
    this.firstTimeLoaded = true;
  },
  methods:{
    async getPullRequestContributionsdata(){
      const queryParams = this.$route.query;
      await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/pullrequest`, { params: queryParams })
        .then((res) => {
          if (res.data.data) {
            const prControbutions = res.data.data;
            if (prControbutions.length > 0) {
              const prControbutionsLables = prControbutions.map(item => item.date);
              const prControbutionsDatasets = [{
                label: 'OPEN',
                data: prControbutions.map(item => item.total_open_prs),
                backgroundColor: this.$constants.CHART.OPEN_PR_COLOR,
              },{
                label: 'MERGED',
                data: prControbutions.map(item => item.total_merged_prs),
                backgroundColor: this.$constants.CHART.MERGED_PR_COLOR,
              },{
                label: 'CLOSED',
                data: prControbutions.map(item => item.total_closed_prs),
                backgroundColor: this.$constants.CHART.CLOSED_PR_COLOR,
              }]
              this.prControbutionsData = {
                labels: prControbutionsLables,
                datasets: prControbutionsDatasets
              }
            }
          } else {
            this.prControbutionsData = {};
          }
        })
        .catch((err) => {
          this.prControbutionsData = {};
          this.$toaster.error(err);
        })
        .finally(() => {
          // After getting data from API
        });
    },
    async getIssueContributionsdata(){
      const queryParams = this.$route.query;
      await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/issue`, { params: queryParams })
        .then((res) => {
          if (res.data.data) {
            const issueContributions = res.data.data;
            if (issueContributions.length > 0) {
              const issueContributionsLables = issueContributions.map(item => item.date);
              const issueContributionsDatasets = [{
                label: 'OPEN',
                data: issueContributions.map(item => item.total_open_issues),
                backgroundColor: this.$constants.CHART.OPEN_ISSUE_COLOR,
              },{
                label: 'CLOSED',
                data: issueContributions.map(item => item.total_closed_issues),
                backgroundColor: this.$constants.CHART.CLOSED_ISSUE_COLOR,
              }]
              this.issueContributionsData = {
                labels: issueContributionsLables,
                datasets: issueContributionsDatasets
              }
            }
          } else {
            this.issueContributionsData = {};
          }
        })
        .catch((err) => {
          this.issueContributionsData = {};
          this.$toaster.error(err);
        })
        .finally(() => {
          // After getting data from API
        });
    }
  }
}
</script>
