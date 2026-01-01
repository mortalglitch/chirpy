# Chirpy Social Media Server
This is small project based off the boot.dev guide for building a microblogging server in Go.

## Requirements for build
- Goose
- Database software (I used postgres so some of the SQL queries are written for it)
- Go

## ENV setup
- DB_URL="postgres://username:@localhost:5432/chirpy?sslmode=disable" 
The key above is an example URL for accessing a local database when debugging.
- PLATFORM="dev"
When platform is set to dev admin access to the RESET url become available allowing quick testing by reseting the database.
- SIGNINGKEY="reallylongkey"
This is an example of a internal key used when establishing JWT's and other auth related functions.
- POLKA_KEY="adifferentlongkey"
This is a key from a fake payment processor to allow access to a dedicated webhook which allows users to "upgrade" their account.


## Endpoints
/app/ - Simple fileserver for checking access

/api/healthz - Health checkpoint to see if the server is running

/admin/reset - dev mode function allowing the database to be reset for testing

/admin/metrics - Shows how many times the app has been hit

/api/users - Handles all user functions
  - Create users with a POST request
  '{
    "email": "person@example.com",
    "password": "123456"
  }'
  - Update users email and password with a PUT request (JWT Token required as Authorization=Bearer [token] in the header)
  '{
    "email": "person@example.com",
    "password": "654321"
  }'

/api/chirps - Main social function endpoint
  - POST - Post a "chirp" (JWT Token required as Authorization=Bearer [token] in the header)
  '{
    "body": "What it is that it is"
  }'
  
  - GET - Get's a list of chirps
  - GET /api/chirps/{chirpID} - Get's a specific chirp
  - GET /api/chirps?author_id={userID} - Get's the chirps from a specific user
  - GET /api/chirps?sort=asc - Get all chirps in ascending order (default)
  - GET /api/chirps?sort=desc - Get all chirps in reverse chronological order
  - DELETE /api/chirps/{chirpID} - Deletes a specific chirp if owned by current user (JWT Token required as Authorization=Bearer [token] in the header)

/api/login - Logs a user in using their username and password and the system will provide a JWT and Refresh token
  '{
    "email": "person@example.com",
    "password": "654321"
  }'

/api/refresh - Checks for valid refresh token and provides updated JWT (Refresh Token required as Authorization=Bearer [token] in the header)

/api/revoke - Sets revoked status to existing valid token used with logging out (Refresh Token required as Authorization=Bearer [token] in the header)

/api/polka/webhooks - webhook for polka service which can upgrade users account to a red "premium" status (Polka Token required as Authorization=ApiKey [token] in the header)
