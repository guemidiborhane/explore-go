import { type ActionFunctionArgs, Link, Outlet, redirectDocument } from "react-router-dom";
import Auth from "~/components/auth";
import { fetchApi } from "~/helpers";

export async function action({ request }: ActionFunctionArgs) {
    console.log(request)
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
