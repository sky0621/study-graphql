# ###################################################################
# step 1: build
# ###################################################################
FROM node:13.6.0 as build-step

# in container
WORKDIR /frontend

# from host-pc to container
COPY . .

# in container
RUN yarn install
RUN yarn build

# ###################################################################
# step 2: launch
# ###################################################################
FROM node:13.6.0

# in container
WORKDIR /frontend

# from host-pc to container
COPY --from=build-step /frontend/.nuxt ./.nuxt
COPY --from=build-step /frontend/package.json ./
COPY --from=build-step /frontend/node_modules ./node_modules

# in container
ENV HOST 0.0.0.0
ENV PORT 3000
EXPOSE 3000
ENTRYPOINT yarn start
