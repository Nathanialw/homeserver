'use strict';

let primary100;
let primary200;
let primary300;
let primary400;
let primary500;

document.addEventListener("DOMContentLoaded", function() {
    // Get the root element
    const root = document.documentElement;

    // Get the computed styles of the root element
    const styles = getComputedStyle(root);

    // Access the CSS custom properties
    primary100 = styles.getPropertyValue('--primary-100').trim();
    primary200 = styles.getPropertyValue('--primary-200').trim();
    primary300 = styles.getPropertyValue('--primary-300').trim();
    primary400 = styles.getPropertyValue('--primary-400').trim();
    primary500 = styles.getPropertyValue('--primary-500').trim();
});

function highlightEpisode(seasonNum, episodeNum) {
    const allEpisodes = document.getElementsByClassName("side-card");
    for (let i = 0; i < allEpisodes.length; i++) {
        allEpisodes[i].classList.remove("highlighted");
    }   

    const episodeElement = document.getElementById(seasonNum + "-" + episodeNum);
    episodeElement.classList.add("highlighted");
}

function highlightSeries(listItem) {
    const allSeries = document.getElementsByClassName("list-group-item");
    for (let i = 0; i < allSeries.length; i++) {
        allSeries[i].classList.remove("highlighted");
    }   

    listItem.classList.add("highlighted")
    console.log("highlighted");
}

function playEpisode(seriesID, seasonNum, episodeNum) {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/selectEpisode", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
        highlightEpisode(seasonNum, episodeNum);
        const response = JSON.parse(xhr.responseText);
        console.log(response);
        const videoElement = document.getElementById("video-player");
        videoElement.src = response.videoURL;
        videoElement.load();
        videoElement.play();          
        }
    };
    xhr.send("seriesID=" + seriesID + "&seasonNum=" + seasonNum + "&episodeNum=" + episodeNum);
}


function SearchInput() {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/updateSeriesSearch", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            //create a list of times to display the recieved arry
            let list = JSON.parse(xhr.responseText);
            if (list.length > 0) {
                const start = document.getElementById("start-title");
                start.textContent = "Select a series...";
            }
            console.log("list", list);

            let listElement = document.getElementById("search-results");
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
    //check to see if the resources already exist first

    //...

    //if not, make a request to the server
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/populateSeries", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
        const response = JSON.parse(xhr.responseText);
        
        console.log(response);
        // //title
        const title = document.getElementById("selected-series-title");
        title.textContent = response[0]
        
        // //synopsis
        const synopsis = document.getElementById("selected-series-synopsis");
        synopsis.textContent = response[1]
        // //image path 1
        
        if (response[2] !== null || response[2] !== undefined || response[2] !== "") {
            const image = document.getElementById("selected-series-image");
            image.src = response[2];
        }
        
        // //image path 2
        // response[3]
        }
    };
    xhr.send("id=" + id);
}

function SelectSeries(seriesID) {
    const selected = document.getElementById("selected-series");
    selected.classList.add("show-selected");

    const start = document.getElementById("show-selected-start");
    start.classList.remove("show-selected-start");


    let code = document.getElementById("imdbCode");
    code.value = seriesID;
}