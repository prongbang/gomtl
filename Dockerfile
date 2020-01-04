FROM mvertes/alpine-mongo

CMD ["mongod", "--bind_ip=0.0.0.0","--smallfiles", "--logpath=/dev/null"]