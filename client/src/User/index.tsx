import React, { useEffect } from 'react'
import { AllUsersDocument, AllUsersQuery, AllUsersQueryVariables, NewUsersDocument, useAddFakeUsersMutation, useAllUsersQuery } from '../generated/graphql'

import { UserList } from './UserList'

export const User: React.FC = () => {
    const { loading, data, refetch, subscribeToMore } = useAllUsersQuery({fetchPolicy: "cache-and-network"})
    const [ addFakeUsersMutation ] = useAddFakeUsersMutation()

    const addFakeUsers = async(count: number) => {
        await addFakeUsersMutation({
            variables: {
                count: count,
            },
            update(cache, {data}) {
                const option = { query: AllUsersDocument }
                const { allUsers, totalUsers } = cache.readQuery<AllUsersQuery, AllUsersQueryVariables>(option) || {}
                if (!allUsers || !totalUsers) {
                    return
                }
                cache.writeQuery<AllUsersQuery, AllUsersQueryVariables>({
                    query: AllUsersDocument,
                    data: {
                        totalUsers: totalUsers + count,
                        allUsers: {
                            ...allUsers,
                            ...data?.addFakeUsers,
                        }
                    }
                })
            }
        })
    }

    const subscribeToNewUsers = (githubLogin: string) => {
        subscribeToMore({
            document: NewUsersDocument,
            variables: {githubLogin},
            updateQuery: (prev, { subscriptionData }) => {
                if (!subscriptionData.data) return prev
                const newUsers = subscriptionData.data
                return Object.assign({}, prev, {
                    allUsers: [newUsers, ...prev.allUsers]
                })
            }
        })
    }

    useEffect(() => {
        subscribeToNewUsers("KEY60228") // TODO: 動的に
    }, [])

    return (
        <>
            {loading &&
                <p>loading users...</p>
            }
            {!loading && data &&
                <UserList count={data.totalUsers} users={data.allUsers} refetch={refetch} addFakeUsers={addFakeUsers} />
            }
        </>
    )
}
