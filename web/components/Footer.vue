<script>
/**
 * Footer Component
 */
export default {};
</script>

<template>
  <!-- Footer Start -->
  <footer class="footer">
    <div class="container-fluid">
      <div class="row">
        <div class="col-md-6">
          {{ new Date().getFullYear() }} &copy; Rights reserved by
          <a href="https://www.improwised.com/" target="_blank">{{
            $constants.PRODUCT_NAME
          }}</a>
        </div>
        <div class="col-md-6">
          <div class="text-md-right footer-links d-none d-sm-block">
            <a href="javascript:void(0);">About Us</a>
            <a href="javascript:void(0);">Help</a>
            <a href="javascript:void(0);">Contact Us</a>
          </div>
        </div>
      </div>
    </div>
  </footer>
  <!-- end Footer -->
</template>
