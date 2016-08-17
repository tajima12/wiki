-- +migrate Up
ALTER TABLE comments ADD user VARCHAR(255) NOT NULL AFTER body;
