<template>
  <div>
    <div class="row">
      <!-- Table -->
      <div class="col-xl-12">
        <Portlet :headertitle="tableTitle">
          <div class="card-body pt-0">
            <div class="table-responsive mb-0">
              <table class="table table-hover table-centered mb-0">
                <thead>
                  <tr>
                    <th>Committer</th>
                    <th>Repository</th>
                    <th>Branch</th>
                    <th>Organization</th>
                    <th>Date</th>
                    <th>Commits</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(tableDetail, index) in tableDetails" :key="index">
                    <td>{{ tableDetail.committer }}</td>
                    <td> <a class="text-secondary" target="_black"
                        :href="`${$constants.GITHUB_URL_PREFIX}/${tableDetail.organization}/${tableDetail.repository}`">
                        {{ tableDetail.repository }}</a></td>
                    <td><a class="text-secondary" target="_black"
                        :href="`${$constants.GITHUB_URL_PREFIX}/${tableDetail.organization}/${tableDetail.repository}/blob/${tableDetail.branch}`">{{
                          tableDetail.branch }}</a></td>
                    <td>{{ tableDetail.organization }}</td>
                    <td>{{ getFormatedDate(tableDetail.date) }}</td>
                    <td><b-button class="btn btn-secondary btn-sm" type="button"
                        @click="LoadCommits(tableDetail.organization, tableDetail.repository, tableDetail.committer)"><i
                          class="fas fa-eye"></i>
                        <WidgetCommitsCard :commit-history-data="commitHistoryData" />
                      </b-button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div class="row">
              <div class="col">
                <div class="float-right">
                  <b-row>
                    <b-col v-show="pageInfo.previous">
                      <button type="button" @click="previousPage"
                        class="btn btn-sm nav-link dropdown-toggle arrow-none waves-effect waves-light"><i
                          class="fas fa-angle-left noti-icon"></i></button>
                    </b-col>
                    <b-col v-show="pageInfo.next">
                      <button type="button" @click="nextPage"
                        class="btn btn-sm nav-link dropdown-toggle arrow-none waves-effect waves-light"><i
                          class="fas fa-angle-right noti-icon"></i></button>
                    </b-col>
                  </b-row>
                </div>
              </div>
            </div>
          </div>
        </Portlet>
      </div>
    </div>
  </div>
</template>

<script>
import WidgetCommitsCard from "~/components/widgets/WidgetCommitsCard.vue"
export default {
  components: {
    WidgetCommitsCard
  },
  props: {
    tableTitle: {
      type: String,
      default: ""
    },
    tableDetails: {
      type: Array,
      default: []
    },
    pageInfo: {
      type: Object,
      default: {
        previous: false,
        next: false
      }
    },
  },
  data() {
    return {
      commitHistoryData: {},
    }
  },
  watch: {
    "$route.query": {
      handler() {
        const queryParams = { ...this.$route.query };
        if (!queryParams.pr_status) {
          this.selectedStatus
        }
      }
    }
  },
  methods: {
    getShortString(str) {
      if (str.length > 20) {
        return str.substring(0, 20) + '...'
      }
      return str.substring(0, 20)
    },
    getFormatedDate(date) {
      return this.$utils.getFormattedTimeStamp(date)
    },
    nextPage() {
      const queryParams = { ...this.$route.query };
      if (!queryParams.commit_page) {
        queryParams.commit_page = 2;
      } else {
        queryParams.commit_page = (Number(queryParams.commit_page) + 1).toString();
      }
      this.$router.push({
        query: queryParams
      })
    },
    previousPage() {
      const queryParams = { ...this.$route.query };
      if (!queryParams.commit_page) {
        queryParams.commit_page = 1;
      } else {
        if (Number(queryParams.commit_page) > 1) {
          queryParams.commit_page = (Number(queryParams.commit_page) - 1).toString();
        }
      }
      this.$router.push({
        query: queryParams
      })
    },
    async LoadCommits(org, repo, user) {
      this.$bvModal.show('commit-model')
      const queryParams = this.$route.query;
      await this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/contributions/organizations/${org}/repository/${repo}/member/${user}`, { params: queryParams })
        .then((res) => {
          if (res.data.data) {
            this.commitHistoryData = res.data.data
          } else {
            this.commitHistoryData = {};
          }
        })
        .catch((err) => {
          this.commitHistoryData = {};
          this.$toaster.error(err);
        })
        .finally(() => {
          // After getting data from API
        });
    }
  }
}
</script>
