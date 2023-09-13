import { Link, Outlet } from "react-router-dom";

export function Component() {
    return (
        <>
            <Link to="/links">Links</Link>
            <Outlet />
        </>
    )
}
