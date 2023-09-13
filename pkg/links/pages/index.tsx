import { useLoaderData } from '~/helpers'
import { Link } from '@/links/types'
import { linksLoader } from '@/links/loaders'
import { Card } from '@/links/components/Card'
import { Link as RouterLink } from 'react-router-dom'

export const loader = linksLoader

export default function IndexLinks() {
    const { links } = useLoaderData<{ links: Link[] }>()

    return (
        <>
            {links.map(link => <p><Card link={link} /></p>)}
            <RouterLink to="/links/new">Add</RouterLink>
        </>
    )
}
