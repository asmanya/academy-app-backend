# Academy App System

A secured Go HTTPS REST API for managing an academy — students, teachers, and admin users — with JWT authentication, role-based access control, and multiple security layers.

---

## Tech Stack

- **Language**: Go 1.24.0
- **Database**: MySQL
- **Auth**: JWT (HMAC-SHA256) via secure httpOnly cookies
- **Password Hashing**: Argon2id
- **XSS Protection**: bluemonday
- **TLS**: Embedded certs via HTTPS

---

## Project Structure

```
academy-app-system/
├── cmd/api/
│   └── server.go               # Entry point
├── internal/
│   ├── api/
│   │   ├── handlers/           # HTTP handlers
│   │   │   ├── students.go
│   │   │   ├── teachers.go
│   │   │   ├── execs.go
│   │   │   └── helpers.go
│   │   ├── middlewares/        # Middleware chain
│   │   │   ├── security_headers.go
│   │   │   ├── cors.go
│   │   │   ├── compression.go
│   │   │   ├── jwt_middleware.go
│   │   │   ├── rate_limiter.go
│   │   │   ├── response_time.go
│   │   │   ├── hpp.go
│   │   │   ├── sanitize.go
│   │   │   └── exclude_routes.go
│   │   └── router/             # Route registration
│   │       ├── router.go
│   │       ├── students_router.go
│   │       ├── teachers_router.go
│   │       └── execs_router.go
│   ├── models/                 # Data structs
│   │   ├── student.go
│   │   ├── teacher.go
│   │   └── exec.go
│   └── repository/
│       └── sqlconnect/         # Database layer
│           ├── sqlconfig.go
│           ├── students_crud.go
│           ├── teachers_crud.go
│           └── execs_crud.go
└── pkg/utils/                  # Shared utilities
    ├── jwt.go
    ├── password.go
    ├── authorize_user.go
    ├── error_handler.go
    ├── pagination.go
    ├── database_utils.go
    └── middlewaresutil.go
```

---

## API Endpoints

### Students `/students`

| Method | Path | Description |
|--------|------|-------------|
| GET | `/students` | List students (paginated, filtered, sorted) |
| GET | `/students/{id}` | Get one student by ID |
| POST | `/students` | Bulk insert students |
| PUT | `/students/{id}` | Full update of a student |
| PATCH | `/students` | Bulk partial update |
| PATCH | `/students/{id}` | Partial update of one student |
| DELETE | `/students/{id}` | Delete one student |
| DELETE | `/students` | Bulk delete students |

### Teachers `/teachers`

| Method | Path | Required Role | Description |
|--------|------|---------------|-------------|
| GET | `/teachers` | admin / manager / exec | List teachers |
| GET | `/teachers/{id}` | admin / manager / exec | Get one teacher |
| GET | `/teachers/{id}/students` | admin / manager / exec | Students in teacher's class |
| GET | `/teachers/{id}/studentcount` | admin / manager / exec | Count of students for teacher |
| POST | `/teachers` | manager / exec | Bulk insert teachers |
| PUT | `/teachers/{id}` | exec | Full update of a teacher |
| PATCH | `/teachers` | exec | Bulk partial update |
| PATCH | `/teachers/{id}` | exec | Partial update of one teacher |
| DELETE | `/teachers/{id}` | admin / manager | Delete one teacher |
| DELETE | `/teachers` | admin / manager | Bulk delete teachers |

### Execs (Admins) `/execs`

| Method | Path | Required Role | Description |
|--------|------|---------------|-------------|
| GET | `/execs` | — | List all admins |
| GET | `/execs/{id}` | — | Get one admin |
| POST | `/execs` | admin / manager | Create admins |
| PATCH | `/execs` | admin / manager | Bulk partial update |
| PATCH | `/execs/{id}` | admin / manager | Partial update of one admin |
| DELETE | `/execs/{id}` | admin | Delete one admin |
| POST | `/execs/login` | — | Login — issues JWT cookie |
| POST | `/execs/logout` | — | Logout — clears JWT cookie |
| POST | `/execs/{id}/updatepassword` | — | Update password |
| POST | `/execs/forgotpassword` | — | Request password reset email |
| POST | `/execs/resetpassword/reset/{resetcode}` | — | Reset password via token |

---

## File-by-File Reference

### Entry Point

**`cmd/api/server.go`**
- `loadEnvFromEmbeddedFile()` — reads embedded `.env` and sets environment variables
- `main()` — initializes HTTPS server, applies middleware chain (security headers, compression, HPP, XSS, JWT, rate limiter, CORS)

