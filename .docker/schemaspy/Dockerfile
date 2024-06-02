FROM --platform=linux/arm64/v8 openjdk:11-jre-slim

ENV SCHEMASPY_VERSION=6.2.4
ENV POSTGRES_DRIVER_VERSION=42.7.3

WORKDIR /app
COPY .docker/schemaspy/schemaspy.properties .

RUN apt-get update && \
    apt-get install -y wget graphviz && \
    apt-get clean

RUN wget https://github.com/schemaspy/schemaspy/releases/download/v${SCHEMASPY_VERSION}/schemaspy-${SCHEMASPY_VERSION}.jar -O /usr/local/bin/schemaspy.jar
RUN wget https://jdbc.postgresql.org/download/postgresql-${POSTGRES_DRIVER_VERSION}.jar -O /usr/local/bin/postgresql.jar

ENTRYPOINT ["java", "-jar", "/usr/local/bin/schemaspy.jar"]
