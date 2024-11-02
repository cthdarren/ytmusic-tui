from textual.app import App, ComposeResult
from textual.containers import HorizontalGroup, Container
from textual.widgets import Static, Label
from api import getSongDetails

from ytmusicapi import YTMusic

# yt = YTMusic('oauth.json')
yt = YTMusic()
search_results = yt.search('Oasis Wonderwall')

class SongsWidget(Static):
    def compose(self) -> ComposeResult:
        # songs = ["Song1", "Song2", "Song3"]
        songs = getSongDetails(search_results)
        for song in songs:
            yield Label(song)

class SearchWidget(Static):
    def compose(self) -> ComposeResult:
        yield Label("Search")

class PlaylistWidget(Static):
    def compose(self) -> ComposeResult:
        songs = ["Song1", "Song2", "Song3"]
        for song in songs:
            yield Label(song)

class NowPlayingWidget(Static):
    def compose(self) -> ComposeResult:
        yield Label("SongName")
        yield Label("Artist")

class MetaWidget(Static):
    def compose(self) -> ComposeResult:
        songs = ["Song1", "Song2", "Song3"]
        for song in songs:
            yield Label(song)

class Sidebar(Container):
    def compose(self) -> ComposeResult:
        yield MetaWidget(id="meta")
        yield PlaylistWidget(id="playlists")

class ContentContainer(HorizontalGroup):
    def compose(self) -> ComposeResult:
        yield Sidebar(id="sidebar")
        yield SongsWidget(id="songs")


class YtMusicTuiApp(App):
    """The main app for ytmusic-tui"""

    CSS_PATH = "layout.tcss"

    BINDINGS = [("d", "toggle_dark", "Toggle dark mode")]

    def compose(self) -> ComposeResult:
        """Create child widgets for the app."""
        yield SearchWidget(id="search")
        yield ContentContainer(id="content")
        yield NowPlayingWidget(id="nowplaying")

    def action_toggle_dark(self) -> None:
        """An action to toggle dark mode."""
        self.dark = not self.dark