---

### Models

**`internal/models/student.go`**
- `Student` struct — ID, FirstName, LastName, Email, Class

**`internal/models/teacher.go`**
- `Teacher` struct — ID, FirstName, LastName, Email, Class, Subject

**`internal/models/exec.go`**
- `Exec` struct — ID, FirstName, LastName, Email, Username, hashed Password, PasswordChangedAt, PasswordResetToken, PasswordTokenExpires, InactiveStatus, Role
- `UpdatePasswordRequest` — request payload for password change
- `UpdatePasswordResponse` — response with new token and `password_updated` flag

---

### Handlers

**`internal/api/handlers/students.go`**
- `GetStudentsHandler()` — paginated, filtered, sortable list of students
- `GetOneStudentHandler()` — single student by ID
- `AddStudentHandler()` — bulk insert with field validation
- `UpdateStudentHandler()` — full update of all fields
- `PatchStudentsHandler()` — bulk partial updates in a transaction
- `PatchOneStudentHandler()` — single partial update using reflection
- `DeleteOneStudentHandler()` — delete by ID
- `DeleteStudentsHandler()` — bulk delete in a transaction

**`internal/api/handlers/teachers.go`**
Same CRUD pattern as students, with role authorization on every route (see endpoints table above), plus:
- `GetStudentsByTeacherId()` — students belonging to a teacher's class
- `GetStudentCountByTeacherId()` — count of students in a teacher's class

**`internal/api/handlers/execs.go`**
- `GetExecsHandler()` — paginated list of admins
- `GetOneExecHandler()` — single admin by ID
- `AddExecsHandler()` — bulk insert with hashed passwords
- `PatchExecsHandler()` / `PatchOneExecHandler()` — partial updates
- `DeleteOneExecHandler()` — delete by ID
- `LoginHandler()` — validates credentials, issues JWT in secure httpOnly cookie
- `LogoutHandler()` — clears the JWT cookie
- `UpdatePasswordHandler()` — verifies old password, hashes new one, issues fresh JWT
- `ForgotPasswordHandler()` — generates SHA256 reset token, stores hash in DB with expiry, sends email
- `ResetPasswordHandler()` — validates token + expiry, hashes new password, clears reset fields

**`internal/api/handlers/helpers.go`**
- `CheckBlankFields()` — reflects over a struct and returns error if any string field is empty
- `GetFieldNames()` — extracts JSON tag names from struct fields (strips `,omitempty`)

---

### Router

**`internal/api/router/router.go`**
- `MainRouter()` — composes students, teachers, and execs sub-routers into one `ServeMux`

**`internal/api/router/students_router.go`**
- `studentsRouter()` — registers all `/students` routes

**`internal/api/router/teachers_router.go`**
- `teachersRouter()` — registers all `/teachers` routes including nested `/students` endpoints

**`internal/api/router/execs_router.go`**
- `execsRouter()` — registers all `/execs` and auth routes

---

### Middlewares

**`internal/api/middlewares/security_headers.go`**
- `SecurityHeader()` — sets HSTS, CSP, X-Frame-Options, X-Content-Type-Options, cache-control, and permissions policy headers

**`internal/api/middlewares/cors.go`**
- `Cors()` — validates Origin against an allowlist, sets CORS response headers
- `isOriginAllowed()` — checks if origin is in the allowed list

**`internal/api/middlewares/compression.go`**
- `Compression()` — wraps response writer with gzip if client supports it
- `gzipResponseWriter.Write()` — writes gzip-compressed bytes to the client

**`internal/api/middlewares/jwt_middleware.go`**
- `JWTMiddleware()` — extracts Bearer cookie, validates JWT signature, injects role/username/userId into request context

**`internal/api/middlewares/rate_limiter.go`**
- `NewRateLimiter()` — creates rate limiter with a visitor IP tracking map
- `resetVisitorCount()` — goroutine that clears visitor counts at a configured interval
- `RLMiddleware()` — returns 429 if a visitor IP exceeds the request limit (default: 5/min)

**`internal/api/middlewares/response_time.go`**
- `ResponseTimeMiddleware()` — measures request duration, adds `X-Response-Time` header, logs method/URL/status/duration

**`internal/api/middlewares/hpp.go`** — HTTP Parameter Pollution protection
- `Hpp()` — returns middleware configured with a param whitelist
- `filterBodyParams()` — removes non-whitelisted or duplicate form body params
- `filterQueryParams()` — same for query string params
- `isWhiteListed()` — checks if a param name is in the whitelist

