'use strict'

function IsSet(value) {
    if (value == "NA") {
        value = ""
    }
    return sanitizeHTML(value)
}

function IsSetRawText(value) {
    if (value == "" || value == "NA") {
        value = ""
    }
    return sanitizeHTML(value)
}

function sanitizeHTML(text) {
    let div = document.createElement('div');
    div.innerText = text;
    return div.innerHTML;
}

async function sanitizeImageUrl(url) {
    return await isValidImageUrl(url) ? url : '';
}

async function isValidImageUrl(url) {
    try {
        const response = await fetch(url, { method: 'HEAD' });
        const contentType = response.headers.get('Content-Type');
        if (!contentType.startsWith('image/')) {
            return false;
        }
    } catch (error) {
        return false;
    }
    return url.endsWith('.png') || url.endsWith('.jpg') || url.endsWith('.jpeg') || url.endsWith('.gif');
}

async function SetHTML (selectedBook) {
    document.getElementById("title").value = IsSet(selectedBook.Title)
    document.getElementById("subtitle").value = IsSet(selectedBook.Subtitle)
    document.getElementById("author").value = IsSet(selectedBook.Author)
    document.getElementById("publish_year").value = IsSet(selectedBook.Publish_year)
    document.getElementById("image-preview").src = await sanitizeImageUrl(selectedBook.Image)
    document.getElementById("image").value = ""
    
    let synopsis = document.getElementById("synopsis")
    synopsis.value = IsSetRawText(selectedBook.Synopsis);
    document.getElementById("synopsis").innerText = IsSetRawText(selectedBook.Synopsis)
    SetTextAreaHieght();

    document.getElementById("link_amazon").value = IsSet(selectedBook.Link_amazon)
    document.getElementById("link_indigo").value = IsSet(selectedBook.Link_indigo)
    document.getElementById("link_pdf").value = IsSet(selectedBook.Link_pdf)
    document.getElementById("link_epub").value = IsSet(selectedBook.Link_epub)
    document.getElementById("link_handmade").value = IsSet(selectedBook.Link_handmade)
    document.getElementById("link_text").value = IsSet(selectedBook.Link_text)
    
    document.getElementById("image").addEventListener("change", function (event) {
        var selectedFile = event.target.files[0];
        var imageUrl = URL.createObjectURL(selectedFile);
        document.getElementById("image-preview").src = imageUrl;
    });
    
    document.getElementById("synopsis").addEventListener("input", SetTextAreaHieght);
}

function SetTextAreaHieght () {
    let numChars = synopsis.value.length;
    let charsPerLine = 70; // adjust this value based on the width of the textarea and the size of the font
    let numLines = Math.ceil(numChars / charsPerLine);
    if (synopsis.value.includes('\n')) {
        numLines += synopsis.value.split('\n').length - 1;
    }
    numLines < 4 ? numLines = 4 : numLines
    synopsis.setAttribute("rows", numLines)
}

//replace the html in the select options with the sorted array
function displayBooks(bookList) {
    //save the current value of the select
    var selectedOption = document.getElementById(bookList).value;
    
    var bookList = document.getElementById(bookList);
    while (bookList.firstChild) {
        bookList.removeChild(bookList.firstChild);
    }
    bookList.innerHTML = "";
    for (var i = 0; i < books.length; i++) {
        var option = document.createElement("option");
        option.value = books[i].Title;
        let text = books[i].Title
        if (books[i].Subtitle != "" && books[i].Subtitle != "NA") {
            text += ": " + books[i].Subtitle
        }
        text += ", by: " + books[i].Author
        option.text = text;
        bookList.add(option);
    }
    //after sorting, set the value of the select to the saved value
    if (selectedOption == "") {
        selectedOption = books[0].Title
        // selectedOption = "Select a book..."
    }
    document.getElementById(bookList.id).value = selectedOption;
    let selectedBook = books.find(book => book.Title === selectedOption)
    SetHTML(selectedBook)
}

displayBooks("bookList")

//get the books from the server
document.getElementById("bookList").addEventListener("change", function () {
    var selectedOption = document.getElementById(bookList.id).value;
    let selectedBook = books.find(book => book.Title === selectedOption)
    SetHTML(selectedBook)
})






