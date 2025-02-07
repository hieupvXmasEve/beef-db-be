/project-root
│── /cmd                # Entry points for different applications
│    ├── /api           # Main API entry point
│    │   ├── main.go    # Initializes and starts the API
│── /internal           # Private application code (not exposed as a module)
│    ├── /config        # Configuration settings (env, database, etc.)
│    ├── /server        # Server setup (HTTP server, middleware, routes)
│    ├── /db            # Database interactions (migrations, setup)
│    ├── /repository    # Data access layer (CRUD operations)
│    ├── /service       # Business logic layer
│    ├── /handler       # HTTP request handlers (controllers)
│    ├── /middleware    # Custom middlewares (auth, logging, etc.)
│    ├── /model         # Data models (struct definitions)
│    ├── /utils         # Utility functions (helpers)
│── pkg                # Public utilities (can be used by other projects)
│   └── utils/
│── /api                # API specifications (OpenAPI, Protobuf, GraphQL schemas)
│   └── swagger.yaml
│── /migrations         # Database migration files
│   └── 001_init.sql
│── /scripts            # Helper scripts for automation (deployment, database, etc.)
│   └── start.sh
│── /configs            # Configuration files
│   └── config.yaml
│── /deployments        # Deployment files
│   ├── Dockerfile
│   └── k8s/
│── /test               # Test files (unit, integration)
│   └── integration/
│── go.mod              # Go module file
│── go.sum              # Go dependencies
│── Makefile            # Automation commands
│── README.md           # Project documentation