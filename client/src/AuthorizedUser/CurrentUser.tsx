import React from 'react'

interface props {
    name: string
    avatar: string,
    logout: () => void,
}

export const CurrentUser: React.FC<props> = props => {
    return (
        <div>
            <img src={props.avatar} width={48} height={48} alt={props.name} />
            <h1>{props.name}</h1>
            <button onClick={props.logout}>logout</button>
        </div>
    )
}
