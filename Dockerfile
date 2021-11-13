# ---- Base Node ----
FROM node:lts-alpine AS base
# install node
RUN apk add --update git build-base python3 python2
## Create a group and user
#RUN addgroup -S appgroup && adduser -S appuser -G appgroup
## Tell docker that all future commands should run as the appuser user
#USER appuser
# set working directory
WORKDIR /usr/src/app
# copy project file
COPY package.json package-lock.json ./

# ---- First stage: build things ----
FROM base AS build
#RUN npm install --only=production
# copy production node_modules aside
#RUN cp -R node_modules prod_node_modules
# install ALL node_modules, including 'devDependencies'
RUN npm install

COPY . .
# Run compilers, code coverage, linters, code analysis and testing tools
RUN npm run build

# ---- Second stage: release ----
FROM base as release

COPY --from=build /usr/src/app/dist ./dist
COPY --from=build /usr/src/app/config ./config
COPY --from=build /usr/src/app/node_modules ./node_modules

# Run the built application when the container starts.
EXPOSE 3000 8000
CMD ["npm", "run", "serve"]
#
