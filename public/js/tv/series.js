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

