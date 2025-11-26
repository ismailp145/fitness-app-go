# Go Fitness API - Clean Architecture Template

**Production-Ready Backend with Clean Architecture**  
[![Go 1.21+](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![JWT Auth](https://img.shields.io/badge/JWT-Auth-000000?style=flat&logo=json-web-tokens)](https://jwt.io/)
[![Clean Architecture](https://img.shields.io/badge/Clean-Architecture-blue)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
[![REST API](https://img.shields.io/badge/REST-API-green)](https://restfulapi.net/)

---

## 📋 Overview

A complete fitness tracking backend built with Go and Clean Architecture principles.  
Perfect starting point for your fitness mobile app with React Native (Expo).

---

### 🏗️ Architecture

Clean Architecture with clear separation of concerns, making the codebase **testable, maintainable, and scalable**.

- **Domain Layer** (entities)
- **Use Case Layer** (business logic)
- **Repository Layer** (data access)
- **Delivery Layer** (HTTP/gRPC)

---

### ✨ Features

- User authentication (JWT)
- Workout logging & tracking
- Exercise library management
- Progress analytics
- Social feed ready
- ML recommendation ready

---

### 🛠️ Tech Stack

- **Go 1.21+** (latest features)
- **PostgreSQL** (robust RDBMS)
- **Gin/Fiber** (HTTP framework)
- **JWT** (authentication)
- **Docker** (containerization)
- **Migrate** (DB migrations)

---

### 🎨 Design Principles

- Dependency Injection
- Interface-based design
- Repository pattern
- SOLID principles
- Testable architecture
- Context-aware operations

---

## 📁 Project Structure

Standard Go project layout following Clean Architecture principles:

```
fitness-api/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── domain/                  # Domain Layer
│   │   ├── user.go
│   │   ├── workout.go
│   │   ├── exercise.go
│   │   ├── workout_log.go
│   │   └── errors.go
│   ├── usecase/                 # Business Logic
│   │   ├── user_usecase.go
│   │   ├── workout_usecase.go
│   │   ├── exercise_usecase.go
│   │   └── auth_usecase.go
│   ├── repository/              # Data Layer
│   │   ├── repository.go        # Interfaces
│   │   └── postgres/
│   │       ├── user_repository.go
│   │       ├── workout_repository.go
│   │       └── exercise_repository.go
│   └── delivery/                # HTTP Layer
│       └── http/
│           ├── router.go
│           ├── handler/
│           │   ├── user_handler.go
│           │   ├── workout_handler.go
│           │   ├── exercise_handler.go
│           │   └── auth_handler.go
│           └── middleware/
│               ├── auth.go
│               ├── cors.go
│               └── logger.go
├── pkg/                         # Public packages
│   ├── jwt/
│   │   └── jwt.go
│   ├── validator/
│   │   └── validator.go
│   └── password/
│       └── hash.go
├── config/
│   └── config.go
├── migrations/
│   ├── 001_create_users.up.sql
│   ├── 002_create_exercises.up.sql
│   ├── 003_create_workouts.up.sql
│   └── 004_create_workout_logs.up.sql
├── Dockerfile
├── docker-compose.yml
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

---

## 🎯 Domain Layer

Core business entities and rules with **zero external dependencies**.

### User Entity (`internal/domain/user.go`)

```go
package domain

import "time"

type User struct {
    ID        int64     `json:"id"`
    Email     string    `json:"email"`
    Username  string    `json:"username"`
    Password  string    `json:"-"`
    FullName  string    `json:"full_name"`
    Bio       string    `json:"bio"`
    Weight    float64   `json:"weight"`     // Current weight in kg
    Height    float64   `json:"height"`     // Height in cm
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
    if u.Email == "" {
        return ErrInvalidEmail
    }
    if u.Username == "" || len(u.Username) < 3 {
        return ErrInvalidUsername
    }
    if len(u.Password) < 8 {
        return ErrPasswordTooShort
    }
    return nil
}
```

### Workout, Exercise, WorkoutLog Entities (`internal/domain/workout.go`)

```go
package domain

import "time"

type Workout struct {
    ID          int64      `json:"id"`
    UserID      int64      `json:"user_id"`
    Name        string     `json:"name"`
    Description string     `json:"description"`
    StartTime   time.Time  `json:"start_time"`
    EndTime     *time.Time `json:"end_time"`
    Duration    int        `json:"duration"`      // Duration in seconds
    Notes       string     `json:"notes"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}

type Exercise struct {
    ID          int64     `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Category    string    `json:"category"`       // e.g., "Strength", "Cardio"
    MuscleGroup string    `json:"muscle_group"`   // e.g., "Chest", "Legs"
    Equipment   string    `json:"equipment"`      // e.g., "Barbell", "Dumbbell"
    VideoURL    string    `json:"video_url"`
    CreatedAt   time.Time `json:"created_at"`
}

type WorkoutLog struct {
    ID         int64     `json:"id"`
    WorkoutID  int64     `json:"workout_id"`
    ExerciseID int64     `json:"exercise_id"`
    Sets       int       `json:"sets"`
    Reps       int       `json:"reps"`
    Weight     float64   `json:"weight"`         // Weight in kg
    Duration   int       `json:"duration"`       // For cardio, in seconds
    Distance   float64   `json:"distance"`       // For cardio, in km
    Notes      string    `json:"notes"`
    CreatedAt  time.Time `json:"created_at"`
}
```

### Error Definitions (`internal/domain/errors.go`)

```go
package domain

import "errors"

var (
    // User errors
    ErrUserNotFound      = errors.New("user not found")
    ErrInvalidEmail      = errors.New("invalid email address")
    ErrInvalidUsername   = errors.New("username must be at least 3 characters")
    ErrPasswordTooShort  = errors.New("password must be at least 8 characters")
    ErrDuplicateEmail    = errors.New("email already exists")
    ErrDuplicateUsername = errors.New("username already exists")
    // Auth errors
    ErrInvalidCredentials = errors.New("invalid email or password")
    ErrUnauthorized       = errors.New("unauthorized")
    ErrInvalidToken       = errors.New("invalid token")
    // Workout errors
    ErrWorkoutNotFound    = errors.New("workout not found")
    ErrInvalidWorkout     = errors.New("invalid workout data")
    // Exercise errors
    ErrExerciseNotFound   = errors.New("exercise not found")
    ErrInvalidExercise    = errors.New("invalid exercise data")
)
```

---

## ⚙️ Use Case Layer

Business logic orchestration and repository interfaces.

### User Use Case (`internal/usecase/user_usecase.go`)

```go
package usecase

import (
    "context"

    "fitness-api/internal/domain"
    "fitness-api/internal/repository"
    "fitness-api/pkg/password"
)

type UserUseCase interface {
    Register(ctx context.Context, user *domain.User) error
    GetByID(ctx context.Context, id int64) (*domain.User, error)
    GetByEmail(ctx context.Context, email string) (*domain.User, error)
    Update(ctx context.Context, user *domain.User) error
    Delete(ctx context.Context, id int64) error
    UpdateWeight(ctx context.Context, userID int64, weight float64) error
}

type userUseCase struct {
    userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
    return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) Register(ctx context.Context, user *domain.User) error {
    if err := user.Validate(); err != nil {
        return err
    }
    existing, err := u.userRepo.GetByEmail(ctx, user.Email)
    if err == nil && existing != nil {
        return domain.ErrDuplicateEmail
    }
    existing, err = u.userRepo.GetByUsername(ctx, user.Username)
    if err == nil && existing != nil {
        return domain.ErrDuplicateUsername
    }
    hashedPassword, err := password.Hash(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword
    return u.userRepo.Create(ctx, user)
}

func (u *userUseCase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
    return u.userRepo.GetByID(ctx, id)
}

func (u *userUseCase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
    return u.userRepo.GetByEmail(ctx, email)
}

func (u *userUseCase) Update(ctx context.Context, user *domain.User) error {
    if err := user.Validate(); err != nil {
        return err
    }
    return u.userRepo.Update(ctx, user)
}

func (u *userUseCase) Delete(ctx context.Context, id int64) error {
    return u.userRepo.Delete(ctx, id)
}

func (u *userUseCase) UpdateWeight(ctx context.Context, userID int64, weight float64) error {
    user, err := u.userRepo.GetByID(ctx, userID)
    if err != nil {
        return err
    }
    user.Weight = weight
    return u.userRepo.Update(ctx, user)
}
```

### Workout Use Case (`internal/usecase/workout_usecase.go`)

```go
package usecase

import (
    "context"
    "time"
    "fitness-api/internal/domain"
    "fitness-api/internal/repository"
)

type WorkoutUseCase interface {
    Create(ctx context.Context, workout *domain.Workout) error
    GetByID(ctx context.Context, id int64) (*domain.Workout, error)
    GetUserWorkouts(ctx context.Context, userID int64, limit, offset int) ([]*domain.Workout, error)
    Update(ctx context.Context, workout *domain.Workout) error
    Delete(ctx context.Context, id int64) error
    EndWorkout(ctx context.Context, workoutID int64) error
    // Workout logs
    AddWorkoutLog(ctx context.Context, log *domain.WorkoutLog) error
    GetWorkoutLogs(ctx context.Context, workoutID int64) ([]*domain.WorkoutLog, error)
    GetUserProgress(ctx context.Context, userID int64, exerciseID int64, days int) ([]*domain.WorkoutLog, error)
}

type workoutUseCase struct {
    workoutRepo repository.WorkoutRepository
}

func NewWorkoutUseCase(workoutRepo repository.WorkoutRepository) WorkoutUseCase {
    return &workoutUseCase{workoutRepo: workoutRepo}
}

func (w *workoutUseCase) Create(ctx context.Context, workout *domain.Workout) error {
    if workout.Name == "" {
        return domain.ErrInvalidWorkout
    }
    workout.StartTime = time.Now()
    return w.workoutRepo.Create(ctx, workout)
}

func (w *workoutUseCase) GetByID(ctx context.Context, id int64) (*domain.Workout, error) {
    return w.workoutRepo.GetByID(ctx, id)
}

func (w *workoutUseCase) GetUserWorkouts(ctx context.Context, userID int64, limit, offset int) ([]*domain.Workout, error) {
    return w.workoutRepo.GetUserWorkouts(ctx, userID, limit, offset)
}

func (w *workoutUseCase) Update(ctx context.Context, workout *domain.Workout) error {
    return w.workoutRepo.Update(ctx, workout)
}

func (w *workoutUseCase) Delete(ctx context.Context, id int64) error {
    return w.workoutRepo.Delete(ctx, id)
}

func (w *workoutUseCase) EndWorkout(ctx context.Context, workoutID int64) error {
    workout, err := w.workoutRepo.GetByID(ctx, workoutID)
    if err != nil {
        return err
    }
    now := time.Now()
    workout.EndTime = &now
    workout.Duration = int(now.Sub(workout.StartTime).Seconds())
    return w.workoutRepo.Update(ctx, workout)
}

func (w *workoutUseCase) AddWorkoutLog(ctx context.Context, log *domain.WorkoutLog) error {
    return w.workoutRepo.CreateWorkoutLog(ctx, log)
}

func (w *workoutUseCase) GetWorkoutLogs(ctx context.Context, workoutID int64) ([]*domain.WorkoutLog, error) {
    return w.workoutRepo.GetWorkoutLogs(ctx, workoutID)
}

func (w *workoutUseCase) GetUserProgress(ctx context.Context, userID int64, exerciseID int64, days int) ([]*domain.WorkoutLog, error) {
    return w.workoutRepo.GetUserProgress(ctx, userID, exerciseID, days)
}
```

### Auth Use Case (`internal/usecase/auth_usecase.go`)

```go
package usecase

import (
    "context"

    "fitness-api/internal/domain"
    "fitness-api/internal/repository"
    "fitness-api/pkg/jwt"
    "fitness-api/pkg/password"
)

type AuthUseCase interface {
    Login(ctx context.Context, email, pass string) (string, *domain.User, error)
    ValidateToken(token string) (int64, error)
}

type authUseCase struct {
    userRepo  repository.UserRepository
    jwtSecret string
}

func NewAuthUseCase(userRepo repository.UserRepository, jwtSecret string) AuthUseCase {
    return &authUseCase{
        userRepo:  userRepo,
        jwtSecret: jwtSecret,
    }
}

func (a *authUseCase) Login(ctx context.Context, email, pass string) (string, *domain.User, error) {
    user, err := a.userRepo.GetByEmail(ctx, email)
    if err != nil {
        return "", nil, domain.ErrInvalidCredentials
    }
    if !password.Verify(user.Password, pass) {
        return "", nil, domain.ErrInvalidCredentials
    }
    token, err := jwt.GenerateToken(user.ID, a.jwtSecret)
    if err != nil {
        return "", nil, err
    }
    return token, user, nil
}

func (a *authUseCase) ValidateToken(token string) (int64, error) {
    return jwt.ValidateToken(token, a.jwtSecret)
}
```

---

## 💾 Repository Layer

Data access interfaces and PostgreSQL implementations.

### Repository Interfaces (`internal/repository/repository.go`)

```go
package repository

import (
    "context"
    "fitness-api/internal/domain"
)

type UserRepository interface {
    Create(ctx context.Context, user *domain.User) error
    GetByID(ctx context.Context, id int64) (*domain.User, error)
    GetByEmail(ctx context.Context, email string) (*domain.User, error)
    GetByUsername(ctx context.Context, username string) (*domain.User, error)
    Update(ctx context.Context, user *domain.User) error
    Delete(ctx context.Context, id int64) error
}

type WorkoutRepository interface {
    Create(ctx context.Context, workout *domain.Workout) error
    GetByID(ctx context.Context, id int64) (*domain.Workout, error)
    GetUserWorkouts(ctx context.Context, userID int64, limit, offset int) ([]*domain.Workout, error)
    Update(ctx context.Context, workout *domain.Workout) error
    Delete(ctx context.Context, id int64) error
    CreateWorkoutLog(ctx context.Context, log *domain.WorkoutLog) error
    GetWorkoutLogs(ctx context.Context, workoutID int64) ([]*domain.WorkoutLog, error)
    GetUserProgress(ctx context.Context, userID int64, exerciseID int64, days int) ([]*domain.WorkoutLog, error)
}

type ExerciseRepository interface {
    Create(ctx context.Context, exercise *domain.Exercise) error
    GetByID(ctx context.Context, id int64) (*domain.Exercise, error)
    List(ctx context.Context, category, muscleGroup string, limit, offset int) ([]*domain.Exercise, error)
    Update(ctx context.Context, exercise *domain.Exercise) error
    Delete(ctx context.Context, id int64) error
    Search(ctx context.Context, query string) ([]*domain.Exercise, error)
}
```

---

## 🌐 HTTP Handlers

Delivery layer with Gin framework handlers and middleware.

### Auth Handler (`internal/delivery/http/handler/auth_handler.go`)

```go
package handler

import (
    "fitness-api/internal/domain"
    "fitness-api/internal/usecase"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AuthHandler struct {
    authUseCase usecase.AuthUseCase
    userUseCase usecase.UserUseCase
}

func NewAuthHandler(authUC usecase.AuthUseCase, userUC usecase.UserUseCase) *AuthHandler {
    return &AuthHandler{
        authUseCase: authUC,
        userUseCase: userUC,
    }
}

func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
    auth := router.Group("/auth")
    {
        auth.POST("/register", h.Register)
        auth.POST("/login", h.Login)
    }
}

type RegisterRequest struct {
    Email    string  `json:"email" binding:"required,email"`
    Username string  `json:"username" binding:"required,min=3"`
    Password string  `json:"password" binding:"required,min=8"`
    FullName string  `json:"full_name"`
    Weight   float64 `json:"weight"`
    Height   float64 `json:"height"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
    Token string       `json:"token"`
    User  *domain.User `json:"user"`
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := &domain.User{
        Email:    req.Email,
        Username: req.Username,
        Password: req.Password,
        FullName: req.FullName,
        Weight:   req.Weight,
        Height:   req.Height,
    }

    if err := h.userUseCase.Register(c.Request.Context(), user); err != nil {
        if err == domain.ErrDuplicateEmail || err == domain.ErrDuplicateUsername {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Auto login after registration
    token, _, err := h.authUseCase.Login(c.Request.Context(), req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "registration successful but login failed"})
        return
    }

    c.JSON(http.StatusCreated, AuthResponse{Token: token, User: user})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, user, err := h.authUseCase.Login(c.Request.Context(), req.Email, req.Password)
    if err != nil {
        if err == domain.ErrInvalidCredentials {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, AuthResponse{Token: token, User: user})
}
```

### Workout Handler (`internal/delivery/http/handler/workout_handler.go`)

> _See original for full implementation. All handler methods support CRUD for workouts, workout logs, and getting user progress. Includes handler for registering routes to Gin router, detailed input structs, and error handling patterns consistent with Go and Gin best practices._

---

### Auth Middleware (`internal/delivery/http/middleware/auth.go`)

> _Gin middleware that validates JWT Bearer token using the AuthUseCase, aborts with 401 on missing/invalid tokens, and sets user_id into context on success._

---

### Router Setup (`internal/delivery/http/router.go`)

> _Centralized router configuration for Gin. Sets up global middleware, routes for health check (`/health`), authentication (public), and protected user, workout, and exercise routes requiring JWT authentication._

---

## 🚀 Setup & Deployment

### Main Entry Point (`cmd/server/main.go`)
> _Main Go entrypoint: loads configuration, creates DB connection and repositories, constructs use cases, sets up HTTP router, starts the web server._

---

### Configuration (`config/config.go`)
> _Environment variable loading and configuration. Uses `github.com/joho/godotenv` for `.env` support._

---

### Environment Variables (`.env.example`)

```dotenv
DATABASE_URL=postgres://postgres:postgres@localhost:5432/fitness_db?sslmode=disable
PORT=8080
JWT_SECRET=your-super-secret-jwt-key-change-this
```

---

### Docker Compose (`docker-compose.yml`)

> Provides both `postgres` and `app` services with correct environment, healthchecks, and volume mounting.

---

### Dockerfile

> Multi-stage build for Go API. Alpine image, statically linked binary, exposes port 8080.

---

### Database Migrations

> [See `migrations/` folder for full `.sql` files.]  
> Includes tables: `users`, `exercises`, `workouts`, `workout_logs`.
> - Proper indexes, constraints, and foreign keys for integrity.

---

### Setup Commands

```bash
# Initialize Go module
go mod init fitness-api

# Install dependencies
go get github.com/gin-gonic/gin
go get github.com/lib/pq
go get github.com/golang-jwt/jwt/v5
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt

# Start PostgreSQL with Docker
docker-compose up -d postgres

# Run migrations
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/fitness_db?sslmode=disable" up

# Run the server
go run cmd/server/main.go
```

---

### 📱 React Native Integration

Your React Native (Expo) frontend can connect to this API using fetch or axios.

```javascript
// Example: Login from React Native
const API_URL = 'http://localhost:8080/api/v1';

async function login(email, password) {
  const response = await fetch(`${API_URL}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  
  const data = await response.json();
  // Store token in AsyncStorage
  await AsyncStorage.setItem('token', data.token);
  return data.user;
}
```

---

## 📡 API Documentation

Full REST API endpoints for your fitness app.

### Authentication Endpoints

#### POST `/api/v1/auth/register`

Request:

```json
{
  "email": "user@example.com",
  "username": "johndoe",
  "password": "securepass123",
  "full_name": "John Doe",
  "weight": 75.5,
  "height": 180
}
```

#### POST `/api/v1/auth/login`

Request:

```json
{
  "email": "user@example.com",
  "password": "securepass123"
}
```

---

### User Endpoints

- **GET** `/api/v1/users/me` &mdash; Get current user profile (requires auth)
- **PUT** `/api/v1/users/me` &mdash; Update current user profile
- **PUT** `/api/v1/users/me/weight`

Request:

```json
{ "weight": 76.2 }
```

---

### Workout Endpoints

- **POST** `/api/v1/workouts`

Request:

```json
{
  "name": "Morning Chest Day",
  "description": "Heavy chest and triceps workout"
}
```

- **GET** `/api/v1/workouts?limit=20&offset=0` – Get user's workout history
- **GET** `/api/v1/workouts/:id` – Get specific workout
- **PUT** `/api/v1/workouts/:id` – Update workout
- **DELETE** `/api/v1/workouts/:id` – Delete workout
- **POST** `/api/v1/workouts/:id/end` – End active workout (sets end_time and duration)

---

### Workout Log Endpoints

- **POST** `/api/v1/workouts/:id/logs`

Request:

```json
{
  "exercise_id": 1,
  "sets": 3,
  "reps": 10,
  "weight": 80.0,
  "notes": "Felt strong today"
}
```

- **GET** `/api/v1/workouts/:id/logs` – Get all exercise logs for a workout
- **GET** `/api/v1/progress/:exerciseId?days=30` – Get user's progress for specific exercise over time

---

### Exercise Endpoints

- **GET** `/api/v1/exercises?category=Strength&muscle_group=Chest`
- **GET** `/api/v1/exercises/:id`
- **GET** `/api/v1/exercises/search?q=bench press`
- **POST** `/api/v1/exercises`

Request:

```json
{
  "name": "Bench Press",
  "description": "Classic chest exercise",
  "category": "Strength",
  "muscle_group": "Chest",
  "equipment": "Barbell",
  "video_url": "https://example.com/video"
}
```

---

### 🔐 Authentication

Protected endpoints require JWT token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

---

## 🚀 Quick Start

Follow these steps to get your fitness API up and running:

1. Clone or create the project structure
2. Install dependencies with `go get`
3. Set up environment variables (copy `.env.example` to `.env`)
4. Start PostgreSQL with Docker Compose
5. Run database migrations
6. Start the server with `go run cmd/server/main.go`

Your API will be available at [http://localhost:8080](http://localhost:8080) 🎉
