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
            :search="search"
            :headers="headers"
            :items="edgesToNodes(edges)"
            :options.sync="options"
            :server-items-length="totalItems"
            fixed-header
          >
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>
  </v-form>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'nuxt-property-decorator'
import {
  // eslint-disable-next-line no-unused-vars
  EdgeOrder,
  // eslint-disable-next-line no-unused-vars
  PageCondition,
  OrderDirection,
  // eslint-disable-next-line no-unused-vars
  Edge,
  TodoOrderKey,
  // eslint-disable-next-line no-unused-vars
  Todo
} from '~/gql-types.d.ts'

import 'vue-apollo'
import todoConnection from '~/apollo/queries/todoConnection.gql'
// eslint-disable-next-line no-unused-vars
import { DataTableHeader } from '@/types/vuetify'

const DEFAULT_ITEMS_PER_PAGE = 10

// v-data-tableにおけるページング・ソート条件値の受け取り用
class DataTableOptions {
  public page: number = 1
  public itemsPerPage: number = DEFAULT_ITEMS_PER_PAGE
  // MEMO: 現状では一度に指定できるソートキーは１つ
  public sortBy: Array<string> = []
  public sortDesc: Array<boolean> = []
  // MEMO: 以下は、v-data-table の options 要素としては備わっているが現時点では未使用
  // public multiSort: Boolean = false
  // public mustSort: Boolean = false
  // public groupBy: Array<String> = []
  // public groupDesc: Array<Boolean> = []
}

@Component({})
export default class TodoPaging extends Vue {
  // 文字列フィルタ用
  private readonly search = ''

  // ヘッダー表示要素の配列
  private readonly headers: DataTableHeader[] = [
    this.tableHeader('ID', 'id', 30),
    this.tableHeader('TODO', 'text', 200),
    this.tableHeader('Done', 'done', 200),
    this.tableHeader('CreatedAt(UnixTimestamp)', 'createdAt', 100),
    this.tableHeader('User', 'user.name', 50)
  ]

  private loading: boolean = false
  private totalItems: number = 0
  private edges: Array<Edge> = []
  private options = new DataTableOptions()
  private startCursor: string | null = null
  private endCursor: string | null = null
  private nowPage: number = 1

  // 文字列フィルタ欄の入力を監視
  @Watch('search')
  watchSearchWord() {
    this.initPageParam()
    this.searchNodes()
  }

  // ページ遷移 or 1ページあたり表示件数変更 or ソートキー・昇順降順変更を監視
  @Watch('options')
  watchOptions() {
    // MEMO: ソートや１ページあたり表示件数の変更時は「1」が渡される。
    if (this.options.page === 1) {
      this.initPageParam()
    }
    this.searchNodes()
  }

  private async searchNodes() {
    try {
      this.loading = true

      const response = await this.$apollo.query({
        query: todoConnection,
        variables: {
          // 文字列フィルタ条件
          filterWord: this.search,
          // ページング条件
          pageCondition: this.createPageCondition(
            this.nowPage,
            this.options.page,
            this.options.itemsPerPage,
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

      if (response && response.data && response.data.todoConnection) {
        const connection = response.data.todoConnection
        if (connection) {
          // １ページ分の検索結果
          this.edges = connection.edges
          // 検索結果総件数
          this.totalItems = connection.totalCount
          // ページ情報
          const pageInfo = connection.pageInfo
          this.startCursor = pageInfo.startCursor
          this.endCursor = pageInfo.endCursor
          this.nowPage = this.options.page
        }
      }
    } catch (error) {
      console.log(error)
      return
    } finally {
      this.loading = false
    }
  }

  private initPageParam(): void {
    this.nowPage = 1
    this.options.page = 1
  }

  private createPageCondition(
    nowPage: number,
    nextPage: number,
    limit: number,
    startCursor: string | null,
    endCursor: string | null
  ): PageCondition {
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
    // MEMO: 現状では一度に指定できるソートキーは１つ
    if (sortBy.length !== 1 || sortDesc.length !== 1) {
      return null
    }
    const direction = sortDesc[0] ? OrderDirection.Desc : OrderDirection.Asc
    switch (sortBy[0]) {
      case 'text':
        return { key: { todoOrderKey: TodoOrderKey.Text }, direction }
      case 'done':
        return { key: { todoOrderKey: TodoOrderKey.Done }, direction }
      case 'createdAt':
        return { key: { todoOrderKey: TodoOrderKey.CreatedAt }, direction }
    }
    return null
  }

  private edgesToNodes(ea: Array<Edge>): Array<Node> {
    return ea.map((e) => e.node as Todo)
  }

  private tableHeader(
    text: string,
    value: string,
    width: number
  ): DataTableHeader {
    return {
      sortable: false,
      text,
      value,
      width
    }
  }
}
</script>
