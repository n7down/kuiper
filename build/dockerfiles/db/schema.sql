CREATE DATABASE device_settings;
use device_settings;

CREATE TABLE bat_cave_settings(device_id VARCHAR(12) NOT NULL, deep_sleep_delay INT NOT NULL, created_at TIMESTAMP, updated_at TIMESTAMP, deleted_at TIMESTAMP, PRIMARY KEY (device_id));
