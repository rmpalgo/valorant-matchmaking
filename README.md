# valorant-matchmaking

# Go Concurrency: Valorant-Inspired Matchmaking

This repository is a practice project demonstrating the use of Go's concurrency features in a FPS game matchmaking scenario.

## Overview

The program simulates a matchmaking system where players queue up for three different game modes: Competitive, Unrated, and Deathmatch. Each game mode requires a different number of players to start a match. The program aims to showcase the use of Go's goroutines, channels, and mutexes in handling this scenario.

Here are the specifics for each game mode:

- Competitive: 10 players required
- Unrated: 10 players required
- Deathmatch: 20 players required

Players are simulated and assigned to queues, after which the matchmaker attempts to create matches with the correct number of players for each game type.

## Running the Project

To run the project:

```bash
go run main.go
