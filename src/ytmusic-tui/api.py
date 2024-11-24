from ytmusicapi import YTMusic
from vlc import MediaPlayer
from player import play_stream_from_url

class Singleton(type):
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super(Singleton, cls).__call__(*args, **kwargs)
        return cls._instances[cls]

class ApiSingleton(metaclass=Singleton):
    def __init__(self):
        self.yt = YTMusic("browser.json")

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

    def get_song_stream(self, vidId) -> str:
        song_meta = self.yt.get_song(vidId)
        stream_data = song_meta.get("streamingData", None)
        if stream_data == None:
            raise ValueError
        stream_url = stream_data.get("serverAbrStreamingUrl", None)
        # formats = stream_data.get("adaptiveFormats", None)
        # if formats == None:
        #     raise ValueError
        # if len(formats) < 0:
        #     raise ValueError
        # stream_url = formats[0].get("url", None)
        if stream_url == None: 
            raise ValueError

        return stream_url

    def play_stream(self, vidId):
        stream_url = self.get_song_stream(vidId)
        play_stream_from_url(stream_url)


    def perform_search(self, query):
        results = self.yt.search(query, filter="songs", limit=20)
        # self.play_stream(results[0]["videoId"])
        return results

    def parse_song_results(self, search_result):
        res = []
        if len(search_result) > 0:
            for result in search_result:
                if not "resultType" in result or not "title" in result:
                    continue
                if result["resultType"] == "song":
                    res.append(result["title"])

        return res
