<template>
  <div class="row justify-content-center">
    <div v-for="(filter, index) in multiselectFilters" :key="index" class="col">
      <MultiSelectDropdownFilter id="select_users" :placeholder="filter.placeholder" :options="filter.options"
        :query="filter.query" />
    </div>
    <div class="col">
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
        this.multiselectFilters.push({
          query: this.$constants.FILTERS.ORG_QP,
          placeholder: this.$constants.FILTERS.ORG_PLACEHOLDER,
          options: res
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
        this.multiselectFilters.push({
          query: this.$constants.FILTERS.REPO_QP,
          placeholder: this.$constants.FILTERS.REPO_PLACEHOLDER,
          options: res
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
        this.multiselectFilters.push({
          query: this.$constants.FILTERS.MEMBER_QP,
          placeholder: this.$constants.FILTERS.MEMBER_PLACEHOLDER,
          options: res
        })
      }).catch((err) => {
        this.$toaster.error(err)
      }).finally(() => {
        // After getting data from API
      })
    }
  }
}
</script>
