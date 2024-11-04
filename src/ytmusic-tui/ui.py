from logging import log
from textual.app import App, ComposeResult
from textual.containers import HorizontalGroup, Container
from textual.widgets import Static, Label, Input
from api import parseSongResults
from ytmusicapi import YTMusic

class SongsWidget(Static):
    def __init__(self, *args, song_list=[], **kwargs):
        super().__init__(*args, **kwargs)
        self.song_list = song_list
    def compose(self) -> ComposeResult:
        for song in self.song_list:
            yield Label(song)

class SearchWidget(Input):
    def __init__(self, update_function, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.update_function = update_function
    def on_input_submitted(self, event: Input.Submitted):
        self.update_function([event.value])

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
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.sidebar = Sidebar(id="sidebar")
        self.song_widget = SongsWidget(id="songs")

    def compose(self) -> ComposeResult:
        yield self.sidebar
        yield self.song_widget


class YtMusicTuiApp(App):
    """The main app for ytmusic-tui"""

    CSS_PATH = "layout.tcss"

    BINDINGS = [("d", "toggle_dark", "Toggle dark mode")]

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.search_widget = SearchWidget(placeholder="Search", type="text", update_function=self.update_song_window)
        self.content_container = ContentContainer(id="content")
        self.now_playing_widget = NowPlayingWidget(id="nowplaying")

    def update_song_window(self, song_list):
        self.content_container.song_widget.song_list = song_list
        self.refresh_content()

    def refresh_content(self):
        self.content_container.song_widget.refresh(recompose=True)

    def compose(self) -> ComposeResult:
        """Create child widgets for the app."""
        yield self.search_widget
        yield self.content_container
        yield self.now_playing_widget
        # yield SearchWidget(placeholder="Search", type="text", update_function=self.update_song_results)
        # yield ContentContainer(id="content", song_list=self.song_results)
        # yield NowPlayingWidget(id="nowplaying")

    def action_toggle_dark(self) -> None:
        """An action to toggle dark mode."""
        self.dark = not self.dark

