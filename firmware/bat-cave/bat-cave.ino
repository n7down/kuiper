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
  Serial.print("Message arrived: ");
  Serial.print(topic);
  Serial.print("- '");
  for (int i=0;i<length;i++) {
    Serial.print((char)payload[i]);
  }
  Serial.print("'");
  Serial.println();
}

void setupWifi(const char* ssid, const char* password)
{
  Serial.print("Wifi connecting...");
  WiFi.begin(ssid, password);
  
  // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print("...");
    // if (count >= 10 && (digitalRead(SLEEPPIN) ) == LOW) {
    //  Serial.println("Back to sleep, try again in 30 sec ");
    //  count = 0;
    //  ESP.deepSleep(30 * 1000000);
    // }
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
  macWithOutColons.toCharArray(mac, 12);

  Serial.print("MAC: ");
  Serial.println(mac);
}

void reconnect() {
  while (!client.connected()) {    
    Serial.print("Attempting MQTT connection (");
    Serial.print(mqtt_server);
    Serial.print(")... ");
    
    // Attempt to connect
    if (client.connect(mqtt_server)) {
      Serial.println("connected");
      
      String topic = "devices/";
      topic += mac;
      char deviceTopic[20];
      topic.toCharArray(deviceTopic, 20);
      client.subscribe(deviceTopic);
      Serial.print("Subscribed to topic: ");
      Serial.println(deviceTopic);
    } else {
      Serial.print("Failed: ");
      Serial.print(client.state());
      Serial.println(" - try again in 5 seconds");
      
      // Wait 5 seconds before retrying
      delay(5000);
    }
  }
  digitalWrite(LED_BUILTIN, HIGH); 
}

void setup() {
  pinMode(LED_BUILTIN, OUTPUT);
  
  digitalWrite(LED_BUILTIN, LOW);
  Serial.begin(115200);
  Serial.println("Starting serial..");
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

  // TODO: send settings
  // TODO: add a delay to wait for new settings to come back
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

  delay(2 * 60 * 1000); // wait 2 mins for response

  Serial.println("Wifi disconnecting");
  client.disconnect();

  ESP.deepSleep(deepSleepDelay * 60 * 1000000);
  // client.loop();
}
  
