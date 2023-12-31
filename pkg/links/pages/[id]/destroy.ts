import { ActionFunctionArgs, redirect } from "react-router-dom";
import { destroyAction } from "~/helpers";

export const protect = true
export async function action({ params }: ActionFunctionArgs) {
    return await destroyAction(`/api/links/${params.id}`, redirect('/links'))
}
