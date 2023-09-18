import { render } from "preact";
import { Router } from "~/router";
import { type ReceiverEvent, subscribe } from "~/websocket";


subscribe({
    channel: 'system',
    receiver: ({ data }: ReceiverEvent) => {
        if (data.message == 'reload') window.location.reload()
    }
})
subscribe({
    channel: 'links',
    receiver: ({ data }: ReceiverEvent) => {
        console.log(data)
    }
})

render(
    <Router />
    , document.getElementById('app')!)
