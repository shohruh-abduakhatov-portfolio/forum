
function setData(data) {
  console.log("data", data);
  if (data == {}) {
    $("#loader").css('visibility', 'hidden');
  }
  // mainContainer.innerHTML = ""
  for (i = 0; i < data.length; i++) {
    card = renderPost(data[i])
    mainContainer.appendChild(card)
  }
  hideLoader()
}


function renderPost(data) {
  function abbr(count) {
    if (count > 999_999) {
      return `${parseInt(count / 1_000_000)} M`
    } else if (count > 999) {
      return `${parseInt(count / 1_000)} K`
    } else {
      return count
    }
  }
  var _category = ``
  if (data.category != null) {
    _category += `
      <a href="/posts-category/${data.category.id}"  target="blank" class="badge badge-primary">#${data.category.name}</a>
    `
  }

  var _img = ``
  if (data.photoId != "") {
    _img += `
    <div class="row">
      <div style="width: 100%; height: 400px;" class="col justify-content-center">
      <div class="img-fluid rounded imgPreview" style="
          background-image: url('${data.photoId}');">
        </div>
      </div>
    </div>`
  }

  var _html = `
  <div class="row justify-content-center">
  <div class="col-lg-12">
    <div class="container">

      <div class="row">
        <div class="col">
          <h3 class="heading mb-0">By: <a href="/user/${data.userId}">${data.user.username}</a></h3>
        </div>
      </div>

      <div class="row">
        <div class="col">
          <a href="/post/${data.id}" target="blank">
            <h3 class="display-4 mb-0">${data.title}</h3></a>
        </div>
      </div>

      <div class="row">
        <div class="col mt-2">
          <p class="" style="text-align: justify;">
          ${data.text}</p>
        </div>
      </div>

      <div class="row mb-2">
        <div class="col">
          ${_category}
        </div>
      </div>

      ${_img}

      <div id="reaction_${data.id}" class="row pt-3 justify-content-center">
        <div class="col col-md-auto">
          <a onclick="likePost(${data.id})" style="cursor: pointer; color: #5e72e4;">
            <i class="fas fa-thumbs-up fa-1x"></i>
              <span>${abbr(data.likeCount)}</span> </a></div>
        <div class="col col-md-auto">
          <a onclick="dislikePost(${data.id})" style="cursor: pointer; color: #5e72e4;">
            <i class="fas fa-thumbs-down fa-1x"> </i>
              <span>${abbr(data.dislikeCount)}</span></a></div>
        <div class="col col-md-auto">
          <a href="/post/${data.id}">
            <i class="fas fa-comment fa-1x"> </i>
              <span>${abbr(data.commentCount)}</span></a></div>
      </div>
    </div>
  </div>
</div>
  `
  var row = document.createElement("div")
  row.className = "py-5 border-top"
  row.innerHTML = _html
  return row
}


function showLoader() {
  $("#loadMore").css('visibility', 'hidden');
  $("#loadMore").css('position', 'absolute');
  $("#loader").css('visibility', 'visible');
  $("#loader").css('position', 'relative');
}

function hideAll() {
  $("#loadMore").css('visibility', 'hidden');
  $("#loadMore").css('position', 'absolute');
  $("#loader").css('visibility', 'hidden');
  $("#loader").css('position', 'absolute');
}

function hideLoader() {
  $("#loadMore").css('visibility', 'visible');
  $("#loadMore").css('position', 'relative');
  $("#loader").css('visibility', 'hidden');
  $("#loader").css('position', 'absolute');
}


function increment(post, postId) {
  console.log(">> " + post);
  // $("i[class~=fa-thumbs-down] ~ span").text(post.dislikeCount)
  // $("i[class~=fa-thumbs-up] ~ span").text(post.likeCount)

  $(`a[onclick="dislikePost(${postId})"] span`).text(post.dislikeCount)
  $(`a[onclick="likePost(${postId})"] span`).text(post.likeCount)

}



function incrementPostDetails(post, postId) {
  console.log(">> " + post);
  $("i[class~=fa-thumbs-down] ~ span").text(post.dislikeCount)
  $("i[class~=fa-thumbs-up] ~ span").text(post.likeCount)

}

function getUrlId() {
  var res = window.location.pathname.split("/");
  var r = res.pop();
  return r
}

function formatDatetime(dt) {
  var _dt = new Date(Date.parse(dt));
  var options = { year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric', second: 'numeric' };
  return _dt.toLocaleDateString("en-US", options)
}

function onNewComment() {
  $("#comment").val("")
  $('#editor').summernote('code', '')
  $("#btnSubmit").prop("disabled", false);
  var commentsCount = $(".fas.fa-comment.fa-1x + span")
  commentsCount.text(parseInt(commentsCount.text()) + 1)

}


const getCategoryData = async () => {
  await fetchPostsByCategory()
    .then(data => setData(data))
    .catch(() => hideAll());
}