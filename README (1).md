
## UDC MANAGEMENT SYSTEM USING GO AND GoFr

The Go CRUD Application is a simple yet robust project built using the GoFr and MySQL database. It facilitates basic CRUD operations (Create, Read, Update, Delete) for managing UDC records. Fork, customize, and extend to meet your specific needs with the flexibility of Go and GoFr.

# Prerequisites
- [Go](https://go.dev/)
- [MySQL](https://www.mysql.com/)
- [Postman](https://www.postman.com/)







## Getting Started

  Run this command line in your terminal
    
    go get gofr.dev

To deploy this project-

1 Clone the repositery

```bash
  Git clone https://github.com/RajHarsh01/GoFr-project-for-Zopsmart
```
2 Download and Verify repo

```bash
  Go mod download
  cat go.sum
```
3 Run project

```bash
  go run main.go
```
4 Open server

```bash
  http://localhost:8080/students
```

# API Endpoints

- GET students
```bash
  get/students
```
![ss](https://github.com/RajHarsh01/GoFr-project-for-Zopsmart/assets/80113516/ec55e5b6-90a2-4d15-9f63-081cc88a5107)

- CREATE students
```bash
  post/students
```
![ss](https://github.com/RajHarsh01/GoFr-project-for-Zopsmart/assets/80113516/72e9d348-32fa-47cf-9116-f0fe2142faaa)

- UPDATE students
```bash
  put/studentID
```
![ss](https://github.com/RajHarsh01/GoFr-project-for-Zopsmart/assets/80113516/7ac741a5-e2dc-4ef7-b521-d2a90dc5fe8b)

- DELETE student
```bash
  delete/studentID
```
![ss](https://github.com/RajHarsh01/GoFr-project-for-Zopsmart/assets/80113516/7df708f5-e6f4-4712-a875-734673e5bc17)

# DATABASE SCHEMA

### Docker image
  copy this command and run in your cmd
  
     docker run --name student-project-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=student_project_db -p 3306:3306 -d mysql:8.0.30

![ss](https://github.com/RajHarsh01/GoFr-project-for-Zopsmart/assets/80113516/ab0fd383-c8d4-4840-9360-e91d607e5e3c)

# UML Diagram

![ss](https://github.com/RajHarsh01/GoFr-project-for-Zopsmart/assets/80113516/c39b1c65-bc94-4f88-8442-3e5dcbddb91b))

# SEQUENCE Diagram

![ss](https://github.com/RajHarsh01/GoFr-project-for-Zopsmart/assets/80113516/0d8c9eec-6883-4913-bbbc-2eede5f33e56)