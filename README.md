# IOT Mositure Sensor

## Purpose
1. Hardware will consist of a IOT device that send the following data
 - Location - a value
 - Tempture
 - Humidity
 - Battery Voltage
2. Data will be sent to an Influx timeseries database
3. Data will be visualized using Kibana dashboard

## Parts
- [ESP8266](https://learn.adafruit.com/adafruit-huzzah-esp8266-breakout)
- [3xAA Batteries]
- [DHT22](https://www.adafruit.com/product/385)
- [Adafruit BMP280 I2C or SPI Barometric Pressure & Altitude Sensor](https://www.adafruit.com/product/2651?gclid=CjwKCAjwm4rqBRBUEiwAwaWjjF3XTMTRwt6PhmwsGnPRPdA7HlE_gyvQVNLfKLg5y95S2kj3FOktUxoCvwYQAvD_BwE)

## Todo
- [ ] Wire up the DHT22 and BMP280 sensor to the ESP8266
- [ ] Create a voltage indicator circuit to the ESP8266
- [ ] Build water proof case

## Possible Parts
- [Adafruit HTU21D-F Temperature & Humidity Sensor Breakout Board](https://www.adafruit.com/product/1899)
- [Analog/Digital MUX Breakout](https://www.sparkfun.com/products/9056)
- [Powerboost 1000c](https://learn.adafruit.com/adafruit-powerboost-1000c-load-share-usb-charge-boost/pinouts)
 - LBO pin will goto 0V when the battery is low

## Notes
- [A battery fed MQTT weatherstation](https://arduinodiy.wordpress.com/2018/02/04/a-battery-fed-mqtt-weatherstation/)
- [Monitoring LiPo battery voltage with Wemos D1 minibattery shield](https://arduinodiy.wordpress.com/2016/12/25/monitoring-lipo-battery-voltage-with-wemos-d1-minibattery-shield-and-thingspeak/)
- [Arduino Battery Voltage Indicator](https://www.instructables.com/id/Arduino-Battery-Voltage-Indicator/)
- [Photon Battery Shield Hookup Guide](https://learn.sparkfun.com/tutorials/photon-battery-shield-hookup-guide/all)
- [ESP8266 on batteries for years](https://www.cron.dk/esp8266-on-batteries-for-years-part-1/)
