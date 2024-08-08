$( document ).ready(function() {
    // выгрузка списка авторов в datalist
    $.ajax({
        type: "GET",
        url: "api/author/hint",
        contentType: 'application/json; charset=utf-8',
        success: function (response) {
            let data = response.hint_list
            let authorCount = data.length
            for (let i=0; i< authorCount; i++) {
                
                let authorId = data[i].id
                let author = data[i].name +" "+ data[i].lastName
                optionTag = '<option name="'+author+'" data-id="'+authorId+'">'+author+'</option>';
                $("#authors").append(optionTag)
            }
        },
        error: function (errorResponse) {
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseJSON.error
            let message = "Ошибка выгрузки списка-подсказки авторов. Статус: " + status + ". Ошибка: " + errorText
            console.log(message)
            alert(message)
        }
    });
});

// кнопка, добавляющая ещё одно поле для автора, если у книги несколько авторов
$("#addAuthorBtn").on("click", function() {
    let authorRow = $(this).parent().parent()
    authorRow.after(`<tr>
        <td></td>
        <td>
            <input type="text" list="authors" class="author">
            <button class="deleteAuthorInput" type="button" class="btn btn-primary">-</button>
        </td>
    </tr>`)

    $(".deleteAuthorInput").on("click", function() {
        $(this).parent().parent().remove()
    });
});

// добавления поля с произведением, если одна физическая кника содержит более одного произведения
// например "Час быка" и "Туманность Андромеды" - два романа Ефремова в одной физической книге
$("#addWorkBtn").on("click", function() {
    let workRow = $(this).parent().parent()
    workRow.after(`<tr>
        <td></td>
        <td>
            <input type="text" class="work">
            <button class="deleteWorkInput" type="button" class="btn btn-primary">-</button>
        </td>
    </tr>`)

    $(".deleteWorkInput").on("click", function() {
        $(this).parent().parent().remove()
    });
});

// сохранение (физической) книги
$("#saveBook").on("click", function(){

    // получение списка id авторов
    let authorList = $(".author")
    authorList.each(function(){
        let authorId = $('[name = "'+$(this).val()+'"]').data("id");
        console.log(authorId)
    });
});

// добавление автора
$("#saveAuthor").on("click", function() {
    let authorData = {
        name: $("#authorName").val(),
        fatherName: $("#authorFatherName").val(),
        lastName: $("#authorLastName").val()
    };

    $.ajax({
        type: "POST",
        url: "api/author/add",
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify(authorData),
        success: function (response) {
        
        // очистка формы
        $("#authorName").val(""),
        $("#authorFatherName").val(""),
        $("#authorLastName").val("")

        // сообщение, что автор добавлен
        let message = "Автор " + authorData.name + " " + authorData.lastName + " добавлен"
        console.log(message)
        $("#addAuthSuccessMessage").show()
        $("#addAuthSuccessMessage").html(message)
        $("#addAuthSuccessMessage").hide(20000)

        // обновление datalist
        let authorId = response.id
        let author = authorData.name +" "+ authorData.lastName
        optionTag = '<option name="'+author+'" data-id="'+authorId+'">'+author+'</option>';
        $("#authors").append(optionTag)

        },
        error: function (errorResponse) {
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseJSON.error
            let message = "Ошибка добавления в БД. Статус: " + status + ". Ошибка: " + errorText
            console.log(message)
            $("#addAuthErrorMessage").show()
            $("#addAuthErrorMessage").html(message)
            $("#addAuthErrorMessage").hide(20000)
        }
    });
});