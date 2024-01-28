<template>
  <div>
    <FilterSection />
    <MatrixSection />
    <ChartSection />
  </div>
</template>

<script>
import FilterSection from '~/components/sections/FilterSection.vue';
import MatrixSection from '~/components/sections/MatrixSection.vue';
import ChartSection from '~/components/sections/ChartSection.vue';
export default {
  name: 'Home',
  components: {
    FilterSection,
    MatrixSection,
    ChartSection
  }
}
</script>
