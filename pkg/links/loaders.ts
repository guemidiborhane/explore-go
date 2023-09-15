import { LoaderFunction } from "react-router-dom";
import { fetchApi } from "~/helpers";
import { Link } from "./types";

export type LinksData = {
    link: Link[]
}
const linksLoader: LoaderFunction = async ({ request }) => {
    return await fetchApi<Link[]>('/api/links', { signal: request.signal })
}

export type LinkData = {
    link: Link
}
const linkLoader: LoaderFunction = ({ params, request }) => {
    return fetchApi<Link>(`/api/links/${params.id}`, { signal: request.signal })
}


export { linksLoader, linkLoader }
