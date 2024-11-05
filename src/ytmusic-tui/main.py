from ui import YtMusicTuiApp

if __name__ == "__main__":
    app = YtMusicTuiApp()
    if app.authorised:
        app.run()
    else:
        print("Please setup oauth by running 'ytmusicapi oauth' in the root folder of the package")
