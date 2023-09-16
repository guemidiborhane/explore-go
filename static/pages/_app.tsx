import { type ActionFunctionArgs, Link, Outlet, redirectDocument } from "react-router-dom";
import Auth from "~/components/auth";
import { fetchApi } from "~/helpers";

import './global.scss'
export async function action({ request }: ActionFunctionArgs) {
    const { signal } = request
    const [, ok] = await fetchApi<{}>('/api/auth/session', { method: 'DELETE', signal })

    if (ok) return redirectDocument('/')

    return null
}

export default function RootLayout() {
    return (
        <>
            <Link to="/links">Links</Link>
            <Auth />
            <Outlet />
        </>
    )
}
