
document.getElementById("calendar").onmousedown = function (e) {
    e = e || window.event;
    var elementId = (e.target || e.srcElement).id;
    console.log(elementId);
}

document.getElementById("calendar").onmouseup = function (e) {
    e = e || window.event;
    var elementId = (e.target || e.srcElement).id;
    console.log(elementId);
}