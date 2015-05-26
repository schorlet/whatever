#whatever challenge

This is the source code of my participation to the [Knowledge Plaza Backend Challenge](https://github.com/whatever-company/challenge/).

##Introduction

This application connects to Twitter using [Application-only authentication](https://dev.twitter.com/oauth/application-only).

So you have to run it with your own credentials.

You can obtain your consumer key and secret by creating a new app on [apps.twitter.com](https://apps.twitter.com/).


Then you have to set the value of the environment variables:

```sh
$ export CONSUMER_KEY="Your consumer key"
$ export CONSUMER_SECRET="Your consumer secret"
```


##Run the client

```sh
# start the client
$ go run client.go "#Arduino"
```

You should see:

```
[
    {
        "text": "RT @InitialState: #Arduino and #RaspberryPi in harmony - As a sensor web server! http://t.co/PeyslaZo1H by @instructables http://t.co/DLQTp…",
        "created_at": "Mon May 25 18:11:03 +0000 2015",
        "retweet_count": 1,
        "favorite_count": 0,
        "user": {
            "name": "Rasbot",
            "screen_name": "RPi2Bot"
        },
        "retweeted_status": {
            "user": {
                "name": "Initial State",
                "screen_name": "InitialState"
            }
        },
        "entities": {
            "hashtags": [
                {
                    "text": "Arduino"
                },
                {
                    "text": "RaspberryPi"
                }
            ]
        }
    },
    ...
```


##Run the server

```sh
# install dependencies listed in bower.json
$ bower install

# start the webapp
$ go run server.go
```

Go to [http://localhost:8000/](http://localhost:8000/) to run the web UI.



