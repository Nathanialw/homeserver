'use strict'

class jj {
    constructor (x,y) {
        this.x = x;
        this.y = y;
    }
}

function sortBooksTitle() {
    books.sort(function (a, b) {
        var titleA = a.Title.toUpperCase(); // ignore upper and lowercase
        var titleB = b.Title.toUpperCase(); // ignore upper and lowercase
        if (titleA < titleB) {
            return -1;
        }
        if (titleA > titleB) {
            return 1;
        }
        // names must be equal
        return 0;
    });
}

function sortBooksAuthor() {
    books.sort(function (a, b) {
        var authorA = a.Author.toUpperCase(); // ignore upper and lowercase
        var authorB = b.Author.toUpperCase(); // ignore upper and lowercase
        if (authorA < authorB) {
            return -1;
        }
        if (authorA > authorB) {
            return 1;
        }
        // names must be equal
        return 0;
    });
}

function sortBooksSubtitle() {
    books.sort(function (a, b) {
        var subtitleA = a.Subtitle.toUpperCase(); // ignore upper and lowercase
        var subtitleB = b.Subtitle.toUpperCase(); // ignore upper and lowercase
        if (subtitleA < subtitleB) {
            return -1;
        }
        if (subtitleA > subtitleB) {
            return 1;
        }
        // names must be equal
        return 0;
    });
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
    }
    document.getElementById(bookList.id).value = selectedOption;
}

displayBooks("book1")

function SortBooks(book) {
    var selectedOption = event.target.value;

    if (selectedOption == "byAuthor") {
        sortBooksAuthor()
    }
    else if (selectedOption == "bySubtitle") {
        sortBooksSubtitle()
    }
    else {
        sortBooksTitle()
    }

    //the display values should be based on the select that was changed so we can type to find the book
    displayBooks(book)
}

// Add an event listener for the change event
document.getElementById('sort1').addEventListener('change', (event) => {
    SortBooks("book1")
});

//add an event listener for the add button that adds another li to the ul
let numBooks = 1
function AddBook() {
    numBooks++
    const book = document.createElement("li");
    book.style.marginTop = "0.2rem";
    book.id = "book"
    const span = document.createElement("span");
    span.classList.add("form-adjacent");
    book.appendChild(span);

    const selectSort = document.createElement("select");
    selectSort.classList.add("form-input");
    selectSort.classList.add("form-row");
    selectSort.classList.add("form-adjacent");
    selectSort.style.textTransform = "capitalize";
    selectSort.style.width = "125px";
    selectSort.style.fontSize = "0.75rem";
    selectSort.id = "sort" + numBooks;
    selectSort.name = "sort" + numBooks;
    selectSort.value = "byTitle";
    selectSort.required = true;
    selectSort.innerHTML = '<option value="byTitle">Sort By Title</option><option value="byAuthor">Sort By Author</option><option value="bySubtitle">Sort By Subtitle</option>';
    span.appendChild(selectSort);  

    
    const select = document.createElement("select");
    select.classList.add("form-input");
    select.classList.add("form-row");
    select.classList.add("form-adjacent");
    select.style.textTransform = "capitalize";
    select.style.fontSize = "0.75rem";
    select.id = "book" + numBooks;
    select.name = "book" + numBooks;
    select.required = true;
    span.appendChild(select);

    //add button to remove book
    const removeButton = document.createElement("button");
    removeButton.classList.add("btn");
    removeButton.innerHTML = "Remove";
    removeButton.id = "remove" + numBooks;
    removeButton.name = "remove" + numBooks;
    removeButton.style.marginLeft = "10px";
    removeButton.style.textTransform = "capitalize";    
    removeButton.style.fontSize = "0.75rem";    
    span.appendChild(removeButton);   
    
    document.getElementById("readingList").appendChild(book);
    SortBooks("book" + numBooks)
    
    const currentNumBooks = numBooks
    document.getElementById("sort" + currentNumBooks).addEventListener('change', (event) => {
        SortBooks("book" + currentNumBooks)
    });
}


document.getElementById('addBook').addEventListener('click', (event) => { 
    AddBook()
});

//when i click the remove button, remove the book from the ul associated with the button includibg the li
document.getElementById('readingList').addEventListener('click', (event) => {
    if (event.target.tagName === 'BUTTON') {
        event.target.parentElement.parentElement.remove()
    }
});







