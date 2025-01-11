import sys
from bs4 import BeautifulSoup
import requests
import json

headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.83 Safari/537.36"
}

savePath = "../../public/images/tv/"
referencePath = "../../images/tv/"

url = "https://www.imdb.com/title/" + sys.argv[1]
html_page = requests.get(url, headers=headers)
soup = BeautifulSoup(html_page.content, 'html.parser')
results = []

#Title
specific_movie_data_1 = soup.find('section', class_='ipc-page-section')
try:
    title = specific_movie_data_1.find('span', class_='hero__primary-text').text
    results.append(title)
except:
    title = None

#Synopsis
try:
    synopsis = specific_movie_data_1.find('span', class_='sc-3ac15c8d-1').text
    results.append(synopsis)
except:
    synopsis = None

# card image
specific_movie_data_2 = soup.find('div', class_='ipc-media')
try:
    image_url  = specific_movie_data_2.find('img',).get('src')
    response = requests.get(image_url, headers=headers)
    with open(savePath + title + '.png', 'wb') as file:
        file.write(response.content)
    results.append(referencePath + title + '.png')
except:
    image = None

# background image
specific_movie_data_2 = soup.find('div', class_='ipc-media')
try:
    image_url = specific_movie_data_2.find('img', ).get('src')
    response = requests.get(image_url, headers=headers)
    with open(savePath + title + '.png', 'wb') as file:
        file.write(response.content)
    results.append(referencePath +title + '.png')
except:
    image = None

print(json.dumps(results))

