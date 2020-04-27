#include "config.h"

#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <PubSubClient.h>

char mac[13];

void callback(char* topic, byte* payload, unsigned int length);

WiFiClient espClient;
PubSubClient client(mqtt_server, 1883, callback, espClient);

void callback(char* topic, byte* payload, unsigned int length) {
  Serial.print("Message arrived: ");
  Serial.print(topic);
  Serial.print(" - '");
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
    
    // Attempt to connect
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

  setupWifi(ssid, password);
  
  //client.setServer(mqtt_server, 1883);
  //client.setCallback(callback);

  if (!client.connected()) {
    reconnect();
  }

  String topic = "devices/";
  topic += mac;
  char deviceTopic[21];
  topic.toCharArray(deviceTopic, 21);
  int result = client.subscribe(deviceTopic, 1);
  Serial.print("Subscribed to topic: ");
  Serial.print(deviceTopic);
  Serial.print(" - Result: ");
  Serial.println(result);
}

void loop() {
  client.loop();
}
  
