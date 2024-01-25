<template>
  <div>
    <FilterSection />
    <MatrixSection />
  </div>
</template>

<script>
import FilterSection from '~/components/sections/FilterSection.vue';
import MatrixSection from '~/components/sections/MatrixSection.vue';
export default {
  name: 'IndexPage',
  components: {
    FilterSection,
    MatrixSection
  }
}
</script>
