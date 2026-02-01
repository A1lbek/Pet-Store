## Assignment 4 – Core System Implementation (Milestone 2)

### Running the System
The backend application is implemented as a monolithic Go service.
The system can be started using the following command:

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
