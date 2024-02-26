<template>
  <!-- Topbar Start -->
  <div class="navbar-custom">
    <div class="container-fluid">
      <ul class="list-unstyled topnav-menu float-right mb-0 mx-2">
        <li class="dropdown d-lg-inline-block ml-2">
          <button type="button" v-b-modal.modal-prevent-closing v-b-tooltip.hover title="Fetch Github Data"
            class="btn btn-sm nav-link dropdown-toggle arrow-none waves-effect waves-light"><i
              class="fas fa-globe noti-icon"></i></button>
        </li>
        <li class="dropdown d-lg-inline-block">
          <button type="button" @click="clearFilter" v-b-tooltip.hover title="Reset Filters"
            class="btn btn-sm nav-link dropdown-toggle arrow-none waves-effect waves-light"><i
              class="fas fa-undo noti-icon"></i></button>
        </li>
        <!-- <li class="dropdown d-none d-lg-inline-block mx-2">
          <a v-b-tooltip.hover title="FullScreen" class="nav-link dropdown-toggle arrow-none waves-effect waves-light" data-toggle="fullscreen" href="#"
            @click="initFullScreen">
            <i class="fas fa-expand noti-icon"></i>
          </a>
        </li> -->
      </ul>
      <FetchGithubData />

      <!-- LOGO -->
      <div class="logo-box">
        <nuxt-link :to="`/`" class="logo logo-dark text-center">
          <span class="logo-sm">
            <span class="logo-lg-text-light">{{
              $constants.PRODUCT_NAME
            }}</span>
          </span>
          <span class="logo-lg">
            <span class="logo-lg-text-light text-capitalize">{{
              $constants.PRODUCT_NAME
            }}</span>
          </span>
        </nuxt-link>

        <nuxt-link :to="`/`" class="logo logo-light text-center">
          <span class="logo-sm">
            <!-- <img src="~/assets/images/logo-sm.png" alt height="24" /> -->
            <span class="logo-lg-text-light">{{
              $constants.PRODUCT_NAME
            }}</span>
          </span>
          <span class="logo-lg">
            <!-- <img src="~/assets/images/logo-light.png" alt height="20" /> -->
            <span class="logo-lg-text-light text-capitalize">{{
              $constants.PRODUCT_NAME
            }}</span>
          </span>
        </nuxt-link>
      </div>

      <ul class="list-unstyled topnav-menu topnav-menu-left m-0">
        <li>
          <button class="button-menu-mobile waves-effect waves-light" @click="toggleMenu">
            <i class="fas fa-bars"></i>
          </button>
        </li>
        <li>
        </li>
      </ul>
      <div class="clearfix"></div>
    </div>
  </div>
  <!-- end Topbar -->
</template>

<script>
import FetchGithubData from '~/components/widgets/WidgetGithubDataCard.vue';
export default {
  components: {
    FetchGithubData
  },
  methods: {
    toggleMenu() {
      this.$parent.toggleMenu();
    },
    initFullScreen() {
      document.body.classList.toggle("fullscreen-enable");
      if (
        !document.fullscreenElement &&
        /* alternative standard method */
        !document.mozFullScreenElement &&
        !document.webkitFullscreenElement
      ) {
        // current working methods
        if (document.documentElement.requestFullscreen) {
          document.documentElement.requestFullscreen();
        } else if (document.documentElement.mozRequestFullScreen) {
          document.documentElement.mozRequestFullScreen();
        } else if (document.documentElement.webkitRequestFullscreen) {
          document.documentElement.webkitRequestFullscreen(
            Element.ALLOW_KEYBOARD_INPUT
          );
        }
      } else if (document.cancelFullScreen) {
        document.cancelFullScreen();
      } else if (document.mozCancelFullScreen) {
        document.mozCancelFullScreen();
      } else if (document.webkitCancelFullScreen) {
        document.webkitCancelFullScreen();
      }
    },
    toggleRightSidebar() {
      this.$parent.toggleRightSidebar();
    },
    horizonalmenu() {
      const element = document.getElementById("topnav-menu-content");
      element.classList.toggle("show");
    },
    clearFilter() {
      this.$router.push({ path: this.$route.path, query: {} })
    }
  }
};
</script>

<style>
.b-dropdown-text {
  padding: 1rem 1.2rem;
  font-weight: bold;
}

.wa-logs-nav-overflow-account {
  max-height: 400px;
  overflow-y: scroll;
}

[data-account-selected] {
  background: #f5f5f5;
}
</style>
