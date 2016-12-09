# Adept
The Adept platform is an extensible gamified quiz application intended to use extrinsic reward to foster intrinsic motivation and increase participation toward the goal of topic mastery.

## Requirements
Golang 1.5+ (in order to modify the code), MySQL

## Bootstrapping
To run locally, modify the database values in the config file at /bin/config/bootstrap-config.json and then run bin/bootstrap or bin/bootstrap.exe.  This will import the courses, quizzes and questions from /bin/raw.  Then move the resulting config.json file into /config.  Run the web server with ./bin/osx/adept, ./bin/linux/adept or bin/win/adept.exe, depending on platform.