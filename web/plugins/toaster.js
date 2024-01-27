const toaster = {
  success(message) {
    this.show("success", message);
  },

  error(message) {
    this.show("error", message);
  },

  warning(message) {
    this.show("warning", message);
  },

  show(_type = "default", message) {
    window.$nuxt.$root.$toasted.show(`${message}`, {
      theme: "bubble",
      position: "top-right",
      duration: 5000,
      type: _type
    });
  }
};

export default ({ app }, inject) => {
  inject("toaster", toaster);
};
