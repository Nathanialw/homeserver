'use strict';

document.addEventListener("DOMContentLoaded", function() {
    const form = document.querySelector(".form");
    form.addEventListener("keydown", function(event) {
        if (event.key === "Enter") {
            event.preventDefault();
        }
    });
});


function highlightSeries(listItem) {
    const allSeries = document.querySelectorAll(".list-group-item");
    for (let i = 0; i < allSeries.length; i++) {
        allSeries[i].classList.remove("highlighted");
    }   

    listItem.classList.add("highlighted")
    console.log("highlighted");
}

let selectedID = "";

function SearchInput(route, previewRoute, Populate) {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", route, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            //create a list of times to display the recieved arry
            let list = JSON.parse(xhr.responseText);
            if (list.length > 0) {
                const searchResults = document.querySelector(".search-container");
                searchResults.classList.add("show-selected");
            }
            console.log("list", list);

            let listElement = document.querySelector(".search-results");
            listElement.innerHTML = "";
            for (let i = 0; i < list.length; i++) {
                let listItem = document.createElement("li");
                listItem.classList.add("list-group-item");
                listItem.innerHTML = list[i][0];
                listElement.appendChild(listItem);

                listItem.addEventListener("click", function() {
                    console.log("clicked");
                    highlightSeries(listItem);
                    SelectSeries(list[i][1]);
                    PreviewMedia(list[i][1], previewRoute, Populate);
                    selectedID = list[i][1];
                })
            }            
        };
    };
    let query = document.getElementById("title").value;
    if (query === "" || query === null || query === undefined || query === " ") {
        return;
    }
    xhr.send("query=" + query);
}

function PreviewMedia(id, route, Populate) {
    //if not, make a request to the server
    const xhr = new XMLHttpRequest();
    xhr.open("POST", route, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            console.log("response", response);
            Populate(response)
        }
    };

    if (id === selectedID) {
        return;
    }
    xhr.send("id=" + id);
}

function SelectSeries(seriesID) {
    const selected = document.querySelector(".selected-container");
    selected.classList.add("show-selected");
    
    let code = document.getElementById("imdbCode");
    code.value = seriesID;
}


function AddPreviewElement(response, id, i) {
    if (response[i] === null || response[i] === undefined || response[i] === ' ') {
        return i + 1;
    }
    const title = document.querySelector(id);
    title.textContent = response[i]
    return i + 1;
}

function AddPreviewElements(response, i)  {
    const genres = document.querySelector(".selected-series-genres");
    while (genres.firstChild) {
        genres.removeChild(genres.firstChild);
    }

    const genresList = response[i].split(",");
    for (let j = 1; j < genresList.length; j++) {
        let listItem = document.createElement("p");
        listItem.classList.add("selected-series-data");
        listItem.innerHTML = genresList[j];
        genres.appendChild(listItem);
    }
    i++

    // //image path 1
    const image0 = document.querySelector(".selected-series-image");
    if (response[i] !== null && response[i] !== undefined && response[i] !== ' ') {
        image0.src = response[i];
    }    
    else {
        // select a random bunny image
        image0.src = 'images/bunnie_1.jpg';
    }
    i++

    console.log("num images", response[i]);
    i++

    const review = document.querySelector(".selected-series-review");
    review.innerHTML = response[i]
}

function PreviewSeries(response) {           
    let i = 0; 
    i = AddPreviewElement(response, ".selected-series-title", i)
    i = AddPreviewElement(response, ".selected-series-synopsis", i)
    i = AddPreviewElement(response, ".selected-series-release", i)
    i = AddPreviewElement(response, ".selected-series-runtime", i)
    i = AddPreviewElement(response, ".selected-series-seasons", i)
    i = AddPreviewElement(response, ".selected-series-rating", i)
    i = AddPreviewElement(response, ".selected-series-ratings", i)
    AddPreviewElements(response, i)
}

function PreviewMovie(response) {            
    let i = 0; 
    i = AddPreviewElement(response, ".selected-series-title", i)
    i = AddPreviewElement(response, ".selected-series-synopsis", i)
    i = AddPreviewElement(response, ".selected-series-release", i)
    i = AddPreviewElement(response, ".selected-series-runtime", i)
    i++
    i = AddPreviewElement(response, ".selected-series-rating", i)
    i = AddPreviewElement(response, ".selected-series-ratings", i)
    AddPreviewElements(response, i)
}