# GoZone
Zoneminder is a great surveilance system for DIY setups and Golang is great for small services within DIY setups. So it just makes sense to bring them together.

## Install
```
go get github.com/bah2830/GoZone
```

## Monitors

### Get All Monitors
```
import zm "github.com/bah2830/GoZone"

client, err := zm.NewClient("http://192.168.1.10/zm", "username", "password")
if err != nil {
    log.Fatal(err)
}

monitors, err := client.GetMonitors()
if err != nil {
    log.Fatal(err)
}

log.Printf("%+v", monitors)
```

### Getting Events For A Monitor
```
import zm "github.com/bah2830/GoZone"

client, err := zm.NewClient("http://192.168.1.10/zm", "username", "password")
if err != nil {
    log.Fatal(err)
}

monitor, err := client.GetMonitorById(1)
if err != nil {
    log.Fatal(err)
}

events, err := monitor.GetEvents(nil)
if err != nil {
    log.Fatal(err)
}

log.Printf("%+v", events)
```

### Monitoring For Events
```
import zm "github.com/bah2830/GoZone"

client, err := zm.NewClient("http://192.168.1.10/zm", "username", "password")
if err != nil {
    log.Fatal(err)
}

monitor, err := client.GetMonitorById(1)
if err != nil {
    log.Fatal(err)
}

eventChan := make(chan zoneminder.Events)
m.MonitorForEvents("Motion", 5, eventChan)

for {
    select {
    case e := <-eventChan:
        log.Printf("\n%+v", e)
    default:
        time.Sleep(30 * time.Second)
        m.StopEventMonitoring()
        return
    }
}
```

## Events
### Get All Events
This can return well over the default 100 pagination limit. Use filters to narrow that down.
```
import zm "github.com/bah2830/GoZone"

client, err := zm.NewClient("http://192.168.1.10/zm", "username", "password")
if err != nil {
    log.Fatal(err)
}

events, err := client.GetEvents(&zm.EventOpts{
    Cause: "Motion",
})

if err != nil {
    log.Fatal(err)
}

log.Printf("%+v", events)
```

### Monitor For Events
```
import zm "github.com/bah2830/GoZone"

client, err := zm.NewClient("http://192.168.1.10/zm", "username", "password")
if err != nil {
    log.Fatal(err)
}

eventChan := make(chan zoneminder.Events)
c.MonitorForEvents(&zm.EventOpts{Cause: "Motion"}, 5, eventChan)

for {
    select {
    case e := <-eventChan:
        log.Printf("\n%+v", e)
    default:
        time.Sleep(30 * time.Second)
        c.StopEventMonitoring()
        return
    }
}
```