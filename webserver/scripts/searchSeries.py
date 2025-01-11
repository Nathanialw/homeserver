import sqlite3
import sqlite_spellfix
from itertools import product, permutations, combinations
from tqdm import tqdm
import sys
import json
import Levenshtein

db = sqlite3.connect('../db/imdb.sqlite3')
db.enable_load_extension(True)
db.load_extension(sqlite_spellfix.extension_path())
cursor = db.cursor()

def search(query):
    # Correcting each query term with spellfix table
    correctedquery = []
    for t in query.split():
        spellfix_query = "SELECT word FROM spellfix_table WHERE word MATCH ? LIMIT 5"
        cursor.execute(spellfix_query, (t,))
        results = cursor.fetchall()
        if results:
            unique_results = list(set([r[0] for r in results]))
            correctedquery.append(unique_results)
        else:
            correctedquery.append([t])  # if no match, keep the word spelled as it is

    # Generate combinations of corrected query terms
    all_combinations = []
    for r in range(1, len(correctedquery) + 1):
        for combo in combinations(correctedquery, r):
            for prod in product(*combo):
                for perm in permutations(prod):
                    all_combinations.append(' '.join(perm))

    all_combinations = list(all_combinations)  # Convert back to list if needed
    # print(all_combinations)

    # Now do the FTS using LIKE for partial matches and order by relevance
    fts_query = '''
    SELECT primaryTitle, tconst, genres, bm25(title_akas) as rank
    FROM title_akas
    WHERE
    primaryTitle LIKE ?
    ORDER BY rank
    LIMIT 5
    '''

    # for q in correctedquery:
    #     print(q)

    weighted_results = []
    for cq in all_combinations:
        # print(cq)
        cursor.execute(fts_query, ('%' + cq + '%',))
        weighted_results.extend(cursor.fetchall())

    #remove duplicates
    weighted_results = list(set(weighted_results))

    # Sort results by Levenshtein distance to the input query
    weighted_results = sorted(weighted_results, key=lambda x: Levenshtein.distance(query, x[0]))

    #only return the top 10 results
    results = weighted_results[:10]

    return results

results = search(sys.argv[1])
# results = search("frieren beyond jorn")
print(json.dumps(results))