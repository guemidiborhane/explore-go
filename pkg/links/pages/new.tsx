import LinkForm from '@/links/components/Form'
import { type ActionFunctionArgs, redirect } from 'react-router-dom'
import { LinkData } from '@/links/loaders'
import { fetchApi, useLoaderData } from '~/helpers'

export async function loader() {
    const [link] = await fetchApi('/api/links/new')

    return { link }
}

export async function action({ request }: ActionFunctionArgs) {
    const body = Object.fromEntries(await request.formData())
    const [, ok,] = await fetchApi('/api/links', request.method, body)

    if (ok) return redirect('/links')

    return null
}

export default function NewLink() {
    const { link } = useLoaderData<LinkData>()

    return (
        <>
            <LinkForm link={link} />
        </>
    )
}
