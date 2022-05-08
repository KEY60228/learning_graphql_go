import React from 'react'
import { useMeQuery } from '../generated/graphql'

import { CurrentUser } from './CurrentUser'

interface props {
    isSingingIn: boolean
    requestCode: () => void
    logout: () => void
}

export const Me: React.FC<props> = props => {
    const { loading, data } = useMeQuery({fetchPolicy: "cache-and-network"})

    return (
        <>
            {data?.me?.githubLogin.length ?
                <CurrentUser {...data.me} logout={props.logout} />
                :
                loading ?
                    <p>loading...</p>
                    :
                    <button onClick={props.requestCode} disabled={props.isSingingIn}>Sign In with GitHub</button>
            }
        </>
    )
}
