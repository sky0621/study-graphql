<template>
  <div>
    <TodoCard />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import 'vue-apollo'
import TodoCard from '~/components/TodoCard.vue'

@Component({
  components: { TodoCard }
})
export default class IndexPage extends Vue {}
</script>
