
<template>
  <div>
    <WidgetTableCard table-name="PR" :table-title="pullRequestTableTitle" :table-details="pullRequestTableDetails" :page-info="pullRequestPaginationInfo" :selected-status="pullRequestSelectedStatus" />
    <WidgetTableCard table-name="Issue" :table-title="issueTableTitle" :table-details="issueTableDetails" :page-info="issuePaginationInfo" :selected-status="issueSelectedStatus" />
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
      pullRequestSelectedStatus: null,
      pullRequestPaginationInfo: {
        previous: false,
        next: false
      },
      issueTableTitle: "Issue Contribution",
      issueTableDetails: [],
      issueSelectedStatus: null,
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
        if (newValue.pr_status){
          this.pullRequestSelectedStatus = newValue.pr_status
        }else{
          this.pullRequestSelectedStatus = null
        }

        if (newValue.issue_status){
          this.issueSelectedStatus = newValue.issue_status
        }else{
          this.issueSelectedStatus = null
        }
      }
    }
   },
   async mounted(){
    await this.getPullRequestContributionTableDeatils();
    await this.getIssueContributionTableDetails();

    if (this.$route.query.pr_status){
      this.pullRequestSelectedStatus = this.$route.query.pr_status
    }else{
      this.pullRequestSelectedStatus = null
    }

    if (this.$route.query.issue_status){
      this.issueSelectedStatus = this.$route.query.issue_status
    }else{
      this.issueSelectedStatus = null
    }
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

