import { LoaderFunction, redirect } from "react-router-dom";
import { fetchApi } from "~/helpers";
import { Link } from "./types";

export type LinksData = {
    link: Link[]
}
const linksLoader: LoaderFunction = async () => {
    const [links] = await fetchApi<Link[]>('/api/links')

    return { links }
}

export type LinkData = {
    link: Link
}
const linkLoader: LoaderFunction = async ({ params }) => {
    const [link, ok] = await fetchApi<Link>(`/api/links/${params.id}`)

    if (ok) return { link }

    return redirect('/404')
}


export { linksLoader, linkLoader }
