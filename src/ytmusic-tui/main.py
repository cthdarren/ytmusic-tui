from textual.app import App, ComposeResult
from textual.containers import HorizontalGroup, Container
from textual.widgets import Header, Footer, Placeholder
from textual.widgets import Button, Footer, Header, Static, Label

class SongsWidget(Static):
    def compose(self) -> ComposeResult:
        songs = ["Song1", "Song2", "Song3"]
        for song in songs:
            yield Label(song)

class SearchWidget(Static):
    DEFAULT_CSS = """
    SearchWidget {
        height: 3;
        dock: top;
    }
    """
    def compose(self) -> ComposeResult:
        yield Label("Search")

class PlaylistWidget(Static):
    def compose(self) -> ComposeResult:
        songs = ["Song1", "Song2", "Song3"]
        for song in songs:
            yield Label(song)

class NowPlayingWidget(Static):
    DEFAULT_CSS = """
    NowPlayingWidget {
        height: 3;
        dock: bottom;
    }
    """
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


if __name__ == "__main__":
    app = YtMusicTuiApp()
    app.run()
