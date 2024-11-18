# Firebase Admin Auth

## Description

This repo contains a simple auth service using Firebase Admin SDK.

App feature:

- Verify ID token

This simple app is an admin app for [firebase-auth](https://github.com/CC-04-24H2/firebase-auth) project.

This project written in [Go](https://go.dev).

## How to Run

1. Create or use existing [Firebase](https://console.firebase.google.com) project
2. Go to `Project settings > Service accounts`
3. On `Firebase Admin SDK` choose `Generate new private key`
4. Place the private key downloaded on `config` directory of this project
5. Create `.env` on the root of this project by copy-paste `.env.example` provided (more on **[Environment Variables](#environment-variables)** section)
6. Install dependencies
   ```bash
   go mod download
   go mod tidy
   ```
7. Run the project
   ```bash
   go run main.go
   ```

### Environment Variables

- `GOOGLE_APPLICATION_CREDENTIALS`: relative/absolute path to the service account JSON file (e.g. `config/service-account.json`)
- `ID_TOKEN`: ID token to be verify, obtained from `user` object on client

## References

- [Admin Auth API](https://firebase.google.com/docs/auth/admin)
- [Admin SDK Setup](https://firebase.google.com/docs/admin/setup)
- [Admin SDK Reference](https://firebase.google.com/docs/reference/admin)
