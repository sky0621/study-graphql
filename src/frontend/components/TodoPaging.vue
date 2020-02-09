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
          <v-data-table
            :search="search"
            :headers="headers"
            :options.sync="options"
            fixed-header
          >
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>
  </v-form>
</template>
<script lang="ts">
import { Component, Vue, Watch } from '~/node_modules/nuxt-property-decorator'
// eslint-disable-next-line no-unused-vars
import { DataTableHeader } from '~/types/vuetify'
import todoConnection from '~/apollo/queries/todoConnection.gql'

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

  // private items: Array<String> = ['A', 'b', 'C', 'D', 'E']

  // eslint-disable-next-line no-use-before-define
  private options = new DataTableOptions()

  @Watch('options')
  watchOptions() {
    this.connection()
  }

  private async connection() {
    try {
      const res = await this.$apollo.query({
        query: todoConnection
      })
      console.log(res)
    } catch (e) {
      // this.$toasted.error(e)
      console.log(e)
    }
  }
}

// v-data-tableにおけるヘッダーの定義用
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

// v-data-tableにおけるページング・ソート条件値の受け取り用
class DataTableOptions {
  public page: number = 1
  public itemsPerPage: number = 10
  // MEMO: 現状では一度に指定できるソートキーは１つ
  public sortBy: Array<string> = []
  public sortDesc: Array<boolean> = []
  // MEMO: 以下は、v-data-table の options 要素としては備わっているが現時点では未使用
  // public multiSort: Boolean = false
  // public mustSort: Boolean = false
  // public groupBy: Array<String> = []
  // public groupDesc: Array<Boolean> = []
  // eslint-disable-next-line no-useless-constructor
  constructor() {}
}
</script>
