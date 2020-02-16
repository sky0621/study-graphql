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
    <!-- 各種ボタン表示エリア -->
    <v-row>
      <v-col col="9">
        <v-btn @click="createCsv">CSVダウンロード</v-btn>
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
import createCsv from '~/apollo/mutations/todo.gql'
// eslint-disable-next-line no-unused-vars
import { Edge, EdgeOrder, PageCondition } from '~/gql-types'

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

  // 今回のページの１番目のレコードを表す識別子
  private startCursor: string | null = null

  // 今回のページの最後のレコードを表す識別子
  private endCursor: string | null = null

  // 現在のページを表す（これも、GraphQLサーバに渡すパラメータとして必要）
  private nowPage: number = 1

  // 文字列フィルタ欄の入力を監視
  @Watch('search')
  watchSearchWord() {
    this.initPageParam()
    this.connection()
  }

  // v-data-tableの状態変更をウォッチし、その変更を契機にconnection関数をコール
  @Watch('options')
  watchOptions() {
    // MEMO: ソートや１ページあたり表示件数の変更時は「1」が渡される。
    if (this.options.page === 1) {
      this.initPageParam()
    }
    this.connection()
  }

  // 初期表示時やページング条件をクリアしたいタイミングでコールする関数
  private initPageParam(): void {
    this.nowPage = 1
    this.options.page = 1
  }

  private async createCsv() {
    try {
      const res = await this.$apollo.mutate({
        mutation: createCsv
      })
      console.log(res)
    } catch (e) {
      console.log(e)
    }
  }

  // Apolloライブラリを使ってGraphQLサーバにクエリ発行
  private async connection() {
    try {
      // $apollo.query()がPromiseを返すのでasync/awaitで受け取り
      const res = await this.$apollo.query({
        query: todoConnection,
        variables: {
          // 文字列フィルタ条件
          filterWord: { filterWord: this.search },
          // ページング条件
          pageCondition: this.createPageCondition(
            this.nowPage, // 現在のページ
            this.options.page, // 遷移先のページ
            this.options.itemsPerPage, // １ページあたりの表示件数指定
            this.startCursor,
            this.endCursor
          ),
          // 並び替え条件
          edgeOrder: this.createEdgeOrder(
            this.options.sortBy,
            this.options.sortDesc
          )
        }
      })

      if (res && res.data && res.data.todoConnection) {
        const conn = res.data.todoConnection

        // 一覧表示するデータを抜き出す
        // edges [ node {id, text, done, ...} ]
        this.items = conn.edges
          .filter((e: Edge) => e.node)
          .map((e: Edge) => e.node)

        // ページングに依らない検索条件に合致する総件数を保持
        this.totalCount = conn.totalCount

        // v-data-tableのoptions変更に影響する各種ページ情報を保持
        const pageInfo = conn.pageInfo
        this.startCursor = pageInfo.startCursor
        this.endCursor = pageInfo.endCursor

        this.nowPage = this.options.page
      } else {
        console.log('no result')
      }
    } catch (e) {
      console.log(e)
    }
  }

  private createPageCondition(
    nowPage: number,
    nextPage: number,
    limit: number,
    startCursor: string | null,
    endCursor: string | null
  ): PageCondition {
    // 現在のページと遷移指示先のページとの比較によって「次へ(forward)」なのか「前へ(backward)」なのか判別
    return {
      forward: nowPage < nextPage ? { first: limit, after: endCursor } : null,
      backward:
        nowPage > nextPage ? { last: limit, before: startCursor } : null,
      nowPageNo: nowPage,
      initialLimit: limit > 0 ? limit : null
    }
  }

  private createEdgeOrder(
    sortBy: Array<string>,
    sortDesc: Array<boolean>
  ): EdgeOrder | null {
    if (sortBy && sortDesc) {
      // MEMO: 現状では一度に指定できるソートキーは１つ
      if (sortBy.length !== 1 || sortDesc.length !== 1) {
        return null
      }
      // TODO: enum値を指定するとビルドが通らなくなるので、やむなく文字列で指定
      const direction = sortDesc[0] ? 'DESC' : 'ASC'
      switch (sortBy[0]) {
        case 'text':
          return { key: { todoOrderKey: 'TEXT' }, direction }
        case 'done':
          return { key: { todoOrderKey: 'DONE' }, direction }
        case 'createdAt':
          return { key: { todoOrderKey: 'CREATED_AT' }, direction }
      }
    }
    return null
  }
}
</script>
