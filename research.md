Architectural Blueprint for High-Concurrency Real-Time Systems: Implementing a Scalable Chat Application in Go1. Executive SummaryThe demand for real-time interaction in modern web applications has necessitated a shift from traditional request-response architectures to persistent, event-driven communication models. This report provides a comprehensive architectural analysis and implementation guide for building a production-grade live chat application using the Go programming language (Golang), the Gin web framework, and PostgreSQL. The system is designed to support Role-Based Access Control (RBAC) with distinct privileges for administrators and members, authenticated via JSON Web Tokens (JWT).A primary focus of this report is the rigorous application of engineering best practices suitable for large-scale deployment. This includes a detailed examination of the "Standard Go Project Layout" adapted for Clean Architecture to ensure maintainability and testability. Furthermore, the report delineates a robust Test-Driven Development (TDD) strategy utilizing containerized integration testing to validate database interactions and concurrency safety without reliance on brittle mocks.By synthesizing analysis of ecosystem libraries—specifically comparing WebSocket implementations and database drivers—and addressing critical production concerns such as race condition prevention, graceful shutdown, and horizontal scalability, this document serves as a definitive reference for engineering teams tasked with delivering high-throughput real-time systems.2. Architectural Design and Project StructureThe foundation of any large-scale software project lies in its directory structure and architectural pattern. In Go, where the compiler enforces strict dependency cycles and visibility rules (e.g., the internal directory), the initial layout dictates the long-term agility of the codebase. For a chat application combining RESTful administrative endpoints with stateful WebSocket connections, a flat structure is insufficient. We advocate for a modular, layered architecture that strictly separates the transport layer (HTTP/WebSockets) from the core business logic and persistence layers.2.1 The Philosophy of Clean Architecture in GoThe proposed architecture adopts the principles of Clean Architecture, often referred to as Hexagonal Architecture or Ports and Adapters in the Go community. The primary objective is to invert dependencies so that the core business rules (the Domain) depend on nothing, while outer layers (Infrastructure, UI, DB) depend on the Domain. This decoupling allows for the independent evolution of the WebSocket handling mechanism from the underlying database technology.2.1.1 Layered Dependency GraphCore/Domain Layer: Contains the enterprise business rules and entities (e.g., User, Message, ChatRoom). These are pure Go structs with no external tags or dependencies.Service/Use Case Layer: Orchestrates the flow of data to and from the domain entities. It relies on interfaces (Ports) rather than concrete implementations.Adapter/Repository Layer: Implements the interfaces defined by the Service layer. This includes the PostgreSQL implementation using pgx.Transport/Delivery Layer: The entry point of the application. This includes Gin HTTP handlers and the WebSocket Hub.2.2 Comprehensive Directory LayoutAdhering to the de facto Standard Go Project Layout, the following structure is recommended to support scalability, distinct separation of concerns, and ease of navigation for large teams./go-chat-system├── cmd/│ └── api/│ └── main.go # Application Composition Root├── internal/│ ├── domain/ # Enterprise Business Rules (Pure Go)│ │ ├── user.go│ │ ├── message.go│ │ └── room.go│ ├── service/ # Application Business Logic│ │ ├── auth_service.go│ │ ├── chat_service.go│ │ └── service.go # Service Interface definitions│ ├── repository/ # Data Persistence Adapters│ │ ├── postgres/│ │ │ ├── user_repo.go│ │ │ ├── message_repo.go│ │ │ └── db_connection.go│ │ └── redis/ # For Pub/Sub scaling (optional)│ ├── transport/ # Delivery Layer│ │ ├── http/│ │ │ ├── handler/│ │ │ │ ├── auth_handler.go│ │ │ │ └── chat_handler.go│ │ │ ├── middleware/│ │ │ │ ├── jwt_auth.go│ │ │ │ ├── rbac.go│ │ │ │ └── logger.go│ │ │ └── router.go│ │ └── websocket/│ │ ├── client.go # Per-connection logic│ │ └── hub.go # Broadcast orchestration│ └── config/ # Configuration struct definitions├── pkg/ # Publicly consumable libraries│ ├── logger/ # Structured logging wrapper (Zap/Slog)│ ├── utils/ # Generic utilities (hashing, time)│ └── validator/ # Custom input validators├── migrations/ # SQL Database Migrations (.sql files)├── test/ # End-to-End and Integration Tests│ ├── integration/│ │ ├── auth_test.go│ │ └── chat_test.go│ └── testdata/├── Dockerfile├── docker-compose.yml├── go.mod├── go.sum└── Makefile # Build and Test Automation2.3 Component AnalysisThe cmd DirectoryThe cmd/api/main.go file serves as the singular entry point. It acts as the "Composition Root," responsible for reading configuration, initializing the database connection pool, instantiating the repository adapters, injecting them into the services, and finally initializing the Gin router. Crucially, main.go should contain no business logic; it is strictly wiring.The internal DirectoryThe Go compiler treats the internal directory specially; code within it cannot be imported by packages outside the module root. This enforces encapsulation, ensuring that the specific implementation details of the chat system cannot be treated as a public library by other projects.domain: Defines the User struct with fields like ID, Role, Email, and Message struct with Content, Timestamp, SenderID. These structs should generally avoid database-specific tags (like GORM tags) if strict decoupling is desired, though pragmatic Go often permits JSON tags here.repository: Contains the concrete implementation of the storage interface. For example, PostgresUserRepository implements the UserRepository interface defined in the domain or service layer.transport/websocket: Isolates the complexity of the WebSocket protocol. The Hub manages the set of active connections, while Client handles the specific reading and writing pumps for a single TCP connection.The pkg DirectoryCode located here is "public" and can be imported by other projects. This is the appropriate location for a generic structured logger wrapper (e.g., configuring uber-go/zap with specific JSON formatting) or generic utility functions that aren't specific to the chat domain.3. Data Persistence StrategyThe performance of a chat application is heavily constrained by its database write throughput and the efficiency of its history retrieval queries. While Go offers a standard database/sql interface, the choice of driver and abstraction layer significantly impacts performance and developer ergonomics.3.1 Driver Ecosystem EvaluationWe evaluated three primary approaches for interacting with PostgreSQL in Go: GORM (a full-featured ORM), database/sql with lib/pq (legacy), and pgx (modern, high-performance).Table 1: Comparative Analysis of Go PostgreSQL DriversFeatureGORM (ORM)pgx (Driver/Toolkit)database/sql + lib/pqAbstraction LevelHigh (Object-Relational)Low to MediumLowPerformanceLow (Reflection overhead)High (Zero-allocation)MediumType SafetyRuntime checksCompile-time (with sqlc)Runtime checksPostgres FeaturesLimited (Generic SQL)Extensive (Copy, Arrays)LimitedDeveloper SpeedFast (for simple CRUD)MediumSlow (Manual scanning)ComplexityHides complexity (Magic)ExplicitVerboseStrategic Recommendation: For a high-frequency chat application, pgx is the unequivocal choice. GORM's reliance on reflection introduces unnecessary CPU overhead per query, which becomes a bottleneck at scale. Furthermore, lib/pq is effectively in maintenance mode. pgx provides direct access to PostgreSQL-specific features (such as COPY for bulk message insertion) and offers a more performant connection pool (pgxpool) implementation than the standard library.3.2 SQL Generation: The Case for sqlcWhile pgx provides the runtime driver, writing raw SQL strings in Go code can be error-prone and lacks type safety. We recommend integrating sqlc, a compiler that generates type-safe Go code from raw SQL queries.Workflow: Developers write SQL queries in .sql files (e.g., -- name: CreateMessage :one). sqlc parses these files and the database schema to generate Go interfaces and structs.Benefit: If a SQL query references a column that doesn't exist or returns a type that doesn't match the Go struct, the build fails. This brings compile-time safety to SQL interactions, a critical feature for maintaining large applications.3.3 Database Schema DesignThe schema must support RBAC and efficient chronological retrieval of messages.users Table:SQLCREATE TYPE user_role AS ENUM ('admin', 'member');

