
#include "DHT.h"
#include <Wire.h>
#include <SPI.h>  //only if you'd want to use the SPI of the BMP
#include <Adafruit_Sensor.h>
#include <Adafruit_BMP280.h>
 
#define DHTPIN D5
#define DHTTYPE DHT22

#define BMP_SCK 13
#define BMP_MISO 12
#define BMP_MOSI 11 
#define BMP_CS 10

DHT dht2(DHTPIN, DHTTYPE);
Adafruit_BMP280 bmp280;

void setup() {
  pinMode(LED_BUILTIN, OUTPUT);
  
  Serial.begin(115200);
  Serial.println("Starting serial..");
  dht2.begin();

  if (!bmp280.begin()) {
    Serial.println("No BMP detected");
    delay(1000);
  }
}

void loop() {
  digitalWrite(LED_BUILTIN, LOW);
  float h = dht2.readHumidity();
  float t = dht2.readTemperature();  
  Serial.print("DHT22 - Temperature: ");
  Serial.print(t);
  Serial.print(" *C");
  Serial.print(" Humidity: ");
  Serial.print(h);
  Serial.println(" %");

  Serial.print("BMP280 - Temperature: ");
  Serial.print(bmp280.readTemperature());
  Serial.print(" *C Pressure: ");
  Serial.print(bmp280.readPressure());
  Serial.println(" Pa");
  delay(2000);

  digitalWrite(LED_BUILTIN, HIGH);  
  delay(2000);

}
  
