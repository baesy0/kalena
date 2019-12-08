
document.getElementById("calendar").onmousedown = function (event) {
    // 크로스 플렛폼을 위해서 아래처럼 이벤트 처리를 한다.
    // https://stackoverflow.com/questions/31544108/what-is-window-event-in-javascript
    e = event || window.event;
    var elementId = (e.target || e.srcElement).id;
    console.log(elementId);
}

document.getElementById("calendar").onmouseup = function (event) {
    e = event || window.event;
    var elementId = (e.target || e.srcElement).id;
    console.log(elementId);
}