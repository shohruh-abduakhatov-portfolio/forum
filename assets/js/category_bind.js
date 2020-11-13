

let globalCategoryList = {}


const callCategoryList = () => {
    console.log("callCategoryList");
    let res = getCategoryList()
        .then(data => {
            globalCategoryList = data
            showCategoryList(data)
        })
        .catch()
}

const categoryAutocompleteController = (e) => { }

const addCategory = () => { }

const showCategoryList = (data) => {
    // <a href="/posts-category/${data.category.id}" target="blank" class="badge badge-primary">#${data.category.name}</a>
    let container = document.getElementById("category-list")
    container.innerHTML = ""
    console.log(globalCategoryList);
    data.forEach(item => {
        container.appendChild(_addTagLink(item))
    });
}

const _addTagLink = (data) => {
    let a = document.createElement('a')
    a.href = "/posts-category/" + data.id
    a.className = "badge badge-primary"
    a.textContent = data.name
    return a
}
