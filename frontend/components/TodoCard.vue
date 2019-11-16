<template>
  <div>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card>
          <v-list two-line subheader>
            <v-list-item v-for="todo in todos" :key="todo.id" link>
              <v-list-item-avatar>
                <v-icon>mdi-gift-outline</v-icon>
              </v-list-item-avatar>
              <v-list-item-content>
                <v-list-item-title>{{ todo.text }}</v-list-item-title>
                <v-list-item-subtitle>{{ todo.done }}</v-list-item-subtitle>
              </v-list-item-content>
              <v-list-item-content>
                <v-list-item-title>
                  {{ todo.user.name }}
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import todos from '~/apollo/queries/todos.gql'
export default {
  data() {
    return {
      todos: []
    }
  },

  apollo: {
    todos: {
      prefetch: true,
      query: todos
    }
  }
}
</script>
