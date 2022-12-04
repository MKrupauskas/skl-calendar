# SKL Calendar

> Note that this was a Saturday midnight project so excuse the quality of the implementation and missing tests.

SKL Calendar is a simple go application that periodically scrapes the https://www.sostineskl.lt/ website and exposes a team's upcoming games in the `.ics` format.

The `/events` endpoint can then be used to subscribe to the events with Google Calendar or other calendar apps.

## Google Calendar

Attaching a team's games to your Google Calendar:

1. Click `+` next to the other calendars section in the sidebar.
2. Click `From URL`.
3. Enter the URL of the `/events?team={team-name}` endpoint of your instance.
4. Click `Add Calendar`.

## Developing

The application is a simple go server.

To run it locally run the command below.

```
go run main.go
```

The part most prone to future failure is the `event/parser` package. It depends on the HTML structure of the website which could change without notice.

Currently, only one team's calendar is exposed but it could be easily parameterized. Contributions are welcome.

## Deploying

https://fly.io/ is the simplest way to deploy the application.

Run the deploy command to deploy the project or release a new version.

```
flyctl deploy
```
