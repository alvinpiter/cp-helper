# cp-helper

`cp-helper` is a web service that that compares user submissions in Codeforces and AtCoder.

The front-end implementation can be seen here: https://github.com/teddypiter/code-compare

The website can be accessed here: https://boiling-tundra-04466.herokuapp.com/

## Endpoint

There are two endpoints served by the service:

### `POST /compare`

Request body:
```
{
    "online_judge": "codeforces",
    "handle": "alvinpiter",
    "rival_handle": "teddypiter",
    "filter": {
        "rating": {
            "minimum": 1000,
            "maximum": 2000
        },
        "tags": {
            "mode": "or",
            "values": ["implementation", "math"]
        }
    }
}
```

A couple of points to note regarding request body:
* There are 2 possible values for `online_judge`, they are `codeforces` and `atcoder`
* `online_judge`, `handle` and `rival_handle` are required
* There are only two possible `filter` at the moment, they are `rating` and `tags` filter
* `minimum` and `maximum` in `rating` filter are optional. If `minimum` is omitted, the minimum will be set as 0, and if `maximum` is omitted, the maximum will be set as 4000
* There are 2 possible values for `mode` in `tags` filter, they are `or` and `and`. `values` is an array of string. If the `mode` is `and`, the problem must have all the given tags. If the `mode` is `or`, the problem must have at least one of the given tags.

Response:
```
{
    "problems":[
        {
            "id":"830C",
            "name":"Bamboo Partition",
            "rating":2300,
            "tags":["brute force","data structures","implementation","math","number theory","sortings","two pointers"],
            "url":"https://codeforces.com/contest/830/problem/C"
        }
    ]
}
```

### `GET /codeforces-problem-tags`

Returns list of codeforces problem tags.
Response:
```
{
    "problem_tags": [
        "constructive algorithms",
        "divide and conquer",
        "dfs and similar",
        "data structures",
        "binary search",
        "2-sat",
        ...
    ]
}
```
