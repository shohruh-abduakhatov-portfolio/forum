
function getData() {
  showLoader()
  console.log(">> getData");
  offset += limit;
  $.ajax({
    type: "GET",
    url: '/api/v1/posts',
    data: {
      "limit": limit,
      "offset": offset
    },
    dataType: "json",
    traditional: true,
    success: function (data) {
      console.log("data", data)
      setData(data)
    },
    error: function (jqXHR, textStatus, errorThrown) {
      console.log("no data")
      hideAll()
    }
  });
}


const fetchPostsByCategory = async () => {
  showLoader()
  console.log(">> fetchPostsByCategory");
  offset += limit;
  return new Promise(
    (res, rej) => {
      $.ajax({
        type: "GET",
        url: '/api/v1/posts-by-category',
        data: {
          "limit": limit,
          "offset": offset,
          "category": parseInt(document.getElementById("category-id").dataset['id'])
        },
        dataType: "json",
        traditional: true,
        success: (data) => res(data),
        error: (jqXHR, textStatus, errorThrown) => rej(errorThrown)
      });
    }
  )
}


function likePost(postId) {
  $.ajax({
    type: "GET",
    url: '/api/v1/like-post',
    data: {
      "postId": postId
    },
    success: function (data) {
      increment(postId, "like")
    },
    error: function (jqXHR, textStatus, errorThrown) {
      console.log("no data like err")
    }
  });
}


function dislikePost(postId) {
  $.ajax({
    type: "GET",
    url: '/api/v1/dislike-post',
    data: {
      "postId": postId
    },
    success: function (data) {
      increment(postId, "dislike")
    },
    error: function (jqXHR, textStatus, errorThrown) {
      console.log("no data dislike err")
    }
  });
}


const getCategoryList = async () =>
  new Promise(
    (res, rej) => {
      console.log(">> getCategoryList");
      $.ajax({
        type: "GET",
        url: '/api/v1/categories',
        success: (data) => {
          res(data)
        },
        error: (jqXHR, textStatus, errorThrown) => {
          console.log("no data like err")
          rej(errorThrown)
        }
      });
    }
  )


