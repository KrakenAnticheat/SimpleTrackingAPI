# SimpleTrackingAPI
 
> Made By Gustavo J.
 
A Simple Tracking API to register and check activity events. Basically is used to save actions that happens on a **application**
 
---
 
![Tests Passing](https://img.shields.io/badge/tests-passing-brightgreen) ![Version 1.0.0](https://img.shields.io/badge/version-1.0.0-orange) ![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white) ![Beta](https://img.shields.io/badge/stability-beta-yellow)
 
---
 
## Usage
 
`go run cmd/server/main.go` to start the **api** /
`GET http://localhost:8080/events`
 
```json
curl -X POST http://localhost:8080/track \
-H "Content-Type: application/json" \
-d '{
"user_id":"123",
"event":"page_view",
"metadata":{"page":"/home"}
}'
```
 
---
 
## Tech Stack
 
| Technology | Description |
| --- | --- |
| **GoLang** | Simple Go syntax |
| **DB** | Fake Database |
 
---
 
## API Reference
 
Base URL: `GET http://localhost:8080/`
 
### `GET http://localhost:8080/events`
 
View events

---
README made with https://www.docora.co/ (By Gustavo J.)
