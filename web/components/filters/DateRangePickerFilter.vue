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
      dateTimeRange: [],
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
  },
  watch:{
    "$route.query":{
      handler(newValue){
        if(Object.keys(newValue).length === 0){
          this.dateTimeRange = [
            new Date(Date.now() - utils.getDaysToMilliSecond(7)), new Date(Date.now())
          ]
        }
      }
    }
  },
  methods:{
    setDateTimeRangeInQueryParams(e) {
      let from = new Date(e[0]);
      let to = new Date(e[1]);
      from.setHours(0, 0, 0, 0);
      to.setHours(23, 59, 59, 999);
      from = from.getTime();
      to = to.getTime();
      const oldQuery = this.$route.query;
      delete [oldQuery.from, oldQuery.to];
      const newQuery = {
        ...oldQuery,
        from,
        to
      };
      this.$router.push({
        query: newQuery
      });
    },
    setDateTimeRangeInDisplay() {
      const from = Number(this.$route.query.from);
      const to = Number(this.$route.query.to);
      this.dateTimeRange = [new Date(from), new Date(to)];
    }
  },
  mounted(){
    this.setDateTimeRangeInDisplay();
  }
}
</script>
<template>
  <date-picker class="vue-date-picker" :shortcuts="shortcuts" v-model="dateTimeRange" type="date" :placeholder="placeholder" :clearable="false"
    :multiple="true" format="DD-MM-YYYY" @change="setDateTimeRangeInQueryParams" append-to-body range />
</template>

<style scoped>
  .vue-date-picker ::v-deep .mx-input{
    padding: 1.2rem;
  }
  .vue-date-picker {
    width: inherit;
  }
</style>
