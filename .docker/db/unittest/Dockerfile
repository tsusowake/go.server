FROM --platform=linux/arm64/v8 mysql:8.0.30

ENV MYSQL_DATABASE=unittest
ENV MYSQL_USER=test
ENV MYSQL_PASSWORD=test
ENV MYSQL_ROOT_PASSWORD=root.local
ENV MYSQL_ALLOW_EMPTY_PASSWORD=""
ENV MYSQL_RANDOM_ROOT_PASSWORD=""
ENV TZ=Asia/Tokyo

EXPOSE 3307

CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_general_ci", "--default-time-zone=+09:00"]
