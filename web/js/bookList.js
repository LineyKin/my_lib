// количество строк на странице
const rowsLimit = 10

// количество кнопок пагинатора
let paginatorItemCount = 0

function buildPaginator() {
    let bookCount = getBookCount()
    let num = bookCount/rowsLimit
    paginatorItemCount = num%10 == 0? num:  Math.trunc(num) + 1
    for (let i=1; i<= paginatorItemCount; i++) {
        let newItem = '<li class="page-item"><a class="page-link" href="#">' + i + '</a></li>'
        $("#bookListPagination").append(newItem)
    }

    $("#bookListPagination .page-item").on("click", function(){
        let paginatorNumber = $(this).find("a").html()

        let sortedBy, sortType
        $("#bookListTable th").each(function(){
            if($(this).attr("isSorted") != undefined) {
                sortedBy = $(this).attr("name")
                sortType = $(this).attr("isSorted")
            }
        })

        getBookList(paginatorNumber, sortedBy, sortType)
    })
}

function getBookList(paginatorNumber, sortedBy, sortType) {

    let offset = rowsLimit * (paginatorNumber - 1)

    console.log("sortedBy " + sortedBy)
    console.log("sortType " + sortType)

    // Выгрузка списка книг по умолчанию
    $.ajax({
        type: "GET",
        data: {
            limit: rowsLimit,
            offset: offset,
            sortedBy: sortedBy,
            sortType: sortType,
        },
        url: "api/book/list",
        success: function (response) {
            buildBookTable(response.book_list)
        },
        error: function (errorResponse) {
            console.log("error")
            console.log(errorResponse)
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseText
            let message = "Ошибка выгрузки списка книг. Статус: " + status + ". Ошибка: " + errorText
            console.log(message)
            alert(message)
        }
    });
}

function buildBookTable(bookListArray) {
    $("#bookListTable .bookRow").remove()
    len = bookListArray.length
    for (let i=0; i < len; i++) {
        let newRow = buildBookRow(bookListArray[i])
        $("#bookListTable").append(newRow)
    }
}

function buildBookRow(obj) {
    return `<tr class="bookRow" data-bookId="`+obj.id+`">
                <td>`+obj.author+`</td>
                <td>`+obj.name+`</td>
                <td>`+obj.publishingHouse+`</td>
                <td>`+obj.publishingYear+`</td>
            </tr>`
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