<template>
  <!-- Begin page -->
  <div id="wrapper">
    <Topbar />
    <div class="content-page">
      <div class="content">
        <div class="container-fluid">
          <Nuxt />
        </div>
      </div>
      <Footer />
    </div>
  </div>
</template>

<script>
export default {
  mounted() {
    document.body.setAttribute("data-layout-mode", "horizontal");
  }
}
</script>

