$.material.init();



new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        tx: false,
        txUser: null,
        ping: null,
        serverOnline: false,
        serverAudioOn: false,
        connectionStatus: false,
        connected: false,
        hideConnectionMsg: false,
    },
    created: function () {
        var self = this;
        this.ws = new ReconnectingWebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function (e) {
            var msg = JSON.parse(e.data);
            self.connectionStatus = msg.connectionStatus;
            self.txUser = msg.txUser;
            if (latencyChart.data.datasets[0].data.length >= 20) {
                latencyChart.data.datasets[0].data.shift();
            }
            if (msg.ping > 500){
                latencyChart.data.datasets[0].data.push(500); // truncate
            } else {
                latencyChart.data.datasets[0].data.push(msg.ping);
            }
            latencyChart.update(0.1);
            self.tx = msg.tx;
            self.serverAudioOn = msg.serverAudioOn;
            self.serverOnline = msg.serverOnline;
        });
        this.ws.addEventListener('open', function () {
            self.connected = true
            setTimeout(function () {
                self.hideConnectionMsg = true;
            }, 1500)
        });
        this.ws.addEventListener('close', function () {
            self.connected = false
            self.hideConnectionMsg = false;
        });

    },
    methods: {
        openWebsocket: function () {

        },
        sendRequestServerAudioOn: function () {
            this.ws.send(
                JSON.stringify({
                    serverAudioOn: !this.serverAudioOn,
                }));
        },
        sendPtt: function () {
            // if (this.serverAudioOn) {
            this.ws.send(
                JSON.stringify({
                    ptt: !this.tx,
                }));
            // }
        },
    }
});

var volumeSlider = document.getElementById('volumeSlider');

noUiSlider.create(volumeSlider, {
    start: [50],
    connect: [true, false],
    // tooltips: [ true ],
    range: {
        'min': 0,
        'max': 100
    },
    pips: { // Show a scale with the slider
        mode: 'steps',
        stepped: true,
        density: 5
    }
});

var data = {
    labels: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
    datasets: [{
        label: "Latency",
        fill: true,
        lineTension: 0.1,
        backgroundColor: "rgba(75,192,192,0.4)",
        borderColor: "rgba(75,192,192,1)",
        borderCapStyle: 'butt',
        borderDash: [],
        borderDashOffset: 0.0,
        borderJoinStyle: 'miter',
        pointBorderColor: "rgba(75,192,192,1)",
        pointBackgroundColor: "#fff",
        pointBorderWidth: 1,
        pointHoverRadius: 5,
        pointHoverBackgroundColor: "rgba(75,192,192,1)",
        pointHoverBorderColor: "rgba(220,220,220,1)",
        pointHoverBorderWidth: 2,
        pointRadius: 1,
        pointHitRadius: 10,
        data: [65],
        spanGaps: false,
    }]
};

var ctx = document.getElementById("latencyChart");
var latencyChart = new Chart(ctx, {
    type: 'line',
    data: data,
    options: {
        legend: {
            display: false,
        },
        // animation:{
        //     duration: 2000,
        //     animation: 'easeInOutQuad',
        // },
        responsive: true,
        layout: {
            padding: {
                left: 10,
                right: 20,
                top: 20
            },
        },
        scales: {
            yAxes: [{
                ticks: {
                    max: 500,
                    min: 0,
                    stepSize: 50
                }
            }],
        }
    }
});


// socket.onopen = function (event) {
//     console.log("Socket opened successfully");
// }

// window.onbeforeunload = function (event) {
//     socket.close();
// }