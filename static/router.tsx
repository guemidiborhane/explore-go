import { Outlet, RouterProvider, createBrowserRouter } from "react-router-dom";
import { StrictMode } from "preact/compat";

const ROUTES = import.meta.glob('/**/pages/**/[a-z[]*.ts(x)?')
const PRESERVED = import.meta.glob('/**/pages/(_app|404).tsx')

const lazyLoader = (promise: any) => {
    return async () => {
        const { default: Component, action, loader } = await promise()

        return { Component, action, loader }
    }
}

const normalize = (str: string): string => {
    return str.replace(/\/pkg\/|\/static\//, '')
        .replace(/pages\//, '')
        .replace(/\/index/, '')
        .replace(/\.ts(x?)$/, '')
        .replace(/\[\.{3}.+\]/, '*')
        .replace(/\[(.+)\]/, ':$1')

}

const preserved: { [key: string]: any } = Object.keys(PRESERVED).reduce((preserved, file) => {
    return { ...preserved, [normalize(file)]: lazyLoader(PRESERVED[file]) }
}, {})

const getPreserved = async (name: string) => {
    const layout = preserved?.[name] && await preserved?.[name]()
    if (layout) return layout

    return { Component: Outlet }
}

const children = Object.keys(ROUTES).reduce((c, file) => {
    let [pkg, ...p] = normalize(file).split('/')
    const path = (p as string[]).join('/')

    const lazy = lazyLoader(ROUTES[file])

    return [...c, { path, lazy, pkg }]
}, [])

const childrenByPkg = children.reduce((group, child) => {
    const { pkg } = child
    group[pkg] = group[pkg] ?? []
    group[pkg].push(child)

    return group
}, {})

const routes = [
    {
        path: '/',
        async lazy() {
            return await getPreserved('_app')
        },
        children: [
            ...Object.keys(childrenByPkg).map((pkg) => {
                return {
                    path: `/${pkg}`,
                    async lazy() {
                        return await getPreserved(`${pkg}/_app`)
                    },
                    children: [
                        ...childrenByPkg[pkg],
                    ]
                }
            }),

            {
                path: '*',
                async lazy() {
                    return await getPreserved('404')
                }
            }

        ]
    }
]

console.log(preserved)

const router = createBrowserRouter(routes)

export function Router() {
    return (
        <StrictMode>
            <RouterProvider router={router} />
        </StrictMode>
    )
}
