function comprar() {
    var total = $('#total_final').html()
    if (total == "0") {
        alert("Error! Tiene el carrito vacio")
    } else {
        $.ajax({
            url: "http://localhost:8000/comprar",
            method: "GET",
            success: function (data) {
                alert("Se ha comprado con exito")
                for (var i = 0; i < 5; i++) {
                    $('#t' + i).html(0)
                }
                $("#total_final").html(0)
            }
        })
    }
}

function calcTotalGet(x) {
    $.ajax({
        url: "http://localhost:8000/calcularTotal",
        method: "GET",
        dataType: 'jsonp',
        complete: function (data) {
            var obj = JSON.parse(data.responseText)
            console.log(obj)
            $.each(obj, function (key, value) {
                if (value instanceof Array) {
                    for (var i = 0; i < value.length; i++) {
                        console.log(value[i])
                        $('#t' + i).html(value[i])
                    }
                } else {
                    $("#total_final").html(value)
                }
            })
        },
    });
}

function onClickHandlerPost(x) {
    var amount = document.getElementById("l1").value;
    var id = x;
    var arr = { amount, id }
    $.ajax({
        url: "http://localhost:8000/callme",
        method: "POST",
        data: arr,
        success: function (data) {
            console.log("Se ha enviado")
        },
    });
}
function onClickHandlerPostLib2(x) {
    var amount = document.getElementById("l2").value;
    var id = x;
    var arr = { amount, id }
    $.ajax({
        url: "http://localhost:8000/callme",
        method: "POST",
        data: arr,
        success: function (data) {
            console.log("Se ha enviado")
        },
    });
}
function onClickHandlerPostLib3(x) {
    var amount = document.getElementById("l3").value;
    var id = x;
    var arr = { amount, id }
    $.ajax({
        url: "http://localhost:8000/callme",
        method: "POST",
        data: arr,
        success: function (data) {
            console.log("Se ha enviado")
        },
    });
}
function onClickHandlerPostLib4(x) {
    var amount = document.getElementById("l4").value;
    var id = x;
    var arr = { amount, id }
    $.ajax({
        url: "http://localhost:8000/callme",
        method: "POST",
        data: arr,
        success: function (data) {
            console.log("Se ha enviado")
        },
    });
}
function onClickHandlerPostLib5(x) {
    var amount = document.getElementById("l5").value;
    var id = x;
    var arr = { amount, id }
    $.ajax({
        url: "http://localhost:8000/callme",
        method: "POST",
        data: arr,
        success: function (data) {
            console.log("Se ha enviado")
        },
    });
}

function addLibOne(x) {
    var amount = 1
    var id = x;
    var arr = { amount, id }
    $.ajax({
        url: "http://localhost:8000/addLibs",
        method: "POST",
        data: arr,
        success: function (data) {
            console.log("Se ha enviado")
        },
    });
}
