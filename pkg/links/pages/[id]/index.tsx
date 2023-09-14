import { Card } from "@/links/components/Card";
import { linkLoader } from "@/links/loaders";
import { Link } from "@/links/types";
import { useLoaderData } from "~/helpers";

export const protect = true
export const loader = linkLoader
export default function Show() {
    const [link] = useLoaderData<Link>()
    return <Card link={link} />
}
