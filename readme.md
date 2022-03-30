# Mech

> Listen, um, I don’t really know what’s going on here, but I’m leaving. I
> don’t know where exactly, but I’m gonna start over.
>
> Come with me.
>
> [Paint it Black (2016)][1]

Download media or send API requests

Some users might want to make anonymous requests, because of privacy or any
number of other reasons. Most API these days only offically support
authenticated access. This is useful for the company providing the API, as they
can use the data for their own purposes (analytics etc). However authentication
really doesnt do anything for the end user. Its just a pointless burden to
getting the information you need for a program you may be writing. Consider
that in many cases, the same information is available via HTML on the primary
website, usually without being logged in. So why can you do that with HTML, but
not with the API? Well you can, using this module.

[1]://f002.backblazeb2.com/file/ql8mlh/Paint.It.Black.2016.mp4

## How to install?

This module works with Windows, macOS or Linux. Check [the releases][2] first.
I dont do a build for every tag, but some tags will have builds available.
First, [download Go][3] and extract archive. Then [download Mech][4] and
extract archive. Then navigate to:

~~~
mech-master/cmd/youtube
~~~

and enter:

~~~
go build
~~~

[2]://github.com/89z/mech/releases
[3]://go.dev/dl
[4]://github.com/89z/mech/archive/refs/heads/master.zip

## Repo

https://github.com/89z/mech
