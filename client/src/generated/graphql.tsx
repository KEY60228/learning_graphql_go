import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  DateTime: any;
};

export type AuthPayload = {
  __typename?: 'AuthPayload';
  token: Scalars['String'];
  user: User;
};

export type Mutation = {
  __typename?: 'Mutation';
  addFakeUsers: Array<User>;
  githubAuth: AuthPayload;
  postPhoto: Photo;
};


export type MutationAddFakeUsersArgs = {
  count?: Scalars['Int'];
};


export type MutationGithubAuthArgs = {
  code: Scalars['String'];
};


export type MutationPostPhotoArgs = {
  input: PostPhotoInput;
};

export type Photo = {
  __typename?: 'Photo';
  category: PhotoCategory;
  created: Scalars['DateTime'];
  description?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  name: Scalars['String'];
  postedBy: User;
  taggedUsers: Array<User>;
  url: Scalars['String'];
};

export enum PhotoCategory {
  Action = 'ACTION',
  Graphic = 'GRAPHIC',
  Landscape = 'LANDSCAPE',
  Portrait = 'PORTRAIT',
  Selfie = 'SELFIE'
}

export type PostPhotoInput = {
  category?: InputMaybe<PhotoCategory>;
  description?: InputMaybe<Scalars['String']>;
  name: Scalars['String'];
  taggedUserIDs: Array<Scalars['String']>;
};

export type Query = {
  __typename?: 'Query';
  allPhotos: Array<Photo>;
  allUsers: Array<User>;
  me?: Maybe<User>;
  totalPhotos: Scalars['Int'];
  totalUsers: Scalars['Int'];
};

export type Subscription = {
  __typename?: 'Subscription';
  newPhoto: Photo;
  newUsers: Array<User>;
};


export type SubscriptionNewPhotoArgs = {
  githubLogin: Scalars['String'];
};


export type SubscriptionNewUsersArgs = {
  githubLogin: Scalars['String'];
};

export type User = {
  __typename?: 'User';
  avatar: Scalars['String'];
  githubLogin: Scalars['ID'];
  name: Scalars['String'];
};

export type GithubAuthMutationVariables = Exact<{
  code: Scalars['String'];
}>;


export type GithubAuthMutation = { __typename?: 'Mutation', githubAuth: { __typename?: 'AuthPayload', token: string } };

export type UserInfoFragment = { __typename?: 'User', githubLogin: string, name: string, avatar: string };

export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = { __typename?: 'Query', me?: { __typename?: 'User', githubLogin: string, name: string, avatar: string } | null };

export type AllUsersQueryVariables = Exact<{ [key: string]: never; }>;


export type AllUsersQuery = { __typename?: 'Query', totalUsers: number, allUsers: Array<{ __typename?: 'User', githubLogin: string, name: string, avatar: string }> };

export type AddFakeUsersMutationVariables = Exact<{
  count: Scalars['Int'];
}>;


export type AddFakeUsersMutation = { __typename?: 'Mutation', addFakeUsers: Array<{ __typename?: 'User', githubLogin: string, name: string, avatar: string }> };

export type NewUsersSubscriptionVariables = Exact<{
  githubLogin: Scalars['String'];
}>;


export type NewUsersSubscription = { __typename?: 'Subscription', newUsers: Array<{ __typename?: 'User', githubLogin: string, name: string, avatar: string }> };

export type NewPhotoSubscriptionVariables = Exact<{
  githubLogin: Scalars['String'];
}>;


export type NewPhotoSubscription = { __typename?: 'Subscription', newPhoto: { __typename?: 'Photo', url: string, category: PhotoCategory, postedBy: { __typename?: 'User', githubLogin: string, name: string, avatar: string } } };

export const UserInfoFragmentDoc = gql`
    fragment userInfo on User {
  githubLogin
  name
  avatar
}
    `;
export const GithubAuthDocument = gql`
    mutation githubAuth($code: String!) {
  githubAuth(code: $code) {
    token
  }
}
    `;
export type GithubAuthMutationFn = Apollo.MutationFunction<GithubAuthMutation, GithubAuthMutationVariables>;

/**
 * __useGithubAuthMutation__
 *
 * To run a mutation, you first call `useGithubAuthMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useGithubAuthMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [githubAuthMutation, { data, loading, error }] = useGithubAuthMutation({
 *   variables: {
 *      code: // value for 'code'
 *   },
 * });
 */
export function useGithubAuthMutation(baseOptions?: Apollo.MutationHookOptions<GithubAuthMutation, GithubAuthMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<GithubAuthMutation, GithubAuthMutationVariables>(GithubAuthDocument, options);
      }
export type GithubAuthMutationHookResult = ReturnType<typeof useGithubAuthMutation>;
export type GithubAuthMutationResult = Apollo.MutationResult<GithubAuthMutation>;
export type GithubAuthMutationOptions = Apollo.BaseMutationOptions<GithubAuthMutation, GithubAuthMutationVariables>;
export const MeDocument = gql`
    query me {
  me {
    ...userInfo
  }
}
    ${UserInfoFragmentDoc}`;

