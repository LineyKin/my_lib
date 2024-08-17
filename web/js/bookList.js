// количество строк на странице
const rowsLimit = 5

// количество кнопок пагинатора
let paginatorItemCount = 0

function buildPaginator() {
    let bookCount = getBookCount()
    paginatorItemCount = Math.trunc(bookCount/rowsLimit) + 1
    for (let i=1; i<= paginatorItemCount; i++) {
        let newItem = '<li class="page-item"><a class="page-link" href="#">' + i + '</a></li>'
        $("#bookListPagination").append(newItem)
    }
}

function getDefaultBookList() {
    let listParams = {
        limit: rowsLimit,
        offset: 0
    }
    // Выгрузка списка книг по умолчанию
    $.ajax({
        type: "GET",
        data: JSON.stringify(listParams),
        url: "api/book/list",
        contentType: 'application/json; charset=utf-8',
        success: function (response) {
            console.log(response)
            console.log(paginatorItemCount)

        },
        error: function (errorResponse) {
            console.log(errorResponse)
            //let status = errorResponse.status + " " + errorResponse.statusText
            //let errorText = errorResponse.responseJSON.error
            //let message = "Ошибка выгрузки списка книг. Статус: " + status + ". Ошибка: " + errorText
            //console.log(message)
            //alert(message)
        }
    });
}

function getBookCount() {
    let count = 0;
    $.ajax({
        type: "GET",
        async: false,
        url: "api/book/count",
        contentType: 'application/json; charset=utf-8',
        success: function (response) {
            count = response.count

        },
        error: function (errorResponse) {
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseJSON.error
            let message = "Ошибка выгрузки количества книг. Статус: " + status + ". Ошибка: " + errorText
            console.log(message)
            alert(message)
        }
    });

    return count
}