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

export { primary100, primary200, primary300, primary400, primary500 };