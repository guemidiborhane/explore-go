import { Card } from "@/links/components/Card";
import { LinkData, linkLoader } from "@/links/loaders";
import { useLoaderData } from "~/helpers";

export const loader = linkLoader
export default function Show() {
    const { link } = useLoaderData<LinkData>()
    return <Card link={link} />
}
