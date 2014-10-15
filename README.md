## Introduction

GoIgneous is a JSON data-store that responds to a RESTful API.  A production version of this service is available at ```igneous.joelf.me```.

## Usage

The RESTful APi provides four basic endpoints with varying return values.

---
#### POST /documents/new
Takes a single argument ```content``` which must be a valid JSON document, and returns an URL from which the document can be fetched using a ```GET``` request.

**Example**
```curl -u igneous:joel -X POST -d 'content="{\"abc\":\"123\"}"' http://igneous.joelf.me/documents/new``` returns ```http://igneous.joelf.me/documents/3```

---
#### GET /documents/:id
Fetches the JSON document with ID ```:id```.  The results of a successful ```POST``` to ```/documents/new```, or ```PUT``` to ```/documents/:id``` are of this form.

##### Example
```curl -u igneous:joel http://igneous.joelf.me/documents/3``` returns ```"{\"abc\":\"123\"}"```

---
#### PUT /documents/:id
Takes a single argument ```content``` which must be a valid JSON document, and updates the JSON document with ID ```:id``` while returning an URL from which the document can be fetched using a ```GET``` request

##### Example
```curl -u igneous:joel -X PUT -d 'content="{\"xyz\":\"789\"}"' http://igneous.joelf.me/documents/3``` returns ```http://igneous.joelf.me/documents/3```

---
#### DELETE /documents/:id
Deletes the JSON document with ID ```:id```.

##### Example
```curl -u igneous:joel -X DELETE http://igneous.joelf.me/documents/3```

---

## Error Handling

All endpoints will send a ```400 BAD REQUEST``` and provide textual details when an error occurs.

## Examples

The test script ```test/server_test.rb``` provides a comprehensive demonstration of the API.  It utilizes the Ruby gem ```rest-client```, a rest client and API written in Ruby.  All dependent gems can be installed using ```bundler``` and running ```bundle install``` from the ```test``` directory.  The script can be run with the ```-p``` flag to test against the production server.