**`internal/api/middlewares/sanitize.go`** — XSS protection
- `XSSMiddlware()` — sanitizes URL path, query params, and JSON body
- `clean()` — recursively handles maps, arrays, and strings
- `sanitizeValue()` — recursively sanitizes nested structures
- `sanitizeString()` — uses bluemonday to strip dangerous HTML

**`internal/api/middlewares/exclude_routes.go`**
- `MiddlewaresExcludePaths()` — skips a middleware for specified paths (e.g., login and password reset bypass JWT check)

---

### Utilities

**`pkg/utils/jwt.go`**
- `SignToken()` — creates and signs a JWT with uid, username, role, and expiry (HMAC-SHA256)

**`pkg/utils/password.go`**
- `HashPassword()` — generates a 16-byte random salt, hashes with argon2id, returns `"salt.hash"` as base64
- `VerifyPassword()` — splits stored hash, re-hashes the input, compares in constant time

**`pkg/utils/authorize_user.go`**
- `ContextKey` type — custom string type for request context keys
- `AuthorizeUser()` — checks request role against an allowed roles list, returns error if unauthorized

**`pkg/utils/error_handler.go`**
- `ErrorHandler()` — logs error with timestamp and file location, returns a formatted error string

**`pkg/utils/pagination.go`**
- `GetPaginationParams()` — parses `page` and `limit` from query params (defaults: page=1, limit=10)

**`pkg/utils/database_utils.go`**
- `GenerateInsertQuery()` — builds INSERT SQL from a struct's `db` tags using reflection (skips `id`)
- `GetStructValues()` — extracts field values in the same order as the generated query
- `AddSorting()` — parses `sortby=field:asc/desc` query param, validates, appends ORDER BY clause
- `AddFilters()` — appends WHERE clauses for name/email/class/subject query filters
- `isValidSortField()` — whitelist check for allowed sort fields
- `isValidSortOrder()` — ensures direction is `asc` or `desc`

**`pkg/utils/middlewaresutil.go`**
- `Middleware` type — function type alias for middleware
- `ApplyMiddlewares()` — wraps a handler with an ordered middleware chain

---

### Database Layer

**`internal/repository/sqlconnect/sqlconfig.go`**
- `ConnectDb()` — reads DB credentials from env vars, builds MySQL DSN, returns `*sql.DB`

**`internal/repository/sqlconnect/students_crud.go`**
- `GetStudentsDBHandler()` — paginated + filtered student query, returns list and total count
- `GetStudentByID()` — single student query by ID
- `AddStudentsDBHandler()` — batch insert with prepared statement, handles FK errors on class field
- `UpdateStudent()` — full row update (all fields)
- `PatchStudents()` — batch partial updates using transaction and reflection
- `PatchOneStudent()` — single partial update using reflection
- `DeleteOneStudent()` — delete by ID, checks rows affected
- `DeleteStudents()` — batch delete in a transaction

**`internal/repository/sqlconnect/teachers_crud.go`**
Same pattern as students, plus:
- `GetStudentsByTeacherIdFromDb()` — joins teachers → students by class
- `GetStudentCountByTeacherIdFromDb()` — count of students for a teacher's class

**`internal/repository/sqlconnect/execs_crud.go`**
- `GetExecsDBHandler()` — lists execs (password field excluded)
- `GetExecByID()` — single exec (password excluded)
- `AddExecsDBHandler()` — batch insert with password hashing
- `PatchExecs()` / `PatchOneExec()` — partial updates via reflection
- `DeleteOneExec()` — delete by ID
- `GetUserByUsername()` — fetches full exec record including password hash (used at login)
- `UpdatePasswordInDb()` — verifies old password, stores new hash + timestamp, issues new JWT
- `ForgotPasswordDbHandler()` — generates 32-byte random token, stores SHA256 hash + expiry, sends reset email
- `ResetPasswordDbHandler()` — validates token + expiry, hashes new password, clears reset fields from DB

---

## Security Features

| Feature | Implementation |
|---------|---------------|
| Authentication | JWT in secure httpOnly cookie |
| Password hashing | Argon2id with random salt |
| XSS protection | bluemonday HTML sanitizer on all inputs |
| HTTP Parameter Pollution | Custom HPP middleware with whitelist |
| Rate limiting | Per-IP request counter with configurable window |
| CORS | Origin whitelist validation |
| Security headers | HSTS, CSP, X-Frame-Options, nosniff, etc. |
| TLS | Embedded certificate and key via HTTPS |
| Authorization | Role-based access control (admin / manager / exec) |
| SQL safety | Prepared statements and transactions throughout |
