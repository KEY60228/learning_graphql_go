import React, { useEffect, useState } from 'react'
import { GithubAuthMutation, useGithubAuthMutation } from '../generated/graphql'

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
        alert(data.githubAuth.token)
    }

    const githubAuth = async(code: string) => {
        await githubAuthMutation({
            variables:  {
                code: code,
            },
        }).then(({data}) => data && authorizationComplete(data))
    }

    useEffect(() => {
        if (window.location.search.match(/code=/)) {
            setIsSigningIn(true)
            const code = window.location.search.replace("?code=", "")
            githubAuth(code)
        }
    }, [])

    return (
        <button onClick={requestCode} disabled={isSigningIn}>Sign In with GitHub</button>
    )
}
