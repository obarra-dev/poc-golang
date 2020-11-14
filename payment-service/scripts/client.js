
let socket = new WebSocket("ws://127.0.0.1:5000/ws/products")
socket.addEventListener("message", function (e){console.log(e);});


socket.addEventListener('open', function (event) {
    socket.send('Hello Server!');
});
