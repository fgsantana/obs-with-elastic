FROM openjdk:11
ARG JAR_FILE=build/libs/*T.jar
COPY ${JAR_FILE} server.jar
ENTRYPOINT ["java","-jar","/server.jar"]