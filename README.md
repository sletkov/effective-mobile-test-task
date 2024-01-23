# HTTP server

## Run application

```sh
make build && make run
```

## Description

### Methods

---

#### Users

- ``GET`` ``params`` ``/api/v1/users`` ``Getting users with filters and limit``

| Name                 | Type   | Description                              |     Constraint                    |
|----------------------|--------|------------------------------------------|-----------------------------------|
| name                 | string | url param for user name                  | 1<=len<=255, Alpha                |
| surname              | string | url param for user nameuser surname      | 1<=len<=255, Alpha                |
| patronymic           | string | url param for user nameuser patronymic   | 1<=len<=255, Alpha                |
| age_from             | int    | url param for user nameuser min age      | >=1, <=100                        |
| age_to               | int    | url param for user nameuser max age      | >=1, <=100                        |
| gender               | string | url param for user nameuser gender       | "male" or female                  |
| nationality          | string | url param for user nameuser nationality  | 2<=len<=2, Alpha                  |
| limit                | int    | url param for user nameuser limit        | >=1, <=50                         |

**Request**

```
```

**Response**

```
[
    ...
    {"id": __, "name": __, "surname": __, "patronymic": __, "age": __, "gender":__, "nationality": __}
    ...
]
```


- ``DELETE`` ``/api/v1/users/{id}`` ``Deleting user by id``

| Name                 | Type   | Description                              |     Constraint                    |
|----------------------|--------|------------------------------------------|-----------------------------------|
| id                   | string | user id                                  | required, >0                      |


**Request**

```
```

**Response**

```
```


- ``PATCH`` ``body`` ``/api/v1/users/{id}`` ``Updating user``

| Name                 | Type   | Description                              |     Constraint                    |
|----------------------|--------|------------------------------------------|-----------------------------------|
| name                 | string | url param for user name                  | 1<=len<=255, Alpha                |
| surname              | string | url param for user nameuser surname      | 1<=len<=255, Alpha                |
| patronymic           | string | url param for user nameuser patronymic   | 1<=len<=255, Alpha                |
| age                  | int    | url param for user nameuser min age      | >0, <100                          |
| gender               | string | url param for user nameuser gender       | "male" or female,                 |
| nationality          | string | url param for user nameuser nationality  | 2<=len<=2, Alpha                  |

**Request**

```
{
    "name": "Ivan",
}
```

**Response**

```
```


- ``POST`` ``body`` ``/api/v1/users`` ``Creating user``

| Name                 | Type   | Description                              |     Constraint                    |
|----------------------|--------|------------------------------------------|-----------------------------------|
| name                 | string | user name                                | required,1<=len<=255, Alpha       |
| surname              | string | user surname                             | required, 1<=len<=255, Alpha      |
| patronymic           | string | user patronymic                          | 1<=len<=255, Alpha                |

**Request**

```
{
    "name": "Ivan",
    "surname": "Ivanov",
    "patronymic": "Ivanovich"
}
```

**Response**

```

```