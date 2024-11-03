from textual.app import App, ComposeResult
from textual.containers import HorizontalGroup, Container
from textual.widgets import Static, Label, Input
from api import parseSongResults
from ytmusicapi import YTMusic

class SongsWidget(Static):
    def __init__(self, *args,song_list, **kwargs):
        super().__init__(*args, **kwargs)
        self.song_list = song_list
    def compose(self) -> ComposeResult:
        for song in self.song_list:
            yield Label(song)

class SearchWidget(Input):
    def compose(self) -> ComposeResult:
        yield Label("Search")
        # yield Input(placeholder="Search", type="text")

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
    def __init__(self, *args, song_list, **kwargs):
        super().__init__(*args, **kwargs)
        self.song_list = song_list
    def compose(self) -> ComposeResult:
        yield Sidebar(id="sidebar")
        yield SongsWidget(id="songs", song_list=self.song_list)


class YtMusicTuiApp(App):
    """The main app for ytmusic-tui"""

    CSS_PATH = "layout.tcss"

    BINDINGS = [("d", "toggle_dark", "Toggle dark mode")]

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.song_results = ['test']

    def compose(self) -> ComposeResult:
        """Create child widgets for the app."""
        yield Input(placeholder="Search", type="text")
        # yield SearchWidget(id="search")
        yield ContentContainer(id="content", song_list=self.song_results)
        yield NowPlayingWidget(id="nowplaying")

    def action_toggle_dark(self) -> None:
        """An action to toggle dark mode."""
        self.dark = not self.dark
