# Payload

Payload is a simple tool for working with production data in your local environment.

## What problem does it solve?

You're working with Cloud SQL in production, but you want to avoid connecting directly (or through cloud_sql_proxy) to your production database lest you accidentally run some destructive commands, or expose customer data. Both of which are considerd bad practice.

`payload` allows you to export the production database, trim customer data — such as by truncating `users` or running other SQL cleanup commands — before downloading the data to your machine and loading it into your local database for testing/local development.

Or you need to inspect real customer data, in which case you can use IAM to grant permission for a given engineer to a specific customer scope, which will not be cleaned up in the export.

This has some neat usecases:

- Test upcoming database migrations against real data
- Help debug customer issues without connecting to production
- Run analytics against production without interfering with real data or running replicas
- Create personal staging environments with real data so that you don't need to manage large collections of seed data that invariably becomes unmaintained

## Security

`payload` is built around IAM in Google Cloud. At the very least you will need to create a service account for it to run as which can export databases, write to Cloud Storage, and fetch objects.

To enable data scrubbing, permission to launch a spot instance is required, which will load your data into a temporary database, scrub certain data, then export the remaining data.

# Usage

```shell
NAME:
   payload - Do it in production, but locally.

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.1

DESCRIPTION:
   Payload is a simple tool for working with production data in your local environment.

COMMANDS:
   create   Dumps the database and stores it on Cloud Storage
   pull     Pulls the latest database dump, creating it if not already stored
   sync     Loads the latest database dump into your local database, pulling it if not already loaded
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

```
