# Elasticsearch Guides
Elasticsearch is a distributed, free and open search and analytics engine for all types of data, including textual, numerical, geo-spatial, structured, and unstructured. Elasticsearch allows you to store, search, and analyze huge volumes of data quickly and in near real-time and give back answers in milliseconds.

Once ES + Kibana is running successfully, open the Kibana dashboard using the following link
[http://localhost:5601](http://localhost:5601/). Now expand the options and go to the dev tools options to run ES queries as mentioned below.

## Document APIs
ES is accessible from a RESTful web service interface and uses schema-less JSON. SPAs can easily use the APIs directly on http://localhost:9200.

#### Create an index
```http
PUT /songs
```

#### Create a document / add data to the index
```http
POST /songs/_doc/1
{
  "title": "Summer of 69",
  "singer": "Bryan Adams",
  "album": "Reckless",
  "release": 1984,
  "genre": "Pop, Rock"
}
```
An index is automatically created if it does not exists.

#### Read a document
```http
GET /songs/_doc/1
```

#### Update a document
```http
POST /songs/_doc/1
{
  "title": "Numb",
  "singer": "Linkin Park",
  "album": "Meteora",
  "release": 2004,
  "genre": "Hip Hop, Rap, Pop"
}
```
Every update of the document will always create a new version. We can also create versions explicitly.

#### Delete a document
```http
DELETE /songs/_doc/1
```

#### Versioning
```http
POST /songs/_doc/1?version=7&version_type=external
{
  "title": "Summer of 69",
}
```

#### Operation Type
Operation type CREATE helps to avoid the overwriting of an existing document. This will throw an error if the document already exists.
```http
POST /songs/_doc/1?op_type=create
{
  "title": "Summer of 69",
}
```

#### Automatic ID Generation
When ID is not specified in index operation, then Elasticsearch automatically generates an ID for that document.
```http
POST /songs/_doc/
{
  "title": "Numb",
}
```

#### Reading only specific fields of a document
```http
GET /songs/_doc/1?_source_includes=title,album
```

## Search APIs

#### Multi-Index Search
We can search on all indexes or specific multiple indexes.
```http
GET /_all/_search?q=album:Reckless
GET /songs,music/_search?q=album:Reckless 
```

#### Multi-Type Search
Search on all fields of multiple types.
```http
GET /songs/_search?q=Bryan 
```

#### Wildcards
```http
GET /son*/_search?q=release:1984 
```