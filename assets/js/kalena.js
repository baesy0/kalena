let render = false;
let startID = 0;
let endID = 0;

document.getElementById("calendar").onmousedown = function (event) {
    // 배경색을 초기화 한다.
    startID = 0;
    endID = 0;
    for (let i = 0; i < 42; i++) { 
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
    // 마우스를 떼면 addSchedule modal을 띄운다.
    $("#addSchedule").modal();
}

document.getElementById("calendar").onmousemove = function () {
    if (render) {
        e = event || window.event;
        let cur = parseInt((e.target || e.srcElement).id, 10);
        let start = startID
        let end = cur
        if (start > end) {
            start = cur
            end = startID
        }
        // 기존에 칠해진 색상을 제거한다.
        for (let i = 0; i < 42; i++) { 
            document.getElementById(i).style.backgroundColor = "#ffffff";
        }
        // 색상을 채운다.
        for (let i = start; i < end+1; i++) {
            document.getElementById(i).style.backgroundColor = "#ffe091";
        }
    }
}


//올해, 이번달 페이지일 경우, 오늘 날짜에 동그라미를 그린다.
window.onload = function highlightToday(){
    let c = document.getElementById("calendar")
    let qYear = c.dataset.queryyear;
    let qMonth = c.dataset.querymonth;
    let offset = c.dataset.offset;
    let today = new Date();
    
    if(qYear == today.getFullYear() && qMonth == today.getMonth()){
        document.getElementById(offset + today.getDate()-1).style.backgroundColor = "red";
        console.log(document.getElementById(offset + today.getDate()-1).outerHTML);
    }
}
