# obs-with-elastic
## Observability with [Elastic](https://www.elastic.co) (APM, Beats, Kibana)

- Monitoring a Kotlin Application with APM
- Nginx log collection with Filebeat
- System metrics (CPU, Memory usage) with Metricbeat
- Services availability with Heartbeat
- Visualization with Kibana

Docker compose has the following services: Elasticsearch, Kibana, APM and Beats agents. There's also two Applications (Kotlin/Go) and an Nginx server used as a proxy.
> An APM agent is [attached](https://www.elastic.co/guide/en/apm/agent/java/master/index.html) to the application to monitor it.

   
- To create and start all containers:
 `docker-compose up -d`