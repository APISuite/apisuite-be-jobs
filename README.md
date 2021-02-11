# API Suite Backend Jobs

This project is part of the API Suite Core and contains some recurrent jobs.

There is a configuration sample file in .env.sample.


## Jobs
- User activation cleaner:
   - defaults to run every 30 minutes
   - removes users that were never activated and were created more than 12h since the job execution time
- Password recovery cleaner
   - defaults to run every 30 minutes
   - removes password recovery tokens that were created more than 2h since the job execution time


## Environment variables

```
APISUITE_JOBS_DB="postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@POSTGRES_ADDRESS:${POSTGRES_PORT_INTERNAL}/${POSTGRES_DB}?sslmode=disable"
APISUITE_JOBS_ACTV_CRON=*/30+*+*+*+*
APISUITE_JOBS_ACTV_TTL=12
APISUITE_JOBS_RECOV_CRON=*/30+*+*+*+*
APISUITE_JOBS_RECOV_TTL=2
```

## Installing

Docker images are available in our [DockerHub](https://hub.docker.com/r/cloudokihub/apisuite-be-jobs).

Every new image is tagged with:
- commit hash
- latest
- semantic version from `package.json`

Depending on your goals, you could use a fixed version like `1.0.4` or 
`latest` to simply get the most recent version every time you pull the image.

## Development

- Commits should follow [conventional commits](https://www.conventionalcommits.org) spec