# Go Fitness API - Clean Architecture Template **Production-Ready Backend with
Clean Architecture** [![Go
1.21+](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![JWT
Auth](https://img.shields.io/badge/JWT-Auth-000000?style=flat&logo=json-web-tokens)](https://jwt.io/)
[![Clean
Architecture](https://img.shields.io/badge/Clean-Architecture-blue)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
[![REST
API](https://img.shields.io/badge/REST-API-green)](https://restfulapi.net/) ##
📋 Overview A complete fitness tracking backend built with Go and Clean
Architecture principles. Perfect starting point for your fitness mobile app with
React Native (Expo). ### 🏗️ Architecture Clean Architecture with clear
separation of concerns, making the codebase testable, maintainable, and
scalable. - Domain Layer (entities) - Use Case Layer (business logic) -
Repository Layer (data access) - Delivery Layer (HTTP/gRPC) ### ✨ Features -
User authentication (JWT) - Workout logging & tracking - Exercise library
management - Progress analytics - Social feed ready - ML recommendation ready
### 🛠️ Tech Stack - Go 1.21+ (latest features) - PostgreSQL (robust RDBMS) -
Gin/Fiber (HTTP framework) - JWT (authentication) - Docker (containerization) -
Migrate (DB migrations) ### 🎨 Design Principles - Dependency Injection -
Interface-based design - Repository pattern - SOLID principles - Testable
architecture - Context-aware operations --- ## 📁 Project Structure Standard Go
project layout following Clean Architecture principles: ``` fitness-api/ ├──
cmd/ │ └── server/ │ └── main.go # Entry point ├── internal/ │ ├── domain/ #
Domain Layer │ │ ├── user.go │ │ ├── workout.go │ │ ├── exercise.go │ │ ├──
workout_log.go │ │ └── errors.go │ ├── usecase/ # Business Logic │ │ ├──
user_usecase.go │ │ ├── workout_usecase.go │ │ ├── exercise_usecase.go │ │ └──
auth_usecase.go │ ├── repository/ # Data Layer │ │ ├── repository.go #
Interfaces │ │ └── postgres/ │ │ ├── user_repository.go │ │ ├──
workout_repository.go │ │ └── exercise_repository.go │ └── delivery/ # HTTP
Layer │ └── http/ │ ├── router.go │ ├── handler/ │ │ ├── user_handler.go │ │ ├──
workout_handler.go │ │ ├── exercise_handler.go │ │ └── auth_handler.go │ └──
middleware/ │ ├── auth.go │ ├── cors.go │ └── logger.go ├── pkg/ # Public
packages │ ├── jwt/ │ │ └── jwt.go │ ├── validator/ │ │ └── validator.go │ └──
password/ │ └── hash.go ├── config/ │ └── config.go ├── migrations/ │ ├──
001_create_users.up.sql │ ├── 002_create_exercises.up.sql │ ├──
003_create_workouts.up.sql │ └── 004_create_workout_logs.up.sql ├── Dockerfile
├── docker-compose.yml ├── .env.example ├── go.mod ├── go.sum └── README.md ```
--- ## 🎯 Domain Layer Core business entities and rules with zero external
dependencies. ### User Entity (internal/domain/user.go) ```go package domain
import "time" type User struct { ID int64 `json:"id"` Email string
`json:"email"` Username string `json:"username"` Password string `json:"-"`
FullName string `json:"full_name"` Bio string `json:"bio"` Weight float64
`json:"weight"` // Current weight in kg Height float64 `json:"height"` // Height
in cm CreatedAt time.Time `json:"created_at"` UpdatedAt time.Time
`json:"updated_at"` } func (u *User) Validate() error { if u.Email == "" {
return ErrInvalidEmail } if u.Username == "" || len(u.Username) < 3 { return
ErrInvalidUsername } if len(u.Password) < 8 { return ErrPasswordTooShort }
return nil } ``` ### Workout Entity (internal/domain/workout.go) ```go package
domain import "time" type Workout struct { ID int64 `json:"id"` UserID int64
`json:"user_id"` Name string `json:"name"` Description string
`json:"description"` StartTime time.Time `json:"start_time"` EndTime *time.Time
`json:"end_time"` Duration int `json:"duration"` // Duration in seconds Notes
string `json:"notes"` CreatedAt time.Time `json:"created_at"` UpdatedAt
time.Time `json:"updated_at"` } type Exercise struct { ID int64 `json:"id"` Name
string `json:"name"` Description string `json:"description"` Category string
`json:"category"` // e.g., "Strength", "Cardio" MuscleGroup string
`json:"muscle_group"` // e.g., "Chest", "Legs" Equipment string
`json:"equipment"` // e.g., "Barbell", "Dumbbell" VideoURL string
`json:"video_url"` CreatedAt time.Time `json:"created_at"` } type WorkoutLog
struct { ID int64 `json:"id"` WorkoutID int64 `json:"workout_id"` ExerciseID
int64 `json:"exercise_id"` Sets int `json:"sets"` Reps int `json:"reps"` Weight
float64 `json:"weight"` // Weight in kg Duration int `json:"duration"` // For
cardio, in seconds Distance float64 `json:"distance"` // For cardio, in km Notes
string `json:"notes"` CreatedAt time.Time `json:"created_at"` } ``` ### Error
Definitions (internal/domain/errors.go) ```go package domain import "errors" var
( // User errors ErrUserNotFound = errors.New("user not found") ErrInvalidEmail
= errors.New("invalid email address") ErrInvalidUsername = errors.New("username
must be at least 3 characters") ErrPasswordTooShort = errors.New("password must
be at least 8 characters") ErrDuplicateEmail = errors.New("email already
exists") ErrDuplicateUsername = errors.New("username already exists") // Auth
errors ErrInvalidCredentials = errors.New("invalid email or password")
ErrUnauthorized = errors.New("unauthorized") ErrInvalidToken =
errors.New("invalid token") // Workout errors ErrWorkoutNotFound =
errors.New("workout not found") ErrInvalidWorkout = errors.New("invalid workout
data") // Exercise errors ErrExerciseNotFound = errors.New("exercise not found")
ErrInvalidExercise = errors.New("invalid exercise data") ) ``` --- ## ⚙️ Use
Case Layer Business logic orchestration and repository interfaces. ### User Use
Case (internal/usecase/user_usecase.go) ```go package usecase import ( "context"
"fitness-api/internal/domain" "fitness-api/internal/repository"
"fitness-api/pkg/password" ) type UserUseCase interface { Register(ctx
context.Context, user *domain.User) error GetByID(ctx context.Context, id int64)
(*domain.User, error) GetByEmail(ctx context.Context, email string)
(*domain.User, error) Update(ctx context.Context, user *domain.User) error
Delete(ctx context.Context, id int64) error UpdateWeight(ctx context.Context,
userID int64, weight float64) error } type userUseCase struct { userRepo
repository.UserRepository } func NewUserUseCase(userRepo
repository.UserRepository) UserUseCase { return &userUseCase{userRepo: userRepo}
} func (u *userUseCase) Register(ctx context.Context, user *domain.User) error {
if err := user.Validate(); err != nil { return err } // Check for duplicate
email existing, err := u.userRepo.GetByEmail(ctx, user.Email) if err == nil &&
existing != nil { return domain.ErrDuplicateEmail } // Check for duplicate
username existing, err = u.userRepo.GetByUsername(ctx, user.Username) if err ==
nil && existing != nil { return domain.ErrDuplicateUsername } // Hash password
hashedPassword, err := password.Hash(user.Password) if err != nil { return err }
user.Password = hashedPassword return u.userRepo.Create(ctx, user) } func (u
*userUseCase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
return u.userRepo.GetByID(ctx, id) } func (u *userUseCase) GetByEmail(ctx
context.Context, email string) (*domain.User, error) { return
u.userRepo.GetByEmail(ctx, email) } func (u *userUseCase) Update(ctx
context.Context, user *domain.User) error { if err := user.Validate(); err !=
nil { return err } return u.userRepo.Update(ctx, user) } func (u *userUseCase)
Delete(ctx context.Context, id int64) error { return u.userRepo.Delete(ctx, id)
} func (u *userUseCase) UpdateWeight(ctx context.Context, userID int64, weight
float64) error { user, err := u.userRepo.GetByID(ctx, userID) if err != nil {
return err } user.Weight = weight return u.userRepo.Update(ctx, user) } ``` ###
Workout Use Case (internal/usecase/workout_usecase.go) ```go package usecase
import ( "context" "fitness-api/internal/domain"
"fitness-api/internal/repository" "time" ) type WorkoutUseCase interface {
Create(ctx context.Context, workout *domain.Workout) error GetByID(ctx
context.Context, id int64) (*domain.Workout, error) GetUserWorkouts(ctx
context.Context, userID int64, limit, offset int) ([]*domain.Workout, error)
Update(ctx context.Context, workout *domain.Workout) error Delete(ctx
context.Context, id int64) error EndWorkout(ctx context.Context, workoutID
int64) error // Workout logs AddWorkoutLog(ctx context.Context, log
*domain.WorkoutLog) error GetWorkoutLogs(ctx context.Context, workoutID int64)
([]*domain.WorkoutLog, error) GetUserProgress(ctx context.Context, userID int64,
exerciseID int64, days int) ([]*domain.WorkoutLog, error) } type workoutUseCase
struct { workoutRepo repository.WorkoutRepository } func
NewWorkoutUseCase(workoutRepo repository.WorkoutRepository) WorkoutUseCase {
return &workoutUseCase{workoutRepo: workoutRepo} } func (w *workoutUseCase)
Create(ctx context.Context, workout *domain.Workout) error { if workout.Name ==
"" { return domain.ErrInvalidWorkout } workout.StartTime = time.Now() return
w.workoutRepo.Create(ctx, workout) } func (w *workoutUseCase) GetByID(ctx
context.Context, id int64) (*domain.Workout, error) { return
w.workoutRepo.GetByID(ctx, id) } func (w *workoutUseCase) GetUserWorkouts(ctx
context.Context, userID int64, limit, offset int) ([]*domain.Workout, error) {
return w.workoutRepo.GetUserWorkouts(ctx, userID, limit, offset) } func (w
*workoutUseCase) Update(ctx context.Context, workout *domain.Workout) error {
return w.workoutRepo.Update(ctx, workout) } func (w *workoutUseCase) Delete(ctx
context.Context, id int64) error { return w.workoutRepo.Delete(ctx, id) } func
(w *workoutUseCase) EndWorkout(ctx context.Context, workoutID int64) error {
workout, err := w.workoutRepo.GetByID(ctx, workoutID) if err != nil { return err
} now := time.Now() workout.EndTime = &now workout.Duration =
int(now.Sub(workout.StartTime).Seconds()) return w.workoutRepo.Update(ctx,
workout) } func (w *workoutUseCase) AddWorkoutLog(ctx context.Context, log
*domain.WorkoutLog) error { return w.workoutRepo.CreateWorkoutLog(ctx, log) }
func (w *workoutUseCase) GetWorkoutLogs(ctx context.Context, workoutID int64)
([]*domain.WorkoutLog, error) { return w.workoutRepo.GetWorkoutLogs(ctx,
workoutID) } func (w *workoutUseCase) GetUserProgress(ctx context.Context,
userID int64, exerciseID int64, days int) ([]*domain.WorkoutLog, error) { return
w.workoutRepo.GetUserProgress(ctx, userID, exerciseID, days) } ``` ### Auth Use
Case (internal/usecase/auth_usecase.go) ```go package usecase import ( "context"
"fitness-api/internal/domain" "fitness-api/internal/repository"
"fitness-api/pkg/jwt" "fitness-api/pkg/password" ) type AuthUseCase interface {
Login(ctx context.Context, email, pass string) (string, *domain.User, error)
ValidateToken(token string) (int64, error) } type authUseCase struct { userRepo
repository.UserRepository jwtSecret string } func NewAuthUseCase(userRepo
repository.UserRepository, jwtSecret string) AuthUseCase { return &authUseCase{
userRepo: userRepo, jwtSecret: jwtSecret, } } func (a *authUseCase) Login(ctx
context.Context, email, pass string) (string, *domain.User, error) { user, err
:= a.userRepo.GetByEmail(ctx, email) if err != nil { return "", nil,
domain.ErrInvalidCredentials } if !password.Verify(user.Password, pass) { return
"", nil, domain.ErrInvalidCredentials } token, err := jwt.GenerateToken(user.ID,
a.jwtSecret) if err != nil { return "", nil, err } return token, user, nil }
func (a *authUseCase) ValidateToken(token string) (int64, error) { return
jwt.ValidateToken(token, a.jwtSecret) } ``` --- ## 💾 Repository Layer Data
access interfaces and PostgreSQL implementations. ### Repository Interfaces
(internal/repository/repository.go) ```go package repository import ( "context"
"fitness-api/internal/domain" ) type UserRepository interface { Create(ctx
context.Context, user *domain.User) error GetByID(ctx context.Context, id int64)
(*domain.User, error) GetByEmail(ctx context.Context, email string)
(*domain.User, error) GetByUsername(ctx context.Context, username string)
(*domain.User, error) Update(ctx context.Context, user *domain.User) error
Delete(ctx context.Context, id int64) error } type WorkoutRepository interface {
Create(ctx context.Context, workout *domain.Workout) error GetByID(ctx
context.Context, id int64) (*domain.Workout, error) GetUserWorkouts(ctx
context.Context, userID int64, limit, offset int) ([]*domain.Workout, error)
Update(ctx context.Context, workout *domain.Workout) error Delete(ctx
context.Context, id int64) error CreateWorkoutLog(ctx context.Context, log
*domain.WorkoutLog) error GetWorkoutLogs(ctx context.Context, workoutID int64)
([]*domain.WorkoutLog, error) GetUserProgress(ctx context.Context, userID int64,
exerciseID int64, days int) ([]*domain.WorkoutLog, error) } type
ExerciseRepository interface { Create(ctx context.Context, exercise
*domain.Exercise) error GetByID(ctx context.Context, id int64)
(*domain.Exercise, error) List(ctx context.Context, category, muscleGroup
string, limit, offset int) ([]*domain.Exercise, error) Update(ctx
context.Context, exercise *domain.Exercise) error Delete(ctx context.Context, id
int64) error Search(ctx context.Context, query string) ([]*domain.Exercise,
error) } ``` ### User Repository
(internal/repository/postgres/user_repository.go) ```go package postgres import
( "context" "database/sql" "fitness-api/internal/domain"
"fitness-api/internal/repository" "time" ) type userRepository struct { db
*sql.DB } func NewUserRepository(db *sql.DB) repository.UserRepository { return
&userRepository{db: db} } func (r *userRepository) Create(ctx context.Context,
user *domain.User) error { query := ` INSERT INTO users (email, username,
password, full_name, bio, weight, height, created_at, updated_at) VALUES ($1,
$2, $3, $4, $5, $6, $7, $8, $9) RETURNING id ` now := time.Now() user.CreatedAt
= now user.UpdatedAt = now err := r.db.QueryRowContext(ctx, query, user.Email,
user.Username, user.Password, user.FullName, user.Bio, user.Weight, user.Height,
user.CreatedAt, user.UpdatedAt, ).Scan(&user.ID) return err } func (r
*userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
query := ` SELECT id, email, username, password, full_name, bio, weight, height,
created_at, updated_at FROM users WHERE id = $1 ` user := &domain.User{} err :=
r.db.QueryRowContext(ctx, query, id).Scan( &user.ID, &user.Email,
&user.Username, &user.Password, &user.FullName, &user.Bio, &user.Weight,
&user.Height, &user.CreatedAt, &user.UpdatedAt, ) if err == sql.ErrNoRows {
return nil, domain.ErrUserNotFound } return user, err } func (r *userRepository)
GetByEmail(ctx context.Context, email string) (*domain.User, error) { query := `
SELECT id, email, username, password, full_name, bio, weight, height,
created_at, updated_at FROM users WHERE email = $1 ` user := &domain.User{} err
:= r.db.QueryRowContext(ctx, query, email).Scan( &user.ID, &user.Email,
&user.Username, &user.Password, &user.FullName, &user.Bio, &user.Weight,
&user.Height, &user.CreatedAt, &user.UpdatedAt, ) if err == sql.ErrNoRows {
return nil, domain.ErrUserNotFound } return user, err } func (r *userRepository)
GetByUsername(ctx context.Context, username string) (*domain.User, error) {
query := ` SELECT id, email, username, password, full_name, bio, weight, height,
created_at, updated_at FROM users WHERE username = $1 ` user := &domain.User{}
err := r.db.QueryRowContext(ctx, query, username).Scan( &user.ID, &user.Email,
&user.Username, &user.Password, &user.FullName, &user.Bio, &user.Weight,
&user.Height, &user.CreatedAt, &user.UpdatedAt, ) if err == sql.ErrNoRows {
return nil, domain.ErrUserNotFound } return user, err } func (r *userRepository)
Update(ctx context.Context, user *domain.User) error { query := ` UPDATE users
SET email = $1, username = $2, full_name = $3, bio = $4, weight = $5, height =
$6, updated_at = $7 WHERE id = $8 ` user.UpdatedAt = time.Now() _, err :=
r.db.ExecContext(ctx, query, user.Email, user.Username, user.FullName, user.Bio,
user.Weight, user.Height, user.UpdatedAt, user.ID, ) return err } func (r
*userRepository) Delete(ctx context.Context, id int64) error { query := `DELETE
FROM users WHERE id = $1` _, err := r.db.ExecContext(ctx, query, id) return err
} ``` ### Workout Repository
(internal/repository/postgres/workout_repository.go) ```go package postgres
import ( "context" "database/sql" "fitness-api/internal/domain"
"fitness-api/internal/repository" "time" ) type workoutRepository struct { db
*sql.DB } func NewWorkoutRepository(db *sql.DB) repository.WorkoutRepository {
return &workoutRepository{db: db} } func (r *workoutRepository) Create(ctx
context.Context, workout *domain.Workout) error { query := ` INSERT INTO
workouts (user_id, name, description, start_time, created_at, updated_at) VALUES
($1, $2, $3, $4, $5, $6) RETURNING id ` now := time.Now() workout.CreatedAt =
now workout.UpdatedAt = now return r.db.QueryRowContext(ctx, query,
workout.UserID, workout.Name, workout.Description, workout.StartTime,
workout.CreatedAt, workout.UpdatedAt, ).Scan(&workout.ID) } func (r
*workoutRepository) GetByID(ctx context.Context, id int64) (*domain.Workout,
error) { query := ` SELECT id, user_id, name, description, start_time, end_time,
duration, notes, created_at, updated_at FROM workouts WHERE id = $1 ` workout :=
&domain.Workout{} err := r.db.QueryRowContext(ctx, query, id).Scan( &workout.ID,
&workout.UserID, &workout.Name, &workout.Description, &workout.StartTime,
&workout.EndTime, &workout.Duration, &workout.Notes, &workout.CreatedAt,
&workout.UpdatedAt, ) if err == sql.ErrNoRows { return nil,
domain.ErrWorkoutNotFound } return workout, err } func (r *workoutRepository)
GetUserWorkouts(ctx context.Context, userID int64, limit, offset int)
([]*domain.Workout, error) { query := ` SELECT id, user_id, name, description,
start_time, end_time, duration, notes, created_at, updated_at FROM workouts
WHERE user_id = $1 ORDER BY start_time DESC LIMIT $2 OFFSET $3 ` rows, err :=
r.db.QueryContext(ctx, query, userID, limit, offset) if err != nil { return nil,
err } defer rows.Close() workouts := []*domain.Workout{} for rows.Next() { w :=
&domain.Workout{} err := rows.Scan( &w.ID, &w.UserID, &w.Name, &w.Description,
&w.StartTime, &w.EndTime, &w.Duration, &w.Notes, &w.CreatedAt, &w.UpdatedAt, )
if err != nil { return nil, err } workouts = append(workouts, w) } return
workouts, rows.Err() } func (r *workoutRepository) Update(ctx context.Context,
workout *domain.Workout) error { query := ` UPDATE workouts SET name = $1,
description = $2, end_time = $3, duration = $4, notes = $5, updated_at = $6
WHERE id = $7 ` workout.UpdatedAt = time.Now() _, err := r.db.ExecContext(ctx,
query, workout.Name, workout.Description, workout.EndTime, workout.Duration,
workout.Notes, workout.UpdatedAt, workout.ID, ) return err } func (r
*workoutRepository) Delete(ctx context.Context, id int64) error { query :=
`DELETE FROM workouts WHERE id = $1` _, err := r.db.ExecContext(ctx, query, id)
return err } func (r *workoutRepository) CreateWorkoutLog(ctx context.Context,
log *domain.WorkoutLog) error { query := ` INSERT INTO workout_logs (workout_id,
exercise_id, sets, reps, weight, duration, distance, notes, created_at) VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id ` log.CreatedAt = time.Now()
return r.db.QueryRowContext(ctx, query, log.WorkoutID, log.ExerciseID, log.Sets,
log.Reps, log.Weight, log.Duration, log.Distance, log.Notes, log.CreatedAt,
).Scan(&log.ID) } func (r *workoutRepository) GetWorkoutLogs(ctx
context.Context, workoutID int64) ([]*domain.WorkoutLog, error) { query := `
SELECT id, workout_id, exercise_id, sets, reps, weight, duration, distance,
notes, created_at FROM workout_logs WHERE workout_id = $1 ORDER BY created_at
ASC ` rows, err := r.db.QueryContext(ctx, query, workoutID) if err != nil {
return nil, err } defer rows.Close() logs := []*domain.WorkoutLog{} for
rows.Next() { log := &domain.WorkoutLog{} err := rows.Scan( &log.ID,
&log.WorkoutID, &log.ExerciseID, &log.Sets, &log.Reps, &log.Weight,
&log.Duration, &log.Distance, &log.Notes, &log.CreatedAt, ) if err != nil {
return nil, err } logs = append(logs, log) } return logs, rows.Err() } func (r
*workoutRepository) GetUserProgress(ctx context.Context, userID int64,
exerciseID int64, days int) ([]*domain.WorkoutLog, error) { query := ` SELECT
wl.id, wl.workout_id, wl.exercise_id, wl.sets, wl.reps, wl.weight, wl.duration,
wl.distance, wl.notes, wl.created_at FROM workout_logs wl JOIN workouts w ON
wl.workout_id = w.id WHERE w.user_id = $1 AND wl.exercise_id = $2 AND
wl.created_at >= NOW() - INTERVAL '$3 days' ORDER BY wl.created_at DESC ` rows,
err := r.db.QueryContext(ctx, query, userID, exerciseID, days) if err != nil {
return nil, err } defer rows.Close() logs := []*domain.WorkoutLog{} for
rows.Next() { log := &domain.WorkoutLog{} err := rows.Scan( &log.ID,
&log.WorkoutID, &log.ExerciseID, &log.Sets, &log.Reps, &log.Weight,
&log.Duration, &log.Distance, &log.Notes, &log.CreatedAt, ) if err != nil {
return nil, err } logs = append(logs, log) } return logs, rows.Err() } ``` ---
## 🌐 HTTP Handlers Delivery layer with Gin framework handlers and middleware.
### Auth Handler (internal/delivery/http/handler/auth_handler.go) ```go package
handler import ( "fitness-api/internal/domain" "fitness-api/internal/usecase"
"net/http" "github.com/gin-gonic/gin" ) type AuthHandler struct { authUseCase
usecase.AuthUseCase userUseCase usecase.UserUseCase } func NewAuthHandler(authUC
usecase.AuthUseCase, userUC usecase.UserUseCase) *AuthHandler { return
&AuthHandler{ authUseCase: authUC, userUseCase: userUC, } } func (h
*AuthHandler) RegisterRoutes(router *gin.RouterGroup) { auth :=
router.Group("/auth") { auth.POST("/register", h.Register) auth.POST("/login",
h.Login) } } type RegisterRequest struct { Email string `json:"email"
binding:"required,email"` Username string `json:"username"
binding:"required,min=3"` Password string `json:"password"
binding:"required,min=8"` FullName string `json:"full_name"` Weight float64
`json:"weight"` Height float64 `json:"height"`
