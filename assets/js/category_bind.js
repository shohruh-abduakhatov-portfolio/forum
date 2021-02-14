

let globalCategoryList = {}


const callCategoryList = (elemId) => {
    console.log("callCategoryList");
    let res = getCategoryList()
        .then(data => {
            globalCategoryList = data;
            showCategoryList(data, elemId);
        })
        .catch()
}


const callNewPostCategoryList = (elemId, selected_cats) => {
    console.log("callCategoryList");
    let res = getCategoryList()
        .then(data => {
            globalCategoryList = data;
            showNewPostCategoryList(data, elemId, selected_cats);
        })
        .catch()
}


const categoryAutocompleteController = (e) => { }

const addCategory = () => { }

const showCategoryList = (data, elemId) => {
    // <a href="/posts-category/${data.category.id}" target="blank" class="badge badge-primary">#${data.category.name}</a>
    let container = document.getElementById(elemId)
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



const showNewPostCategoryList = (data, elemId, selected_cats) => {
    // <a class="badge badge-primary">#${data.category.name}</a>
    let container = document.querySelector(elemId)
    container.innerHTML = ""
    console.log(globalCategoryList);
    data.forEach(item => {
        let span = document.createElement('span')
        span.innerHTML = `
            <input class="custom-control-input" type="checkbox" name="category-list"
                    id="${item.id}" value="${item.id}" ${selected_cats.includes(String(item.id)) ? 'checked' : ''}>
            <label class="custom-control-label" for="${item.id}">${item.name}</label>
        `
        container.appendChild(span);
    });
}

// const newPostCategoryClicked = (id) => {
//     let elem = document.querySelector('[data-id="' + id + '"]');
//     let added = document.getElementById('selected-categories-added')
//     let removed = document.getElementById('selected-categories')
//     if (elem.getAttribute('data-added') == '0') {
//         console.log(id, elem, '0', added, removed);
//         added.appendChild(elem)
//         elem.dataset.added = 1;
//         elem.style.backgroundColor = 'lightcoral';
//     } else {
//         console.log(id, elem, '1');
//         removed.appendChild(elem)
//         elem.style.backgroundColor = '#324cdd';
//         elem.dataset.added = 0;
//     }
// }



