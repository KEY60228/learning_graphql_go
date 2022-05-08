import { ApolloClient, ApolloConsumer } from '@apollo/client'
import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useGithubAuthMutation, MeDocument, MeQuery, MeQueryVariables } from '../generated/graphql'

import { Me } from './Me'

export const AuthorizedUser: React.FC = () => {
    const [isSigningIn, setIsSigningIn] = useState<boolean>(false)
    const [githubAuthMutation] = useGithubAuthMutation()
    const navigate = useNavigate()

    const requestCode = () => {
        const clientID = process.env.REACT_APP_GITHUB_CLIENT_ID
        window.location.href = `https://github.com/login/oauth/authorize?client_id=${clientID}&scope=user`
    }

    const githubAuth = async(code: string) => {
        await githubAuthMutation({
            variables:  { code },
            update(cache, {data}) {
                if (data) {
                    localStorage.setItem('token', data.githubAuth.token)
                    navigate("/", {replace: true})
                    setIsSigningIn(false)
                }
            },
            refetchQueries: [MeDocument],
        })
    }

    const logout = (client: ApolloClient<object>) => {
        localStorage.removeItem("token")
        client.writeQuery<MeQuery, MeQueryVariables>({ query: MeDocument, data: {me: null} })
    }

    useEffect(() => {
        const url = new URL(window.location.href)
        const params = url.searchParams
        const code = params.get("code")
        if (code) {
            setIsSigningIn(true)
            githubAuth(code)
        }
    }, [])

    return (
        <ApolloConsumer>
            {client => <Me isSingingIn={isSigningIn} requestCode={requestCode} logout={() => logout(client)} />}
        </ApolloConsumer>
    )
}
