import { createBrowserRouter } from "react-router-dom";
import Root from "./layouts/root"
import { routes as linksRoutes } from "@links/index"

export const router = createBrowserRouter([
    {
        path: '/',
        element: <Root />,
        children: [linksRoutes].flat()
    }
])
