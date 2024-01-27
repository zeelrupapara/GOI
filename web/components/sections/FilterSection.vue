<template>
  <div class="row justify-content-center">
    <div v-for="(filter, index) in multiselectFilters" :key="index" class="col col-3">
      <MultiSelectDropdownFilter id="select_users" :placeholder="filter.placeholder" :options="filter.options"
        :query="filter.query" :selected-options="filter.selectedOptions" />
    </div>
    <div class="col col-3">
      <DateRangePickerFilter :placeholder="timeRangeSelectPlaceholder" />
    </div>
  </div>
</template>

<script>
import MultiSelectDropdownFilter from '~/components/filters/MultiSelectDropdownFilter.vue';
import DateRangePickerFilter from '~/components/filters/DateRangePickerFilter.vue';
export default {
  components: {
    MultiSelectDropdownFilter,
    DateRangePickerFilter
  },
  data() {
    return {
      multiselectFilters: [],
      timeRangeSelectPlaceholder: this.$constants.FILTERS.DATETIME_PLACEOLDER
    }
  },
  async mounted() {
    await this.getOrgnizationFilterOptions()
    await this.getMemberFilterOptions()
    await this.getRepositoryFilterOptions()
  },
  methods: {
    getOrgnizationFilterOptions() {
      this.$axios.get(`${this.$constants.API_URL_PREFIX}/filters/organization`).then((res) => {
        res = res.data.data;
        let selectedOrgs = null
        if (this.$route.query.orgs){
          selectedOrgs = JSON.parse(this.$route.query.orgs)
          selectedOrgs = selectedOrgs.map(org => res.find(o => o.key === org))
        }
        this.multiselectFilters.push({
          query: this.$constants.FILTERS.ORG_QP,
          placeholder: this.$constants.FILTERS.ORG_PLACEHOLDER,
          options: res,
          selectedOptions: selectedOrgs
        })

      }).catch((err) => {
        this.$toaster.error(err)
      }).finally(() => {
        // After getting data from API
      })
    },
    getRepositoryFilterOptions() {
      this.$axios.get(`${this.$constants.API_URL_PREFIX}/filters/repository`).then((res) => {
        res = res.data.data;
        let selectedRepos = null
        if (this.$route.query.repos){
          selectedRepos = JSON.parse(this.$route.query.repos)
          selectedRepos = selectedRepos.map(repo => res.find(o => o.key === repo))
        }
        this.multiselectFilters.push({
          query: this.$constants.FILTERS.REPO_QP,
          placeholder: this.$constants.FILTERS.REPO_PLACEHOLDER,
          options: res,
          selectedOptions: selectedRepos
        })
      }).catch((err) => {
        this.$toaster.error(err)
      }).finally(() => {
        // After getting data from API
      })
    },
    getMemberFilterOptions() {
      this.$axios.get(`${this.$constants.API_URL_PREFIX}/filters/member`).then((res) => {
        res = res.data.data;
        let selectedMembs = null
        if (this.$route.query.membs) {
          selectedMembs = JSON.parse(this.$route.query.membs)
          selectedMembs = selectedMembs.map(memb => res.find(o => o.key === memb))
        }
        this.multiselectFilters.push({
          query: this.$constants.FILTERS.MEMBER_QP,
          placeholder: this.$constants.FILTERS.MEMBER_PLACEHOLDER,
          options: res,
          selectedOptions: selectedMembs
        })
      }).catch((err) => {
        this.$toaster.error(err)
      }).finally(() => {
        // After getting data from API
      })
    },
  }
}
</script>
