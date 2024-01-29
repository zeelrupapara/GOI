
<template>
  <div>
    <WidgetTableCard table-name="PR" :table-title="pullRequestTableTitle" :table-details="pullRequestTableDetails" :page-info="pullRequestPaginationInfo" />
    <WidgetTableCard table-name="Issue" :table-title="issueTableTitle" :table-details="issueTableDetails" :page-info="issuePaginationInfo" />
  </div>
</template>

<script>
import WidgetTableCard from '@/components/widgets/WidgetTableCard.vue';
export default {
   components:{
    WidgetTableCard
   },
   data(){
    return {
      pullRequestTableTitle: "Pull Request Contribution",
      pullRequestTableDetails: [],
      pullRequestPaginationInfo: {
        previous: false,
        next: false
      },
      issueTableTitle: "Issue Contribution",
      issueTableDetails: [],
      issuePaginationInfo: {
        previous: false,
        next: false
      },
    }
   },
   watch: {
    "$route.query": {
      async handler(newValue){
        await this.getPullRequestContributionTableDeatils();
        await this.getIssueContributionTableDetails();
      }
    }
   },
   async mounted(){
    await this.getPullRequestContributionTableDeatils();
    await this.getIssueContributionTableDetails();
   },
   methods:{
    async getPullRequestContributionTableDeatils(){
      const queryParams = this.$route.query;
      await await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/pullrequest/details`, { params: queryParams })
        .then((res) => {
            const prContributionDetails = res.data.data;
            this.pullRequestTableDetails = prContributionDetails.details;
            this.pullRequestPaginationInfo = prContributionDetails.page_info;
        })
        .catch((err) => {
          this.$toaster.error(err);
        })
        .finally(() => {
        });
      },
      async getIssueContributionTableDetails(){
        const queryParams = this.$route.query;
        await await this.$axios
          .get(`${this.$constants.API_URL_PREFIX}/contributions/issue/details`, { params: queryParams })
          .then((res) => {
              const issueContributionDetails = res.data.data;
              this.issueTableDetails = issueContributionDetails.details;
              this.issuePaginationInfo = issueContributionDetails.page_info;
          })
          .catch((err) => {
            this.$toaster.error(err);
          })
          .finally(() => {
          });
      }
    }
}
</script>

