export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string,
  String: string,
  Boolean: boolean,
  Int: number,
  Float: number,
};

export type Mutation = {
   __typename?: 'Mutation',
  createTodo: Scalars['ID'],
  createUser: Scalars['ID'],
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

export type Query = {
   __typename?: 'Query',
  todos: Array<Todo>,
  todo: Todo,
  users: Array<User>,
  user: User,
};


export type QueryTodoArgs = {
  id: Scalars['ID']
};


export type QueryUserArgs = {
  id: Scalars['ID']
};

export type Todo = {
   __typename?: 'Todo',
  id: Scalars['ID'],
  text: Scalars['String'],
  done: Scalars['Boolean'],
  user: User,
};

export type User = {
   __typename?: 'User',
  id: Scalars['ID'],
  name: Scalars['String'],
  todos: Array<Todo>,
};

