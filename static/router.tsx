import { ActionFunctionArgs, LoaderFunctionArgs, Outlet, RouterProvider, createBrowserRouter } from "react-router-dom";
import { StrictMode, type FC } from "preact/compat";
import Protected from "@/auth/components/protected";

type LazyComponent = {
    default: FC;
    protect: boolean;
    action: (args: ActionFunctionArgs) => Response | null;
    loader: (args: LoaderFunctionArgs) => Response;
};

type PreservedRoutes = Record<string, () => Promise<LazyComponent>>;

// @ts-ignore
const ROUTES = import.meta.glob('/**/pages/**/[a-z[]*.ts(x)?');
// @ts-ignore
const PRESERVED = import.meta.glob('/**/pages/(_app|404).tsx');

const lazyLoader = (promise: () => Promise<LazyComponent>) => {
    return async () => {
        const { default: Component, action, loader, protect } = await promise();
        // @ts-ignore
        const ProtectedComponent = () => <Protected><Component /></Protected>;

        return { Component: protect ? ProtectedComponent : Component, action, loader };
    };
};

const normalize = (str: string): string => {
    return str
        .replace(/\/pkg\/|\/static\//, "")
        .replace(/pages\//, "")
        .replace(/\/index/, "")
        .replace(/\.ts(x?)$/, "")
        .replace(/\[\.{3}.+\]/, "*")
        .replace(/\[(.+)\]/, ":$1");
};

const preserved: PreservedRoutes = Object.keys(PRESERVED).reduce((preserved, file) => {
    return { ...preserved, [normalize(file)]: lazyLoader(PRESERVED[file]) };
}, {});

const getPreserved = async (name: string) => {
    const layout = preserved[name] && (await preserved[name]());
    if (layout) return layout;

    return { Component: Outlet };
};

const children = Object.keys(ROUTES).map((file) => {
    let [pkg, ...p] = normalize(file).split('/');
    const path = p.join('/');

    const lazy = lazyLoader(ROUTES[file]);

    return { path, lazy, pkg };
});

const childrenByPkg: Record<string, any[]> = children.reduce((group, child) => {
    const { pkg } = child;
    group[pkg] = group[pkg] ?? [];
    group[pkg].push(child);

    return group;
}, {} as Record<string, any>);

const routes = [
    {
        path: '/',
        async lazy() {
            return await getPreserved('_app');
        },
        children: [
            ...Object.keys(childrenByPkg).map((pkg) => {
                return {
                    path: `/${pkg}`,
                    async lazy() {
                        return await getPreserved(`${pkg}/_app`);
                    },
                    children: [
                        ...childrenByPkg[pkg],
                    ],
                };
            }),

            {
                path: '*',
                async lazy() {
                    return await getPreserved('404');
                },
            },
        ],
    },
];

const router = createBrowserRouter(routes);

export function Router() {
    return (
        <StrictMode>
            <RouterProvider router={router} />
        </StrictMode>
    );
}
