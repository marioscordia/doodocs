# File handler

This is a REST API that allows user to see zip file's detailed information, compress files into a zip file and also sending file to provided emails.

## Technologies

Project was built using Fiber framework for handling server requests, standard library for working with zip files and jordan-wright/email for sending by email.

## Notes

    These notes are in case if you want to run application on your computer.

    1. You have to create **config.json** file in **./config/** directory with email and password fields.
    2. Since gmail server is used for email sending, you need to go into your gmail account and create application password, which must be put to password field accordingly.
    3. Lastly, you need to uncomment everything in **./config/config.go** file and also it's calling in **./cmd/main.go** file in order to set env. variables.

## Running application

Without Docker:

```bash
make start
```

With Docker:

```bash
make build
```

```bash
make run
```
