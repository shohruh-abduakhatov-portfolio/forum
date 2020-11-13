function bindComments(data) {
    if (data == null) {
        hideLoader()
        return
    }
    var commentsContainer = document.getElementById('commentsContainer');
    commentsContainer.innerHTML = ""
    for (i = 0; i < data.length; i++) {
        card = renderComments(data[i])
        commentsContainer.appendChild(card)
    }
    hideLoader()
}

function bindNewComment(data) {
    if (data == null) {
        return
    }
    card = renderComments(data);
    document.getElementById('commentsContainer').appendChild(card);
}

function renderComments(data) {
    _html = `
        <div class="row pb-2">
            <div class="col">
                <h3 class="heading mb-0">By:
                <a href="/user/${data.user.id}" target="blank">
                    @${data.user.username}</a></h3>
            </div>
        </div>

        <div class="row ml-5">
            <div class="col ">
                <mark class="text-muted">${formatDatetime(data.created_at)}</mark>
            </div>
        </div>

        <div class="row ml-5">
            <div class="col mt-2">
                <blockquote class="blockquote text-left">
                <p class="mb-0">
                    ${data.comment}
                </p>
                </blockquote>
            </div>
        </div>
    `
    var row = document.createElement("div")
    row.className = "mb-5"
    row.innerHTML = _html
    return row
}