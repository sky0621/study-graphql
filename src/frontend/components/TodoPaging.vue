<template>
  <v-form>
    <v-row>
      <v-col col="5">
        <v-card class="pa-4">
          <v-text-field v-model="search" label="Search"></v-text-field>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col col="9">
        <v-card>
          <v-data-table
            :headers="headers"
            :items="todos"
            :search="search"
            :items-per-page="itemsPerPage"
            :sort-by="sortBy"
            :sort-desc="sortDesc"
            fixed-header
          >
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>
  </v-form>
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
export default class TodoPaging extends Vue {
  todos: Todo[] = []
  search: string = ''
  itemsPerPage: number = 10
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
        sortable: true,
        text: 'TODO',
        value: 'text'
      },
      {
        sortable: true,
        text: 'Done',
        value: 'done'
      },
      {
        sortable: true,
        text: 'CreatedAt(UnixTimestamp)',
        value: 'createdAt'
      },
      {
        sortable: true,
        text: 'User',
        value: 'user.name'
      }
    ]
  }
}
</script>
