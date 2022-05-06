import React from 'react'
import { useAllUsersQuery } from '../generated/graphql'

import { UserList } from './UserList'

export const User: React.FC = () => {
    const { loading, data, refetch } = useAllUsersQuery()

    return (
        <>
            {loading &&
                <p>loading users...</p>
            }
            {!loading && data &&
                <UserList count={data.totalUsers} users={data.allUsers} refetch={refetch} />
            }
        </>
    )
}
