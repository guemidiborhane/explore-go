import { render } from 'preact'
import { StrictMode } from 'preact/compat'
import { RouterProvider } from 'react-router-dom'
import { router } from './router.tsx'

render(
    <StrictMode>
        <RouterProvider router={router} />
    </StrictMode>
    , document.getElementById('app')!)
