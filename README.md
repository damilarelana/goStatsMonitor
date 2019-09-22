# goStatsMonitor

## This is a monitoring application that displays a Youtube channel's statistic

```
Views count i.e. number of views
Subscribers count i.e. number of subscribers
Videos count i.e. number of videos
```

### The application consists of a frontend (index.html) that is feed with JSON data by a backend (youtube.go) through a websocket connection (facilitated by websocket.go) between the frontend/backend.

### The implementation leverages

* [Youtube's API](https://www.googleapis.com/youtube/v3/channels) [consumed via youtube.go]
* Docker image [backend generated via Dockerfile]
* [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux) [to handle endpoints resquest/dispatcher]
