# Concepts Lab
This is a playground repository to learn new technologies and develop simple POCs related to implementation of robust and scalable services architecture.

#### Contents
1. *elk-logging* : Implement ELK stack to centralize logging for services, ingest Kafka messages and CSV data into searchable indices. Create dashboards to visualize CSV data and analyze data from log and Kafka messages.
2. *elk-monitoring* : Implement ELK + Metricbeat for enabling metric collection and visualization for system, Docker and services running inside containers such as Kafka.
3. *prom-monitoring* : Implement a Grafana dashboard that uses Prometheus metrics to monitor HTTP REST APIs, providing insights into request rates, response times, errors, and status codes for tracking performance and identifying issues.
4. *kafka-lag-monitoring* : A Dockerized Kafka monitoring stack featuring KRaft mode, Prometheus, Grafana, and a Go-based lag simulator for real-time consumer lag visualization.