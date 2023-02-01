# Quote Generator Documentation

## **Create quotes**
1. Method: ```POST```
2. Endpoint: ```https://quote-generator-iks2.onrender.com/api/```
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



## **Get random quotes**
1. Method : ```GET```
2. Endpoint: ```https://quote-generator-iks2.onrender.com/api/```
3. Params: 
    * *number* (required) : int
3. Result: 
```
{
    "message": "Get random quotes successfully",
    "error": false,
    "data": {
        "quotes": [
            {
                "text": "Genius is one percent inspiration and ninety-nine percent perspiration.",
                "author": "Thomas Edison"
            }
        ]
    }
}
```