let render = false;
let startID = 0;
let endID = 0;

document.getElementById("calendar").onmousedown = function (event) {
    // 배경색을 초기화 한다.
    startID = 0;
    endID = 0;
    let i;
    for (i = 0; i < 42; i++) { 
        document.getElementById(i).style.backgroundColor = "#ffffff";
    }
    // 시작점의 배경을 칠한다.
    // 크로스 플렛폼을 위해서 아래처럼 이벤트 처리를 한다.
    // https://stackoverflow.com/questions/31544108/what-is-window-event-in-javascript
    e = event || window.event;
    startID = parseInt((e.target || e.srcElement).id,10);
    document.getElementById(startID).style.backgroundColor = "#ffe091";
    render = true;
}

document.getElementById("calendar").onmouseup = function (event) {
    e = event || window.event;
    endID = parseInt((e.target || e.srcElement).id, 10);
    document.getElementById(endID).style.backgroundColor = "#ffe091";
    render = false;
}

document.getElementById("calendar").onmousemove = function () {
    if (render) {
        e = event || window.event;
        currentID = parseInt((e.target || e.srcElement).id, 10);
        let i;
        let start = startID
        let end = currentID
        if (start > end) {
            start = currentID
            end = startID
        }
        // 기존에 칠해진 색상을 제거한다.
        for (i = 0; i < 42; i++) { 
            document.getElementById(i).style.backgroundColor = "#ffffff";
        }
        // 색상을 채운다.
        for (i = start; i < end+1; i++) {
            document.getElementById(i).style.backgroundColor = "#ffe091";
        }
    }
}