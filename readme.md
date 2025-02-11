# 🚀 TODOアプリ (Go + PostgreSQL + Solid.js)

このプロジェクトは、**Go (Golang) + PostgreSQL + Solid.js** を使用した **TODO管理アプリ** です。  
バックエンドは **Go (Gin) + GORM**、フロントエンドは **Solid.js** で構築されます。  
データベースには **PostgreSQL** を使用し、`docker-compose` で管理します。  

---

## **📌 環境**
- **バックエンド:** Go (Gin, GORM)
- **データベース:** PostgreSQL
- **フロントエンド:** Solid.js
- **認証:** (未定 / JWT / Firebase など)
- **API:** REST API

---

## **📍 1. 環境構築手順**
### **✅ 1.1 リポジトリをクローン**
```bash
git clone https://github.com/your-repo/todo-app.git
cd todo-app
```

### **✅ 1.2 Docker で PostgreSQL を起動**
```bash
cd docker
docker-compose up -d
```
📌 **データベース `mydb` が `localhost:5432` で起動します。**

🔹 **コンテナの状態を確認**
```bash
docker ps
```

🔹 **PostgreSQL に接続 (手動確認)**
```bash
docker exec -it postgres_container psql -U user -d mydb
```

---

## **📍 2. バックエンドのセットアップ**
### **✅ 2.1 Go の環境構築**
```bash
cd backend
go mod tidy
```

### **✅ 2.2 Go バックエンドの起動**
```bash
go run main.go
```
📌 **成功すると `http://localhost:8080/` で API が起動！**

🔹 **動作確認**
```bash
curl -X GET "http://localhost:8080/"
```
✅ **レスポンス:**
```json
{"message": "API is running!"}
```

---

## **📍 3. フロントエンドのセットアップ**
(準備中)

---

## **📍 4. API エンドポイント**
| HTTP | エンドポイント | 説明 |
|------|--------------|------|
| **GET** | `/tasks` | **全タスクの取得** |
| **GET** | `/tasks/{task_id}` | **タスクの詳細取得** |
| **POST** | `/tasks` | **新規タスクの作成** |
| **PUT** | `/tasks/{task_id}` | **タスクの更新** |
| **DELETE** | `/tasks/{task_id}` | **タスクの削除** |
| **GET** | `/tasks/{task_id}/subtasks` | **タスクのサブタスク取得** |
| **POST** | `/tasks/{task_id}/subtasks` | **サブタスクの追加** |
| **DELETE** | `/tasks/{task_id}/subtasks/{subtask_id}` | **サブタスクの削除** |
| **GET** | `/tasks/{task_id}/comments` | **タスクのコメント取得** |
| **POST** | `/tasks/{task_id}/comments` | **コメントの追加** |
| **DELETE** | `/tasks/{task_id}/comments/{comment_id}` | **コメントの削除** |

📌 **詳細は `docs/api-spec.md` に記載予定。**

---

## **📍 5. DB テーブル (GORM)**
| テーブル名 | 説明 |
|-----------|----------------|
| `tasks` | タスク管理 |
| `users` | ユーザー管理 |
| `comments` | タスクへのコメント |
| `task_subtasks` | 親タスクとサブタスクの紐付け |
| `task_comments` | タスクとコメントの紐付け |

---

## **📍 6. 開発・ビルド**
### **✅ Go バックエンドのビルド**
```bash
go build -o backend-app main.go
./backend-app
```

### **✅ フロントエンドのビルド (準備中)**
```bash
cd frontend
npm install
npm run dev
```

---

## **📍 7. よくある問題**
❌ **データベースに接続できない**
```bash
docker-compose restart
```
🔹 `docker ps` でコンテナの状態を確認

❌ **Go の依存関係エラー**
```bash
go mod tidy
```

---

## **📍 8. 今後のアップデート**
- [ ] **フロントエンド (Solid.js) の開発**
- [ ] **認証 (JWT or Firebase Auth) の実装**
- [ ] **タスクのステータス履歴管理**
- [ ] **リアルタイム更新 (WebSocket / SSE)**
- [ ] **CI/CD パイプラインの導入**

---

## **📜 ライセンス**
MIT License

---

## **🚀 次のアクション**
1️⃣ **この `README.md` で OK か？ 追記・修正したい点はあるか？**  
2️⃣ **`Task` モデルを作成し、`GORM AutoMigrate()` を適用する？**  
3️⃣ **`POST /tasks` の実装に進む？**
