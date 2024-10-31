from layout import manager, playlistwindow
import pytermgui as ptg
from ytmusicapi import YTMusic

if __name__ == "__main__":
    test = ptg.Label('Test')
    test.pos = (0,0)
    playlistwindow.__add__(test)
    # yt = YTMusic('oauth.json')
    # playlistId = yt.create_playlist('test', 'test description')
    # search_results = yt.search('Oasis Wonderwall')
    manager.run()
