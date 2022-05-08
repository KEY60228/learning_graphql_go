import React, { useEffect, useState } from 'react'
import { GithubAuthMutation, useGithubAuthMutation } from '../generated/graphql'

import { Me } from './Me'

export const AuthorizedUser: React.FC = () => {
    const [isSigningIn, setIsSigningIn] = useState<boolean>(false)
    const [githubAuthMutation] = useGithubAuthMutation()

    const requestCode = () => {
        const clientID = process.env.REACT_APP_GITHUB_CLIENT_ID
        window.location.href = `https://github.com/login/oauth/authorize?client_id=${clientID}&scope=user`
    }

    const authorizationComplete = (data: GithubAuthMutation) => {
        localStorage.setItem('token', data.githubAuth.token)
        window.location.replace("/")
        setIsSigningIn(false)
    }

    const githubAuth = async(code: string) => {
        await githubAuthMutation({
            variables:  {
                code: code,
            },
        }).then(({data}) => data && authorizationComplete(data))
    }

    const logout = () => {
        localStorage.removeItem("token")
        window.location.reload()
    }

    useEffect(() => {
        if (window.location.search.match(/code=/)) {
            setIsSigningIn(true)
            const code = window.location.search.replace("?code=", "")
            githubAuth(code)
        }
    }, [])

    return (
        <Me isSingingIn={isSigningIn} requestCode={requestCode} logout={logout} />
    )
}
