
var canvas = document.getElementById("example"),
ctx = canvas.getContext('2d');
canvas.height = 400;
canvas.width  = 600;

ctx.fillStyle = "black"
ctx.fillRect(0, 0, 600, 400);

ctx.strokeStyle = "#5088B3";
ctx.lineWidth = 1;

for (var i = 0; i < 600; i+=50) {
    ctx.beginPath();
    ctx.moveTo(i, 0);
    ctx.lineTo(i, 400);
    ctx.stroke();
}

for (var i = 0; i < 400; i+=50) {
    ctx.beginPath();
    ctx.moveTo(0, i);
    ctx.lineTo(600, i);
    ctx.stroke();
}

ctx.stroke(); 


function loadData(x,y) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var my_json = JSON.parse(this.responseText)
            console.log(my_json)

            for (var i = 0; i < 10; i++) {
                draw(my_json.x[i],my_json.y[i])
            }
       }
    };
    var params = 'x='+ x + '&y=' + y
    xhttp.open("GET", "http://127.0.0.1:9000/?"+params, true);
    xhttp.send(); 
}

window.onload = function(e){ 
    console.log("window.onload", e, Date.now() ,window.tdiff, 
    (window.tdiff[1] = Date.now()) && window.tdiff.reduce(fred) ); 
}


function draw(x,y){
    ctx.strokeStyle = "green";
    ctx.lineWidth = 1;

    ctx.beginPath();
    ctx.arc(x,y,4,0,360,false)
    ctx.fillStyle = "white";
    ctx.fill();
    ctx.stroke();
}

function handleMouseDown(event) {
    var e = window.event;

    var posX = e.clientX;
    var posY = e.clientY;

    draw(posX,posY)
    
    loadData(posX,posY)
}


canvas.addEventListener('mousedown', handleMouseDown, false)
canvas.addEventListener('mouseup', handleMouseUp, false)
