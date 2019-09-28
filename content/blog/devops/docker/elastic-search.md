---
title: "Elastic Search"
description: "Elastic Search"
date: 2019-05-16T12:55:33+08:00
draft: false
toc: true
---

`docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:5`

## 可视化

docker run -p 9100:9100 mobz/elasticsearch-head:5

kibana

docker run -p 5601:5601 docker.elastic.co/kibana/kibana:5.6.16

docker run -e "XPACK_GRAPH_ENABLED=false" --link goofy_curie:elasticsearch -e "ELASTICSEARCH_URL=http://localhost:9200" -p 5601:5601  docker.elastic.co/kibana/kibana:5.6.16

docker run --name kibana  --link goofy_curie:elasticsearch -e "ELASTICSEARCH_URL=http://10.222.0.1:9200" --net default -p 5601:5601 kibana:5

docker run -p 5601:5601 -p 9200:9200 -p 5044:5044  sebp/elk:es241_l240_k461