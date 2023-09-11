import { Outlet } from "react-router-dom"

export function Component() {
    return (
        <>
            <h1>Links Layout</h1>
            <Outlet />
        </>
    )
}

Component.displayName = 'LinksLayout'
