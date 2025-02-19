FROM openjdk:17-jdk-slim

WORKDIR /app

CMD ["sh", "-c", "java -jar main.jar 2>&1 | awk '{ print strftime(\"[%Y-%m-%d %H:%M:%S]\"), $0 }' | tee -a /logs/java.log"]

EXPOSE 8081
