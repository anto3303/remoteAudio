new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        ptt: false,
        txUser: null,
        ping: null,
        serverAudioOn: false,
        connectionStatus: false,
    },
    created: function () {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function (e) {
            var msg = JSON.parse(e.data);
            self.connectionStatus = msg.connectionStatus
            self.txUser = msg.txUser
            self.ping = msg.ping
            self.ptt = msg.ptt
            self.serverAudioOn = msg.serverAudioOn
        });
        this.ws.addEventListener('open', function () {
            console.log("websocket opened")
        })
    },
    methods: {
        sendRequestServerAudioOn: function () {
            this.ws.send(
                JSON.stringify({
                    ptt: !this.serverAudioState,
                }));
        },
        sendPtt: function () {
            this.ws.send(
                JSON.stringify({
                    ptt: !this.ptt,
                }));
        },
    }
})

// socket.onopen = function (event) {
//     console.log("Socket opened successfully");
// }

// window.onbeforeunload = function (event) {
//     socket.close();
// }