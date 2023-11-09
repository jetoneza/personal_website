# My Personal Website

## Prerequisites

1. Install [Go](https://go.dev/doc/install).
2. Install [Node.js](https://nodejs.org/en/).
3. Create a `.env` file in the project root and add the environment variables.

```
PORT=3000
SVELTE_PORT=3003
BACKEND_URL=http://127.0.0.1:3000
APP_NAME="Development Website"
TOKEN_KEY=<your_token_key_here>
AUTH_USER=<auth_user>
AUTH_PASSWORD=<auth_password>
```

## Running the application

1. Run the backend server

```
go run ./cmd/build run
```

2. Open another terminal and run the frontend

```
go run ./cmd/build run:web
```

3. Go to `http://localhost:5173/`

## Adding content

1. To add content, an admin user must log in. Use the `POST /api/v1/auth/register` endpoint to register a user. This uses basic authentication. Ensure you use the values provided in the environment variables `AUTH_USER` and `AUTH_PASSWORD`.

```
curl -X POST http://localhost:3000/api/v1/auth/register
   -H "Content-Type: application/x-www-form-urlencoded"
   -d "name=Admin&email=test@email.com&password=testpassword"
   --user "<auth_user>:<auth_password>"
```

2. Go to `http://localhost:5173/admin/login`. Login the newly created user.
3. Click create post.
4. Populate the form and add a markdown text in the content textarea.
5. After creating the post, you will be redirected to the Posts page. Click view.
