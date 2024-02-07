<template>
  <b-modal size="lg" id="commit-model" hide-footer hide-header>
    <div class="d-block text-center">
      <!-- title section -->
      <div class="media">
        <img :src="commitHistoryData.avatar_url" alt="Profile Image" height="54" class="d-flex mr-3 rounded-circle">
        <div class="media-body">
          <h5 class="mt-0"><a href="#" class="text-dark">{{ commitHistoryData.login }}</a></h5>
          <p v-if="commitHistoryData.email" class="font-5"><b>Email:</b> <span><a href="#" class="text-muted">{{
            commitHistoryData.email }}</a></span></p>
        </div>
      </div>
      <hr>
      <h5><b>Commit History:</b> <span class="text-muted">{{ $utils.getFormattedTimeStamp(commitHistoryData.start_time)
      }}</span>~<span class="text-muted">{{ $utils.getFormattedTimeStamp(commitHistoryData.end_time) }}</span></h5>
      <div class="row">
        <div class="col-12">
          <div class="wrap" style="height: 600px; overflow: auto;" ref="wrap">
            <div class="list">
              <div class="timeline" dir="ltr">
                <article v-for="(data, index) in commitHistoryData.commit_history" :key="index"
                  :class="{ 'timeline-item-left': `${index % 2}` == 0 }" class="timeline-item">
                  <div v-if="data.button">
                    <h2 class="m-0 d-none">&nbsp;</h2>
                    <div v-if="data.button" class="time-show mt-3">
                      <a href="javascript: void(0);" class="btn btn-primary width-lg">{{ data.button }}</a>
                    </div>
                  </div>
                  <div class="timeline-desk">
                    <div v-if="!data.button" class="timeline-box">
                      <span :class="{
                        'arrow-alt': `${index % 2}` == 0,
                        'arrow': `${index % 2}` != 0,
                      }"></span>
                      <span class="timeline-icon">
                        <i class="mdi mdi-adjust"></i>
                      </span>
                      <h4 class="mt-0 font-16">{{ $utils.getFormattedTimeStamp(data.committed_date) }}</h4>
                      <p class="text-muted">
                        <small>{{ $utils.getFormattedDateTime(data.committed_date) }}</small>
                      </p>
                      <p class="mb-0">{{ formatedMessage(data.commit_message) }}</p>
                    </div>
                  </div>
                </article>
              </div>
            </div>
            <mugen-scroll scroll-container="wrap">

            </mugen-scroll>
          </div>

        </div>
      </div>
    </div>
      <b-button class="mt-3" block @click="closeModel">Close</b-button>
  </b-modal>
</template>

<script>
export default {
  props: {
    commitHistoryData: {
      type: Object,
      default: () => { }
    }
  },
  methods: {
    closeModel() {
      this.$bvModal.hide('commit-model')
    },
    formatedMessage(message) {
      const messages = message.split('*')
      if (message.length > 0) {
        return messages[0]
      }
    }
  }
}
</script>
