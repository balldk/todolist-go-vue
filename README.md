# TodoList

A todolist webapp built on top of Vue.js (Frontend) and Go (Backend)

### Configuration

- Manually run `/server/database/model.sql` script (Mysql)
- Create `.env` file inside `/server` folder with this structure
```.env
# Secret key for server's data encryption
SECRET= 

# Database Source (e.g. user:password@tcp(hostname:3306)/schema_name)
DB_SOURCE= 

# Server Port (e.g. 3000)
PORT=
```

### Usage
- Run Following script in terminal
```
cd client
yarn build
cd ../server
go run main.go
```