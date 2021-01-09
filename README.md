# mqtt-homekit-light

![goreleaser](https://github.com/adamcstephens/mqtt-homekit-light/workflows/goreleaser/badge.svg)

This creates a bridge between virtual homekit light and mqtt, allowing one to toggle an mqtt topic using homekit.

You can configure options using command line flags, or environment variables using the prefix `MHL_`.

```bash
Usage of mqtt-homekit-light:
  -debug
        Enable debugging
  -manufacturer string
        homekit accessory manufacturer (default "hc")
  -mqtt string
        mqtt url (default "mqtt://localhost:1883")
  -mqtt-client-id string
        mqtt client ID (default "homekit-mqtt")
  -mqtt-topic string
        topic to listen for state, will call <mqtt-topic>/set when homekit triggers (default "homekit/light")
  -name string
        homekit accessory name (default "Light")
  -pin string
        homekit PIN for pairing (default "32191123")
  -storage-path string
        where to store persistent files (default "./")
```

## Links

* [brutella/hc](https://github.com/brutella/hc)
