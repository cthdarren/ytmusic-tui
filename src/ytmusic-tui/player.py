import json
from pydub import AudioSegment
from pydub.playback import play
import requests
from io import BytesIO

def play_stream_from_url(stream_url:str):
    with open("browser.json", "r") as file:
        cookie_data = json.load(file)
        response = requests.get(stream_url, stream=True, cookies=cookie_data)

        audio = AudioSegment.from_file(BytesIO(response.content), format="mp4")

        play(audio)
