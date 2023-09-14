import { linksLoader } from '@/links/loaders'
import { Card } from '@/links/components/Card'
import { Link as RouterLink } from 'react-router-dom'
import { useLoaderData } from '~/helpers'
import { Link } from '@/links/types'

export const loader = linksLoader
export const protect = true

export default function IndexLinks() {
    const [links] = useLoaderData<Link[]>()

    return (
        <>
            {links.map(link => <p><Card link={link} /></p>)}
            <RouterLink to="/links/new">Add</RouterLink>
        </>
    )
}
