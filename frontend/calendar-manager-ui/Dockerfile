# Stage 1: build
FROM node:16 as build
WORKDIR /app
COPY package*.json ./
RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build --prod


# Stage 2: Serve
FROM nginx:alpine
COPY --from=build /app/dist/calendar-manager-ui /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
