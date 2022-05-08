import React from 'react'
import { useAddFakeUsersMutation, useAllUsersQuery } from '../generated/graphql'

import { UserList } from './UserList'

export const User: React.FC = () => {
    const { loading, data, refetch } = useAllUsersQuery()
    const [ addFakeUsersMutation ] = useAddFakeUsersMutation()

    const addFakeUsers = async(count: number) => {
        await addFakeUsersMutation({
            variables: {
                count: count,
            },
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
