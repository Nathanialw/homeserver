'use strict'

function FetchAccountData(data) {
    console.log(data)
    fetch(data, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            // 'Authorization': 'Bearer ' + token // if you use token authentication
        },
        // body: JSON.stringify(data), // if you're sending data
    })
    .then(response => response.text())
    .then(data => {
        // Here you can use the data to update your webpage
        // For example, if you have an element with id 'history', you can do:
        var parser = new DOMParser();
        var doc = parser.parseFromString(data, 'text/html');

        // Get the element you want to update
        var historyElement = document.getElementById('account-data');

        // Remove any existing content
        while (historyElement.firstChild) {
            historyElement.removeChild(historyElement.firstChild);
        }

        // Append the new elements
        while (doc.body.firstChild) {
            historyElement.appendChild(doc.body.firstChild);
        }
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

document.getElementById("FavouritedBooks").addEventListener("click", () => {FetchAccountData("/favouritedbooks")})
document.getElementById("ReadingLists").addEventListener("click", () => {FetchAccountData("/readinglists")})
document.getElementById("ReadingHistory").addEventListener("click", () => {FetchAccountData("/readinghistory")})
document.getElementById("QueuedBooks").addEventListener("click", () => {FetchAccountData("/queuedbooks")})
document.getElementById("PostHistory").addEventListener("click", () => {FetchAccountData("/posthistory")})



