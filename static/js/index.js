/*function onClickHandler(x) {
    $.ajax({
        url: "http://localhost:8000/callme",
        method: "GET",
        success: function (data) {
            $("#response").html(data + x);
        },
    });
}*/

function onClickHandlerPost(x) {
    var amount = document.getElementById("l1").value;
    var id = x;
    var arr = {amount,id}
    $.ajax({
        url: "http://localhost:8000/callme",
        method: "POST",
        data:arr,
        success: function (data) {
            console.log("Se ha enviado")
        },
    });
}