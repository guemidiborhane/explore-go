import { useLoaderData } from '~/helpers'
import { Link } from '@/links/types'
import { linksLoader } from '@/links/loaders'
import { Card } from '@/links/components/Card'
import { Link as RouterLink } from 'react-router-dom'

export const loader = linksLoader

export function Component() {
    const { links } = useLoaderData<{ links: Link[] }>()

    return (
        <>
            {links.map(link => <p><Card link={link} /></p>)}
            <RouterLink to="/links/new">Add</RouterLink>
        </>
    )
}
Component.displayName = "IndexLinks"
