from tkinter import Tk, BOTH, Canvas


def window():
    janela = Tk()
    data = {
        "width": 800,
        "height": 600,
        "title": "minha janela",
        "bg": "black",
        "line_color": "red",
        "line_size": 3,
    }
    janela.title = data["title"]
    janela.minsize(data["width"], data["height"])
    dark = Canvas(janela, width=data["width"], height=data["height"], bg=data["bg"])
    dark.pack()
    # border
    dark.create_line(
        0,
        0,
        data["width"],
        0,
        fill=data["line_color"],
        width=data["line_size"],
    )
    dark.create_line(
        0,
        0,
        0,
        data["height"],
        fill=data["line_color"],
        width=data["line_size"],
    )
    dark.create_line(
        data["width"],
        0,
        data["width"],
        data["height"],
        fill=data["line_color"],
        width=data["line_size"],
    )
    dark.create_line(
        0,
        data["height"],
        data["width"],
        data["height"],
        fill=data["line_color"],
        width=data["line_size"],
    )
    janela.mainloop()


if __name__ == "__main__":
    window()
