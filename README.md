# AlignSpace Backend — Golang (Fiber) API

This repository contains the backend for the **AlignSpace** platform — a centralized system for workrooms, PRDs, dashboards, documents, and workspace management.
The backend is built using **Golang + Fiber**, chosen for performance, clean structure, and maintainability.

---

## 🚀 Tech Stack

### **Backend Language**
- **Golang 1.22+**

### **Framework**
- **Fiber** (Fast, Express-like, easy for Node developers switching to Go)

### **Database**
- PostgreSQL (Main transactional DB)  
- Redis (Caching, sessions, rate-limiting — optional but recommended)

### **ORM / DB Layer**
- GORM **or** sqlc (depending on final decision — both supported)

### **Authentication**
- JWT-based authentication  
- Role-based access control (RBAC)  
- Workspace/project-level permissions

### **Architecture Style**
- Modular service-layer architecture  
- Clean folder structure (`handlers`, `services`, `repositories`, `models`)  
- DTO-based request/response validation  
- Centralized error handling middleware  
- RESTful API design

---

## 📁 Project Modules

The platform contains these major functional areas:

### **1. Dashboard**
- High-level overview of workspaces and projects  
- Recently accessed PRDs, documents, workrooms  

### **2. Workspace Module**
- Workspace listing  
- Create/edit/delete workspace  
- Clicking workspace → opens Workspace Workroom  
- Workspace contains PRDs, documents, projects  

### **3. Projects Module**
- Global project list view  
- Clicking a project → opens **PRD View** page  
- Project meta, discussions, version history  

### **4. Documents Module**
- Knowledgebase storage (no search for now)  
- Add/remove documents  
- Access from workspace or global listing  

### **5. PRD Workroom**
- Core place for editing and collaborating on PRDs  
- Components:  
  - Overview  
  - Requirements  
  - Discussions  
  - Tasks  
  - Versioning  
  - Attachments  
- Real-time collaboration planned (Phase 2)
