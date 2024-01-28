<template>
  <div>
    <FilterSection />
    <MatrixSection />
    <ChartSection />
    <TableSection />
  </div>
</template>

<script>
import FilterSection from '~/components/sections/FilterSection.vue';
import MatrixSection from '~/components/sections/MatrixSection.vue';
import ChartSection from '~/components/sections/ChartSection.vue';
import TableSection from '~/components/sections/TableSection.vue';
export default {
  name: 'Home',
  components: {
    FilterSection,
    MatrixSection,
    ChartSection,
    TableSection
  }
}
</script>
