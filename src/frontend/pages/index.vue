<template>
  <div>
    <TodoPaging />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import 'vue-apollo'
import TodoPaging from '~/components/TodoPaging.vue'

@Component({
  components: { TodoPaging }
})
export default class IndexPage extends Vue {}
</script>
