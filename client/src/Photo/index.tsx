import React from 'react'
import { useAllPhotosQuery } from '../generated/graphql'

export const Photo: React.FC = () => {
    const {loading, data} = useAllPhotosQuery({fetchPolicy: "cache-and-network"})

    return (
        <>
            {loading &&
                <p>loading...</p>
            }
            {!loading && data &&
                data.allPhotos.map(photo => <img key={photo.id} src={photo.url} alt={photo.name} width={350} />)
            }
        </>
    )
}
