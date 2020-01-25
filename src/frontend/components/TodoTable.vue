<template>
  <v-data-table
    :headers="headers"
    :items="todos"
    :search="search"
    :items-per-page="itemsPerPage"
    :sort-by="sortBy"
    :sort-desc="sortDesc"
  >
  </v-data-table>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import 'vue-apollo'
import todos from '~/apollo/queries/todos.gql'
// eslint-disable-next-line no-unused-vars
import { Todo } from '~/gql-types'
// eslint-disable-next-line no-unused-vars
import { DataTableHeader } from '@/types/vuetify'

@Component({
  apollo: {
    todos: {
      prefetch: true,
      query: todos
    }
  }
})
export default class TodoCard extends Vue {
  todos: Todo[] = []
  search: string = ''
  itemsPerPage: number = 3
  sortBy: string = 'createdAt'
  sortDesc: boolean = true

  get headers(): DataTableHeader[] {
    return [
      {
        sortable: false,
        text: 'ID',
        value: 'id'
      },
      {
        sortable: false,
        text: 'TODO',
        value: 'text'
      },
      {
        sortable: false,
        text: 'Done',
        value: 'done'
      },
      {
        sortable: false,
        text: 'CreatedAt(UnixTimestamp)',
        value: 'createdAt'
      },
      {
        sortable: false,
        text: 'User',
        value: 'user.name'
      }
    ]
  }
}
</script>
