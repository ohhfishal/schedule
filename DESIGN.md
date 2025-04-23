
## Schedule Schema Planning
```json
# Event
{
  "id": int64,
  "name": string,
  "description": Optional[string],
  "start": DATE,
  "end": Optional[Date],
  "status": Enum("confirmed", "tenative", "cancelled"),
  "created": time.Time,
  "updated": time.Time,
}
# Date
type Date struct {
  DateTime time.Time
  Date string
  Timezone *time.Location
}
```

## Inspired from Google's API
```json
{
  "id": string,
  "status": string,
  "created": datetime,
  "updated": datetime,
  "summary": string,
  "description": string,
  "location": string, ??
  "colorId": string, !!
  "creator": {
    "id": string,
    "email": string, ??
    "displayName": string, ??
    "self": boolean ??
  },
  "start": {
    "date": date,
    "dateTime": datetime,
    "timeZone": string
  },
  "end": {
    "date": date,
    "dateTime": datetime,
    "timeZone": string
  },
  "recurrence": {
    ...
    "terminates": boolean,
  },
  "recurringEventId": string, // ID of the instance this is
  "originalStartTime": {
    "date": date,
    "dateTime": datetime,
    "timeZone": string
  },
  "locked": boolean,
  "reminders": {
    "useDefault": boolean,
    "overrides": [
      {
        "method": string,
        "minutes": integer
      }
    ]
  },
  "source": {
    "url": string,
    "title": string
  },
  "attachments": [
    {
      "fileUrl": string,
      "title": string,
      "mimeType": string,
      "iconLink": string,
      "fileId": string
    }
  ],
  "birthdayProperties": {
    "contact": string,
    "type": string,
    "customTypeName": string
  },
  "eventType": string
}
```
