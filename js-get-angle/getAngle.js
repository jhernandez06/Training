
var layout = document.getElementById('layout');
var changeLayout = document.getElementById('button-change');
var send = document.getElementById('button-send');
var badRequest = document.getElementById('bad-request');
var formulario = document.forms['layout'];
var hours = formulario['hours'];
var minutes = formulario['minutes'] ;

function toggleFlex(){
    layout.classList.toggle('grid-container');
}
function showBadRequest(){
    badRequest.classList.remove('d-none')
}
function removeBadRequest(){
    badRequest.classList.add('d-none')
}
changeLayout.addEventListener('click', function(){
    toggleFlex();
})
send.addEventListener('click', function(){
    getAngle();
})

function getAngle(){

    var hoursInt = parseInt(hours.value);
    var minutesInt = parseInt(minutes.value);
    if (isNaN(hoursInt) || isNaN(minutesInt)){
        showBadRequest();
        document.getElementById('info').innerHTML = `Add hours and minutes to get an angle.`
        document.getElementById('angle').innerHTML = `0º`;
    }
    else if(hoursInt > 12 || minutesInt > 60){
        showBadRequest();
        document.getElementById('info').innerHTML = `Add hours and minutes to get an angle.`
        document.getElementById('angle').innerHTML = `0º`;
        
    }else if(hoursInt < 1 || minutesInt < 0){
        showBadRequest();
        document.getElementById('info').innerHTML = `Add hours and minutes to get an angle.`
        document.getElementById('angle').innerHTML = `0º`;
        
    }else{
        removeBadRequest();
        document.getElementById('info').innerHTML = `At <span class = "span"> ${hours.value}Hours </span> with <span class = "span">${minutes.value}Minutes</span> there's an angle of:` 
        
        var angle = 0;
        angleHours = 30*(hoursInt);
        angleMinutes = 5.5*(minutesInt);
        
        if(hoursInt == 12){
            angle = angleMinutes;
        }else {
            angle = angleHours - angleMinutes;
            if(angle < 0){
                angle = angle * (-1);
            }
        }
        document.getElementById('angle').innerHTML = `${angle}º`;
    }
}


