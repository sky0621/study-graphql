export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string,
  String: string,
  Boolean: boolean,
  Int: number,
  Float: number,
  /** カーソル（１レコードをユニークに特定する識別子） */
  Cursor: any,
};

/** 前ページ遷移条件 */
export type BackwardPagination = {
  /** 取得件数 */
  last: Scalars['Int'],
  /** 取得対象識別用カーソル（※前ページ遷移時にこのカーソルよりも前にあるレコードが取得対象） */
  before?: Maybe<Scalars['Cursor']>,
};

/** ページングを伴う結果返却用 */
export type Connection = {
  /** ページ情報 */
  pageInfo: PageInfo,
  /** 結果一覧（※カーソル情報を含む） */
  edges: Array<Edge>,
  /** 検索結果の全件数 */
  totalCount: Scalars['Int'],
};


/** 検索結果一覧（※カーソル情報を含む） */
export type Edge = {
  /** Nodeインタフェースを実装したtypeなら代入可能 */
  node?: Maybe<Node>,
  cursor: Scalars['Cursor'],
};

/** 並び替え条件 */
export type EdgeOrder = {
  /** 並べ替えキー項目 */
  key: OrderKey,
  /** ソート方向 */
  direction: OrderDirection,
};

/** 次ページ遷移条件 */
export type ForwardPagination = {
  /** 取得件数 */
  first: Scalars['Int'],
  /** 取得対象識別用カーソル（※次ページ遷移時にこのカーソルよりも後ろにあるレコードが取得対象） */
  after?: Maybe<Scalars['Cursor']>,
};

/** マッチングパターン種別（※要件次第で「前方一致」や「後方一致」も追加） */
export enum MatchingPattern {
  /** 部分一致 */
  PartialMatch = 'PARTIAL_MATCH',
  /** 完全一致 */
  ExactMatch = 'EXACT_MATCH'
}

export type Mutation = {
   __typename?: 'Mutation',
  noop?: Maybe<NoopPayload>,
  createTodo: Scalars['ID'],
  createUser: Scalars['ID'],
};


export type MutationNoopArgs = {
  input?: Maybe<NoopInput>
};


export type MutationCreateTodoArgs = {
  input: NewTodo
};


export type MutationCreateUserArgs = {
  input: NewUser
};

export type NewTodo = {
  text: Scalars['String'],
  userId: Scalars['String'],
};

export type NewUser = {
  name: Scalars['String'],
};

export type Node = {
  id: Scalars['ID'],
};

export type NoopInput = {
  clientMutationId?: Maybe<Scalars['String']>,
};

export type NoopPayload = {
   __typename?: 'NoopPayload',
  clientMutationId?: Maybe<Scalars['String']>,
};

/** 並べ替え方向 */
export enum OrderDirection {
  /** 昇順 */
  Asc = 'ASC',
  /** 降順 */
  Desc = 'DESC'
}

/** 
 * 並べ替えのキー
 * 汎用的な構造にしたいが以下はGraphQLの仕様として不可だった。
 * ・enum・・・汎化機能がない。
 * ・interface・・・inputには実装機能がない。
 * ・union・・・inputでは要素に持てない。
 * とはいえ、並べ替えも共通の仕組みとして提供したく、結果として機能毎に enum フィールドを列挙
 */
export type OrderKey = {
  /** TODO一覧の並べ替えキー */
  todoOrderKey?: Maybe<TodoOrderKey>,
};

/** ページング条件 */
export type PageCondition = {
  /** 前ページ遷移条件 */
  backward?: Maybe<BackwardPagination>,
  /** 次ページ遷移条件 */
  forward?: Maybe<ForwardPagination>,
  /** 現在ページ番号（今回のページング実行前の時点のもの） */
  nowPageNo: Scalars['Int'],
  /** １ページ表示件数 */
  initialLimit?: Maybe<Scalars['Int']>,
};

/** ページ情報 */
export type PageInfo = {
   __typename?: 'PageInfo',
  /** 次ページ有無 */
  hasNextPage: Scalars['Boolean'],
  /** 前ページ有無 */
  hasPreviousPage: Scalars['Boolean'],
  /** 当該ページの１レコード目 */
  startCursor: Scalars['Cursor'],
  /** 当該ページの最終レコード */
  endCursor: Scalars['Cursor'],
};

export type Query = {
   __typename?: 'Query',
  node?: Maybe<Node>,
  todo: Todo,
  /** Relay準拠ページング対応検索によるTODO一覧取得 */
  todoConnection?: Maybe<TodoConnection>,
  users: Array<User>,
  user: User,
};


export type QueryNodeArgs = {
  id: Scalars['ID']
};


export type QueryTodoArgs = {
  id: Scalars['ID']
};


export type QueryTodoConnectionArgs = {
  filterWord?: Maybe<TextFilterCondition>,
  pageCondition?: Maybe<PageCondition>,
  edgeOrder?: Maybe<EdgeOrder>
};


export type QueryUserArgs = {
  id: Scalars['ID']
};

/** 文字列フィルタ条件 */
export type TextFilterCondition = {
  /** フィルタ文字列 */
  filterWord: Scalars['String'],
  /** マッチングパターン（※オプション。指定無しの場合は「部分一致」となる。） */
  matchingPattern?: Maybe<MatchingPattern>,
};

export type Todo = Node & {
   __typename?: 'Todo',
  /** ID */
  id: Scalars['ID'],
  /** TODO */
  text: Scalars['String'],
  /** 済みフラグ */
  done: Scalars['Boolean'],
  /** 作成日時 */
  createdAt: Scalars['Int'],
  /** ユーザー情報 */
  user: User,
};

/** ページングを伴う結果返却用 */
export type TodoConnection = Connection & {
   __typename?: 'TodoConnection',
  /** ページ情報 */
  pageInfo: PageInfo,
  /** 検索結果一覧（※カーソル情報を含む） */
  edges: Array<TodoEdge>,
  /** 検索結果の全件数 */
  totalCount: Scalars['Int'],
};

/** 検索結果一覧（※カーソル情報を含む） */
export type TodoEdge = Edge & {
   __typename?: 'TodoEdge',
  node?: Maybe<Todo>,
  cursor: Scalars['Cursor'],
};

/** TODO並べ替えキー */
export enum TodoOrderKey {
  /** TODO */
  Text = 'TEXT',
  /** 済みフラグ */
  Done = 'DONE',
  /** 作成日時（初期表示時のデフォルト） */
  CreatedAt = 'CREATED_AT',
  /** ユーザー名 */
  UserName = 'USER_NAME'
}

export type User = {
   __typename?: 'User',
  id: Scalars['ID'],
  name: Scalars['String'],
  todos: Array<Todo>,
};

