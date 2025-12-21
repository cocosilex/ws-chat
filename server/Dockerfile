# ----- BUILD
FROM gradle:9.2.1-jdk25 AS build

WORKDIR /app

COPY . .

RUN gradle shadowjar -x test

EXPOSE 8080

ENTRYPOINT ["java", "-jar", "app/build/libs/app-all.jar"]
