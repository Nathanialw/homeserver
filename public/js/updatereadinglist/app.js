'use strict'

function  IsSet(value) {
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


async function SetHTML(selectedReadingList) {
    document.getElementById("name").value = IsSet(selectedReadingList.Name)
    document.getElementById("image-preview").src = await sanitizeImageUrl(selectedReadingList.Chart)
    document.getElementById("category").value = IsSet(selectedReadingList.Category)
    
    let description = document.getElementById("description")
    description.value = IsSet(selectedReadingList.Description)
    SetTextAreaHieght()

    document.getElementById("image").addEventListener("change", function (event) {
        var selectedFile = event.target.files[0];
        var imageUrl = URL.createObjectURL(selectedFile);
        document.getElementById("image-preview").src = imageUrl;
    });
    
    document.getElementById("description").addEventListener("input", SetTextAreaHieght);
}

function SetTextAreaHieght(selectedReadingList) {
    let numChars = description.value.length;
    let charsPerLine = 70; // adjust this value based on the width of the textarea and the size of the font
    let numLines = Math.ceil(numChars / charsPerLine);
    if (description.value.includes('\n')) {
        numLines += description.value.split('\n').length - 1;
    }
    numLines < 4 ? numLines = 4 : numLines
    description.setAttribute("rows", numLines)
}

//replace the html in the select options with the sorted array
function displayReadingList(readingList) {
    //save the current value of the select
    var selectedOption = document.getElementById(readingList).value;
    
    var readingList = document.getElementById(readingList);
    while (readingList.firstChild) {
        readingList.removeChild(readingList.firstChild);
    }
    readingList.innerHTML = "";
    for (var i = 0; i < readinglists.length; i++) {
        var option = document.createElement("option");
        option.value = readinglists[i].Name;
        let text = readinglists[i].Name
        // if (readinglists[i].Subtitle != "" && readinglists[i].Subtitle != "NA") {
        //     text += ": " + readinglists[i].Subtitle
        // }
        text += ", " + readinglists[i].Category
        option.text = text;
        readingList.add(option);
    }
    //after sorting, set the value of the select to the saved value
    if (selectedOption == "") {
        selectedOption = readinglists[0].Name
        // selectedOption = "Select a book..."
    }
    document.getElementById(readingList.id).value = selectedOption;
    let selectedReadingList = readinglists.find(readinglist => readinglist.Name === selectedOption)
    SetHTML(selectedReadingList)
}

displayReadingList("readingLists")

//get the books from the server
document.getElementById("readingLists").addEventListener("change", function () {
    var selectedOption = document.getElementById("readingLists").value;
    let selectedReadingList = readinglists.find(readinglist => readinglist.Name === selectedOption)
    SetHTML(selectedReadingList)
})




