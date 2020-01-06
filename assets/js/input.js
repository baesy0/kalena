function addScheduleModal(){
    let startdate = document.getElementById("startdate").value;
    let enddate = document.getElementById("enddate").value;
    let title = document.getElementById("title").value;
    let layer = document.getElementById("layer").value;

    $.ajax({
        url:"/api/schedule",
        type: "post",
        data:{
            collection: "bae",
            start: startdate,
            end: enddate,
            title: title,
            layer: layer,
        },
        dataType: "json",
        success: function(data){
            //추후 해당 스케쥴 렌더링
            //지금은 알림창 정도
            console.log("success");
            alert("success!");
        },
        error: function(request,status,error){
            alert("code:"+request.status+"\n"+"status:"+status+"\n"+"Msg:"+request.responseText+"\n"+"error:"+error);
        }
    });
}