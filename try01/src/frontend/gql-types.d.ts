export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Int64: any;
  /** カーソル（１レコードをユニークに特定する識別子） */
  Cursor: any;
};

/** 検索結果一覧（※カーソル情報を含む） */
export type TodoEdge = Edge & {
  __typename?: 'TodoEdge';
  node: Todo;
  cursor: Scalars['Cursor'];
};

/** TODO並べ替えキー */
export enum TodoOrderKey {
  /** ID */
  Id = 'ID',
  /** TODO */
  Task = 'TASK',
  /** ユーザー名 */
  UserName = 'USER_NAME'
}

/** 検索結果一覧（※カーソル情報を含む） */
export type Edge = {
  /** Nodeインタフェースを実装したtypeなら代入可能 */
  node: Node;
  cursor: Scalars['Cursor'];
};

/** 検索結果一覧（※カーソル情報を含む） */
export type CustomerEdge = Edge & {
  __typename?: 'CustomerEdge';
  node: Customer;
  cursor: Scalars['Cursor'];
};

/** 並び替え条件 */
export type EdgeOrder = {
  /** 並べ替えキー項目 */
  key: OrderKey;
  /** ソート方向 */
  direction: OrderDirection;
};

/** ページング条件 */
export type PageCondition = {
  /** 前ページ遷移条件 */
  backward?: Maybe<BackwardPagination>;
  /** 次ページ遷移条件 */
  forward?: Maybe<ForwardPagination>;
  /** 現在ページ番号（今回のページング実行前の時点のもの） */
  nowPageNo: Scalars['Int64'];
  /** １ページ表示件数 */
  initialLimit: Scalars['Int64'];
};

/** ページングを伴う結果返却用 */
export type Connection = {
  /** ページ情報 */
  pageInfo: PageInfo;
  /** 結果一覧（※カーソル情報を含む） */
  edges: Array<Edge>;
  /** 検索結果の全件数 */
  totalCount: Scalars['Int64'];
};

/** ページングを伴う結果返却用 */
export type CustomerConnection = Connection & {
  __typename?: 'CustomerConnection';
  /** ページ情報 */
  pageInfo: PageInfo;
  /** 検索結果一覧（※カーソル情報を含む） */
  edges: Array<CustomerEdge>;
  /** 検索結果の全件数 */
  totalCount: Scalars['Int64'];
};

/** TODO */
export type Todo = Node & {
  __typename?: 'Todo';
  /** ID */
  id: Scalars['ID'];
  /** タスク */
  task: Scalars['String'];
  /** ユーザー情報 */
  customer: Customer;
};

/** ページ情報 */
export type PageInfo = {
  __typename?: 'PageInfo';
  /** 次ページ有無 */
  hasNextPage: Scalars['Boolean'];
  /** 前ページ有無 */
  hasPreviousPage: Scalars['Boolean'];
  /** 当該ページの１レコード目 */
  startCursor: Scalars['Cursor'];
  /** 当該ページの最終レコード */
  endCursor: Scalars['Cursor'];
};

/**
 * 並べ替えのキー
 * 
 * 【検討経緯】
 * 汎用的な構造、かつ、タイプセーフにしたく、interface で定義の上、機能毎に input ないし enum で実装しようとした。
 * しかし、input は interface を実装できない仕様だったので諦めた。
 * enum に継承機能があればよかったが、それもなかった。
 * union で TodoOrderKey や（増えたら）他の機能の並べ替えのキーも | でつなぐ方法も考えたが、
 * union を input に要素として持たせることはできない仕様だったので、これも諦めた。
 * とはいえ、並べ替えも共通の仕組みとして提供したく、結果として機能毎の enum フィールドを共通の input 内に列挙していく形にした。
 */
export type OrderKey = {
  /** ユーザー一覧の並べ替えキー */
  customerOrderKey?: Maybe<CustomerOrderKey>;
  /** TODO一覧の並べ替えキー */
  todoOrderKey?: Maybe<TodoOrderKey>;
};

/** ユーザー */
export type Customer = Node & {
  __typename?: 'Customer';
  /** ID */
  id: Scalars['ID'];
  /** 名前 */
  name: Scalars['String'];
  /** 年齢 */
  age: Scalars['Int'];
  /** Todo */
  todos: Array<Todo>;
};

/** 文字列フィルタ条件 */
export type TextFilterCondition = {
  /** フィルタ文字列 */
  filterWord: Scalars['String'];
  /** マッチングパターン */
  matchingPattern: MatchingPattern;
};

/** ページングを伴う結果返却用 */
export type TodoConnection = Connection & {
  __typename?: 'TodoConnection';
  /** ページ情報 */
  pageInfo: PageInfo;
  /** 検索結果一覧（※カーソル情報を含む） */
  edges: Array<TodoEdge>;
  /** 検索結果の全件数 */
  totalCount: Scalars['Int64'];
};


/** 並べ替え方向 */
export enum OrderDirection {
  /** 昇順 */
  Asc = 'ASC',
  /** 降順 */
  Desc = 'DESC'
}

/** 次ページ遷移条件 */
export type ForwardPagination = {
  /** 取得件数 */
  first: Scalars['Int64'];
  /** 取得対象識別用カーソル（※次ページ遷移時にこのカーソルよりも後ろにあるレコードが取得対象） */
  after: Scalars['Cursor'];
};

export type Query = {
  __typename?: 'Query';
  node?: Maybe<Node>;
  /** Relay準拠ページング対応検索によるTODO一覧取得 */
  customerConnection?: Maybe<CustomerConnection>;
  /** Relay準拠ページング対応検索によるTODO一覧取得 */
  todoConnection?: Maybe<TodoConnection>;
};


export type QueryNodeArgs = {
  id: Scalars['ID'];
};


export type QueryCustomerConnectionArgs = {
  pageCondition?: Maybe<PageCondition>;
  edgeOrder?: Maybe<EdgeOrder>;
  filterWord?: Maybe<TextFilterCondition>;
};


export type QueryTodoConnectionArgs = {
  pageCondition?: Maybe<PageCondition>;
  edgeOrder?: Maybe<EdgeOrder>;
  filterWord?: Maybe<TextFilterCondition>;
};

/** マッチングパターン種別（※要件次第で「前方一致」や「後方一致」も追加） */
export enum MatchingPattern {
  /** 部分一致 */
  PartialMatch = 'PARTIAL_MATCH',
  /** 完全一致 */
  ExactMatch = 'EXACT_MATCH'
}

/** 前ページ遷移条件 */
export type BackwardPagination = {
  /** 取得件数 */
  last: Scalars['Int64'];
  /** 取得対象識別用カーソル（※前ページ遷移時にこのカーソルよりも前にあるレコードが取得対象） */
  before: Scalars['Cursor'];
};

/** ユーザー並べ替えキー */
export enum CustomerOrderKey {
  /** ID */
  Id = 'ID',
  /** ユーザー名 */
  Name = 'NAME',
  /** 年齢 */
  Age = 'AGE'
}


export type Node = {
  id: Scalars['ID'];
};
