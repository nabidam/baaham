# Watch Together – Real-Time Media & Whiteboard Platform

A real-time **watch-together / listen-together** platform with chat and shared whiteboard, designed for **two-user rooms**.  
Users can watch movies, listen to music, chat, and draw together in a synchronized room.

---

## Features

- 🔐 **Authentication**

  - Pre-registered users
  - JWT-based authentication

- 🏠 **Rooms**

  - Rooms created by admin
  - Exactly **2 users per room**
  - Isolated real-time environment per room

- 💬 **Real-Time Chat**

  - WebSocket-based chat
  - Room-scoped messages

- 🎬 **Watch Together**

  - Play movies pre-uploaded to the server
  - Synchronized play / pause / seek
  - Server-authoritative state

- 🎵 **Listen Together**

  - Play songs pre-uploaded to the server
  - Synchronized playback
  - Playlist support (private & public)

- 🧾 **Playlists**

  - Create playlists from existing songs
  - Private or public visibility

- 🎨 **Shared Whiteboard**
  - Real-time collaborative canvas
  - Draw, erase, clear
  - Changes synced via WebSocket events

---

## Tech Stack

### Backend

- Go
- HTTP API (REST)
- WebSockets (real-time sync)
- JWT authentication
- PostgreSQL (users, rooms, playlists)
- File storage for media

### Frontend

- Vite
- React + TypeScript
- Tailwind CSS
- WebSocket client
- HTML5 Audio / Video
- Canvas API (whiteboard)

---

## Architecture Overview

- **Monorepo**

  - `backend/` → Go server
  - `frontend/` → React app
  - `shared/` → WebSocket event contracts

- **Server-authoritative**

  - Playback state lives on the server
  - Clients only emit intents (play, pause, seek)
  - Server broadcasts synchronized state

- **Room-scoped WebSockets**
  - One WebSocket connection per room per user
  - All events are isolated to the room

---

## Room Rules

- Each room has **exactly 2 users**
- Users must be added to a room by admin
- Only room members can connect to its WebSocket
- All real-time events are broadcast to both users

---

## WebSocket Event Types (High Level)

- `CHAT_MESSAGE`
- `MEDIA_PLAY`
- `MEDIA_PAUSE`
- `MEDIA_SEEK`
- `SONG_CHANGE`
- `WHITEBOARD_DRAW`
- `WHITEBOARD_CLEAR`
- `USER_JOIN`
- `USER_LEAVE`

Event schemas are defined in `shared/`.

---

## Media Handling

- Movies and songs are uploaded to the server in advance
- Clients request available media via API
- Playback sync is done via WebSocket events
- Media files are streamed over HTTP

---

## Development Setup

```bash
# start backend
cd backend
go run cmd/server/main.go

# start frontend
cd frontend
npm install
npm run dev
```
