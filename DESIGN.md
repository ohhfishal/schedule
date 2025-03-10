
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
