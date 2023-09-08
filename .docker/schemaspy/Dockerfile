FROM --platform=linux/arm64/v8 openjdk:8u212-jdk-alpine

ENV MYSQL_VERSION=8.0.30
ENV SCHEMASPY_VERSION=6.2.4
ENV DRIVER_URL https://repo1.maven.org/maven2/mysql/mysql-connector-java/${MYSQL_VERSION}/mysql-connector-java-${MYSQL_VERSION}.jar
ENV APP_URL https://github.com/schemaspy/schemaspy/releases/download/v${SCHEMASPY_VERSION}/schemaspy-${SCHEMASPY_VERSION}.jar

WORKDIR /app
COPY schemaspy.properties .

RUN apk --update add graphviz ttf-dejavu && \
    apk --update add --virtual .builddep tzdata wget libressl && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    mkdir /drivers && \
    wget -O /drivers/mysql-connector-java.jar ${DRIVER_URL} && \
    wget -O schemaspy.jar ${APP_URL} && \
    apk del .builddep && \
    rm -rf /var/cache/apk/*
