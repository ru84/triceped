# Contributing to triceped

Thank you for your interest in contributing to **triceped**! We welcome contributions from the community to help improve and enhance this project. This document provides guidelines and information to help you get started.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Building the Code](#building-the-code)
- [Contributing Code](#contributing-code)
  - [Fork the Repository](#fork-the-repository)
  - [Create a Branch](#create-a-branch)
  - [Make Changes](#make-changes)
  - [Commit Messages](#commit-messages)
  - [Push Changes](#push-changes)
  - [Submit a Pull Request](#submit-a-pull-request)
- [Contributing Issues](#contributing-issues)
  - [Bug Reports](#bug-reports)
  - [Feature Requests](#feature-requests)
- [Coding Guidelines](#coding-guidelines)
- [Code Review Process](#code-review-process)
- [Community Guidelines](#community-guidelines)
- [License](#license)

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: triceped is written in Go. Install [Go 1.23+](https://golang.org/dl/).
- **Git**: Version control system for cloning and contributing to the repository.
- **Make** (optional): For using the provided `Makefile` to simplify build tasks.

### Building the Code

1. **Clone the Repository**

   ```bash
   git clone https://github.com/ru84/triceped.git
   cd triceped
   ```

2. **Install Dependencies**

   triceped uses Go modules. Dependencies are managed automatically.

   ```bash
   go mod tidy
   ```

3. **Build the Project**

   **Using Makefile:**

   ```bash
   make build
   ```

   **Or directly with Go:**

   ```bash
   go build -o bin/triceped cmd/triceped/main.go
   ```

4. **Run Tests**

   ```bash
   make test
   ```

   **Or directly with Go:**

   ```bash
   go test ./...
   ```

## Contributing Code

We appreciate code contributions! Please follow these steps to contribute:

### Fork the Repository

Click the "Fork" button on the [triceped repository](https://github.com/ru84/triceped) page to create a copy under your GitHub account.

### Create a Branch

Create a new branch for your work:

```bash
git checkout -b feature/your-feature-name
```

### Make Changes

Implement your feature or bug fix, adhering to the [Coding Guidelines](#coding-guidelines).

### Commit Messages

- Use clear and descriptive commit messages.
- Follow the format: `<type>(<scope>): <subject>`

  **Example:**

  ```bash
  git commit -m "feat(converter): add support for resource groups"
  ```

  **Types:**

  - **feat**: A new feature
  - **fix**: A bug fix
  - **docs**: Documentation only changes
  - **style**: Changes that do not affect the meaning of the code (white-space, formatting, etc.)
  - **refactor**: A code change that neither fixes a bug nor adds a feature
  - **test**: Adding missing tests or correcting existing tests
  - **chore**: Changes to the build process or auxiliary tools and libraries

### Push Changes

Push your branch to your forked repository:

```bash
git push origin feature/your-feature-name
```

### Submit a Pull Request

1. Navigate to the original repository's [Pull Requests](https://github.com/ru84/triceped/pulls) page.
2. Click "New pull request".
3. Select your branch and create the pull request.
4. Fill out the template with details about your changes.

## Contributing Issues

### Bug Reports

If you encounter a bug, please open an issue and include:

- **Description**: A clear and concise description of the problem.
- **Steps to Reproduce**: Instructions to reproduce the issue.
- **Expected Behavior**: What you expected to happen.
- **Actual Behavior**: What actually happened.
- **Environment Details**: OS, Go version, triceped version.

[Open a Bug Report](https://github.com/ru84/triceped/issues/new?assignees=&labels=bug&template=bug_report.md&title=)

### Feature Requests

We welcome ideas for new features! When submitting a feature request, please include:

- **Problem Statement**: What problem does this feature solve?
- **Proposal**: Describe the solution you'd like.
- **Alternatives**: Any alternative solutions or features you've considered.

[Submit a Feature Request](https://github.com/ru84/triceped/issues/new?assignees=&labels=enhancement&template=feature_request.md&title=)

## Coding Guidelines

- **Style**: Follow Go's standard formatting conventions. Use `go fmt` before committing.
- **Linting**: Ensure your code passes lint checks. Use tools like `golint` and `go vet`.
- **Testing**: Write unit tests for new functionality and ensure all tests pass.
- **Documentation**: Update or add comments for public functions and exported types.

## Code Review Process

All submissions are reviewed before being merged. Here's what to expect:

1. **Review**: Project maintainers will review your pull request for correctness and compliance with guidelines.
2. **Feedback**: You may receive comments requesting changes.
3. **Approval**: Once approved, your changes will be merged into the main branch.
4. **Release**: Your contribution will be included in the next release.

## Community Guidelines

We are committed to fostering a welcoming and inclusive environment. By participating, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md).

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).

---

Thank you for contributing to triceped! Your efforts help make this project better for everyone.
