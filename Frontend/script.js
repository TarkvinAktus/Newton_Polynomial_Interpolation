
var canvas = document.getElementById("example");
ctx = canvas.getContext('2d');
canvas.height = 700;
canvas.width  = 600;

let array_x = [];
let array_y = [];

var posX = 0;
var posY = 0;

let next_x = 0;
let prev_x = 0;
let line = 0;

ctx.fillStyle = "black"
ctx.fillRect(0, 0, 600, 700);

ctx.strokeStyle = "#5088B3";
ctx.lineWidth = 1;

for (var i = 0; i < 600; i+=50) {
    ctx.beginPath();
    ctx.moveTo(i, 0);
    ctx.lineTo(i, 700);
    ctx.stroke();
}

for (var i = 0; i < 700; i+=50) {
    ctx.beginPath();
    ctx.moveTo(0, i);
    ctx.lineTo(600, i);
    ctx.stroke();
}

ctx.stroke(); 


reload.onclick = erase;




function erase_fast(){
    ctx.fillStyle = "rgba(0,0,0,"+ 1+")"
    ctx.fillRect(0, 0, 600, 700);

    ctx.strokeStyle = "rgba(80,136,179,"+ 1+")"
    ctx.lineWidth = 1;

    for (var i = 0; i < 600; i+=50) {
        ctx.beginPath();
        ctx.moveTo(i, 0);
        ctx.lineTo(i, 700);
        ctx.stroke();
    }

    for (var i = 0; i < 700; i+=50) {
        ctx.beginPath();
        ctx.moveTo(0, i);
        ctx.lineTo(600, i);
        ctx.stroke();
    }
}

function erase(){
    j = 0
    
    function my_loop(){
        setTimeout(function(){
            ctx.fillStyle = "rgba(0,0,0,"+ 0.05*(j+1)+")"
            ctx.fillRect(0, 0, 600, 700);

            ctx.strokeStyle = "rgba(80,136,179,"+ 0.05*(j+1)+")"
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
            j++
            if (j < 20){
                my_loop();
            }
            }, 50);
    }
        
    my_loop();  
    
}

function loadData(x,y) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var my_json = JSON.parse(this.responseText)
            erase_fast()
            
            console.log(my_json.pol)
            
            var my_y = 0;
            for (var i = 0; i < 600; i++) {
                my_y = 0;
                var pol_count = my_json.pol.length
                for(var j = 0;j<pol_count;j++){
                    my_y += my_json.pol[j] * Math.pow(i,j)
                }
                draw(i,my_y)
        
            }
        /*
           for (var i = 0; i < 600; i++){
                draw(i,my_json.pol[i])
           }*/

       }
    };

    var params = 'x='+ array_x + '&y=' + array_y
    xhttp.open("GET", "http://127.0.0.1:9000/?"+params, true);
    xhttp.send(); 
}
/*
window.onload = function(e){ 
    console.log("window.onload", e, Date.now() ,window.tdiff, 
    (window.tdiff[1] = Date.now()) && window.tdiff.reduce(fred) ); 
}*/

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
    posX = e.clientX;
    posY = e.clientY;
    console.log(posX + "-"+posY)

    draw(posX,posY);
    array_x.push(posX);
    array_y.push(posY);
    
    loadData(posX,posY)
}

//canvas.addEventListener('mouseup', handleMouseUp, false)
canvas.addEventListener('mousedown', handleMouseDown, false)
//canvas.addEventListener('mouseup', handleMouseUp, false)
