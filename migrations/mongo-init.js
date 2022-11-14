db.createCollection("accounts");

db.accounts.createIndex({email: 1}, {unique: true})

db.accounts.insertMany([
    {
        email: "alex",
        password: "$2a$14$.lAgcS4UiAiJKiD9kbvlHObowOCdV5mEmreUvG/HewCVlm0heqNN6" //kutsenko
    },
    {
        email: "roma",
        password: "$2a$14$Q1NW5.b/kWan0Y.m1Zjy5uApp9PtuTKDOdmppL.5V8/j/AczOLlHq" //chach
    }
])
