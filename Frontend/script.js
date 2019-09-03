
var canvas = document.getElementById("canvas");
ctx = canvas.getContext('2d');

document.getElementById("reload").onclick = fluent_erase;

//Some global vars
//Sorry about that (._.)
canvas.height = 700;
canvas.width  = 600;

let array_x = [];
let array_y = [];

let posX = 0;
let posY = 0;

function first_draw(){
    ctx.fillStyle = "black"
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    ctx.strokeStyle = "#5088B3";
    ctx.lineWidth = 1;

    for (var i = 0; i < canvas.width; i+=50) {
        ctx.beginPath();
        ctx.moveTo(i, 0);
        ctx.lineTo(i, canvas.height);
        ctx.stroke();
    }

    for (var i = 0; i < canvas.height; i+=50) {
        ctx.beginPath();
        ctx.moveTo(0, i);
        ctx.lineTo(canvas.width, i);
        ctx.stroke();
    }

    ctx.stroke(); 
}

function erase_fast(){
    ctx.fillStyle = "rgb(0,0,0)"
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    ctx.strokeStyle = "rgb(80,136,179)"
    ctx.lineWidth = 1;

    for (var i = 0; i < canvas.width; i+=50) {
        ctx.beginPath();
        ctx.moveTo(i, 0);
        ctx.lineTo(i, canvas.height);
        ctx.stroke();
    }

    for (var i = 0; i < canvas.height; i+=50) {
        ctx.beginPath();
        ctx.moveTo(0, i);
        ctx.lineTo(canvas.width, i);
        ctx.stroke();
    }
}

function fluent_erase(){
    //Flush x and y data
    array_x = [];
    array_y = [];

    //Some magic recursion cycle 
    //To provide fluency 
    j = 0
    function my_loop(){
        setTimeout(function(){
            ctx.fillStyle = "rgba(0,0,0,"+ 0.05*(j+1)+")"
            ctx.fillRect(0, 0, canvas.width, canvas.height);

            ctx.strokeStyle = "rgba(80,136,179,"+ 0.05*(j+1)+")"
            ctx.lineWidth = 1;

            for (var i = 0; i < canvas.width; i+=50) {
                ctx.beginPath();
                ctx.moveTo(i, 0);
                ctx.lineTo(i, canvas.height);
                ctx.stroke();
            }

            for (var i = 0; i < canvas.height; i+=50) {
                ctx.beginPath();
                ctx.moveTo(0, i);
                ctx.lineTo(canvas.width, i);
                ctx.stroke();
            }
            j++
            if (j < 20){
                my_loop();
            }
            }, 50);
    }
    my_loop();  
}

function loadData() {
    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var my_json = JSON.parse(this.responseText);
            erase_fast();
            
            console.log(my_json.pol);
            //Draw function by response set of polynomial coefficents
            var my_y = 0;
            for (var i = 0; i < canvas.width; i++) {
                my_y = 0;
                var pol_count = my_json.pol.length;
                for(var j = 0;j<pol_count;j++){
                    my_y += my_json.pol[j] * Math.pow(i,j);
                }
                draw(i,my_y);
            }
       }
    }

    var params = 'x='+ array_x + '&y=' + array_y;
    xhttp.open("GET", "http://127.0.0.1:9000/?"+params, true);
    xhttp.send(); 
}

function draw(x,y){
    ctx.strokeStyle = "green";
    ctx.lineWidth = 1;

    ctx.beginPath();
    ctx.arc(x,y,4,0,360,false);
    ctx.fillStyle = "white";
    ctx.fill();
    ctx.stroke();
}

function drawPointAndSendData() {
    var e = window.event;

    //var canvas_pos = canvas.getBoundingClientRect();
    //console.log(rect.top, rect.right, rect.bottom, rect.left);
    var canvas_col = document.getElementById("canvas-col")
    //console.log(canvas_col.offsetLeft + canvas_col.offsetTop)

    posX = e.clientX-canvas_col.offsetLeft;
    posY = e.clientY-canvas_col.offsetTop;

    console.log(posX + "-"+posY)
    console.log(e.clientX + "-"+e.clientY)

    draw(posX,posY);
    array_x.push(posX);
    array_y.push(posY);
    
    loadData(posX,posY);
}

window.onload = function(){ 
    console.log("window.onload");
    first_draw(); 
}
canvas.addEventListener('mouseup', drawPointAndSendData, false)

