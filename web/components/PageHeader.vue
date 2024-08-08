<script>
export default {
  components: {},
  props: {
    title: {
      type: String,
      default: '',
    },
    items: {
      type: Array,
      default: () => {
        return []
      },
    },
  },
}
</script>

<template>
  <div class="page-title-full-width page-title-box">
    <div class="container-fluid">
      <div class="row">
        <div>
          <div class="page-title-right">
            <b-breadcrumb :items="items" class="m-0"></b-breadcrumb>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
