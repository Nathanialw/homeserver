import sys
from bs4 import BeautifulSoup
import requests
import json
import os

headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.83 Safari/537.36"
}

tv_ratings = [
    "TV-Y",
    "TV-Y7",
    "TV-Y7-FV",
    "TV-G",
    "TV-PG",
    "TV-14",
    "TV-MA",
    "G",
    "PG",
    "PG-13",
    "R",
    "NC-17",
    "NR",
    "UR",
    "Unrated",
    "Not Rated",
    "12",
    "12A",
    "15",
    "15A",
    "18",
    "18A",
    "14",
    "M"
]

def scrape_series_preview(key):
    savePath = "../../public/images/tv/" + key + "/"
    referencePath = "../../images/tv/" + key + "/"
    url = "https://www.imdb.com/title/" + key
    html_page = requests.get(url, headers=headers)
    soup = BeautifulSoup(html_page.content, 'html.parser')
    results = []

    #Title
    title = ""
    try:
        parent = soup.find('section', class_='ipc-page-section')
        title = parent.find('span', class_='hero__primary-text').text
        value = title
        result = [value]
        results.append(result)
    except:
        value = "not available"
        result = [value]
        results.append(result)

    # Synopsis
    try:
        parent = soup.find('section', class_='sc-9a2a0028-2')
        value = parent.find('span', class_='sc-3ac15c8d-2').text
        result = [value]
        results.append(result)
    except:
        value = "not available"
        result = [value]
        results.append(result)

    # release date
    try:
        parent = soup.find('li', {'data-testid': 'title-details-releasedate'})
        value = parent.find('a', class_='ipc-metadata-list-item__list-content-item').text
        result = [value]
        results.append(result)
    except:
        value = "not available"
        result = [value]
        results.append(result)

    # runtime
    try:
        parent = soup.find('section', {'data-testid': 'TechSpecs'})
        value = parent.find('div', class_='ipc-metadata-list-item__content-container').text
        result = [value]
        results.append(result)
    except:
        value = "not available"
        result = [value]
        results.append(result)

    # seasons
    try:
        parent = soup.find('div', {'data-testid': 'episodes-browse-episodes'})
        value = parent.find('label', class_='ipc-simple-select__label').text
        result = [value]
        results.append(result)
    except:
        value = " "
        result = [value]
        results.append(result)

    # rating
    try:
        parent = soup.find('section', {'data-testid': 'hero-parent'})
        parent = parent.find('ul', class_='sc-ec65ba05-2')
        list = parent.find_all('a')
        value = ["not available"]
        results.append(value)
        for i in list:
            for j in tv_ratings:
                if i.text == j:
                    value = [i.text]
                    results.pop()
                    results.append(value)
                    break
    except:
        value = "not available"
        result = [value]
        results.append(result)

    # ratings score
    try:
        parent = soup.find('div', {'data-testid': 'hero-rating-bar__aggregate-rating__score'})
        value = parent.find('span').text
        result = [value]
        results.append(result)
    except:
        value = "not available"
        result = [value]
        results.append(result)

    # genres
    try:

        parent = soup.find('div', class_='ipc-chip-list__scroller')
        list = parent.find_all('span')
        value = []
        for i in list:
            value.append(i.text)
        results.append(value)
    except:
        value = "not available"
        result = [value]
        results.append(result)

    # card image
    try:
        parent = soup.find('div', class_='ipc-media')
        image_url = parent.find('img',).get('src')
        response = requests.get(image_url, headers=headers)
        if not os.path.exists(savePath):
            os.makedirs(savePath)
        with open(savePath + title + '.png', 'wb') as file:
            file.write(response.content)
        result = [referencePath + title + '.png']
        results.append(result)
    except:
        image = " "
        result = [image]
        results.append(result)

    # extra images
    try:        
        parent = soup.find('section', {'data-testid': 'Photos'})
        list = parent.find_all('img', )
        if not os.path.exists(savePath):
            os.makedirs(savePath)

        num = 0
        for image_url in list:
            response = requests.get(image_url.get('src'), headers=headers)
            append = "_" + str(num)
            with open(savePath + title  + append +  '.png', 'wb') as file:
                file.write(response.content)
            num += 1
        results.append([str(num)])
    except:
        image = " "
        result = [image]
        results.append(result)

    # review
    try:    
        parent = soup.find('section', {'data-testid': 'UserReviews'})
        value = parent.find('div', class_='ipc-html-content-inner-div').text
        result = [value]
        results.append(result)
    except:
        image = " "
        result = [image]
        results.append(result)

    return results

if __name__ == "__main__":
    # if len(sys.argv) != 2:
    #     print("Usage: python scrapeSeriesPreview.py <key>")
    #     sys.exit(1)

    key = sys.argv[1]
    # key = "tt22248376" #frieren
    # key = "tt0944947" #GoT
    results = scrape_series_preview(key)

    # Print the results as JSON
    print(json.dumps(results))
