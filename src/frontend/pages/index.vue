<template>
  <div>
    <TodoTable />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import 'vue-apollo'
import TodoTable from '~/components/TodoTable.vue'

@Component({
  components: { TodoTable }
})
export default class IndexPage extends Vue {}
</script>
