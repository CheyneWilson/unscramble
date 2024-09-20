# Unscramble

This is (to be) a scrabble-like game.


# Architecture

Outline of the intended architecture. Note, the Bag API and CLI ser

## Bag Library

The `Bag` library represents a bag of letter tiles. It comes with a default tile-set. 
It also supports choosing a differing an initial tile-set, drawing and removing or returning tiles, adding tiles, 
counting the remaining tiles, and listing the contents. Tiles can be drawn at random, or specific tiles can be sought.

### Bag CLI

The `BagCLI` is a command line program which contains a `Bag`

### Bag API

The `Bag API` is a RESTful API (OpenAPI/Swagger) of a `Bag` server. 


