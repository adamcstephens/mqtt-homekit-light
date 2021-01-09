# mqtt-homekit-light

![goreleaser](https://github.com/adamcstephens/mqtt-homekit-light/workflows/goreleaser/badge.svg)

This creates a bridge between virtual homekit light and mqtt, allowing one to toggle an mqtt topic using homekit.

You can configure options using command line flags, environment variables using the prefix `MHL_`, or with a JSON config file.

Packages are available for RPM and DEB, on am64 and aarch64 architectures.

## CLI

```bash
Usage of mqtt-homekit-light:
  -config string
        Config file to load settings from
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

## Environment Variables

Environment variables match the CLI flags; prefixed with `MHL_`, in all caps, and underscores replacing hyphens.

```bash
MHL_MQTT_CLIENT_ID=myclient mqtt-homekit-light
```

## Configuration File

```json
{
      "mqtt-client-id": "myclient"
}
```

## Running with systemd

Requires systemd 235 or user due to the use of [dynamic users](http://0pointer.net/blog/dynamic-users-with-systemd.html). Tested on versions 246 and 247.

The systemd service can be used to run multiple instances using a systemd template unit.

1. Create a JSON config file in `/etc/mqtt-homekit-light/<name>.json`
2. Enable and start the service, `systemctl enable --now mqtt-homekit-light@<name>.service`

## Links

* [brutella/hc](https://github.com/brutella/hc)
