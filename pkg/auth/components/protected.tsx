import { type ReactNode, useEffect, useState } from "preact/compat";
import { Navigate } from "react-router-dom";
import { fetchApi } from "~/helpers";
import { User } from "@/auth/types";

export async function checkUser(): Promise<boolean> {
    const [, ok] = await fetchApi<User>('/api/auth')

    return ok
}

export default function Protected({ children }: { children: ReactNode }) {
    const [signedIn, setSignedIn] = useState<boolean | undefined>(undefined)

    useEffect(() => {
        (async () => {
            setSignedIn(await checkUser())
        })()
    }, [])

    if (signedIn == undefined) return
    if (signedIn == false) return <Navigate to="/auth/new" replace />

    return children
}
