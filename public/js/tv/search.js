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

function SearchInput() {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/updateSeriesSearch", true);
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
                    PreviewSeries(list[i][1]);
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

function PreviewSeries(id) {
    //set as loading

    //check to see if the resources already exist first

        //check the db for the series
        //...
        //if it exists, populate the fields
        //...
        //return

    //if not, make a request to the server
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/populateSeries", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            
            console.log(response);
            // //title
            const title = document.querySelector(".selected-series-title");
            title.textContent = response[0]
            
            // //synopsis
            const synopsis = document.querySelector(".selected-series-synopsis");
            synopsis.textContent = response[1]
            
            const release = document.querySelector(".selected-series-release");
            release.textContent = response[2]

            const runtime = document.querySelector(".selected-series-runtime");
            runtime.textContent = response[3]
            
            const seasons = document.querySelector(".selected-series-seasons");
            seasons.textContent = response[4]

            const rating = document.querySelector(".selected-series-rating");
            rating.textContent = response[5]
            
            const ratings = document.querySelector(".selected-series-ratings");
            ratings.textContent = response[6] + "/10"

            const genres = document.querySelector(".selected-series-genres");
            while (genres.firstChild) {
                genres.removeChild(genres.firstChild);
            }

            const genresList = response[7].split(",");
            for (let i = 1; i < genresList.length; i++) {
                let listItem = document.createElement("p");
                listItem.classList.add("selected-series-data");
                listItem.innerHTML = genresList[i];
                genres.appendChild(listItem);
            }

            // //image path 1
            const image0 = document.querySelector(".selected-series-image");
            if (response[8] !== null && response[8] !== undefined && response[8] !== ' ') {
                image0.src = response[8];
            }    
            else {
                // select a random bunny image
                image0.src = 'images/bunnie_1.jpg';
            }
      
            console.log("num images", response[9]);

            const review = document.querySelector(".selected-series-review");
            review.innerHTML = response[10]
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