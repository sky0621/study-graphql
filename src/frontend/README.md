# frontend

> My ace Nuxt.js project

## Build Setup

``` bash
# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev

# build for production and launch server
$ yarn build
$ yarn start

# generate static project
$ yarn generate
```

For detailed explanation on how things work, check out [Nuxt.js docs](https://nuxtjs.org).

## yarn

### type script codegen

<pre>
yarn codegen
</pre>

## Dockerfile

### build

<pre>
sudo docker login
sudo docker build . -t sky0621dhub/study-graphql-frontend:v0.1
</pre>

### start

<pre>
sudo docker run sky0621dhub/study-graphql-frontend:v0.1
</pre>

### push to docker-hub

<pre>
sudo docker login
sudo docker push sky0621dhub/study-graphql-frontend
</pre>

### pull from docker-hub

<pre>
sudo docker pull sky0621dhub/study-graphql-frontend:v0.1
</pre>
