# =============== BASE ===============
FROM liquibase/liquibase:4.33-alpine AS base

WORKDIR /liquibase/changelog

COPY database/changelog.xml .
COPY database/migrations/ ./migrations/

ENV LIQUIBASE_CHANGELOG_FILE=changelog.xml

FROM base AS runner

ENV \
  API_DATABASE_HOST=postgres \
  API_DATABASE_PORT=5432 \
  API_DATABASE_USERNAME=postgres \
  #API_DATABASE_PASSWORD=postgres \
  API_DATABASE_DATABASE=postgres \
  API_DATABASE_CHANGELOG=_changelog \
  API_DATABASE_CHANGELOG_LOCK=_changelog_lock \
  LIQUIBASE_TAG=latest

ENTRYPOINT ["bash", "-c"]
CMD ["liquibase \
  --url=jdbc:postgresql://${API_DATABASE_HOST}:${API_DATABASE_PORT}/${API_DATABASE_DATABASE} \
  --changeLogFile=${LIQUIBASE_CHANGELOG_FILE} \
  --username=${API_DATABASE_USERNAME} \
  --password=${API_DATABASE_PASSWORD} \
  --databaseChangeLogTableName=${API_DATABASE_CHANGELOG} \
  --databaseChangeLogLockTableName=${API_DATABASE_CHANGELOG_LOCK} \
  updateToTag ${LIQUIBASE_TAG}"]
