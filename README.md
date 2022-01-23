# go-gcs-to-pg-data-importer
Data import from GCS to PostgreSQL by golang

# Setup
## Run docker containers
1. Run docker-compose

```bash
$ docker-compose up
```

2. connect GCS fake server:

```bash
$ docker-compose exec app curl --insecure https://gcs:4443/storage/v1/b/tables/o
#=> {"kind":"storage#objects","items":[{"kind":"storage#object","name":"users.json","id":"tables/users.json","bucket":"tables","size":"63","contentType":"application/json","crc32c":"VuKJWg==","md5Hash":"2lbYdXcwrlOlP6Q7efN/pQ==","timeCreated":"2022-01-23T08:05:22.41268Z","timeDeleted":"0001-01-01T00:00:00Z","updated":"2022-01-23T08:05:22.412714Z","generation":"0"}]}
```

3. connect PostgreSQL server:

```bash
$ docker-compose exec app curl http://postgresql:5432
#=> curl: (52) Empty reply from server
```

## Database
Create inital data:

```bash
$ dc exec postgresql psql -f /init/create_users.sql -U postgres
#=> CREATE TABLE
```

# Docker images

## GCS images

Using [fake-gcs-server](https://github.com/fsouza/fake-gcs-server)
