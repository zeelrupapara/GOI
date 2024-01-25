<template>
  <div class="row justify-content-center">
    <div v-for="select in multiselectData" :key="select.id" class="col">
      <MultiSelectDropdownFilter id="select_users" :placeholder="select.placeholder" />
    </div>
    <div class="col">
      <DateRangePickerFilter />
    </div>
  </div>
</template>

<script>
import MultiSelectDropdownFilter from '~/components/filters/MultiSelectDropdownFilter.vue';
import DateRangePickerFilter from '~/components/filters/DateRangePickerFilter.vue';
export default {
  components: {
    MultiSelectDropdownFilter,
    DateRangePickerFilter
  },
  data() {
    return {
      multiselectData: [
        {
          id: 1,
          placeholder: "Select Users"
        },
        {
          id: 1,
          placeholder: "Select Organizations"
        },
        {
          id: 1,
          placeholder: "Select Repositories"
        },
      ],
      timeRangeSelectPlaceholder: "Select Date Range"
    }
  }
}
</script>