CREATE TABLE users (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
username VARCHAR(50) NOT NULL UNIQUE,
email VARCHAR(255) NOT NULL UNIQUE,
password_hash VARCHAR(255) NOT NULL,
role user_role NOT NULL DEFAULT 'member',
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
Insight: Using UUID for user IDs prevents resource enumeration attacks. The user_role ENUM enforces RBAC at the database level, preventing invalid roles from ever being persisted.messages Table:SQLCREATE TABLE messages (
id BIGSERIAL PRIMARY KEY,
room_id UUID NOT NULL,
user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
content TEXT NOT NULL,
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Compound index for efficient history retrieval per room
CREATE INDEX idx*messages_room_created ON messages (room_id, created_at DESC);
Insight: While Users use UUIDs, Messages use BIGSERIAL (int64). This is because B-Tree indexes on sequential integers are more compact and performant for insertion than random UUIDs. Since messages are created frequently and rarely looked up by ID externally, the sequential ID offers a significant performance advantage for write-heavy chat logs.4. Authentication and Role-Based Access Control (RBAC)Security is paramount. The system requires a mechanism to authenticate users via REST endpoints and maintain that identity across persistent WebSocket connections.4.1 JSON Web Token (JWT) StrategyStateless authentication via JWT is the industry standard for this use case. It avoids the need for shared session storage (like Redis) for authentication, simplifying the architecture.4.1.1 Token Lifecycle and StructureSigning Algorithm: We employ HMAC-SHA256 (HS256) for signing tokens. As the application is monolithic (or logically grouped), a symmetric key is efficient. If the architecture splits into microservices where auth verification happens externally, switching to RSA (RS256) would allow asymmetric verification.Payload Claims:sub: The UUID of the user.role: The user's role (admin or member) to facilitate rapid RBAC checks without database lookups.iat: Issued At timestamp.exp: Expiration timestamp (short-lived, e.g., 15 minutes).jti: Unique Token ID (useful for revocation lists if implemented).4.1.2 The Refresh Token PatternTo balance security with user experience, we implement a dual-token system:Access Token: Short-lived (15 min). Sent in the HTTP Authorization header.Refresh Token: Long-lived (7 days). Stored in an HttpOnly, Secure, SameSite=Strict cookie. This token is used to request new access tokens.Implication: If an access token is stolen, it is only valid for minutes. The refresh token, which is harder to steal (due to HttpOnly), can be revoked on the server side (by deleting it from a database whitelist).4.2 RBAC Implementation: Middleware vs. Policy EnginesThe requirement specifies two roles: admin and member. We evaluated two implementation paths:Casbin: A powerful authorization library supporting complex models (PERM: Policy, Effect, Request, Matchers). It loads policies from configuration files or databases.Custom Middleware: A lightweight Go function that inspects the JWT claims.Architectural Decision: For a system with binary roles (admin/member), Custom Middleware is the professional choice. Casbin introduces significant complexity (policy loading, caching, adapter maintenance) that is unnecessary for this scope. Custom middleware is easier to debug, test, and performant as it requires no external lookups.4.2.1 Gin Middleware ImplementationThe authentication pipeline consists of two stages:Stage 1: Identity Verification (AuthMiddleware)This middleware extracts the Bearer token, validates the signature, parses the claims, and injects the userID and role into the Gin Context.Stage 2: Access Enforcement (RBACMiddleware)This middleware runs after AuthMiddleware. It checks the Gin Context for the role.Gofunc RequireRole(allowedRoles...string) gin.HandlerFunc {
return func(c \*gin.Context) {
userRole := c.GetString("role")
for *, role := range allowedRoles {
if userRole == role {
c.Next()
return
}
}
c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient privileges"})
}
}
This allows declarative route protection:GoadminRoutes := r.Group("/admin")
adminRoutes.Use(authMiddleware, RequireRole("admin"))
adminRoutes.POST("/broadcast", broadcastHandler) 5. Real-Time Communication Engine (WebSockets)Implementing WebSockets in Go requires navigating a landscape of libraries and managing the inherent complexities of concurrency.5.1 Library Selection: Gorilla vs. CoderThe Go ecosystem offers two primary contenders for WebSocket implementations.Table 2: WebSocket Library ComparisonFeatureGorilla WebSocket (gorilla/websocket)Coder/Nhooyr (coder/websocket)StatusMature, Battle-tested, Community StandardModern, Idiomatic, "Zero-copy" claimsAPI DesignLower-level, explicit controlContext-driven, net.Conn compatiblePerformanceHighVery High (Optimized for low allocation)DocumentationExtensive examplesGood, conciseMaintenancePreviously paused, now activeActive (fork of nhooyr)Recommendation: We recommend gorilla/websocket for this project. While coder/websocket is more modern, Gorilla remains the industry standard with the widest array of community examples, middleware integrations, and documentation coverage, which is crucial for the "learning points" requirement of this report. The patterns used (Hub, Pump) are transferable.5.2 The Hub Architectural PatternTo manage concurrency safely, we employ the Hub Pattern. This creates a centralized manager for all active connections, decoupling the WebSocket transport from the application logic.5.2.1 The ComponentsClient: Represents a single connected user. It holds the *websocket.Conn and a buffered channel send chanbyte.Hub: A struct maintaining the set of all registered clients (map[*Client]bool) and channels for register, unregister, and broadcast.5.2.2 Concurrency Model: The "Pump" StrategyA critical production learning point is that websocket.Conn allows only one concurrent reader and one concurrent writer.Read Pump: Each Client spawns a goroutine dedicated to reading from the WebSocket. It pumps messages from the socket to the Hub.Write Pump: Each Client spawns a goroutine dedicated to writing to the WebSocket. It pumps messages from the send channel to the socket.This isolation ensures that no two goroutines ever write to the socket simultaneously, preventing the common "concurrent write to websocket connection" panic.5.3 Handling Race Conditions and DeadlocksThe Hub manages the state of the system (the map of clients). To prevent race conditions, the Hub must not export its map or use a Mutex that blocks the entire server during broadcasts.Instead, the Hub runs a single event loop (a for-select block) in its own goroutine.Gofunc (h \*Hub) Run() {
for {
select {
case client := <-h.register:
h.clients[client] = true
case client := <-h.unregister:
if \_, ok := h.clients[client]; ok {
delete(h.clients, client)
close(client.send)
}
case message := <-h.broadcast:
for client := range h.clients {
select {
case client.send <- message:
default:
// CRITICAL: Prevent Head-of-Line Blocking
close(client.send)
delete(h.clients, client)
}
}
}
}
}
Architectural Insight: The nested select with a default case inside the broadcast loop is non-negotiable for production. If a client's network is slow and their send channel fills up, attempting to write to it without a default case would block the entire Hub, freezing the chat for all users. The default case drops the slow client, prioritizing the health of the system over a single connection.6. Test-Driven Development (TDD) StrategyTDD in Go, especially for infrastructure-heavy applications involving DBs and WebSockets, requires a strategy that moves beyond simple unit tests with mocks. Mocks often lead to "testing the implementation" rather than behavior.6.1 The Pyramid of Testing for ChatUnit Tests (Business Logic): Focus on the Service layer. Use pure unit tests to verify rules (e.g., "Members cannot post in read-only rooms"). Mocks are acceptable here for the Repository interfaces.Integration Tests (Persistence): Focus on the Repository layer. Use Testcontainers.End-to-End Tests (Transport): Focus on Handlers. Use httptest and real WebSocket clients.6.2 Integration Testing with TestcontainersTesting database interactions with mocks (go-sqlmock) is brittle. It verifies that specific SQL strings are sent, not that the data is actually stored or returned correctly.Strategy:Use the testcontainers-go package.In the TestMain or setup function, programmatically spin up a Docker container running the exact version of PostgreSQL used in production.Apply migrations to this ephemeral database.Run the repository tests against this real database.Teardown the container.This guarantees that SQL syntax, constraints, and data types are valid.6.3 TDD for WebSocketsTesting WebSockets via TDD is complex due to their asynchronous nature.Step 1: Create a test HTTP server using httptest.NewServer that exposes the WebSocket handler.Step 2: Use the gorilla/websocket dialer in the test function to connect to this test server (ws:// URL).Step 3: Send a message from the test client.Step 4: Read from the test client and assert that the echoed message (or broadcast) matches expectations.Race Detection:The most critical part of the TDD strategy for this app is the Go Race Detector.Command: go test -race./...This instrumented build flag detects unsynchronized access to shared memory at runtime. It is the only reliable way to verify that the Hub and Pump implementations are truly concurrency-safe.7. Recommended Package EcosystemTo ensure stability and maintainability, we recommend a curated list of packages that are widely adopted and actively maintained.Table 3: Recommended Package ManifestCategoryPackageJustificationWeb Frameworkgithub.com/gin-gonic/ginHigh performance, minimal allocation, huge middleware ecosystem.WebSocketgithub.com/gorilla/websocketThe industry standard; stable API and extensive documentation.Database Drivergithub.com/jackc/pgx/v5High performance, native Postgres feature support.Authenticationgithub.com/golang-jwt/jwt/v5Standard compliant JWT implementation.Configurationgithub.com/spf13/viperHandles environment variables, config files, and defaults seamlessly.Logginggo.uber.org/zapStructured logging with zero-allocation overhead (critical for high-throughput chat).Migrationsgithub.com/golang-migrate/migrateRobust CLI and library for versioned database schema changes.Testinggithub.com/testcontainers/testcontainers-goProgrammatic Docker containers for real integration tests.Assertionsgithub.com/stretchr/testifyProvides assert and require packages to make tests readable.Docsgithub.com/swaggo/swagGenerates Swagger/OpenAPI 2.0 docs from Go comments.8. Critical Production Learning PointsTransitioning a chat application from a development environment to production introduces specific challenges related to resource management and system reliability.8.1 Graceful ShutdownIn a stateful system like a WebSocket server, you cannot simply kill the process. Doing so severs connections immediately, potentially losing in-flight messages and causing client-side errors.Mechanism: Implement a signal handler for SIGINT and SIGTERM.Workflow:Intercept signal.Stop the HTTP server from accepting new upgrades.Iterate through the Hub's registered clients.Send a websocket.CloseMessage (Control Frame) to each client.Wait for a short timeout (e.g., 5 seconds) for clients to close their side or for the write pumps to finish.Terminate the process.8.2 File Descriptor LimitsEvery WebSocket connection consumes a file descriptor (FD). Default Linux limits are often set to 1024.Impact: The server will crash or reject connections after ~1000 users.Fix: Configure the deployment environment (systemd, Docker ulimit) to increase the FD limit to a value higher than the expected concurrent user count (e.g., 65535).8.3 Horizontal Scalability (The Redis Adapter)The Hub implementation described is monolithic; it works for a single server instance. If the application scales to multiple instances behind a load balancer, User A on Server 1 cannot message User B on Server 2 because the memory space is isolated.Solution: Introduce Redis Pub/Sub.Implementation: When the Hub receives a broadcast message, instead of just iterating local clients, it publishes the message to a Redis channel. Every Hub instance subscribes to this Redis channel. When a message arrives via Redis, the Hub delivers it to its local WebSocket clients. This allows the system to scale linearly with the number of server instances.8.4 Security ConsiderationsCross-Site WebSocket Hijacking (CSWSH): Unlike standard HTTP, WebSockets are not restricted by Same-Origin Policy (SOP). A malicious site can open a socket to your server.Mitigation: The Upgrader struct in gorilla/websocket has a CheckOrigin field. In production, this function must verify that the Origin header matches your domain. Do not simply return true (allow all) as often seen in tutorials.9. Detailed Implementation Guide9.1 Database Layer with pgx and testcontainersRepository Implementation (internal/repository/postgres/user_repo.go):Gopackage postgres

import (
"context"
"fmt"
"github.com/jackc/pgx/v5/pgxpool"
"my-chat-app/internal/domain"
)

type UserRepo struct {
db \*pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {
query := `
		INSERT INTO users (username, email, password_hash, role) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at`
err := r.db.QueryRow(ctx, query, user.Username, user.Email, user.Password, user.Role).
Scan(&user.ID, &user.CreatedAt)
if err!= nil {
return fmt.Errorf("failed to create user: %w", err)
}
return nil
}
Integration Test (test/integration/user_repo_test.go):Gopackage integration

import (
"context"
"testing"
"time"

    "github.com/stretchr/testify/assert"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/modules/postgres"
    "github.com/testcontainers/testcontainers-go/wait"
    "my-chat-app/internal/domain"
    repo "my-chat-app/internal/repository/postgres"
    "github.com/jackc/pgx/v5/pgxpool"

)

func TestUserRepo_Create(t \*testing.T) {
ctx := context.Background()

    // spin up container
    pgContainer, err := postgres.RunContainer(ctx,
    	testcontainers.WithImage("postgres:15-alpine"),
    	postgres.WithDatabase("testdb"),
    	postgres.WithUsername("user"),
    	postgres.WithPassword("password"),
    	testcontainers.WithWaitStrategy(
    		wait.ForLog("database system is ready to accept connections").
    			WithOccurrence(2).
    			WithStartupTimeout(5*time.Second)),
    )
    if err!= nil {
    	t.Fatal(err)
    }
    defer pgContainer.Terminate(ctx)

    // connect
    connStr, _ := pgContainer.ConnectionString(ctx, "sslmode=disable")
    pool, err := pgxpool.New(ctx, connStr)
    if err!= nil {
        t.Fatal(err)
    }

    // run migration (simplified for example)
    _, err = pool.Exec(ctx, `CREATE TABLE users (...)`)
    if err!= nil {
        t.Fatal(err)
    }

    r := repo.NewUserRepo(pool)
    u := &domain.User{
    	Username: "testuser",
    	Email:    "test@example.com",
    	Password: "hashedpassword",
    	Role:     domain.RoleMember,
    }

    err = r.Create(ctx, u)
    assert.NoError(t, err)
    assert.NotEmpty(t, u.ID)

}
9.2 The WebSocket Hub Implementationinternal/transport/websocket/hub.go:Gopackage websocket

import (
"sync"
)

type Hub struct {
// Registered clients.
clients map[*Client]bool

    // Inbound messages from the clients.
    broadcast chanbyte

    // Register requests from the clients.
    register chan *Client

    // Unregister requests from clients.
    unregister chan *Client

    mu sync.RWMutex

}

func NewHub() *Hub {
return &Hub{
broadcast: make(chanbyte),
register: make(chan *Client),
unregister: make(chan *Client),
clients: make(map[*Client]bool),
}
}

func (h \*Hub) Run() {
for {
select {
case client := <-h.register:
h.mu.Lock()
h.clients[client] = true
h.mu.Unlock()

    	case client := <-h.unregister:
            h.mu.Lock()
    		if _, ok := h.clients[client]; ok {
    			delete(h.clients, client)
    			close(client.send)
    		}
            h.mu.Unlock()

    	case message := <-h.broadcast:
            h.mu.RLock()
    		for client := range h.clients {
    			select {
    			case client.send <- message:
    			default:
                    // Buffer is full; disconnect slow client to protect Hub
    				close(client.send)
    				delete(h.clients, client)
    			}
    		}
            h.mu.RUnlock()
    	}
    }

}
internal/transport/websocket/client.go:Gopackage websocket

import (
"time"
"github.com/gorilla/websocket"
)

const (
// Time allowed to write a message to the peer.
writeWait = 10 \* time.Second

    // Time allowed to read the next pong message from the peer.
    pongWait = 60 * time.Second

    // Send pings to peer with this period. Must be less than pongWait.
    pingPeriod = (pongWait * 9) / 10

)

type Client struct {
Hub *Hub
Conn *websocket.Conn
send chanbyte
}

func (c \*Client) readPump() {
defer func() {
c.Hub.unregister <- c
c.Conn.Close()
}()

    c.Conn.SetReadLimit(512) // Limit message size
    c.Conn.SetReadDeadline(time.Now().Add(pongWait))
    c.Conn.SetPongHandler(func(string) error {
        c.Conn.SetReadDeadline(time.Now().Add(pongWait))
        return nil
    })

    for {
    	_, message, err := c.Conn.ReadMessage()
    	if err!= nil {
    		break
    	}
    	c.Hub.broadcast <- message
    }

}

func (c \*Client) writePump() {
ticker := time.NewTicker(pingPeriod)
defer func() {
ticker.Stop()
c.Conn.Close()
}()

    for {
    	select {
    	case message, ok := <-c.send:
    		c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
    		if!ok {
    			// The Hub closed the channel.
    			c.Conn.WriteMessage(websocket.CloseMessage,byte{})
    			return
    		}

    		w, err := c.Conn.NextWriter(websocket.TextMessage)
    		if err!= nil {
    			return
    		}
    		w.Write(message)

    		if err := w.Close(); err!= nil {
    			return
    		}

    	case <-ticker.C:
    		c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
    		if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err!= nil {
    			return
    		}
    	}
    }

}
9.3 Middleware Implementation (RBAC)internal/transport/http/middleware/rbac.go:Gopackage middleware

import (
"net/http"
"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles...string) gin.HandlerFunc {
return func(c \*gin.Context) {
// Assumes AuthMiddleware has already run and set "role"
userRole := c.GetString("role")
if userRole == "" {
c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User role not determined"})
return
}

    	// Admin superuser override
    	if userRole == "admin" {
    		c.Next()
    		return
    	}

    	for _, role := range allowedRoles {
    		if userRole == role {
    			c.Next()
    			return
    		}
    	}

    	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
    }

} 10. ConclusionThis report has outlined a complete architectural roadmap for a production-ready real-time chat application. By adhering to Clean Architecture, we ensure that the business rules remain isolated from the transport mechanism, facilitating future upgrades or refactors. The choice of pgx over GORM ensures the database layer can handle high throughput, while the Hub pattern with gorilla/websocket provides a thread-safe foundation for real-time communication.Crucially, the adoption of Test-Driven Development using testcontainers allows the engineering team to verify database integrity without relying on production environments, and the strict use of the race detector during testing guarantees that the concurrent complexities of the system are managed correctly. This blueprint not only meets the immediate functional requirements but also establishes a foundation for a scalable, maintainable, and robust distributed system.
