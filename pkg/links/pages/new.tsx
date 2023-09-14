import LinkForm from '@/links/components/Form'
import { type ActionFunctionArgs, redirect } from 'react-router-dom'
import { fetchApi, useLoaderData } from '~/helpers'
import { Link } from '../types'

export async function loader() {
    return fetchApi('/api/links/new')
}

export const protect = true

export async function action({ request }: ActionFunctionArgs) {
    const body = Object.fromEntries(await request.formData())
    const [, ok,] = await fetchApi('/api/links', {
        method: request.method,
        signal: request.signal,
        // @ts-ignore
        body
    })

    if (ok) return redirect('/links')

    return null
}

export default function NewLink() {
    const [link] = useLoaderData<Link>()

    return (
        <>
            <LinkForm link={link} />
        </>
    )
}
