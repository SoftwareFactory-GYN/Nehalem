# Setup instructions for local development 

## Required software: 
- go (https://golang.org/doc/install)
- dep (https://golang.github.io/dep/docs/installation.html)
- Docker (https://docs.docker.com/install/)

--- 

### 01. Run inside the desired package 'rest_api, web_app, ect.'
`dep ensure`

### 02. Start the cassandra cluster:
` docker run --name dev_cassandra -d -p 9042:9042 -p 7000:7000 -p 7001:7001 -p 7199:7199 -p 9160:9160 cassandra`

### Wait 30s for db to start 

### 03. Then start the cli session:
` docker run -it --link dev_cassandra:cassandra --rm cassandra cqlsh cassandra`

### 04. Run the following commands in the db cli session to create database and tables:
```sql
CREATE KEYSPACE nehalem WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};
USE nehalem;
CREATE TABLE users(
    id UUID PRIMARY KEY,
    username text,
    password text,
 ); 
```

---
# TODO List:

- [x] Create a salt and hash function [See this link for more help](https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72)
- [x] Implement Token auth with user via POST on the Login Handler
- [x] Create a utils package 
- [x] Create a program to seed the database with users in the Utils package
- [ ] Implement client login and home page
- [ ] Implement repository fetch list on rest_api
- [ ] Implement client repository list 
- [ ] Implement repository creation on rest_api
- [ ] Implement repository creation on client
- [ ] Implement repository download
- [ ] Implement client list of commits in repository
- [ ] Implement Docker run lint in files in specific commit
- [ ] Implement store of lint for specific commit
- [ ] Implement client statistics page for specific commit
- [ ] Implement client statistics for commit ranges, ie last 10 commits, last 50 commits, ect
  


