# backend

## gqlgen

### init
<pre>
go run github.com/99designs/gqlgen init
</pre>

### update
<pre>
go run github.com/99designs/gqlgen -v
</pre>

## Dockerfile

### build

<pre>
sudo docker build . -t sky0621dhub/study-graphql-backend:v0.1
</pre>

### start

<pre>
sudo docker run sky0621dhub/study-graphql-backend:v0.1
</pre>

### push to docker-hub

<pre>
sudo docker login
sudo docker push sky0621dhub/study-graphql-backend
</pre>

### pull from docker-hub

<pre>
sudo docker pull sky0621dhub/study-graphql-backend:v0.1
</pre>
