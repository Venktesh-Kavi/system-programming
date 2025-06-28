## MongoDB Learnings

> Notes to capture learning on mongodb

### Embeds and One to Many Relationships

* When deciding relationships between objects in mongodb. Based on the querying patterns, choose embeds or typical (ref) based relationships.
* Embeds can be easily represented as follows in java:

```
class Publisher {
    private String publisherName;
    private String publisherId;
    private List<Book> books;
}

class Book {
    private String isbn;
    private String bookName;
}
```
* The above code embeds the books documents in the publisher document.
* Obviously this is an mongodb anti-pattern modelling, because the books document is a growing array.
* Document can be referenced via @DocRef and @DocumentReference constructrs from mongodb.
* The difference between the two is @DocRef allows referencing document from another collection via only the `_id` field.
* @DocumentReference allows referencing by field on the correspoding document and is latest construcut as per mongodb documentation.
* Depending on the querying patterns, thought needs to be put on growing arrays, embedding large documents and other patterns mentioned here. [MongoDB Schema Anti Patterns](https://www.mongodb.com/developer/products/mongodb/schema-design-anti-pattern-summary/)


