<template>
  <CustomerPaging />
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import CustomerPaging from '~/components/CustomerPaging.vue'
@Component({
  components: { CustomerPaging }
})
export default class IndexPage extends Vue {}
</script>
