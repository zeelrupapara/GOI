<script>
export default {
  data() {
    return {
      loading: false
    };
  },
  methods: {
    start() {
      this.loading = true;
    },
    finish() {
      this.loading = false;
    }
  }
};
</script>

<template>
  <div v-if="loading"
    class="w-100 cursor-progress position-fixed fixed-top fixed-bottom fixed-left fixed-right d-flex justify-content-center align-items-center zindex-tooltip wal-loading-transparent">
    <b-spinner label="spinner" variant="primary"></b-spinner>
  </div>
</template>

<style>
.wal-loading-transparent {
  background: rgb(0 0 0 / 15%);
}

.cursor-progress {
  cursor: progress;
}
</style>
