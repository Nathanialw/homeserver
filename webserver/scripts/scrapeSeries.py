import sys
from bs4 import BeautifulSoup
import requests

headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.83 Safari/537.36"
}

path = "../../public/images/"

url = "https://www.imdb.com/title/" + sys.argv[1]
print(url)
html_page = requests.get(url, headers=headers)
print(html_page)
soup = BeautifulSoup(html_page.content, 'html.parser')

#Title
specific_movie_data_1 = soup.find('section', class_='ipc-page-section')
try:
    title = specific_movie_data_1.find('span', class_='hero__primary-text').text
    print(title)
except:
    title = None
    print(title)

#Synopsis
try:
    synopsis = specific_movie_data_1.find('span', class_='sc-3ac15c8d-1').text
    print(synopsis)
except:
    synopsis = None
    print(synopsis)

# card image
specific_movie_data_2 = soup.find('div', class_='ipc-media')
try:
    image_url  = specific_movie_data_2.find('img',).get('src')
    print(image_url)
    response = requests.get(image_url, headers=headers)
    with open(path + title + '.png', 'wb') as file:
        file.write(response.content)
    print("card Image downloaded successfully.")
except:
    image = None
    print(image)

# background image
specific_movie_data_2 = soup.find('div', class_='ipc-media')
try:
    image_url = specific_movie_data_2.find('img', ).get('src')
    print(image_url)
    response = requests.get(image_url, headers=headers)
    with open(path + title + '.png', 'wb') as file:
        file.write(response.content)
    print("background Image downloaded successfully.")
except:
    image = None
    print(image)

