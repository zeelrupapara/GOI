<script>
import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';
import utils from "@/utils/utils.js";
export default {
  components: {
    DatePicker
  },
  props: {
    placeholder: {
      type: String,
      default: "Select Options"
    }
  },
  data() {
    return {
      value: null,
      shortcuts: [
        {
          text: "Today",
          onClick: () => [new Date(Date.now()), new Date(Date.now())]
        }, {
          text: "7 Days",
          onClick: () => [
            new Date(Date.now() - utils.getDaysToMilliSecond(7)), new Date(Date.now())
          ]
        },
        {
          text: "30 Days",
          onClick: () => [
            new Date(Date.now() - utils.getDaysToMilliSecond(30)),
            new Date(Date.now())
          ]
        }
      ]
    }
  }
}
</script>
<template>
  <date-picker :shortcuts="shortcuts" v-model="value" type="date" :placeholder="placeholder" :clearable="false"
    :multiple="true" format="DD-MM-YYYY" append-to-body range />
</template>



