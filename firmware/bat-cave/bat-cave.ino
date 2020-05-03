#include "config.h"
#include "DHT.h"

#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <Wire.h>
#include <SPI.h>
#include <PubSubClient.h>
#include <Adafruit_Sensor.h>
#include <ArduinoJson.h>
 
#define DHTPIN D5
#define SENDDATAPIN D6
#define DHTTYPE DHT22

struct {
  uint32_t deepSleepDelay;
} rtcData;

const char dht22Topic[] = "sensor/dht22";
const char statsTopic[] = "sensor/stats";
const char settingsTopic[] = "bc/settings";
const int subscriptionWaitTime = 2; // in min

char mac[13];

// default settings
int deepSleepDelay = 15; // in min

void callback(char* topic, byte* payload, unsigned int length);

DHT dht22(DHTPIN, DHTTYPE);
WiFiClient espClient;
PubSubClient client(mqtt_server, 1883, callback, espClient);

void callback(char* topic, byte* payload, unsigned int length) {
  StaticJsonDocument<20> doc;
  Serial.print("Message arrived: ");
  String buf;
  for (int i=0;i<length;i++) {
    buf += (char)payload[i];
  }
  Serial.println(buf);

  char json[20];
  buf.toCharArray(json, 20);
  DeserializationError error = deserializeJson(doc, json);
  if (error) {
    Serial.print("Error: deserializeJson() failed - ");
    Serial.println(error.c_str());
    return;
  }

  bool hasChanges = false;
  if (doc.containsKey("s")) {
    rtcData.deepSleepDelay = doc["s"];
    Serial.print("Deep sleep set to: ");
    Serial.println(deepSleepDelay);
    hasChanges = true;
  }
  
  if (hasChanges) {

    // write data to rtc data
    ESP.rtcUserMemoryWrite(0, (uint32_t*) &rtcData, sizeof(rtcData));
    Serial.println("Saved changes to RTC memeory");
  }
}

void setupWifi(const char* ssid, const char* password)
{
  Serial.print("Wifi connecting...");
  WiFi.begin(ssid, password);
  
  // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print("...");
  }
  Serial.print("Connected to ");
  Serial.println(ssid);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  String m = WiFi.macAddress();
  String macWithOutColons;
  for (int i = 0; i < m.length(); i++) {
    char currentChar = m[i];
    if (currentChar != ':') {
      macWithOutColons += currentChar;
    }
  }
  macWithOutColons.toLowerCase();
  macWithOutColons.toCharArray(mac, 13);
  Serial.print("MAC: ");
  Serial.println(mac);
}

void reconnect() {
  while (!client.connected()) {    
    Serial.print("Attempting MQTT connection (");
    Serial.print(mqtt_server);
    Serial.print(")... ");
    
    // attempt to connect
    if (client.connect(mqtt_server)) {
      Serial.println("connected");
    } else {
      Serial.print("Failed: ");
      Serial.print(client.state());
      Serial.println(" - try again in 5 seconds");
      
      // Wait 5 seconds before retrying
      delay(5000);
    }
  }
}

void setup() {
  Serial.begin(115200);
  Serial.println("Starting serial..");
  dht22.begin();
  setupWifi(ssid, password);
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
  Serial.print("Sent message: ");
  Serial.print(dht22Message);
  Serial.print(" - Result: ");
  Serial.println(result);
  
  char elapsedTimeString[40];
  sprintf(elapsedTimeString, "%u", elapsedTime);
  StaticJsonDocument<100> statsRoot;
  statsRoot["m"] = mac;
  statsRoot["v"] = String(battV);
  statsRoot["c"] = String(elapsedTimeString);

  char statsMessage[100];
  serializeJson(statsRoot, statsMessage); 
  
  result = client.publish(statsTopic, statsMessage);
  Serial.print("Sent message: ");
  Serial.print(statsMessage);
  Serial.print(" - Result: ");
  Serial.println(result);

  StaticJsonDocument<100> settingsRoot;
  settingsRoot["m"] = mac;
  settingsRoot["s"] = deepSleepDelay;

  char settingsMessage[100];
  serializeJson(settingsRoot, settingsMessage); 
  
  result = client.publish(settingsTopic, settingsMessage);
  Serial.print("Sent message: ");
  Serial.print(settingsMessage);
  Serial.print(" - Result: ");
  Serial.println(result);
 
  // FIXME: topic should be created in setup
  String topic = "devices/";
  topic += mac;
  char deviceTopic[21];
  topic.toCharArray(deviceTopic, 21);
  result = client.subscribe(deviceTopic, 1);
  Serial.print("Subscribed to topic: ");
  Serial.print(deviceTopic);
  Serial.print(" - Result: ");
  Serial.println(result);

  // run client.loop() for 2 mins to get messages
  Serial.println("Waiting for messages");
  unsigned long subscriptionStartTime = millis();
  while ((millis() - subscriptionStartTime) < 1000 * 60 * subscriptionWaitTime) { // 2 mins
    client.loop();
  }
  Serial.println("Finished waiting for messages");

  Serial.print("Unsubscribing to topic: ");
  Serial.print(deviceTopic);
  result = client.unsubscribe(deviceTopic);
  Serial.print(" - Result: ");
  Serial.println(result);

  Serial.println("Wifi disconnecting");
  client.disconnect();

  // read from rtc data and use the setting
  ESP.rtcUserMemoryRead(0, (uint32_t*) &rtcData, sizeof(rtcData));
  ESP.deepSleep(1000000 * 60 * rtcData.deepSleepDelay);
}

  
