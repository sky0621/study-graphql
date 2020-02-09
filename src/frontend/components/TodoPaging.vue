<template>
  <v-form>
    <!-- 「文字列フィルタ」テキストボックス表示エリア -->
    <v-row>
      <v-col col="5">
        <v-card class="pa-4">
          <v-text-field v-model="search" label="Search"></v-text-field>
        </v-card>
      </v-col>
    </v-row>
    <!-- ページング込みの一覧テーブル表示エリア -->
    <v-row>
      <v-col col="9">
        <v-card>
          <v-data-table :search="search" :headers="headers" fixed-header>
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>
  </v-form>
</template>
<script lang="ts">
import { Component, Vue } from '@/node_modules/nuxt-property-decorator'
// eslint-disable-next-line no-unused-vars
import { DataTableHeader } from '@/types/vuetify'

@Component({})
export default class TodoPaging extends Vue {
  // 文字列フィルタ入力値の受け口
  private readonly search = ''

  // 一覧テーブルのヘッダー表示要素の配列
  private readonly headers: DataTableHeader[] = [
    // eslint-disable-next-line no-use-before-define
    new DataTableHeaderImpl('ID', 'id', false, 50),
    // eslint-disable-next-line no-use-before-define
    new DataTableHeaderImpl('TODO', 'text', true, 50),
    // eslint-disable-next-line no-use-before-define
    new DataTableHeaderImpl('Done', 'done', true, 50),
    // eslint-disable-next-line no-use-before-define
    new DataTableHeaderImpl('CreatedAt(UnixTimestamp)', 'createdAt', true, 50),
    // eslint-disable-next-line no-use-before-define
    new DataTableHeaderImpl('User', 'user.name', false, 50)
  ]
}

class DataTableHeaderImpl implements DataTableHeader {
  text: string
  value: string
  sortable: boolean
  width: number
  constructor(text: string, value: string, sortable: boolean, width: number) {
    this.text = text
    this.value = value
    this.sortable = sortable
    this.width = width
  }
}
</script>
