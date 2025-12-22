# =============== BASE ===============
FROM liquibase/liquibase:4.33-alpine AS base

WORKDIR /liquibase/changelog

COPY database/changelog.xml .
COPY database/migrations/ ./migrations/

ENV LIQUIBASE_CHANGELOG_FILE=changelog.xml

FROM base AS runner

ENV \
  DATABASE_HOST=postgres \
  DATABASE_PORT=5432 \
  DATABASE_USERNAME=postgres \
  #DATABASE_PASSWORD=postgres \
  DATABASE_DATABASE=postgres \
  DATABASE_CHANGELOG=_changelog \
  DATABASE_CHANGELOG_LOCK=_changelog_lock \
  LIQUIBASE_TAG=latest

ENTRYPOINT ["bash", "-c"]
CMD ["liquibase \
  --url=jdbc:postgresql://${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_DATABASE} \
  --changeLogFile=${LIQUIBASE_CHANGELOG_FILE} \
  --username=${DATABASE_USERNAME} \
  --password=${DATABASE_PASSWORD} \
  --databaseChangeLogTableName=${DATABASE_CHANGELOG} \
  --databaseChangeLogLockTableName=${DATABASE_CHANGELOG_LOCK} \
  updateToTag ${LIQUIBASE_TAG}"]
