# IOT Weather Station

## Purpose
1. Hardware will consist of a IOT device that send the following data
 - Location
 - Hardware version
 - Device number - each device gets a number (incremented value)
 - Temperature
 - Humidity
 - Battery Voltage
2. Data will be sent to an Influx timeseries database
3. Data will be visualized using a Kibana dashboard

## Hardware
### Parts
- [ESP8266](https://learn.adafruit.com/adafruit-huzzah-esp8266-breakout)
- [DHT22](https://www.adafruit.com/product/385)
- [Adafruit BMP280 I2C or SPI Barometric Pressure & Altitude Sensor](https://www.adafruit.com/product/2651?gclid=CjwKCAjwm4rqBRBUEiwAwaWjjF3XTMTRwt6PhmwsGnPRPdA7HlE_gyvQVNLfKLg5y95S2kj3FOktUxoCvwYQAvD_BwE)
- [Lithium Ion Cylindrical Battery - 3.7v 2200mAh](https://www.adafruit.com/product/1781) - not sure if this will work with the charging board
- [D1 Mini Single Lithium Battery Charging Board](https://www.amazon.com/WINGONEER-Single-Lithium-Battery-Charging/dp/B077VNW5RP/ref=sr_1_3?keywords=d1+mini+battery+shield&qid=1565500842&s=gateway&sr=8-3)

## Todo
## Hardware
- [x] Wire up the DHT22 and BMP280 sensor to the ESP8266
- [ ] Create a voltage indicator circuit to the ESP8266
- [ ] Build water proof case
## Software
- [ ] IOT weather stations send message to mosquitto mqtt server
- [ ] Setup passwords on mosquitto mqtt server
- [ ] Go app subscribes to topic and receives message and logs the data to influx database
## Docker
- [ ] Dockerize influx database
- [ ] Dockerize golang app
- [ ] Dockerize mosquitto mqtt server

## Notes
- [A battery fed MQTT weatherstation](https://arduinodiy.wordpress.com/2018/02/04/a-battery-fed-mqtt-weatherstation/)
- [Monitoring LiPo battery voltage with Wemos D1 minibattery shield](https://arduinodiy.wordpress.com/2016/12/25/monitoring-lipo-battery-voltage-with-wemos-d1-minibattery-shield-and-thingspeak/)
- [Arduino Battery Voltage Indicator](https://www.instructables.com/id/Arduino-Battery-Voltage-Indicator/)
- [Photon Battery Shield Hookup Guide](https://learn.sparkfun.com/tutorials/photon-battery-shield-hookup-guide/all)
- [ESP8266 on batteries for years](https://www.cron.dk/esp8266-on-batteries-for-years-part-1/)
- [Solar Wifi Weather Station (v2.0)](https://www.danilolarizza.com/stazione-meteo-solare-wifi-v2-0/)
- [Mosquitto MQTT](https://www.switchdoc.com/2018/02/tutorial-installing-and-testing-mosquitto-mqtt-on-raspberry-pi/)
- [Mosquitto Username and Password](http://www.steves-internet-guide.com/mqtt-username-password-example/)
