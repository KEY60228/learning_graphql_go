import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { PostPhotoInput, PhotoCategory, usePostPhotoMutation, AllPhotosDocument, AllPhotosQuery, AllPhotosQueryVariables } from '../generated/graphql'

export const PostPhoto: React.FC = () => {
    const navigate = useNavigate()
    const [ postPhotoMutation ] = usePostPhotoMutation()

    const [photo, setPhoto] = useState<PostPhotoInput>({
        name: "",
        description: "",
        category: PhotoCategory.Portrait,
        file: null,
        taggedUserIDs: ["KEY60228"], // TODO
    })

    const postPhoto = async() => {
        await postPhotoMutation({
            variables: {
                input: photo,
            },
            update: (cache, {data}) => {
                const option = { query: AllPhotosDocument }
                const { allPhotos, totalPhotos } = cache.readQuery<AllPhotosQuery, AllPhotosQueryVariables>(option) || {}
                if (!allPhotos || !totalPhotos) {
                    return
                }
                cache.writeQuery<AllPhotosQuery, AllPhotosQueryVariables>({
                    query: AllPhotosDocument,
                    data: {
                        totalPhotos: totalPhotos + 1,
                        allPhotos: {
                            ...allPhotos,
                            ...data?.postPhoto,
                        }
                    }
                })
            }
        })
        navigate("/")
    }

    return (
        <form
            onSubmit={e => e.preventDefault()}
            style={{
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'flex-start',
                alignItems: 'flex-start',
            }}
        >
            <h1>Post a Photo</h1>
            <input
                type="text"
                style={{
                    margin: '10px',
                }}
                placeholder="photo name..."
                value={photo.name}
                onChange={e => setPhoto(prev => ({...prev, name: e.target.value}))}
            />
            <textarea
                style={{
                    margin: "10px",
                }}
                placeholder="photo description..."
                defaultValue={photo.description ?? ""}
                onChange={e => setPhoto(prev => ({...prev, description: e.target.value}))}
            ></textarea>
            <select
                style={{
                    margin: "10px",
                }}
                onChange={e => {
                    const category = e.target.value as unknown as PhotoCategory // TODO: どうするんがいいんか
                    setPhoto(prev => ({...prev, category}))
                }}
            >
                <option value="PORTRAIT">PORTRAIT</option>
                <option value="LANDSCAPE">LANDSCAPE</option>
                <option value="ACTION">ACTION</option>
                <option value="GRAPHIC">GRAPHIC</option>
            </select>
            <input
                type="file"
                style={{
                    margin: "10px",
                }}
                accept="image/jpeg"
                onChange={e => setPhoto(prev => ({...prev, file: e.target.files?.length ? e.target.files[0] : ''}))}
            />
            <div style={{ margin: "10px" }}>
                <button onClick={postPhoto}>Post Photo</button>
                <button onClick={() => navigate("/")}>Cancel</button>
            </div>
        </form>
    )
}
