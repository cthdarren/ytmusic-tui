from ytmusicapi import YTMusic

class Singleton(type):
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super(Singleton, cls).__call__(*args, **kwargs)
        return cls._instances[cls]

class ApiSingleton(metaclass=Singleton):
    def __init__(self):
        self.yt = YTMusic("oauth.json")

    def get_playlist_names(self, playlist_results: list[dict]) -> list[str]:
        playlist_title_list = []
        for playlist in playlist_results:
            playlist_title = playlist.get("title", None)
            if playlist_title == None:
                continue
            playlist_title_list.append(playlist_title)

        return playlist_title_list

    def get_playlists(self) -> list[str]:
        return self.get_playlist_names(self.yt.get_library_playlists())

    def perform_search(self, query):
        return self.yt.search(query)

    def parse_song_results(self, search_result):
        res = []
        if len(search_result) > 0:
            for result in search_result:
                if not "resultType" in result or not "title" in result:
                    continue
                if result["resultType"] == "song":
                    res.append(result["title"])

        return res
