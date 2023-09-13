import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { routes as linksRoutes } from '@/links'
import { StrictMode } from "preact/compat";

const router = createBrowserRouter([
    {
        path: '/',
        lazy: () => import('./layouts'),
        children: [linksRoutes].flat()
    }
])

export function Router() {
    return (
        <StrictMode>
            <RouterProvider router={router} />
        </StrictMode>
    )
}
