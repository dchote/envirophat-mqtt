# envirophat-mqtt

A simple go app that reads sensor information from the Pimoroni Enviro pHAT and publishes it to the specified MQTT endpoint at the specified interval.

```
Usage: envirophat-mqtt [options]
Options:
  --server=<server>          MQTT server host/IP [default: 127.0.0.1]
  --port=<port>              MQTT server port [default: 1883]
  --topic=<topic>            MQTT topic prefix [default: envirophat]
  --clientid=<clientid>      MQTT client identifier [default: envirohat]
  --interval=<interval>      Poll interval (seconds) [default: 5]
  -h, --help                 Show this screen.
  -v, --version              Show version.
```
