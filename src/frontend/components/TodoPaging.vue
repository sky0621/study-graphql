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
            :items="items"
            :options.sync="options"
            :server-items-length="totalCount"
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
}

@Component({})
export default class TodoPaging extends Vue {
  // 文字列フィルタ入力値の受け口
  private readonly search = ''

  // 一覧テーブルのヘッダー表示要素の配列
  private readonly headers: DataTableHeader[] = [
    new DataTableHeaderImpl('ID', 'id', false, 50),
    new DataTableHeaderImpl('TODO', 'text', true, 50),
    new DataTableHeaderImpl('Done', 'done', true, 50),
    new DataTableHeaderImpl('CreatedAt(UnixTimestamp)', 'createdAt', true, 50),
    new DataTableHeaderImpl('User', 'user.name', false, 50)
  ]

  // 一覧テーブルのデータ（v-data-tableの状態変更をウォッチし、その変更を契機にGraphQLクエリ発行→結果を格納）
  // eslint-disable-next-line no-array-constructor
  private items = new Array<Node>()

  // v-data-tableの状態変更をウォッチするための受け皿
  private options = new DataTableOptions()

  // ページングに依らない検索条件に合致する総件数を保持
  private totalCount: number = 0

  // v-data-tableの状態変更をウォッチし、その変更を契機にconnection関数をコール
  @Watch('options')
  watchOptions() {
    this.connection()
  }

  // Apolloライブラリを使ってGraphQLサーバにクエリ発行
  private async connection() {
    try {
      // $apollo.query()がPromiseを返すのでasync/awaitで受け取り
      // まずは、ページング・並べ替え条件等を指定せず、単純にクエリを叩く
      const res = await this.$apollo.query({
        query: todoConnection
      })

      if (res && res.data && res.data.todoConnection) {
        const conn = res.data.todoConnection

        // 一覧表示するデータを抜き出す
        // edges [ node {id, text, done, ...} ]
        this.items = conn.edges.filter((e) => e.node).map((e) => e.node)

        // ページングに依らない検索条件に合致する総件数を保持
        this.totalCount = conn.totalCount
      } else {
        console.log('no result')
      }
    } catch (e) {
      console.log(e)
    }
  }
}
</script>
