# Quote Generator Documentation
## **Create quotes**
1. Method: ```POST```
2. URL: ```/api/```
3. Body: 
```
[
    {
    "text": "Genius is one percent inspiration and ninety-nine percent perspiration.",
    "author": "Thomas Edison"
  },
  {
    "text": "You can observe a lot just by watching.",
    "author": "Yogi Berra"
  }
]

```
4. Result:
```
{
    "message": "Create quotes successfully",
    "error": false,
    "data": [
        {
            "text": "Genius is one percent inspiration and ninety-nine percent perspiration.",
            "author": "Thomas Edison"
        },
        {
            "text": "You can observe a lot just by watching.",
            "author": "Yogi Berra"
        }
    ]
}
```

## **Get a random quotes**
1. Method : ```GET```
2. URL: ```/api/```
3. Result: 
```
{
    "message": "Get random quote successfully",
    "error": false,
    "data": {
        "quote": {
            "text": "Genius is one percent inspiration and ninety-nine percent perspiration.",
            "author": "Thomas Edison"
        }
    }
}
```