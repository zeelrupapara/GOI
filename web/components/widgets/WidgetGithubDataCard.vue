<template>
  <b-modal id="modal-prevent-closing" ref="modal" title="Fetch Github Data By Given Time Range" @show="resetModal"
    @hidden="resetModal" @ok="handleOk">
    <form class="text-center" ref="form" @submit.stop.prevent="handleSubmit">
      <date-picker class="vue-date-picker" v-model="dateTimeRange" type="date" placeholder="Select Date Range"
        :clearable="false" :multiple="true" format="DD-MM-YYYY" range />
    </form>
  </b-modal>
</template>

<script>
import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';
export default {
  components: {
    DatePicker
  },
  data() {
    return {
      dateTimeRange: [
        new Date(Date.now() - this.$utils.getDaysToMilliSecond(7)), new Date(Date.now())
      ],
    }
  },
  methods: {
    resetModal() {
      this.dateTimeRange = [
        new Date(Date.now() - this.$utils.getDaysToMilliSecond(7)), new Date(Date.now())
      ]
    },
    async handleOk(bvModalEvent) {
      // Prevent modal from closing
      bvModalEvent.preventDefault()
      // Trigger submit handler
      await this.handleSubmit()
    },
    async handleSubmit() {
      this.$nextTick(() => {
        this.$bvModal.hide('modal-prevent-closing')
      })
      await this.fetchGithubData()
    },


    async fetchGithubData(){
      await this.$axios.post(`${this.$constants.API_URL_PREFIX}/github/data`, {
        start_time: this.dateTimeRange[0].getTime(),
        end_time: this.dateTimeRange[1].getTime()
      }).then((res) => {
        if (res.data.status === 'success'){
          this.$toaster.success(this.$constants.MESSAGES.GITHUB_DATA_FETCH_SUCCESS)
        }else{
          this.$toaster.error(this.$constants.MESSAGES.GITHUB_DATA_FETCH_ERROR)
        }
      }).catch((err) => {
        this.$toaster.error(this.$constants.MESSAGES.GITHUB_DATA_FETCH_ERROR)
      })
    }
  }
}
</script>
