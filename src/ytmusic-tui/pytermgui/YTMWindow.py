import pytermgui as ptg
from pytermgui.widgets.boxes import SINGLE
from typing import Any

class YTMWindow(ptg.Window):
    is_static = True
    is_noresize = True

    def __init__(self, *widgets: Any, **attrs: Any) -> None:
        super().__init__(*widgets, **attrs)
        self.box = SINGLE
