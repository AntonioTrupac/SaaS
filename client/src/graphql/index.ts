import { GraphQLClient } from 'graphql-request';
import { RequestInit } from 'graphql-request/dist/types.dom';
import { useMutation, useQuery, useInfiniteQuery, UseMutationOptions, UseQueryOptions, UseInfiniteQueryOptions, QueryFunctionContext } from 'react-query';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };

function fetcher<TData, TVariables>(client: GraphQLClient, query: string, variables?: TVariables, headers?: RequestInit['headers']) {
  return async (): Promise<TData> => client.request<TData, TVariables>(query, variables, headers);
}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Any: any;
};

export type Actor = {
  age: Scalars['Int'];
  email: Scalars['String'];
  firstName: Scalars['String'];
  lastName: Scalars['String'];
  password: Scalars['String'];
  phone: Scalars['String'];
};

export type Address = {
  __typename?: 'Address';
  addressLine: Scalars['String'];
  city: Scalars['String'];
  country: Scalars['String'];
  id: Scalars['Int'];
  postalCode: Scalars['Int'];
  userAuthId?: Maybe<Scalars['Int']>;
};

export type AddressInput = {
  addressLine: Scalars['String'];
  city: Scalars['String'];
  country: Scalars['String'];
  postalCode: Scalars['Int'];
};

export type AuthOps = {
  __typename?: 'AuthOps';
  login: Scalars['Any'];
  refreshToken: Scalars['Any'];
  register: Scalars['Any'];
};


export type AuthOpsLoginArgs = {
  email: Scalars['String'];
  password: Scalars['String'];
};


export type AuthOpsRefreshTokenArgs = {
  input: RefreshTokenInput;
};


export type AuthOpsRegisterArgs = {
  input: NewAuthUser;
};

export type AuthUser = Actor & {
  __typename?: 'AuthUser';
  address: Array<Address>;
  age: Scalars['Int'];
  email: Scalars['String'];
  firstName: Scalars['String'];
  id: Scalars['Int'];
  lastName: Scalars['String'];
  password: Scalars['String'];
  phone: Scalars['String'];
};

export type Category = {
  __typename?: 'Category';
  id: Scalars['Int'];
  name: Scalars['String'];
};

export type CategoryInput = {
  name: Scalars['String'];
};

export type HabitSettings = {
  __typename?: 'HabitSettings';
  complete: Scalars['Int'];
  currentStreak: Scalars['Int'];
  failed: Scalars['Int'];
  habitId: Scalars['Int'];
  id: Scalars['Int'];
  skipped: Scalars['Int'];
  total: Scalars['Int'];
};

export type HabitSettingsInput = {
  complete: Scalars['Int'];
  currentStreak: Scalars['Int'];
  failed: Scalars['Int'];
  skipped: Scalars['Int'];
  total: Scalars['Int'];
};

export type Habits = {
  __typename?: 'Habits';
  habitSetting?: Maybe<HabitSettings>;
  id: Scalars['Int'];
  name: Scalars['String'];
  userId: Scalars['String'];
};

export type HabitsInput = {
  habitSetting?: InputMaybe<HabitSettingsInput>;
  name: Scalars['String'];
};

export type Image = {
  __typename?: 'Image';
  id: Scalars['Int'];
  productId: Scalars['Int'];
  url: Scalars['String'];
};

export type ImageInput = {
  url: Scalars['String'];
};

export type MoodType = {
  __typename?: 'MoodType';
  id: Scalars['Int'];
  name: Scalars['String'];
};

export type MoodTypeInput = {
  name: Scalars['String'];
};

export type Moods = {
  __typename?: 'Moods';
  id: Scalars['Int'];
  moodTypeId: Scalars['Int'];
  notes: Scalars['String'];
  userId: Scalars['Int'];
};

export type MoodsInput = {
  notes: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  auth: AuthOps;
  createHabits: Habits;
  createMoodTypes?: Maybe<MoodType>;
  createMoods?: Maybe<Moods>;
  createProducts: Product;
};


export type MutationCreateHabitsArgs = {
  input: HabitsInput;
};


export type MutationCreateMoodTypesArgs = {
  input?: InputMaybe<MoodTypeInput>;
};


export type MutationCreateMoodsArgs = {
  input: MoodsInput;
  typeId: Scalars['Int'];
};


export type MutationCreateProductsArgs = {
  input: ProductInput;
};

