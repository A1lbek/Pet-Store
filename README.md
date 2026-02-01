## Assignment 4 – Core System Implementation (Milestone 2)

### Running the System
The backend application is implemented as a monolithic Go service.
The system can be started using the following command:

''
go run ./cmd/app

After startup, the HTTP server listens on port 8080 and is ready to accept requests.

---

### Implemented Core Features
This milestone focuses on implementing a working core system rather than a complete feature set.
The following core features related to the Pet domain have been implemented:

1. Create Pet  
   Endpoint: POST /pets  
   The system accepts JSON input describing a pet and creates a new pet entity in memory.
   A unique identifier and default status are assigned automatically.

2. Get All Pets  
   Endpoint: GET /pets  
   Returns a list of all pets currently stored in the system in JSON format.

3. Get Pet by ID  
   Endpoint: GET /pets/{id}  
   Returns a single pet based on its identifier, demonstrating basic routing and data lookup.

These endpoints demonstrate that the system can accept input, process data, and return results as required.

---

### Relation to Assignment 3 Design
The implementation strictly follows the approved design from Assignment 3.

Architecture:
- The system is implemented as a monolithic backend application.
- A layered architecture is used, separating handlers, services, and repositories.

Data Model:
- The Pet domain model matches the ERD defined in Assignment 3.
- All attributes (id, name, species, age, price, status) correspond directly to the original data model.

Module Responsibility:
- The Pet module is implemented as an independent domain module.
- Business logic is handled in the service layer, while data access is isolated in the repository layer.

Scope Control:
- Only core functionality is implemented in this milestone, as planned.
- Authentication, payments, UI, and advanced error handling are intentionally excluded and deferred to later milestones.

---

### API Demonstration
API endpoints can be tested using a REST client.
A prepared .http file is provided to demonstrate all implemented endpoints using the VS Code REST Client.

### Data Model & Storage

This section describes the implementation of the **core data structures** and the **in-memory storage** for the Pet Store backend.  
It covers the responsibilities of **Participant 2**, including concurrency handling.

#### Implemented Domain Models

All entities from the ERD in Assignment 3 are represented in the code, ensuring full alignment with the approved design:

- **User**
  - Fields: ID, Name, Email, Password, Role
  - Represents a system user who can place orders.
- **Pet**
  - Fields: ID, Name, Species, Age, Price, Status ✅
  - Full CRUD operations implemented.
- **Order**
  - Fields: ID, UserID, TotalPrice, Status, CreatedAt
  - Represents a user's order.
- **Product**
  - Fields: ID, Name, Category, Price, Stock
  - Represents items available for sale.
- **OrderItem**
  - Fields: ID, OrderID, ProductID, Quantity, Price
  - Represents individual products included in an order.

> Note: Only **Pet** has fully implemented CRUD operations in this milestone. Other entities are defined as structs to satisfy the ERD and ensure data model consistency.

#### Pet Repository Implementation

- **In-Memory Storage:**  
  Pets are stored in a `map[int]Pet`, with a mutex (`sync.Mutex`) to ensure thread-safe access.

- **CRUD Methods:**
  1. `Create` — Adds a new pet with an auto-incremented ID and default status.
  2. `GetAll` — Returns all pets.
  3. `GetByID` — Retrieves a pet by ID.
  4. `Update` — Updates an existing pet by ID.
  5. `Delete` — Deletes a pet by ID.

- **Concurrency Handling:**  
  - `sync.Mutex` ensures that concurrent access to the pet map is safe.  
  - A `logCh` channel and a background goroutine (`backgroundLogger`) asynchronously log newly added pets, demonstrating safe concurrent operations.

#### Running and Testing

- The Pet repository can be tested through the HTTP endpoints described in the previous section.  
- All operations are thread-safe and consistent with the ERD.  
- Example: Adding a pet via POST /pets triggers asynchronous logging without blocking other operations.
