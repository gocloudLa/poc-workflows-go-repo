# Go Service Example

This is a minimal Go service example demonstrating the centralized CI/CD workflow system.

## Structure

- `main.go`: Main service file with simple functions
- `main_test.go`: Unit tests using Go's testing package
- `Dockerfile`: Container image definition
- `.github/workflows/`: GitHub Actions workflows that consume shared workflows

## Workflows

This repository uses the centralized workflows from `gocloudLa/poc-workflows-shared`:

- **Pull Request Events**: Validates code on PRs to `develop` branch
- **Develop Push Events**: Builds, tags, and deploys to dev environment
- **Main Push Events**: Creates releases, deploys to staging, validates, and deploys to production

## Configuration

### Required Secrets

- `ECR_REPOSITORY_URL`: ECR repository URL
- `AWS_ROLE_ARN_SHARED`: IAM role for ECR access
- `AWS_ROLE_ARN_DEV`: IAM role for dev environment
- `AWS_ROLE_ARN_STAGING`: IAM role for staging environment
- `AWS_ROLE_ARN_PROD`: IAM role for production environment
- `AWS_REGION`: AWS region (optional, defaults to us-east-1)

### GitHub Environments

Configure the following environments in GitHub:
- `dev`: Automatic deployment (no protection rules)
- `staging`: Automatic deployment (no protection rules)
- `prd`: Requires manual approval (PCI compliance)

## Running Locally

```bash
# Install dependencies
go mod download

# Run tests
go test -v -coverprofile=coverage.out ./...

# View coverage
go tool cover -func=coverage.out

# Run linter (if golangci-lint is installed)
golangci-lint run

# Build
go build -o service .

# Run
./service serve
```

## Docker

```bash
# Build Docker image
docker build -t go-service .

# Run container
docker run -p 8080:8080 go-service
```

## CI/CD Flow

1. **Feature Branch → PR to develop**: Runs linters, security scans, unit tests, coverage validation
2. **develop branch (push)**: Security scan → Build & Tag → Deploy to Dev → Verification → Health Check
3. **main branch (push)**: Release → Build & Tag → Deploy to Staging → Validation → **Approval** → Deploy to Production

