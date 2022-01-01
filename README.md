# Notice
Widevine is currently revoking a lot of keys.  This program won't work unless you have your own Widevine key and device information (I do not have any working keys).

# Hulu Downloader
The code in this repository allows you to download videos unencumbered with DRM from Hulu.  The code in `widevine` is in general independent of the Hulu related code and can be used for Widevine license generation/decryption.  The code in `client` is also standalone but only implements a handful of Hulu API endpoints that are basically only useful for a tool of this nature.

## Prerequisites
The code in this repository by itself does not require any external libraries or tools to be installed.  It merely finds the video URLs and decryption keys. The only dependencies required are cryptographic libraraies specified in go.mod but Go should handle these automatically.  However, to actually perform MP4 decryption, Bento4 (and specifically its `mp4decrypt` tool) are required.  Bento4 is an open source library for MP4 manipulation.  Binary releases of its tools can be downloaded [here](https://www.bento4.com/downloads/).  [`yt-dlp`](https://github.com/yt-dlp/yt-dlp) is also required to download the MPD playlist files to mp4s.  Technically, this could be implemented rather easily in this repository but I want to keep this repository simple and avoid rewriting code to deal with segment merging or quality selection menus.

## Retrieving Hulu Session Cookie
Hulu requires Recaptcha for authentication so just passing account credentials is not possible without captcha solving services.  To work around this, this tool simply takes a Hulu session cookie.

> Note: Ensure you are signed in before following these steps.

### Chrome
Visit [https://hulu.com](https://hulu.com).  Click the lock icon in the URL bar.  Then select the item labelled Cookies.  Then find hulu.com in the list, select it, and expand the "Cookies" list with an icon that looks like a folder.  Then select the cookie titled `_hulu_session`.  Chrome will then show various attributes of this cookie.  Right click the area labelled "Content", press select all and then right click again and press copy.  The value of the Hulu session cookie is now on your clipboard.  A demonstration can be found [here](https://www.cookieyes.com/wp-content/uploads/2021/10/chrome2.mp4).

### Firefox
Visit [https://hulu.com](https://hulu.com).  Right click and then click Inspect.  Then visit the Storage tab.  Now, under the cookies pane on the left, select hulu.com.  Then retrieve the value of `_hulu_session` from the list of cookies.  A demonstration can be found [here](https://www.cookieyes.com/wp-content/uploads/2021/10/firefox1.mp4).

## Demonstration
Say we want to download an episode of M\*A\*S\*H.

    $ go install github.com/chris124567/hulu # The rest of these commands assume $GOPATH/bin is in your PATH.  If it is not, just cd to $GOPATH/bin and run "./hulu" instead of "hulu"
    $ HULU_SESSION="abc" hulu search -query="m*a*s*h"
    Title                           ID
    M*A*S*H                         ae94231d-0f04-482a-b9ee-9911e339e3ed
    MASH (1970)                     42f7eefe-2448-4ed5-87cb-6233c89c20f6
    American Psycho (2000)          404a410c-ef36-469d-8fcd-1f93ec44a5c0
    American Horror Story           a67a233c-fcfe-4e8e-b000-052603ddd616
    Hitman: Agent 47 (2015)         a4d96c8d-ba7d-4d99-b4b3-942ecde47282
    Ma (2019)                       dbb13a18-79d2-4567-8ed4-e2eddbec9492
    The Martian (2015)              e52328e3-6e2b-4565-91d5-2f7ee7c846ab
    HBO Max                         1b3523c1-3090-4c27-a1e8-a04d33867c34
    ...

We want the M\*A\*S\*H TV show, so we choose the ID "ae94231d-0f04-482a-b9ee-9911e339e3ed."  We want to look at the first season so we specify that the season number equals 1.

    $ HULU_SESSION="abc" hulu season -id="ae94231d-0f04-482a-b9ee-9911e339e3ed" -number=1
    Title                           ID
    Pilot                           4045ee04-07e8-4c33-94a6-4244b7b67c5f
    To Market, to Market            7a43d075-2b47-4c94-8767-8531e20bab81
    Requiem for a Lightweight       2ccd2cf5-a013-4501-a689-1ed6b94a9549
    Chief Surgeon Who?              112b061b-1c18-4f15-bed8-042d44919735
    The Moose                       688e10d3-6db4-47ba-a99b-bfd8aacd6c7a
    Yankee Doodle Doctor            3ff14d70-e2ac-4bc2-83c6-87b9cf132c13
    ...

Now to get the episode we want (the pilot), pass the ID of the episode to the `download` subcommand.

> Note: If we wanted to download a movie, instead of getting the episode list (which movies don't have) and selecting the specific episode ID, just pass the original ID from the search results above to `download`.


    $ HULU_SESSION="abc" hulu download -id="4045ee04-07e8-4c33-94a6-4244b7b67c5f"
    MPD URL:  https://manifest-dp.hulustream.com/OMITTED
    Decryption command:  mp4decrypt input.mp4 output.mp4 --key OMITTED:OMITTED

Now we have the URL and the keys.  First, let's see what formats are available:

    $ yt-dlp --allow-unplayable-formats -F "https://manifest-dp.hulustream.com/OMITTED"
    WARNING: You have asked for unplayable formats to be listed/downloaded. This is a developer option intended for debugging. 
             If you experience any issues while using this option, DO NOT open a bug report
    [generic] xxxxxxxx: Requesting header
    WARNING: [generic] Falling back on generic information extractor.
    [generic] xxxxxxxx: Downloading webpage
    [generic] xxxxxxxx: Extracting information
    [info] Available formats for xxxxxxxx:
    ID               EXT RESOLUTION |   TBR PROTO | VCODEC        VBR ACODEC     ABR  ASR    MORE INFO
    ---------------- --- ---------- - ----- ----- - ----------- ----- --------- ---- ------- --------------------------
    132545434.add-0  m4a audio only |   68k https |                   mp4a.40.5  68k 48000Hz [en], DASH audio, m4a_dash
    132545434.add-1  m4a audio only |   68k https |                   mp4a.40.5  68k 48000Hz [en], DASH audio, m4a_dash
    ...
    132545434.add-11 m4a audio only |   68k https |                   mp4a.40.5  68k 48000Hz [en], DASH audio, m4a_dash
    132545134.vdd-0  mp4 512x288    |  460k https | avc1.640015  460k                        DASH video, mp4_dash
    132545134.vdd-1  mp4 512x288    |  460k https | avc1.640015  460k                        DASH video, mp4_dash
    ...
    132545134.vdd-11 mp4 512x288    |  460k https | avc1.640015  460k                        DASH video, mp4_dash

Let's get the audio first.  We will choose `132545434.add-0`, the lowest quality format, for this example. Download it with:

    $ yt-dlp --allow-unplayable-formats -f "132545434.add-0" "https://manifest-dp.hulustream.com/OMITTED" -o audio.mp4

Next we will get the video.  We will also just take the lowest quality format (`132545134.vdd-0`) here.

    $ yt-dlp --allow-unplayable-formats -f "132545134.vdd-0" "https://manifest-dp.hulustream.com/OMITTED" -o video.mp4

Now we should have two mp4 files, one for the video and one for the audio.  We ultimately will merge these, but first we need to decrypt them.

Remember the mp4decrypt command from above?  Specifically look at the `--key OMITTED:OMITTED` part.  The decryption key is the same for both the video and the audio.  So we can run:

    $ mp4decrypt audio.mp4 audio_dec.mp4 --key OMITTED:OMITTED
    $ mp4decrypt video.mp4 video_dec.mp4 --key OMITTED:OMITTED

Finally, we can merge the two sources (this command does not do any reencoding):

    $ ffmpeg -i video_dec.mp4 -i audio_dec.mp4 -acodec copy -vcodec copy merged.mp4

And now merged.mp4 will be a DRM free mp4 file straight from Hulu!  It is possible to automate these steps by writing a simple script.

## TODO
- Subtitles
- Storing authentication cookie in a text file to avoid having to pass it for every command

## Credits
The bulk of the Widevine related code was ported from `pywidevine` which is a library floating around the Internet of unknown provenance.