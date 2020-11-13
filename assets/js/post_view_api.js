function likePost(postId) {
    $.ajax({
        type: "GET",
        url: '/api/v1/like-post',
        data: {
            "postId": getUrlId()
        },
        success: function (data) {
            increment(data)
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log("no data like err")
        }
    });
}

function dislikePost() {
    $.ajax({
        type: "GET",
        url: '/api/v1/dislike-post',
        data: {
            "postId": getUrlId()
        },
        success: function (data) {
            increment(data)
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log("no data dislike err")
        }
    });
}

