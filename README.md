# mqtt-homekit-light

This creates a bridge between homekit and mqtt, specifically for a homekit light.

```bash
 -> go run ./ -help
Usage of /var/folders/3p/q9nkzynn60j6cbfhcykhh3340000z8/T/go-build214451856/b001/exe/mqtt-homekit-light:
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
exit status 2
```

## Links

* [brutella/hc](https://github.com/brutella/hc)
