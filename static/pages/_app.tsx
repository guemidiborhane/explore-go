import { Link, Outlet } from "react-router-dom";

export default function RootLayout() {
    return (
        <>
            <Link to="/links">Links</Link>
            <Outlet />
        </>
    )
}
