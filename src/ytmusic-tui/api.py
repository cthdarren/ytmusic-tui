

def getSongDetails(searchResult):
    res = []
    if len(searchResult) > 0:
        for result in searchResult:
            if not "resultType" in result or not "title" in result:
                continue
            if result["resultType"] == "song":
                res.append(result["title"])

    return res

                


