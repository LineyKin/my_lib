$("#addAuthorBtn").on("click", function() {
    let authorRow = $(this).parent().parent()
    authorRow.after(`<tr>
        <td></td>
        <td>
            <input type="text" class="author">
            <button class="deleteAuthorInput" type="button" class="btn btn-primary">-</button>
        </td>
    </tr>`)

    $(".deleteAuthorInput").on("click", function() {
        $(this).parent().parent().remove()
    });
});

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