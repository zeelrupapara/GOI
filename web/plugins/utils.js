import utils from "@/utils/utils.js";

export default function (context, inject) {
  inject("utils", utils);
  context.$utils = utils;
}
