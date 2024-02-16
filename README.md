# `unjn`: Upwork New Jobs Notifier

`unjn` (pronounced "un-gin" or like "engine" but with an "un") is a simple program written in [Go](https://go.dev/) that sends you notifications via [ntfy](https://ntfy.sh/) when new jobs are posted on Upwork.

## Usage

First, install the ntfy app on your phone. You can find links to the app stores [here](https://ntfy.sh/).

Download `unjn` from the [FIXME: releases page](/), and run like so:

```
./unjn ntfy_topic upwork_rss_url_1 upwork_rss_url_2
```

## FIXME: dev notes

https://github.com/mmcdole/gofeed
https://pkg.go.dev/heckel.io/ntfy@v1.31.0/client#Client.Publish
