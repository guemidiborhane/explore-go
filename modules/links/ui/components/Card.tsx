import { Link } from '@links/types';
import { Form, Link as RouterLink, useSubmit } from 'react-router-dom';
import { SubmitTarget } from 'react-router-dom/dist/dom';

export const Card = ({ link }: { link: Link; }) => {
    const submit = useSubmit()
    const handleDelete = (e: Event) => {
        if (confirm('Are you sure?')) {
            submit(e.currentTarget as SubmitTarget)
        }
    }
    return (
        <>
            <h3>{link.link}</h3>
            <RouterLink key={link.ID} to={`/links/${link.ID}/edit`}>Edit</RouterLink>
            <Form method="DELETE" action={`/links/${link.ID}/destroy`} onSubmit={handleDelete}>
                <button>Delete</button>
            </Form>
        </>
    );
};
Card.displayName = "Card";

