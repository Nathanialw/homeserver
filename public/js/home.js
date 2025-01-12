'use strict'

function updateSearchEngine() {
    const searchEngine = document.getElementById("search-engine").value;
    const form = document.getElementById("search-form");
    const queryInput = document.getElementById("search-query");

    switch (searchEngine) {
        case "brave":
            form.action = "https://search.brave.com/search";
            queryInput.name = "q";
            break;
        case "google":
            form.action = "https://www.google.com/search";
            queryInput.name = "q";
            break;
        case "yandex":
            form.action = "https://yandex.com/search/";
            queryInput.name = "text";
            break;
        default:
            form.action = "https://search.brave.com/search";
            queryInput.name = "q";
    }
}



document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("search-query").focus();
});
