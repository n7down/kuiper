
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

DHT dht22(DHTPIN, DHTTYPE);
Adafruit_BMP280 bmp280;
WiFiClient espClient;
PubSubClient client(espClient); 

const char ssid[] = "";
const char password[] = "";
const char mqtt_server[] = "";
const char type[] = "bc";
const char id[] = "1";
const char humidityTopic[] = "sensor/humidity";
const char tempTopic[] = "sensor/temp";
const char pressureTopic[] = "sensor/pressure";
const char voltageTopic[] = "sensor/voltage";

const int minutes = 1;
const int readDelay = 1000 * 60 * minutes;

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

  float h = dht22.readHumidity();
  float t = dht22.readTemperature();
  float tt = bmp280.readTemperature();
  float p = bmp280.readPressure();  

  unsigned int batt = analogRead(A0);
  double battV = batt * (4.2 / 1023);

  char label[10];
  strcpy(label, type);
  strcat(label, label);

  StaticJsonDocument<50> humidityRoot;
  humidityRoot["id"] = label;
  humidityRoot["dht22hum"] = String(h); // %

  char humidityMessage[50];
  serializeJson(humidityRoot, humidityMessage); 
  
  int result = client.publish(humidityTopic, humidityMessage);
  Serial.print("Sent message: ");
  Serial.print(humidityMessage);
  Serial.print(" - Result: ");
  Serial.println(result);
  
  StaticJsonDocument<100> tempRoot;
  tempRoot["id"] = label;
  tempRoot["dht22temp"] = String(t); // in *C
  tempRoot["bmp280temp"] = String(tt); // in *C

  char tempMessage[100];
  serializeJson(tempRoot, tempMessage); 
  
  result = client.publish(tempTopic, tempMessage);
  Serial.print("Sent message: ");
  Serial.print(tempMessage);
  Serial.print(" - Result: ");
  Serial.println(result);

  StaticJsonDocument<50> pressureRoot;
  pressureRoot["id"] = label;
  pressureRoot["bmp280pres"] = String(p); // in Pa

  char pressureMessage[50];
  serializeJson(pressureRoot, pressureMessage); 
  
  result = client.publish(pressureTopic, pressureMessage);
  Serial.print("Sent message: ");
  Serial.print(pressureMessage);
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

  Serial.println("disconnecting");
  client.disconnect();
 
  delay(readDelay);

  digitalWrite(LED_BUILTIN, LOW);  
  delay(readDelay);

}
  
