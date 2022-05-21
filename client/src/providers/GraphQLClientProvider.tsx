import { GraphQLClient } from 'graphql-request';
import React, { createContext } from 'react';

type Props = {
  children: React.ReactNode;
  defaultState?: GraphQLClientState;
};

export type GraphQLClientState = {
  graphQLClient: GraphQLClient;
};

export type GraphQLClientProviderState = GraphQLClientState | null;

export const GraphQLClientContext =
  createContext<GraphQLClientProviderState>(null);

const initialState: GraphQLClientState = {
  graphQLClient: new GraphQLClient('http://localhost:8080/query'),
};

{
  // headers: {
  //   // authorization: accessToken ? `bearer ${accessToken}` : '',
  // },
}

export function GraphQLClientProvider({ children, defaultState }: Props) {
  return (
    <GraphQLClientContext.Provider value={defaultState || initialState}>
      {children}
    </GraphQLClientContext.Provider>
  );
}
