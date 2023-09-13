import { useLoaderData as loaderHook } from "react-router-dom";

type Body = {
    ID?: string | number
    [key: string]: any
}

export async function fetchApi<T>(url: string, method: string = 'GET', body: Body | undefined = undefined): Promise<[T, boolean, number]> {
    if (method !== 'GET' && body && body.ID && typeof body.ID === 'string') {
        body.ID = parseInt(body.ID);
    }
    return fetch(url, {
        method,
        body: JSON.stringify(body),
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
    }).then(async (response) => [await response.json(), response.ok, response.status]);
}

export function useLoaderData<T>(): T {
    return loaderHook() as T
}

export async function destroyAction(url: string, response: Response): Promise<Response | null> {
    const [, ok,] = await fetchApi(url, 'DELETE')

    if (ok) return response

    return null
}