/**
 * __useMeQuery__
 *
 * To run a query within a React component, call `useMeQuery` and pass it any options that fit your needs.
 * When your component renders, `useMeQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMeQuery({
 *   variables: {
 *   },
 * });
 */
export function useMeQuery(baseOptions?: Apollo.QueryHookOptions<MeQuery, MeQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<MeQuery, MeQueryVariables>(MeDocument, options);
      }
export function useMeLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<MeQuery, MeQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<MeQuery, MeQueryVariables>(MeDocument, options);
        }
export type MeQueryHookResult = ReturnType<typeof useMeQuery>;
export type MeLazyQueryHookResult = ReturnType<typeof useMeLazyQuery>;
export type MeQueryResult = Apollo.QueryResult<MeQuery, MeQueryVariables>;
export const AllUsersDocument = gql`
    query allUsers {
  totalUsers
  allUsers {
    githubLogin
    name
    avatar
  }
}
    `;

/**
 * __useAllUsersQuery__
 *
 * To run a query within a React component, call `useAllUsersQuery` and pass it any options that fit your needs.
 * When your component renders, `useAllUsersQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAllUsersQuery({
 *   variables: {
 *   },
 * });
 */
export function useAllUsersQuery(baseOptions?: Apollo.QueryHookOptions<AllUsersQuery, AllUsersQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<AllUsersQuery, AllUsersQueryVariables>(AllUsersDocument, options);
      }
export function useAllUsersLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<AllUsersQuery, AllUsersQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<AllUsersQuery, AllUsersQueryVariables>(AllUsersDocument, options);
        }
export type AllUsersQueryHookResult = ReturnType<typeof useAllUsersQuery>;
export type AllUsersLazyQueryHookResult = ReturnType<typeof useAllUsersLazyQuery>;
export type AllUsersQueryResult = Apollo.QueryResult<AllUsersQuery, AllUsersQueryVariables>;
export const AddFakeUsersDocument = gql`
    mutation addFakeUsers($count: Int!) {
  addFakeUsers(count: $count) {
    githubLogin
    name
    avatar
  }
}
    `;
export type AddFakeUsersMutationFn = Apollo.MutationFunction<AddFakeUsersMutation, AddFakeUsersMutationVariables>;

/**
 * __useAddFakeUsersMutation__
 *
 * To run a mutation, you first call `useAddFakeUsersMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddFakeUsersMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addFakeUsersMutation, { data, loading, error }] = useAddFakeUsersMutation({
 *   variables: {
 *      count: // value for 'count'
 *   },
 * });
 */
export function useAddFakeUsersMutation(baseOptions?: Apollo.MutationHookOptions<AddFakeUsersMutation, AddFakeUsersMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddFakeUsersMutation, AddFakeUsersMutationVariables>(AddFakeUsersDocument, options);
      }
export type AddFakeUsersMutationHookResult = ReturnType<typeof useAddFakeUsersMutation>;
export type AddFakeUsersMutationResult = Apollo.MutationResult<AddFakeUsersMutation>;
export type AddFakeUsersMutationOptions = Apollo.BaseMutationOptions<AddFakeUsersMutation, AddFakeUsersMutationVariables>;
export const NewUsersDocument = gql`
    subscription newUsers($githubLogin: String!) {
  newUsers(githubLogin: $githubLogin) {
    githubLogin
    name
    avatar
  }
}
    `;

/**
 * __useNewUsersSubscription__
 *
 * To run a query within a React component, call `useNewUsersSubscription` and pass it any options that fit your needs.
 * When your component renders, `useNewUsersSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useNewUsersSubscription({
 *   variables: {
 *      githubLogin: // value for 'githubLogin'
 *   },
 * });
 */
export function useNewUsersSubscription(baseOptions: Apollo.SubscriptionHookOptions<NewUsersSubscription, NewUsersSubscriptionVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useSubscription<NewUsersSubscription, NewUsersSubscriptionVariables>(NewUsersDocument, options);
      }
export type NewUsersSubscriptionHookResult = ReturnType<typeof useNewUsersSubscription>;
export type NewUsersSubscriptionResult = Apollo.SubscriptionResult<NewUsersSubscription>;
export const NewPhotoDocument = gql`
    subscription newPhoto($githubLogin: String!) {
  newPhoto(githubLogin: $githubLogin) {
    url
    category
    postedBy {
      githubLogin
      name
      avatar
    }
  }
}
    `;

/**
 * __useNewPhotoSubscription__
 *
 * To run a query within a React component, call `useNewPhotoSubscription` and pass it any options that fit your needs.
 * When your component renders, `useNewPhotoSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useNewPhotoSubscription({
 *   variables: {
 *      githubLogin: // value for 'githubLogin'
 *   },
 * });
 */
export function useNewPhotoSubscription(baseOptions: Apollo.SubscriptionHookOptions<NewPhotoSubscription, NewPhotoSubscriptionVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useSubscription<NewPhotoSubscription, NewPhotoSubscriptionVariables>(NewPhotoDocument, options);
      }
export type NewPhotoSubscriptionHookResult = ReturnType<typeof useNewPhotoSubscription>;
export type NewPhotoSubscriptionResult = Apollo.SubscriptionResult<NewPhotoSubscription>;