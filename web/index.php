<html>
    <head>
        <title>My first PHP page</title>
    </head>
    <body>
        <?php
            echo "Hello World!";
        ?>
    </body>
    <script src="websocket_client.js"></script>
    <script>
        const main = () => {
            // get topic from url get parameter
            let topic = new URLSearchParams(window.location.search).get('topic');
            console.log("topic: " + topic);

            if (!topic) {
                return;
            }

            const webSocketClient = new WebsocketClient(
            "ws://localhost" + "/ws/sub/" + topic,
                () => {
                    console.log("onOpen")
                },
                (messageEvent) => {
                    console.log("message received")
                    console.log(messageEvent)
                },
                (closeEvent) => {
                    console.log("disconnected")
                    console.log(closeEvent)
                },
                (errorEvent) => {
                    console.log("error")
                    console.log(errorEvent)
                }
            );
        }
        main()
    </script>
</html>