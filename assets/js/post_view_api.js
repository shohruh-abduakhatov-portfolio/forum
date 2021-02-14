function likePostDetails() {
    let postId = getUrlId();
    $.ajax({
        type: "GET",
        url: '/api/v1/like-post',
        data: {
            "postId": postId
        },
        success: function (data) {
            incrementPostDetails(data, postId)
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log("no data like err")
        }
    });
}

function dislikePostDetails() {
    let postId = getUrlId();
    $.ajax({
        type: "GET",
        url: '/api/v1/dislike-post',
        data: {
            "postId": postId
        },
        success: function (data) {
            incrementPostDetails(data, postId)
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log("no data dislike err")
        }
    });
}

