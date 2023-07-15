# Golang, GORM and GIN Framework CRUD Application

### We use golang GIN framework and GORM for this project.

## Steps to run the project
1. Clone the repo
2. Run the command
``` go mod download ```
3. Rename the .env.example file to .env and change variable's value
4. Run the migration command 
``` go run migrate/migrate.go ```
5. Open Postman and send request using following routes

   POST: http://localhost:3000/posts (Create new post)

   GET: http://localhost:3000/posts (Fetch all post)

   GET: http://localhost:3000/posts/1 (Find/Edit a post)

   PUT: http://localhost:3000/posts/1 (Update Post)

   DELETE: http://localhost:3000/posts/1 (Delete Post)

   ## Happy Coding
