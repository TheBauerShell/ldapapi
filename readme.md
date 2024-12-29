## Personen API
Eine einfache REST API zum Verwalten von Personendaten. Die API stellt CRUD-Operationen (Create, Read, Update, Delete) für Personendatensätze bereit.

### Endpunkte
### Alle Personen abrufen
 ```
 GET /persons
 ```
Liefert eine Liste aller Personen zurück.

#### Eine einzelne Person abrufen
```
GET /persons/{id}
```
Liefert die Personendaten für die angegebene ID.

#### Neue Person anlegen
```
POST /person
```
Legt einen neuen Personendatensatz an.

#### Request Body:
```json
{
    "id": "string",
    "first_name": "string", 
    "last_name": "string",
    "email": "string",
    "phone": "string"
}
```

#### Person aktualisieren
```
PUT /person/{id}
```
Aktualisiert die Daten einer bestehenden Person.

#### Request Body:
```json
{
    "first_name": "string",
    "last_name": "string", 
    "email": "string",
    "phone": "string"
}
```

#### Person löschen
```
DELETE /person/{id}
```
Löscht den Personendatensatz mit der angegebenen ID.

#### Alle Personen zurücksetzen
```
POST /reset-persons
```
Setzt alle Personendaten auf den Ausgangszustand zurück.

### Installation
```sh
# Repository klonen
git clone https://github.com/user/repo.git

# Abhängigkeiten installieren  
go mod download

# Server starten
go run cmd/main.go
```
Der Server läuft standardmäßig auf Port 8080.

### Entwicklung
Die API ist mit Go und dem gorilla/mux Router implementiert. Die Personendaten werden in einer JSON-Datei gespeichert.

Lizenz
MIT

### Hier sind die verschiedenen Möglichkeiten zum Testen der API:

#### 1. Mit cURL
```sh 
# Alle Personen abrufen
curl -X GET http://localhost:8080/persons

# Eine einzelne Person abrufen
curl -X GET http://localhost:8080/persons/1

# Neue Person anlegen
curl -X POST http://localhost:8080/person \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Max",
    "last_name": "Mustermann", 
    "email": "max@example.com",
    "phone": "+491234567890"
  }'

# Person aktualisieren
curl -X PUT http://localhost:8080/person/1 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Maximilian",
    "last_name": "Mustermann",
    "email": "max@example.com", 
    "phone": "+491234567890"
  }'

# Person löschen
curl -X DELETE http://localhost:8080/person/1

# Alle Personen zurücksetzen
curl -X POST http://localhost:8080/reset-persons
```
#### 2. Mit Postman
Öffne Postman
Erstelle eine neue Collection für die API
Füge Requests für die verschiedenen Endpunkte hinzu:
GET http://localhost:8080/persons
GET http://localhost:8080/persons/{id}
POST http://localhost:8080/person
PUT http://localhost:8080/person/{id}
DELETE http://localhost:8080/person/{id}
POST http://localhost:8080/reset-persons
Füge bei POST/PUT Requests den entsprechenden JSON-Body hinzu
Klicke auf "Send" um die Requests auszuführen
#### 3. Mit dem Browser
Für einfache GET Requests kannst du auch direkt den Browser nutzen:
```
http://localhost:8080/persons
http://localhost:8080/persons/1
```
#### 4. PowerShell Testing mit Invoke-RestMethod
```powershell
# Alle Personen abrufen
Invoke-RestMethod -Uri "http://localhost:8080/persons" -Method Get

# Eine einzelne Person abrufen
Invoke-RestMethod -Uri "http://localhost:8080/persons/1" -Method Get

# Neue Person anlegen
$body = @{
    first_name = "Max"
    last_name = "Mustermann"
    email = "max@example.com"
    phone = "+491234567890"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/person" `
    -Method Post `
    -Body $body `
    -ContentType "application/json"

# Person aktualisieren
$updateBody = @{
    first_name = "Maximilian"
    last_name = "Mustermann"
    email = "max@example.com"
    phone = "+491234567890"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/person/1" `
    -Method Put `
    -Body $updateBody `
    -ContentType "application/json"

# Person löschen
Invoke-RestMethod -Uri "http://localhost:8080/person/1" -Method Delete

# Alle Personen zurücksetzen
Invoke-RestMethod -Uri "http://localhost:8080/reset-persons" -Method Post
```

Stelle sicher, dass der Server läuft (go run cmd/main.go) bevor du die Tests ausführst.