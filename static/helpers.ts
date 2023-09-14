import { useLoaderData as loaderHook } from "react-router-dom";

type APIResponse<T> = [T, boolean, number, AbortController]

type Params = RequestInit & {
    body?: BodyInit | { [k: string]: FormDataEntryValue } | null
}

export async function fetchApi<T>(url: string, params: Params = { method: 'GET' }): Promise<APIResponse<T>> {

    const body = params.body
    /* @ts-ignore */
    if (params.method !== 'GET' && body && body.id && typeof body.id === 'string') {
        /* @ts-ignore */
        body.id = parseInt(body?.id);
    }

    const controller = new AbortController();
    if (!params.signal) {
        params.signal = controller.signal
    }

    return fetch(url, {
        ...params,
        body: JSON.stringify(body),
        credentials: 'include',
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
    }).then(async (response) => [await response.json(), response.ok, response.status, controller]);
}

export function useLoaderData<T>(): APIResponse<T> {
    return loaderHook() as APIResponse<T>
}

export async function destroyAction(url: string, response: Response): Promise<Response | null> {
    const [, ok,] = await fetchApi(url, { method: 'DELETE' })

    if (ok) return response

    return null
}
