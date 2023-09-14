import { ActionFunctionArgs, redirect } from "react-router-dom"
import { linkLoader } from '@/links/loaders'
import { fetchApi, useLoaderData } from '~/helpers'
import LinkForm from '@/links/components/Form'
import { Link } from "@/links/types"

export const protect = true
export const loader = linkLoader
export async function action({ request, params }: ActionFunctionArgs) {
    const body = Object.fromEntries(await request.formData())

    const [, ok,] = await fetchApi(`/api/links/${params.id}`, {
        method: request.method,
        signal: request.signal,
        // @ts-ignore
        body
    })
    if (ok) return redirect('/links')

    return null
}

export default function EditLink() {
    const [link] = useLoaderData<Link>()

    return (
        <>
            <LinkForm link={link} method="PATCH" />
        </>
    )
}
