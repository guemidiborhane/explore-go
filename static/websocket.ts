export interface Message {
    channel: string;
    message: string
}

export type ReceiverEvent = { isTrusted: boolean, data: Message }
type Receiver = (e: ReceiverEvent) => void
type Subscription = { receiver: Receiver, channel: string }

export default class Consumer {
    protected connection?: WebSocket = undefined
    private subscriptions: Subscription[] = []
    subscribe: (s: Subscription) => void
    ping_timeout: number = 5000
    private ping_interval: any
    private reconnect_interval: any
    private ping: string = JSON.stringify({ channel: 'health', message: 'ping' })

    constructor() {
        this.subscribe = this.addSubscription.bind(this)
    }

    private connect() {
        if (this.connection && this.connection.readyState === WebSocket.OPEN) {
            return; // Already connected
        }

        if (this.reconnect_interval) clearInterval(this.reconnect_interval)
        const url = "ws://" + window.location.host + "/ws";
        const connection = new WebSocket(url)
        connection.onmessage = this.onMessage.bind(this)
        connection.onopen = this.onOpen.bind(this)
        connection.onclose = this.onClose.bind(this)
        connection.onerror = () => connection.close()

        this.connection = connection
    }

    private onMessage(event: MessageEvent) {
        const data = JSON.parse(event.data)
        this.subscriptions
            .filter(s => s.channel == data.channel)
            .forEach(s => s.receiver({ ...event, data }))
    }

    private addSubscription(subscription: Subscription) {
        this.connection?.close()
        this.subscriptions.push(subscription)
        this.connect()
    }

    private onClose() {
        clearInterval(this.ping_interval)
        const connect = this.connect.bind(this)
        this.reconnect_interval = setInterval(() => {
            connect()
        }, this.ping_timeout)
    }

    private onOpen() {
        // to Keep the connection alive
        this.ping_interval = setInterval(() => {
            this.connection?.send(this.ping);
        }, this.ping_timeout)
    }
}

export const subscribe = (new Consumer).subscribe
