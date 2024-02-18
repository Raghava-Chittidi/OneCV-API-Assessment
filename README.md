# OneCV API Assessment - Chittidi Raghava Naidu

## Deployed API Link

```
https://onecv-assessment.onrender.com
```

## Project Structure

```
.
├── Dockerfile
├── README.md
├── cmd
│   └── server
│       └── main.go
├── db_table_images
│   ├── registered.png
│   ├── students.png
│   └── teachers.png
├── docker-compose.yml
├── go.mod
├── go.sum
└── internal
    ├── api
    │   └── api.go
    ├── dataaccess
    │   ├── student.go
    │   └── teacher.go
    ├── database
    │   └── database.go
    ├── errors
    │   └── errors.go
    ├── handlers
    │   ├── commonstudents.go
    │   ├── notifications.go
    │   ├── register.go
    │   └── suspend.go
    ├── models
    │   ├── student.go
    │   └── teacher.go
    ├── router
    │   └── router.go
    ├── routes
    │   └── routes.go
    ├── tests
    │   ├── commonstudents_test.go
    │   ├── main_test.go
    │   ├── notifications_test.go
    │   ├── register_test.go
    │   ├── student_test.go
    │   ├── suspend_test.go
    │   ├── teacher_test.go
    │   └── util_test.go
    └── util
        ├── migration.go
        └── util.go
```

## Setting up the project

(Note: I am currently not using .gitignore for env files for easier setup of the project)

1. Clone the repository using this command

```bash
git clone https://github.com/Raghava-Chittidi/OneCV-API-Assessment.git
```

2. Ensure that the .env file contains DB_URL variable. If not, please add this into the .env file

```bash
DB_URL=postgres://mqjqotqy:Q9cQLXuVXQr7efbpWDAv9GuBVOLi9hDE@topsy.db.elephantsql.com/mqjqotqy
```

## Running the project locally

There are 2 ways to run this project. The first is using docker-compose while the second is manually starting the server. Visit any api endpoint at http://localhost:8000

### Using Docker-Compose

1. Go to the database.go file in the database folder in the internal directory. In the ConnectDB function, ensure that lines 7 and lines 15-19 are commented and save changes made.

2. Execute this command in the terminal. Ctrl + C to stop the container

```bash
docker-compose up --build
```

### Manually running the server

1. Go to the database.go file in the database folder in the internal directory. In the ConnectDB function, ensure that lines 7 and lines 15-19 are uncommented and save changes made.

2. Execute these commands in the terminal.

```bash
go mod download
cd cmd/server
go run main.go
```

## Running unit tests

1. Go to the database.go file in the database folder in the internal directory. In the ConnectDB function, ensure that lines 7 and lines 15-19 are uncommented and save changes made.

2. Change directory to the tests folder in the internal directory. Execute these commands in the terminal.

```bash
cd internal/tests
go test
```

2. It should take about 1 min - 1.5mins for all the tests to run.

## Current database tables and data

![students table](./db_table_images/students.png)
![teachers table](./db_table_images/teachers.png)
![teachers_registered_students table](./db_table_images/registered.png)

## API Endpoints

1. Endpoint: POST /api/register - Registers all the students whose emails are provided to the teacher. All the student emails provided must be valid and there must be an existing student in the database with this email. The teacher email provided must be valid and there must be an existing teacher in the database with this email

```
Request Body:

{
  "teacher": "teacherken@gmail.com"
  "students":
    [
      "studentemily@gmail.com",
      "studentmary@gmail.com"
    ]
}

Success response status: HTTP 204
Error response status: HTTP 400 - When invalid request body is provided.
Error response status: HTTP 404 - When an invalid email of student or teacher is provided
Error response status: HTTP 422 - When no student emails are provided

```

2. Endpoint: GET /api/commonstudents - Gets emails of common students under the teachers given. All the teacher emails provided must be valid and there must be an existing teacher in the database with this email.

```
/api/commonstudents?teacher=teacherken%40gmail.com&teacher=teacherjoe%40gmail.com

Success response:

{
  "students":
    [
      "studentjon@gmail.com",
      "studenthon@gmail.com"
    ]
}

Success Response status: HTTP 200
Error response status: HTTP 404 - When an invalid email of teacher is provided
Error response status: HTTP 422 - When no teacher emails are provided

```

3. Endpoint: POST /api/suspend - Suspends the student whose email is provided. The student email provided must be valid and there must be an existing student in the database with this email.

```
Request Body:

{
  "student": "studentmary@gmail.com"
}

Success response status: HTTP 204
Error response status: HTTP 400 - When invalid request body is provided.
Error response status: HTTP 404 - When an invalid email of student is provided

```

4. Endpoint: POST /api/retrievefornotifications - Returns a list conataining all the emails of the students who should be notified. All the student emails provided must be valid and there must be an existing student in the database with this email. The teacher email provided must be valid and there must be an existing teacher in the database with this email

```
Request Body:

{
  "teacher":  "teacherken@gmail.com",
  "notification": "Hello students! @studentemily@gmail.com @studentmary@gmail.com"
}

Success response:

{
  "recipients":
    [
        "studentjon@gmail.com",
        "studenthon@gmail.com",
        "studentemily@gmail.com",
        "studentmary@gmail.com"
    ]

}

Success response status: HTTP 200
Error response status: HTTP 400 - When invalid request body is provided.
Error response status: HTTP 404 - When an invalid email of student or teacher is provided

```
