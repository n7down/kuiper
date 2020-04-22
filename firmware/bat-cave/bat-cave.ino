#include "config.h"
#include "DHT.h"

#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <Wire.h>
#include <SPI.h>
#include <PubSubClient.h>
#include <Adafruit_Sensor.h>
#include <ArduinoJson.h>
 
#define DEBUG 1
#define DHTPIN D5
#define SENDDATAPIN D6
#define DHTTYPE DHT22

const char dht22Topic[] = "sensor/dht22";
const char statsTopic[] = "sensor/stats";
const char settingsTopic[] = "bc/settings";

char mac[12];

// settings
int deepSleepDelay = 15; // in min

DHT dht22(DHTPIN, DHTTYPE);
WiFiClient espClient;
PubSubClient client(espClient); 

void callback(char* topic, byte* payload, unsigned int length) {
  // TODO: parse out binary message
  // 2 types of settings
  // 1. how often to send data
  // 2. how often to send heartbeat - once/twice/three/four times a day - send back the voltage with the heartbeat
  // when done sleeping - set the setting, send back a ack
  // send back on 'bc/ack' with the type (SUCCESS or FAILURE) the device (bc1) and the command that was sent - send that to influx
#ifndef DEBUG
  Serial.print("Message arrived: ");
  Serial.print(topic);
  Serial.print("- '");
  for (int i=0;i<length;i++) {
    Serial.print((char)payload[i]);
  }
  Serial.print("'");
  Serial.println();
#endif
}

void setupWifi(const char* ssid, const char* password)
{
#ifndef DEBUG
  Serial.print("Wifi connecting...");
#endif
  WiFi.begin(ssid, password);
  
  // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
#ifndef DEBUG
    Serial.print("...");
#endif
    // if (count >= 10 && (digitalRead(SLEEPPIN) ) == LOW) {
    //  Serial.println("Back to sleep, try again in 30 sec ");
    //  count = 0;
    //  ESP.deepSleep(30 * 1000000);
    // }
  }
#ifndef DEBUG
  Serial.print("Connected to ");
  Serial.println(ssid);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());
#endif

  String m = WiFi.macAddress();
  String macWithOutColons;
  for (int i = 0; i < m.length(); i++) {
    char currentChar = m[i];
    if (currentChar != ':') {
      macWithOutColons += currentChar;
    }
  }

  macWithOutColons.toLowerCase();
  macWithOutColons.toCharArray(mac, 12);

#ifndef DEBUG
  Serial.print("MAC: ");
  Serial.println(mac);
#endif
}

void reconnect() {
  while (!client.connected()) {    
#ifndef DEBUG
    Serial.print("Attempting MQTT connection (");
    Serial.print(mqtt_server);
    Serial.print(")... ");
#endif
    // Attempt to connect
    if (client.connect(mqtt_server)) {
#ifndef DEBUG
      Serial.println("connected");
#endif
      String deviceTopic = "devices/"
      deviceTopic += mac
      client.subscribe(deviceTopic);
    } else {
#ifndef DEBUG
      Serial.print("Failed: ");
      Serial.print(client.state());
      Serial.println(" - try again in 5 seconds");
#endif
      
      // Wait 5 seconds before retrying
      delay(5000);
    }
  }
  digitalWrite(LED_BUILTIN, HIGH); 
}

void setup() {
  pinMode(LED_BUILTIN, OUTPUT);
  
  digitalWrite(LED_BUILTIN, LOW);
#ifndef DEBUG
  Serial.begin(115200);
  Serial.println("Starting serial..");
#endif
  dht22.begin();

  setupWifi(ssid, password);
  
  client.setServer(mqtt_server, 1883);
  client.setCallback(callback);
}

void loop() {
  unsigned long startTime = millis();

  if (!client.connected()) {
    reconnect();
  }

  unsigned long elapsedTime = millis() - startTime;

  float h = dht22.readHumidity();
  float t = dht22.readTemperature();

  unsigned int batt = analogRead(A0);
  double battV = batt * (4.2 / 1023);

  StaticJsonDocument<100> dht22Root;
  dht22Root["m"] = mac;
  dht22Root["h"] = String(h); // % 
  dht22Root["t"] = String(t); // %

  char dht22Message[100];
  serializeJson(dht22Root, dht22Message); 
  
  int result = client.publish(dht22Topic, dht22Message);
#ifndef DEBUG
  Serial.print("Sent message: ");
  Serial.print(dht22Message);
  Serial.print(" - Result: ");
  Serial.println(result);
#endif
  
  char elapsedTimeString[40];
  sprintf(elapsedTimeString, "%u", elapsedTime);
#ifndef DEBUG
  Serial.print("Sending stats: ");
  Serial.println(elapsedTimeString);
#endif
  
  StaticJsonDocument<100> statsRoot;
  statsRoot["m"] = mac;
  statsRoot["v"] = String(battV);
  statsRoot["c"] = String(elapsedTimeString);

  char statsMessage[100];
  serializeJson(statsRoot, statsMessage); 
  
  result = client.publish(statsTopic, statsMessage);
#ifndef DEBUG
  Serial.print("Sent message: ");
  Serial.print(statsMessage);
  Serial.print(" - Result: ");
  Serial.println(result);
#endif

  // TODO: send settings
  // TODO: add a delay to wait for new settings to come back

#ifndef DEBUG
  Serial.println("Wifi disconnecting");
#endif
  client.disconnect();

  ESP.deepSleep(deepSleepDelay * 60 * 1000000);
  // delay(5 * 60 * 1000); // 5 minutes
  // client.loop();
}
  
