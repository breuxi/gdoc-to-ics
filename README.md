# GDoc-To-Ics

GDocToIcs is an easy to use, high speed Google Docs Spreadsheets to Ics Converter. It turns your spreadsheets in subscribeable calendar feeds!

## Origin Story

On the 5th of November 2022, someone from the german vtuber community on twitter started to collect all vtuber birthdays in a google spreadsheet. I was amazed and thankful for the idea, but instantly thought "I will never be able to open this every day, I am already bad at keeping dates". To organize my life, I had build a few syncable virtual calendar services, but I was way too lazy to create calendar events for the whole spreadsheet. So I kept this thought in my mind for some time, and suddenly remembered that there was a way to share calendar data as an ICal/Ics file and that some people (e.g. my university) used this, to create subcribeable online calendars, that are easy to integrate in your existing setup. As I was looking forward to learn some more go, this seemed like a good side project. Now, here it is. Enjoy!

## Usage

If you have no development or sysadmin experience, I recommend using my hosted service at https://gdoc-to-ics.breuxi.de/.

Currently it is only able to read normal calendars in the format of the one in the "Origin Story", but I will maybe extend it to be much
more versatile.

To use it, you simply have to enter the url (e.g. https://gdoc-to-ics.breuxi.de/) followed by "/ical/" and then the ID of the Spreadsheet you want to use in your preferred calendar software. You can also specify a filename (some programs need that to name the calendar) by adding "/\<Your Filename>.ics" and a specific sheet with its id (by default 0, google calls it gid in the url).

### Basic Example

```
https://gdoc-to-ics.breuxi.de/ical/1iinbiLCiivIYpdDDHftpeYaZZYU7_ItRn8I8u9ZAIWA
```

### Filename Example

```
https://gdoc-to-ics.breuxi.de/ical/1iinbiLCiivIYpdDDHftpeYaZZYU7_ItRn8I8u9ZAIWA/GreatCalendar.ics
```

### Sheet Example

```
https://gdoc-to-ics.breuxi.de/ical/1iinbiLCiivIYpdDDHftpeYaZZYU7_ItRn8I8u9ZAIWA?sheet_id=1965686716
```

These urls can be put into almost all current calendar softwares.

### Subscribe in Thunderbird

https://support.mozilla.org/en-US/kb/creating-new-calendars#w_on-the-network-connect-to-your-online-calendars

### Subscribe in IOS Calendar

https://support.apple.com/guide/iphone/use-multiple-calendars-iph3d1110d4/ios

### Subscribe on Android using ICSx5

On android, you can use another great open source project to subscribe: [ICSx⁵](https://github.com/bitfireAT/icsx5)!

## Installing/Selfhosting

I recommend using Docker to run this app. You can set the preferred ip and port by exporting `GDTICS_LISTEN="\<Ip>:\<Port>"`.

### Docker

### Manual Install

## Contribute

I am happy for every pull request! But please create an issue first, so we can evaluate if it fits to the project!

## Roadmap

- Frontend to help users configure their feeds
- Reminder Support
- Color Support
