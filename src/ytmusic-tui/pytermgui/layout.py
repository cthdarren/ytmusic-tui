import pytermgui as ptg
from YTMWindow import YTMWindow

# searchwindow = ptg.Window().set_title("Search")
searchwindow = YTMWindow().set_title("Search")
playlistwindow = YTMWindow().set_title("Playlists")
songswindow = YTMWindow().set_title("Songs")
nowplayingwindow = YTMWindow().set_title("Now Playing")

manager = ptg.WindowManager()
manager.layout.add_slot("Header", height=3)
manager.layout.add_break()
manager.layout.add_slot("Playlists", width=30)
manager.layout.add_slot("Main")
manager.layout.add_break()
manager.layout.add_slot("NowPlaying", height=5)
manager.add(searchwindow)
manager.add(playlistwindow)
manager.add(songswindow)
manager.add(nowplayingwindow)
