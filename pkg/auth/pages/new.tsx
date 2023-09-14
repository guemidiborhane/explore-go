import { ActionFunctionArgs, Form, redirectDocument } from "react-router-dom";
import { fetchApi } from "~/helpers";
import { User } from "@/auth/types";

export async function action({ request }: ActionFunctionArgs) {
    const { method, signal } = request
    const body = Object.fromEntries(await request.formData())
    // @ts-ignore
    const [, ok] = await fetchApi<User>('/api/auth/session', { method, signal, body })

    if (ok) return redirectDocument('/')

    return null
}

export default function LoginPage() {
    return (
        <Form method="POST">
            <input type="text" name="username" id="username" />
            <input type="password" name="password" id="password" />
            <button>Login</button>
        </Form>
    )
}
