class WebsocketClient {
	constructor(url = "ws://localhost:8080", onOpen, onMessage, onClose, onError) {
		this.socket = new WebSocket(url);
		this.socket.onopen = onOpen();
        this.socket.onmessage = onMessage;
        this.socket.onclose = onClose;
        this.socket.onerror = onError;
    }
}