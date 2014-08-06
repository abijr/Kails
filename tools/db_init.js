// A file to initialize mongodb database.
// Creates database and collections, and adds initial documents,
// also creates indexes.

use kails;

db.users.ensureIndex({"name": 1}, {"unique": true});
db.users.ensureIndex({"email": 1}, {"unique": true});

default_user = {
    name: "default_user",
    email: "default@email.com",
    pass: "hashed_password",
    salt: "salt_for_hash"
    lang: "spanish",
    since: new Date(),
    levels: {},
}
