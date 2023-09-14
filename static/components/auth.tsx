import { checkUser } from "@/auth/components/protected";
import { useEffect, useState } from "preact/hooks";
import { Form, Link, useSubmit } from "react-router-dom";
import { type SubmitTarget } from "react-router-dom/dist/dom";


export default function Auth() {
    const submit = useSubmit()
    const handleSignout = (e: Event) => {
        e.preventDefault()
        e.stopPropagation()
        if (confirm('Are you sure?')) {
            submit(e.currentTarget as SubmitTarget)
        }
    }
    const [signedIn, setSignedIn] = useState(false)

    useEffect(() => {
        (async () => {
            setSignedIn(await checkUser())
        })()
    }, [])

    if (!signedIn) return <Link to="/auth/new">Sign In</Link>


    return (
        <Form action="/" method="DELETE" onSubmit={handleSignout}>
            <button>Sign Out</button>
        </Form>
    )
}
