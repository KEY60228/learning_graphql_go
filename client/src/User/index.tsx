import React from 'react'
import { AllUsersDocument, AllUsersQuery, AllUsersQueryVariables, useAddFakeUsersMutation, useAllUsersQuery } from '../generated/graphql'

import { UserList } from './UserList'

export const User: React.FC = () => {
    const { loading, data, refetch } = useAllUsersQuery({fetchPolicy: "cache-and-network"})
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
