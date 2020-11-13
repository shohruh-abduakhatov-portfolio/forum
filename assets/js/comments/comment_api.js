function getComments() {
    showLoader()
    $.ajax({
        type: "GET",
        url: '/api/v1/comments',
        data: {
            "postId": getUrlId()
        },
        dataType: "json",
        traditional: true,
        success: function (data) {
            console.log("data", data)
            try {
                bindComments(data)
            } catch {
            }
            hideLoader()
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log("no data")
            hideLoader()
        }
    });
}

$(function () {
    $('form').on('submit', function (e) {

        let comment = $('#editor').summernote('code').trim();
        if (comment.length === 0) {
            return
        }
        $('#comment').val(comment)
        $("#btnSubmit").prop("disabled", true);
        e.preventDefault();
        $.ajax({
            type: 'post',
            url: '/api/v1/comment/new',
            data: $('form').serialize(),
            success: function (data) {
                console.log(data);
                onNewComment();
                bindNewComment(data);
            },
            error: function (jqXHR, textStatus, errorThrown) {
                console.log("coulnot post comment")
                $("#btnSubmit").prop("disabled", false);
            }
        });

    });

});