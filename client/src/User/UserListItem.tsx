import React from 'react'

interface props {
    name: string
    avatar: string
}

export const UserListItem: React.FC<props> = props => {
    return (
        <li>
            <img src={props.avatar} width={48} height={48} alt={props.name} />
            { props.name }
        </li>
    )
}
