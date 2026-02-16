# Glofox
Glofox is a SaaS platform for boutiques, studios, and gyms that allows business owners to manage their courses, classes, members, memberships, and related operations.

- boutiques, studios, and gyms owners to create classes  
- Members to book classes  
- In-memory storage

## To Run
go mod tidy
go run cmd/server/main.go

📌 API Endpoints

1️⃣ Create Class
POST /classes
Request Body:
{
  "name": "Pilates",
  "start_date": "2026-12-01",
  "end_date": "2026-12-20",
  "capacity": 20
}
Curl:
curl -X POST http://localhost:8080/classes \
-H "Content-Type: application/json" \
-d '{"name":"Pilates","start_date":"2026-12-01","end_date":"2026-12-20","capacity":20}'

2️⃣ Book a Class
POST /bookings
Request Body:
{
  "name": "Ashish",
  "date": "2026-12-16",
  "class_name": "Pilates"
}
Curl:
curl -X POST http://localhost:8080/bookings \
-H "Content-Type: application/json" \
-d '{"name":"Ashish","date":"2026-12-16","class_name":"Pilates"}'
