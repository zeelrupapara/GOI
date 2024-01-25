<template>
  <div class="widget-rounded-circle card">
    <div class="card-body">
      <div class="row">
        <div class="col-6">
          <div :class="`avatar-lg rounded-circle`">
            <img :src="image" alt="" srcset="">
          </div>
        </div>
        <div class="col-6">
          <div class="text-right">
            <h3 class="text-dark mt-1">
              <span>
                <countTo :end-val="count" :duration="3000"></countTo>
              </span>
            </h3>
            <p class="text-muted mb-1 text-truncate">{{ title }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import countTo from 'vue-count-to'
export default {
  components: {
    countTo
  },
  props: {
    title: {
      type: String,
      default: "Not Found"
    },
    count: {
      type: Number,
      default: 0
    },
    image: {
      type: String,
      default: ""
    },
    imageAlt: {
      type: String,
      default: ""
    }
  }
}
</script>
