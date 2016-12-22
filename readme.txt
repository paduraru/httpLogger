Simple http logger
==================

Usage:

docker-compose up

Will start the logger on port 80

Sample usage
============

Let's say your app is making a request to example.com/test

Enter the following line in your /etc/hosts:

127.0.0.1 example.com

Start the logger. Any request that your app is making to that address will now go to the logger. The logger will print some request info and return 200.