import { fetchApi } from "@/helpers"
import { RouteObject, redirect } from "react-router-dom"

const routes: RouteObject[] = [
    {
        path: '/links',
        lazy: () => import('./_layout'),
        children: [
            {
                path: '',
                lazy: () => import('./pages/index')
            },
            {
                path: 'new',
                lazy: () => import('./pages/create')
            },
            {
                path: ':id/edit',
                lazy: () => import('./pages/edit')
            },
            {
                path: ':id/destroy',
                async action({ params }) {
                    const url = `/api/links/${params.id}`
                    const [, ok,] = await fetchApi(url, 'DELETE')

                    if (ok) return redirect('/links')

                    return null
                }
            }
        ]
    }
]

export { routes }
