FROM mongo:6.0

CMD ["sh", "-c", "mongod --config /etc/mongod.conf | awk '{ print strftime(\"[%Y-%m-%d %H:%M:%S]\"), $0 }' | tee -a /logs/mongodb.log"]

EXPOSE 27017
