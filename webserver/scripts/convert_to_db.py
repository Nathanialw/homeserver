import subprocess
import os
import gzip
import shutil
import sqlite3
import sqlite_spellfix
from tqdm import tqdm
import sys

# Download text file from IMDb
subprocess.run([
    'wget', '-r', '-l1', '-np', '-nd', '-P', '',
    "https://datasets.imdbws.com/title.basics.tsv.gz"
])

# Unzip the file
with gzip.open('title.basics.tsv.gz', 'rb') as f_in:
    with open('title.basics.tsv', 'wb') as f_out:
        shutil.copyfileobj(f_in, f_out)

# Connect to SQLite database (or create it)
conn = sqlite3.connect('../db/imdb.sqlite3')
conn.enable_load_extension(True)
conn.load_extension(sqlite_spellfix.extension_path())
cursor = conn.cursor()

# Create table
cursor.execute('''
CREATE VIRTUAL TABLE IF NOT EXISTS title_akas USING fts5(
    tconst, titleType, primaryTitle, originalTitle, isAdult, startYear, endYear, runtimeMinutes, genres
)
''')

# Copy data to the database

with open('title.basics.tsv', 'r', encoding='utf-8') as file:
    next(file)  # Skip the header row
    # Use tqdm to show progress
    for line in tqdm(file, desc="Inserting data into title_akas"):
        fields = line.strip().split('\t')
        if fields[1] != 'tvSeries':
            continue
        cursor.execute('''
        INSERT INTO title_akas (tconst, titleType, primaryTitle, originalTitle, isAdult, startYear, endYear, runtimeMinutes, genres)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
        ''', fields)

# Create spellfix table
cursor.execute("CREATE VIRTUAL TABLE IF NOT EXISTS spellfix_table USING spellfix1")

# Populate the spellfix1 table with individual words from primaryTitle
cursor.execute("SELECT primaryTitle FROM title_akas")
titles = cursor.fetchall()

# Use tqdm to show progress
for title in tqdm(titles, desc="Populating spellfix_table"):
    cursor.execute("INSERT INTO spellfix_table(word) VALUES (?)", title)
    words = title[0].split()
    # Skip titles with only one word
    if words == 1:
        continue
    for word in words:
        cursor.execute("INSERT INTO spellfix_table(word) VALUES (?)", (word,))

# Commit and close the database connection
conn.commit()
conn.close()

# Delete the text file
os.remove('title.basics.tsv')
os.remove('title.basics.tsv.gz')