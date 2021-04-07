FROM node:alpine

# set working directory
WORKDIR /web

# add `/app/node_modules/.bin` to $PATH
ENV PATH /web/node_modules/.bin:$PATH

# copy and install react dependencies
COPY ./web/package.json ./
COPY ./web/package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

# copy everything else
COPY ./web .

# start app. not necessary since Go serves the react files
CMD ["npm", "start"]