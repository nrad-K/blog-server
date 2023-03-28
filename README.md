# Blog API

## 概要

この API は、ブログ投稿のための RESTful なエンドポイントを提供します。ユーザーは、ブログの作成、編集、削除、および閲覧を行うことができます。

## 実行方法

### ディレクトリの移動
```
cd [blog-serverのpath] 
```

### golang環境構築
※ asdfが入っていない場合は、インストールを行う
【参考】 [asdf](https://asdf-vm.com/)
```
asdf install
```

### .env ファイルの用意

以下XXXXの値を任意の値を設定してください

```
ROOTUSER=XXXX
ROOTPASSWD=XXXX
DATABASE=blog_db
USERNAME=XXXX
USERPASSWD=XXXX
DB_CONTAINER_NAME=XXXX
HOSTNAME=localhost:3306
APP_CONTAINER_NAME=XXXX
DB_PORT=3306
APP_PORT=8080
```

### Run Applications

```
docker-compose up -d
```

###

## エンドポイント

## 記事の詳細を取得

### **`GET /articles/{article_id}`**

Response body

```json
{
  "article_id": 1,
  "title": "First Title",
  "contents": "First Content",
  "user_name": "John Smith",
  "like": 2,
  "comments": [
    {
      "comment_id": 1,
      "article_id": 1,
      "message": "First Message",
      "created_at": "YYYY-MM-DD HH:mm:ss"
    },
    {
      "comment_id": 2,
      "article_id": 1,
      "message": "Second Message",
      "created_at": "YYYY-MM-DD HH:mm:ss"
    }
  ],
  "created_at": "YYYY-MM-DD HH:mm:ss"
}
```

## 記事の一覧を取得

### **`GET /articles/list`**

Response body

```json
[
  {
    "article_id": 1,
    "title": "First Title",
    "contents": "First Content",
    "user_name": "John Smith",
    "like": 0,
    "comments": null,
    "created_at": "YYYY-MM-DD HH:mm:ss"
  },
  {
    "article_id": 2,
    "title": "Second Title",
    "contents": "Second Content",
    "user_name": "John Smith",
    "like": 3,
    "comments": null,
    "created_at": "YYYY-MM-DD HH:mm:ss"
  }
]
```

## 記事を投稿する

### **`POST /articles`**

Request body

```json
{
  "title": "First Title",
  "contents": "First Content",
  "user_name": "John Smith"
}
```

Response body

```json
{
  "article_id": 3,
  "title": "First Title",
  "contents": "First Content",
  "user_name": "John Smith",
  "like": 0,
  "comments": null,
  "created_at": "YYYY-MM-DD HH:mm:ss"
}
```

## いいねをする

### **`POST /articles/like`**

Request body

```json
{
  "article_id": 3,
  "title": "First Title",
  "contents": "First Content",
  "user_name": "John Smith"
}
```

Response body

```json
{
  "article_id": 3,
  "title": "First Title",
  "contents": "First Content",
  "user_name": "John Smith",
  "like": 1,
  "comments": null,
  "created_at": "YYYY-MM-DD HH:mm:ss"
}
```

## 記事を更新する

### **`PUT /articles`**

Request body

```json
{
  "article_id": 1,
  "title": "First Title",
  "contents": "Update Content",
  "user_name": "John Smith"
}
```

Response body

```json
{
  "article_id": 1,
  "title": "First Title",
  "contents": "First Content",
  "user_name": "John Smith",
  "like": 0,
  "comments": null,
  "created_at": "YYYY-MM-DD HH:mm:ss"
}
```

## 記事を削除する

### **`DELETE /articles/{article_id}`**

Response body

```json
{
  "article_id": 3
}
```

## コメントをする

### **`POST /comment`**

Request body

```json
{
  "article_id": 1,
  "message": "First Comment"
}
```

Response body

```json
{
  "comment_id": 13,
  "article_id": 1,
  "message": "First Message",
  "created_at": "YYYY-MM-DD HH:mm:ss"
}
```
