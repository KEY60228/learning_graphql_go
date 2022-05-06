import React from 'react'
import { ApolloQueryResult } from '@apollo/client'
import { AllUsersQuery, User } from '../generated/graphql'

import { UserListItem } from './UserListItem'

interface props {
    count: number
    users: User[]
    refetch: () => Promise<ApolloQueryResult<AllUsersQuery>>
}

export const UserList: React.FC<props> = props => {
    return (
        <div>
            <p>{props.count} users</p>
            <button onClick={() => props.refetch()}>Refetch Users</button>
            <ul>
                {props.users.map(user => 
                    <UserListItem key={user.githubLogin} name={user.name} avatar={user.avatar} />
                )}
            </ul>
        </div>
    )
}