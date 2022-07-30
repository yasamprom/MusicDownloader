import sys, os
from youtube_dl import YoutubeDL
from pytube import YouTube


def search(arg):
    """
    :param arg: query
    :return: link to most relevant video
    """
    ydl_opts = {
        'quiet': True,
        'format': 'bestaudio',
        'noplaylist': 'True'
    }
    with YoutubeDL(ydl_opts) as ydl:
        try:
            video = ydl.extract_info(f"ytsearch:{arg}", download=False)['entries'][0]
        except Exception:
            return "", ""
    return video["title"], video["webpage_url"]


def main():
    title, link = search(sys.argv[1])
    if title == "":
        print("")
        exit(0)
    path = "data/" + title.replace('/', ' ').replace('.', '') + '.mp3'

    yt = YouTube(link)
    video = yt.streams.filter(only_audio=True).first()
    out_file = video.download()
    base, ext = os.path.splitext(out_file)
    new_file = base + '.mp3'
    os.rename(out_file, new_file)
    print(new_file)
    exit(0)


if __name__ == "__main__":
    main()
