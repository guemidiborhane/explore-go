import { ActionFunctionArgs, redirect } from "react-router-dom"
import { linkLoader, type LinkData } from '@/links/loaders'
import { fetchApi, useLoaderData } from '~/helpers'
import LinkForm from '@/links/components/Form'

export const loader = linkLoader
export async function action({ request, params }: ActionFunctionArgs) {
    const body = Object.fromEntries(await request.formData())
    const [, ok,] = await fetchApi(`/api/links/${params.id}`, request.method, body)

    if (ok) return redirect('/links')

    return null
}

export default function EditLink() {
    const { link } = useLoaderData<LinkData>()

    return (
        <>
            <LinkForm link={link} method="PATCH" />
        </>
    )
}