export type NewAuthUser = {
  address: Array<AddressInput>;
  age: Scalars['Int'];
  email: Scalars['String'];
  firstName: Scalars['String'];
  lastName: Scalars['String'];
  password: Scalars['String'];
  phone: Scalars['String'];
};

export type Node = {
  id: Scalars['Int'];
};

export type Product = {
  __typename?: 'Product';
  category: Array<Category>;
  description: Scalars['String'];
  id: Scalars['Int'];
  image: Array<Image>;
  name: Scalars['String'];
  price: Scalars['Float'];
  rating: Scalars['Float'];
  stock: Scalars['Int'];
};

export type ProductInput = {
  categories: Array<CategoryInput>;
  description: Scalars['String'];
  images: Array<ImageInput>;
  name: Scalars['String'];
  price: Scalars['Float'];
  rating: Scalars['Float'];
  stock: Scalars['Int'];
};

export type Query = {
  __typename?: 'Query';
  getHabits: Array<Habits>;
  getHabitsById?: Maybe<Habits>;
  getMoodById: Moods;
  getMoodTypes: Array<MoodType>;
  getMoods: Array<Moods>;
  getProductById?: Maybe<Product>;
  getProducts: Array<Product>;
  getUser: Scalars['Any'];
};


export type QueryGetHabitsByIdArgs = {
  id: Scalars['Int'];
};


export type QueryGetMoodByIdArgs = {
  id: Scalars['Int'];
};


export type QueryGetProductByIdArgs = {
  id: Scalars['Int'];
};


export type QueryGetUserArgs = {
  id: Scalars['Int'];
};

export type RefreshTokenInput = {
  token: Scalars['String'];
};

export type RegisterVariables = Exact<{
  input: NewAuthUser;
}>;


export type Register = { __typename?: 'Mutation', auth: { __typename?: 'AuthOps', register: any } };

export type LoginVariables = Exact<{
  email: Scalars['String'];
  password: Scalars['String'];
}>;


export type Login = { __typename?: 'Mutation', auth: { __typename?: 'AuthOps', login: any } };

export type GetMoodTypesVariables = Exact<{ [key: string]: never; }>;


export type GetMoodTypes = { __typename?: 'Query', getMoodTypes: Array<{ __typename?: 'MoodType', id: number, name: string }> };


export const RegisterDocument = `
    mutation Register($input: NewAuthUser!) {
  auth {
    register(input: $input)
  }
}
    `;
export const useRegister = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<Register, TError, RegisterVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<Register, TError, RegisterVariables, TContext>(
      ['Register'],
      (variables?: RegisterVariables) => fetcher<Register, RegisterVariables>(client, RegisterDocument, variables, headers)(),
      options
    );
export const LoginDocument = `
    mutation Login($email: String!, $password: String!) {
  auth {
    login(email: $email, password: $password)
  }
}
    `;
export const useLogin = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<Login, TError, LoginVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<Login, TError, LoginVariables, TContext>(
      ['Login'],
      (variables?: LoginVariables) => fetcher<Login, LoginVariables>(client, LoginDocument, variables, headers)(),
      options
    );
export const GetMoodTypesDocument = `
    query GetMoodTypes {
  getMoodTypes {
    id
    name
  }
}
    `;
export const useGetMoodTypes = <
      TData = GetMoodTypes,
      TError = unknown
    >(
      client: GraphQLClient,
      variables?: GetMoodTypesVariables,
      options?: UseQueryOptions<GetMoodTypes, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<GetMoodTypes, TError, TData>(
      variables === undefined ? ['GetMoodTypes'] : ['GetMoodTypes', variables],
      fetcher<GetMoodTypes, GetMoodTypesVariables>(client, GetMoodTypesDocument, variables, headers),
      options
    );
export const useInfiniteGetMoodTypes = <
      TData = GetMoodTypes,
      TError = unknown
    >(
      pageParamKey: keyof GetMoodTypesVariables,
      client: GraphQLClient,
      variables?: GetMoodTypesVariables,
      options?: UseInfiniteQueryOptions<GetMoodTypes, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useInfiniteQuery<GetMoodTypes, TError, TData>(
      variables === undefined ? ['GetMoodTypes.infinite'] : ['GetMoodTypes.infinite', variables],
      (metaData) => fetcher<GetMoodTypes, GetMoodTypesVariables>(client, GetMoodTypesDocument, {...variables, ...(metaData.pageParam ?? {})}, headers)(),
      options
    );
