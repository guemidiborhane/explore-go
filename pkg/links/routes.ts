import { destroyAction } from "~/helpers"
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
                lazy: () => import('./pages/new')
            },
            {
                path: ':id/edit',
                lazy: () => import('./pages/edit')
            },
            {
                path: ':id/destroy',
                async action({ params }) {
                    return await destroyAction(`/api/links/${params.id}`, redirect('/links'))
                }
            }
        ]
    }
]

export { routes }
