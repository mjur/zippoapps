# ZippoApps Task
## description
  The app exposes an API endpoint to retirieve a main sku based on the package.
  The main sku returned is chosen randomly from a list of valid main skus based on 
  the percentile_min and percentile_max properties relation to the random number selected.
  If the are multiple valid skus for the random number the first one found is returned.
  The skus are stored in a postgres db and migrations and seeding is done on startup


## Enviroment variables
| Variable          |Description                            |
| ------------------|:-------------------------------------:|
| HOST              | The host to run the app on            |
| PORT              | The port to run the app on            |
| TIMEOUT           | Timeout of the server                 |
| DATABASE_HOST     | The host the db runs on               |
| DATABASE_PORT     | The port the db runs on               |
| DATABASE_USERNAME | The db username                       |
| DATABASE_PASSWORD | The db password                       |
| DATABASE_NAME     | The db name                           |
| CACHE_TTL         | The TTL of the cache                  |



## Example request


```bash
  curl -X GET -H "X-Appengine-Country: US" -d @req.json http://localhost:8080/configuration/com.softinit.iquitos.mainapp
  ```

# Running the app

### Docker Compose
run 
 ```bash
  make run
  ```
  or 
 ```bash
   docker-compose up --build -d
  ```
To terminate the container run
```bash
  make stop
  ```
  or 
 ```bash
  docker-compose down
  ```
 ### Natively
 run 
 ```bash
  make run-local
  ```



# Running tests
### Docker
run
```bash
docker exec zippo-apps-task-app make test
```
or
```bash
docker exec zippo-apps-task-app go test -race -cover ./pkg/...
```
### Natively
run 
```bash
make test
```

# Deploying the app
Build the app using the provided Dockerfile
```bash
docker build -t <image-name> .
```
or 
```bash
make docker
```

Run the app in the orchestration system of your choice

## Dependencies
The app requires a Postgres instance to run 
To configure see the Enviroment variables section


# Proposed Architecture
## Requirements
  The app should be:
   - Low latency
   - Globaly spread
   - Scalable
   - We need to identify the country of origin of the request

## Deployment
Deploy the app on Google Cloud engine, which allows for out of the box scaling and 
identifying the country of origin using the X-Appengine-Country header in the request
The deployment should be done to multiple regions to account for the geographical distribution 
of the clients, which will cover the latency issues and spreading the app globaly.

The current store used is postgres and the dbs should either be replicated to ensure data consistency 
across regions. An alternative would be to manage the data separately for each region, 
but this leads to greater complexity of maangement.

We could also replace Postgres with a highly-availible data store such as Google Cloud Bigtable.
Note this would require a code change.


## Diagram

![Alt text](./diagram.png?raw=true "Diagram")


