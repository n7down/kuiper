#include "DHT.h"
#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <Wire.h>
#include <SPI.h>
#include <PubSubClient.h>
#include <Adafruit_Sensor.h>
#include <Adafruit_BMP280.h>
#include <ArduinoJson.h>
 
#define DHTPIN D5
#define DHTTYPE DHT22

#define BMP_SCK 13
#define BMP_MISO 12
#define BMP_MOSI 11 
#define BMP_CS 10

const char ssid[] = "";
const char password[] = "";
const char mqtt_server[] = "";
const char type[] = "bc";
const char id[] = "1";
const char dht22Topic[] = "sensor/dht22";
const char bmp280Topic[] = "sensor/bmp280";
const char voltageTopic[] = "sensor/voltage";
const char timeTopic[] = "time/utc";
const int hours = 1;
// bool receivedCurrentTime = false;

DHT dht22(DHTPIN, DHTTYPE);
Adafruit_BMP280 bmp280;
WiFiClient espClient;
PubSubClient client(espClient); 

// void callback(char* topic, byte* payload, unsigned int length) {
// receivedCurrentTime = true;
// TODO: get the current timestamp
// }

void setupWifi(const char* ssid, const char* password)
{
  Serial.println("WiFi connecting...");
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
}

void setup() {
  pinMode(LED_BUILTIN, OUTPUT);
  
  Serial.begin(115200);
  Serial.println("Starting serial..");
  dht22.begin();

  if (!bmp280.begin()) {
    Serial.println("No BMP detected");
    delay(1000);
  }
  setupWifi(ssid, password);
  client.setServer(mqtt_server, 1883);
  // client.setCallback(callback);

  char subscriptionTopic[20];
  strcpy(subscriptionTopic, timeTopic);
  strcat(subscriptionTopic, "/");  
  strcat(subscriptionTopic, type);
  strcat(subscriptionTopic, id);
  client.subscribe(subscriptionTopic);
}

void loop() {

  digitalWrite(LED_BUILTIN, HIGH);
  while (!client.connected()) {
		Serial.print("Attempting MQTT connection (");
		Serial.print(mqtt_server);
		Serial.print(")... ");
		// Create a random client ID
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

  char label[10];
  strcpy(label, type);
  strcat(label, id);
    
  // if (receivedCurrentTime) {
	  float h = dht22.readHumidity();
	  float t = dht22.readTemperature();
	  float tt = bmp280.readTemperature();
	  float p = bmp280.readPressure();  

	  unsigned int batt = analogRead(A0);
	  double battV = batt * (4.2 / 1023);

	  StaticJsonDocument<100> dht22Root;
	  dht22Root["id"] = label;
	  dht22Root["humidity"] = String(h); // % 
	  dht22Root["temp"] = String(t); // %

	  char dht22Message[100];
	  serializeJson(dht22Root, dht22Message); 
	  
	  int result = client.publish(dht22Topic, dht22Message);
	  Serial.print("Sent message: ");
	  Serial.print(dht22Message);
	  Serial.print(" - Result: ");
	  Serial.println(result);
	  
	  StaticJsonDocument<100> bmp280Root;
	  bmp280Root["id"] = label;
	  bmp280Root["temp"] = String(t); // in *C
    bmp280Root["pres"] = String(p); // in Pa

	  char bmp280Message[100];
	  serializeJson(bmp280Root, bmp280Message); 
	  
	  result = client.publish(bmp280Topic, bmp280Message);
	  Serial.print("Sent message: ");
	  Serial.print(bmp280Message);
	  Serial.print(" - Result: ");
	  Serial.println(result);

	  StaticJsonDocument<50> voltageRoot;
	  voltageRoot["id"] = label;
	  voltageRoot["voltage"] = String(battV);

	  char voltageMessage[50];
	  serializeJson(voltageRoot, voltageMessage); 
	  
	  result = client.publish(voltageTopic, voltageMessage);
	  Serial.print("Sent message: ");
	  Serial.print(voltageMessage);
	  Serial.print(" - Result: ");
	  Serial.println(result);

  // } else {    
  //  StaticJsonDocument<100> timeRoot;
  //  timeRoot["deviceName"] = label;
  
  //  char timeMessage[100];
  //  serializeJson(timeRoot, timeMessage); 
    
  //  int result = client.publish(timeTopic, timeMessage);
    
  //  Serial.print("Sent message: ");
  //  Serial.print(timeMessage);
  //  Serial.print(" - Result: ");
  //  Serial.println(result);
  // }
  
  Serial.println("disconnecting");
  client.disconnect();
 
  digitalWrite(LED_BUILTIN, LOW); 
  
  ESP.deepSleep(hours * 60 * 60 * 1000000);
}
  
