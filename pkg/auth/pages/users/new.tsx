import { redirectDocument, type ActionFunctionArgs, Form } from "react-router-dom";
import { fetchApi } from "~/helpers";

export async function action({ request }: ActionFunctionArgs) {
    const { method, signal } = request
    const body = Object.fromEntries(await request.formData())
    // @ts-ignore
    const [, ok] = await fetchApi<User>('/api/auth', { method, signal, body })

    if (ok) {
        // @ts-ignore
        const [, ok] = await fetchApi<User>('/api/auth/session', { method, signal, body })

        if (ok) return redirectDocument('/')
    }

    return null
}

export default function RegisterPage() {
    return (
        <Form method="POST">
            <input type="text" name="name" id="name" />
            <input type="text" name="username" id="username" />
            <input type="password" name="password" id="password" />
            <button>Sign up</button>
        </Form>
    )
}
