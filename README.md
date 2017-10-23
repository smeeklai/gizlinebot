# GIZ Line Bot (App Engine)

## Setting up
### Install and initialize Google Cloud SDK
**You can skip this part if you prefer to use [Google Cloud Shell](https://cloud.google.com/shell/docs/quickstart) instead**

Follow instruction [here](https://cloud.google.com/appengine/docs/standard/go/download)

### Create your new yaml file
```
cd myapp/
cp .app-example.yaml YOURNAME-app.yaml
```
### Config YOURNAME-app.yaml
- Specify your `service` name
- Update your `CHANNEL_SECRET` and `CHANNEL_TOKEN` from the line developer area [Messaging API](https://developers.line.me/)

## Run the line bot locally
### Setup mysql server on your local machine
1. Install via [Homebrew](https://gist.github.com/nrollr/3f57fc15ded7dddddcc4e82fe137b58e) or Download and install mysql community server [here](https://dev.mysql.com/downloads/mysql/)
2. Import the SQL Schema from assets/init.sql

### Setup ngrok
Install [ngrok](https://ngrok.com/download)

### Start the bot
1. Config parameters in YOURNAME-app.yaml
    - Must Set `LOCAL_DB` in YOURNAME-app.yaml to `True`
    - Set your SQL parameters
2. Start an ngrok tunnel, run `ngrok http 127.0.0.1:8888`
3. start the bot
    - Make sure you've **installed and started mysql server** on your local machine
    - Run `dev_appserver.py YOURNAME-app.yaml --port 8888`

## Deploy the line bot on app engine
1. Replace sql parameters in YOURNAME-app.yaml with below parameters
```
LOCAL_DB: False
GIZLB_SQL_USER: gizlinebot
GIZLB_SQL_PASS: 'Mpmc3EzwUU06Pq9hq8T55fEnaN2okglRd5CPS2i4fcA'
GIZLB_SQL_HOST: vagabonddataninjas:asia-east1:gizsurvey
GIZLB_SQL_PORT: ''
GIZLB_SQL_DB: gizsurvey
```
2. Run `gcloud app deploy YOURNAME-app.yaml --version YOUR-VERSION`
