import { Link } from '@links/types'
import { Form } from 'react-router-dom'

export default function LinkForm({ link, method = "POST" }: { link: Link, method?: string }) {
    return (
        <Form method={method}>
            <input type="text" name="link" value={link.link} />
            <button>Save</button>
        </Form>
    )
}
