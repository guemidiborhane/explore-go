import { Link } from '@/links/types';
import { Form, Link as RouterLink, useSubmit } from 'react-router-dom';
import { SubmitTarget } from 'react-router-dom/dist/dom';

import styles from './index.module.scss'

export const Card = ({ link }: { link: Link; }) => {
    const submit = useSubmit()
    const handleDelete = (e: Event) => {
        e.preventDefault()
        e.stopPropagation()

        if (confirm('Are you sure?')) {
            submit(e.currentTarget as SubmitTarget)
        }
    }
    return (
        <div className={styles.Card}>
            <h3>
                <RouterLink to={`/links/${link.id}`}>
                    {link.link}
                </RouterLink>
            </h3>
            <RouterLink key={link.id} to={`/links/${link.id}/edit`}>Edit</RouterLink>
            <Form method="DELETE" action={`/links/${link.id}/destroy`} onSubmit={handleDelete}>
                <button>Delete</button>
            </Form>
        </div>
    );
};
Card.displayName = "Card";
