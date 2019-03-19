function onClickHandler(x) {
    $(document).ready(function () {
        $("#"+x).on('click', function () {
            $.ajax({
                url: "http://localhost:8000/callme",
                method: "GET",
                success: function (data) {
                    $("#response").html(data+x);
                },
            });
        });
    });
}