-- +migrate Up
ALTER TABLE articles ADD user VARCHAR(255) NOT NULL AFTER body;
