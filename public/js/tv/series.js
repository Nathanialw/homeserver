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

function toggleReview() {
    const reviewElement = document.querySelector('.selected-series-review-container');
    const isExpanded = reviewElement.classList.contains('expanded');

    if (isExpanded) {
        // Collapse the review
        reviewElement.style.maxHeight = '7rem';
    } else {
        // Expand the review
        reviewElement.style.maxHeight = reviewElement.scrollHeight + 'px';
    }

    reviewElement.classList.toggle('expanded');
}