import constants from "@/config/constants.js";

export default function (context, inject) {
  inject("constants", constants);
  context.$constants = constants;
}
