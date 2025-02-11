# API 設計

## API 一覧

| No. | API機能No. | 種別 | API名 | 機能概要 |
| --- | --- | --- | --- | --- |
| 1 | [TASK-001](#task-001) | API | タスク一覧取得 | タスク一覧を取得する |
| 2 | [TASK-002](#task-002) | API | タスクを取得する | task_id に紐づくタスクを取得する |
| 3 | [TASK-003](#task-003) | API |  新規タスクの作成 | タスクを新規作成する |
| 4 | [TASK-004](#task-004) | API | タスクの更新 | タスクを更新する（全体） |
| 5 | [TASK-005](#task-005) | API | タスクの削除 | タスクを削除する |
| 6 | [TASK-006](#task-006) | API | サブタスクの追加 | タスクにサブタスクを追加する |
| 7 | [TASK-007](#task-007) | API | コメントの追加 | タスクにコメントを追加する |
| 8 | [TASK-008](#task-008) | API | コメントの編集 | タスクのコメントを編集する |
| 9 | [TASK-009](#task-009) | API | コメントの削除 | タスクのコメントを削除する |
| 10 | [USER-001](#user-001) | API | ユーザー一覧取得 | ユーザー一覧を取得する |
| 11 | [USER-002](#user-002) | API | ユーザーを取得する | user_id に紐づくユーザーを取得する |
| 12 | [USER-003](#user-003) | API | 新規ユーザーの作成 | ユーザーを新規作成する |
| 13 | [USER-004](#user-004) | API | ユーザーの更新 | ユーザーを更新する（全体） |
| 14 | [USER-005](#user-005) | API | ユーザーの削除 | ユーザーを削除する |
| 15 | [USER-006](#user-006) | API | ユーザーの権限変更 | ユーザーの権限を変更する |

## API 詳細

### TASK-001 

| API機能No. | TASK-001 |
| --- | --- |
| API名 | タスク一覧取得 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | タスク一覧を取得するAPI |
| METHOD | GET |
| URI | `/api/v1/tasks` | 


#### リクエスト

リクエスト例

1. 条件なしで全取得する

```bash
curl -X GET "http://localhost:8080/tasks"
```

2. タイトルで検索する

```bash
curl -X GET "http://localhost:8080/tasks?title=新しいタスク"
```

3. 期限日で検索する

```bash
curl -X GET "http://localhost:8080/tasks?due_date=2024-02-20"
```

4. サブタスクでないもののみ取得する

```bash
curl -X GET "http://localhost:8080/tasks?is_subtask=false"
```

#### レスポンス


| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| task_id | int | 10 | ○ | × | タスクのUUID |
| title | string | 255 | ○ | × | タスクのタイトル |
| description | string | 255 | × | × | タスクの説明 |
| priority | string | 255 | × | × | タスクの優先度 |
| status | string | 255 | ○ | × | タスクのステータス |
| reporter | object | 255 | ○ | × | タスクの報告者 |
| assignee | object | 255 | × | × | タスクの担当者 |
| created_at | date | 255 | ○ | × | タスクの作成日 |
| start_date | date | 255 | × | × | タスクの開始日 |
| due_date | date | 255 | × | × | タスクの期限日 |
| story_points | int | 10 | × | × | タスクのストーリーポイント |
| category | string | 255 | × | × | タスクのカテゴリ |
| is_subtask | boolean | 1 | × | × | サブタスクかどうか |
| sub_tasks | array | 255 | × | ○ | サブタスクのリスト |
| comments | array | 255 | × | ○ | コメントのリスト |

レスポンス例

```json
{
  "task_id": 1,
  "title": "新しいタスク",
  "description": "新しいタスクの説明",
  "priority": "高",
  "status": "未着手",
  "reporter": {
    "id": "user-1",
    "name": "山田 太郎",
    "icon": "https://example.com/icons/user1.png"
  },
  "assignee": {
    "id": "user-2",
    "name": "鈴木 花子",
    "icon": "https://example.com/icons/user2.png"
  },
  "created_at": "2024-02-20",
  "start_date": "2024-02-20",
  "due_date": "2024-02-20",
  "story_points": 3,
  "category": "開発",
  "is_subtask": false,
  "sub_tasks": [
    {
      "task_id": 2,
      "title": "サブタスク1",
      "description": "サブタスク1の説明",
      "priority": "中",
      "status": "未着手",
      "reporter": "山田太郎",
      "assignee": "鈴木花子",
      "created_at": "2024-02-20",
      "start_date": "2024-02-20",
      "due_date": "2024-02-20",
      "story_points": 1,
      "category": "開発",
      "is_subtask": true,
      "sub_tasks": [],
      "comments": [
        {
          "comment_id": 1,
          "comment": "サブタスク1のコメント",
          "created_at": "2024-02-20",
          "created_by": "山田太郎"
        }
      ]
    }
  ],
  "comments": [
    {
      "comment_id": 1,
      "comment": "新しいタスクのコメント",
      "created_at": "2024-02-20",
      "created_by": "山田太郎"
    }
  ]
}
```

### TASK-002

| API機能No. | TASK-002 |
| --- | --- |
| API名 | タスクを取得する |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | task_id に紐づくタスクを取得するAPI |
| METHOD | GET |
| URI | `/api/v1/tasks/{task_id}` |

#### リクエスト

```bash
curl -X GET "http://localhost:8080/tasks/1"
```

#### レスポンス

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| task_id | int | 10 | ○ | × | タスクのUUID |
| title | string | 255 | ○ | × | タスクのタイトル |
| description | string | 255 | × | × | タスクの説明 |
| priority | string | 255 | × | × | タスクの優先度 |
| status | string | 255 | ○ | × | タスクのステータス |
| reporter | string | 255 | ○ | × | タスクの報告者 |
| assignee | string | 255 | × | × | タスクの担当者 |
| created_at | date | 255 | ○ | × | タスクの作成日 |
| start_date | date | 255 | × | × | タスクの開始日 |
| due_date | date | 255 | × | × | タスクの期限日 |
| story_points | int | 10 | × | × | タスクのストーリーポイント |
| category | string | 255 | × | × | タスクのカテゴリ |
| is_subtask | boolean | 1 | × | × | サブタスクかどうか |
| sub_tasks | array | 255 | × | ○ | サブタスクのリスト |
| comments | array | 255 | × | ○ | コメントのリスト |

レスポンス例

```json
{
  "task_id": 1,
  "title": "新しいタスク",
  "description": "新しいタスクの説明",
  "priority": "高",
  "status": "未着手",
  "reporter": {
    "id": "user-1",
    "name": "山田 太郎",
    "icon": "https://example.com/icons/user1.png"
  },
  "assignee": {
    "id": "user-2",
    "name": "鈴木 花子",
    "icon": "https://example.com/icons/user2.png"
  },
  "created_at": "2024-02-20",
  "start_date": "2024-02-20",
  "due_date": "2024-02-20",
  "story_points": 3,
  "category": "開発",
  "is_subtask": false,
  "sub_tasks": [
    {
      "task_id": 2,
      "title": "サブタスク1",
      "description": "サブタスク1の説明",
      "priority": "中",
      "status": "未着手",
      "reporter": "山田太郎",
      "assignee": "鈴木花子",
      "created_at": "2024-02-20",
      "start_date": "2024-02-20",
      "due_date": "2024-02-20",
      "story_points": 1,
      "category": "開発",
      "is_subtask": true,
      "sub_tasks": [],
      "comments": [
        {
          "comment_id": 1,
          "comment": "サブタスク1のコメント",
          "created_at": "2024-02-20",
          "created_by": "山田太郎"
        }
      ]
    }
  ],
  "comments": [
    {
      "comment_id": 1,
      "comment": "新しいタスクのコメント",
      "created_at": "2024-02-20",
      "created_by": "山田太郎"
    }
  ]
}
```

### TASK-003

| API機能No. | TASK-003 |
| --- | --- |
| API名 | 新規タスクの作成 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | タスクを新規作成するAPI |
| METHOD | POST |
| URI | `/api/v1/tasks` |

#### リクエスト

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| title | string | 255 | ○ | × | タスクのタイトル |
| description | string | 255 | × | × | タスクの説明 |
| priority | string | 255 | × | × | タスクの優先度 |
| status | string | 255 | ○ | × | タスクのステータス |
| reporter | string | 255 | ○ | × | タスクの報告者 |
| assignee | string | 255 | × | × | タスクの担当者 |
| start_date | date | 255 | × | × | タスクの開始日 |
| due_date | date | 255 | × | × | タスクの期限日 |
| story_points | int | 10 | × | × | タスクのストーリーポイント |
| category | string | 255 | × | × | タスクのカテゴリ |
| is_subtask | boolean | 1 | ○ | × | サブタスクかどうか |


リクエスト例

- 最小限の項目を指定して新規タスクを作成する

```bash
curl -X POST "http://localhost:8080/tasks" -d '{
  "title": "新しいタスク",
  "status": "未着手",
  "reporter": "user_001",
  "is_subtask": false
}'

```


- 全ての項目を指定して新規タスクを作成する
```bash
curl -X POST "http://localhost:8080/tasks" -d '{
  "title": "新しいタスク",
  "description": "新しいタスクの説明",
  "priority": "高",
  "status": "未着手",
  "reporter": "user_001",
  "assignee": "user_002",
  "start_date": "2024-02-20",
  "due_date": "2024-02-20",
  "story_points": 3,
  "category": "開発",
  "is_subtask": false
}'
```

#### レスポンス

```json
{
  "message": "ok"
}
```

### TASK-004

| API機能No. | TASK-004 |
| --- | --- |
| API名 | タスクの更新 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | タスクを更新するAPI |
| METHOD | PUT |
| URI | `/api/v1/tasks/{task_id}` |

#### リクエスト

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| title | string | 255 | ○ | × | タスクのタイトル |
| description | string | 255 | × | × | タスクの説明 |
| priority | string | 255 | × | × | タスクの優先度 |
| status | string | 255 | ○ | × | タスクのステータス |
| assignee | string | 255 | × | × | タスクの担当者 |
| start_date | date | 255 | × | × | タスクの開始日 |
| due_date | date | 255 | × | × | タスクの期限日 |
| story_points | int | 10 | × | × | タスクのストーリーポイント |
| category | string | 255 | × | × | タスクのカテゴリ |


リクエスト例

- タイトルなど多くの項目を変更する

```bash
curl -X PUT "http://localhost:8080/tasks/1" -d '{
  "title": "変更後のタイトル"
  "description": "新しいタスクの説明",
  "priority": "高",
  "status": "未着手",
  "assignee": "user_002",
  "start_date": "2024-02-20",
}'
```

- ステータスを変更する

```bash
curl -X PUT "http://localhost:8080/tasks/1" -d '{
  "status": "完了"
}'
```

#### レスポンス

```json
{
  "message": "ok"
}
```

### TASK-005

| API機能No. | TASK-005 |
| --- | --- |
| API名 | タスクの削除 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | タスクを削除するAPI |
| METHOD | DELETE |
| URI | `/api/v1/tasks/{task_id}` |

#### リクエスト

```bash
curl -X DELETE "http://localhost:8080/tasks/1" -d '{
  "user_id": "user_001"
}'
```

引数の `user_id` は、削除対象のタスクの削除権限を持つユーザーのIDかの確認を行うために使用する。

#### レスポンス

```json
{
  "message": "ok"
}
```

### TASK-006

| API機能No. | TASK-006 |
| --- | --- |
| API名 | サブタスクの追加 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | タスクにサブタスクを追加するAPI |
| METHOD | POST |
| URI | `/api/v1/tasks/{task_id}/subtasks` |

#### リクエスト

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| title | string | 255 | ○ | × | タスクのタイトル |
| description | string | 255 | × | × | タスクの説明 |
| priority | string | 255 | × | × | タスクの優先度 |
| status | string | 255 | ○ | × | タスクのステータス |
| reporter | string | 255 | ○ | × | タスクの報告者 |
| assignee | string | 255 | × | × | タスクの担当者 |
| start_date | date | 255 | × | × | タスクの開始日 |
| due_date | date | 255 | × | × | タスクの期限日 |
| story_points | int | 10 | × | × | タスクのストーリーポイント |
| category | string | 255 | × | × | タスクのカテゴリ |

リクエスト例

最小限の項目を指定して新規サブタスクを作成する場合：

```bash
curl -X POST "http://localhost:8080/tasks/1/subtasks" -d '{
  "title": "サブタスク1",
  "status": "未着手",
  "reporter": "user_001",
  "created_at": "2024-02-20",
}'
```

#### レスポンス例

```json
{
  message: "ok"
}
```

### TASK-007

| API機能No. | TASK-008 |
| --- | --- |
| API名 | コメントの追加 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | タスクにコメントを追加するAPI |
| METHOD | POST |
| URI | `/api/v1/tasks/{task_id}/comments` |

#### リクエスト

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| comment | string | 255 | ○ | × | コメントの内容 |
| created_by | string | 255 | ○ | × | コメントの作成者 |

リクエスト例

```bash
curl -X POST "http://localhost:8080/tasks/1/comments" -d '{
  "comment": "新しいタスクのコメント",
  "created_by": "user_001"
}'
```

#### レスポンス

```json
{
  "message": "ok"
}
```

### TASK-008

| API機能No. | TASK-009 |
| --- | --- |
| API名 | コメントの編集 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | タスクのコメントを編集するAPI |
| METHOD | PUT |
| URI | `/api/v1/tasks/{task_id}/comments/{comment_id}` |

#### リクエスト

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| comment | string | 255 | ○ | × | コメントの内容 |
| created_by | string | 255 | ○ | × | コメントの作成者 |

リクエスト例

```bash
curl -X PUT "http://localhost:8080/tasks/1/comments/1" -d '{
  "comment": "新しいタスクのコメント",
  "created_by": "user_001"
}'

```

`created_by` は、コメントの作成者がコメントを編集する権限を持っているか確認するために使用する。
持っていたとき、（コメント本人の場合、コメントの内容を変更する）


### TASK-009

| API機能No. | TASK-010 |
| --- | --- |
| API名 | コメントの削除 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | タスクのコメントを削除するAPI |
| METHOD | DELETE |
| URI | `/api/v1/tasks/{task_id}/comments/{comment_id}` |

#### リクエスト

リクエスト前に、自分の権限で削除できるコメントか確認する。

```bash
curl -X DELETE "http://localhost:8080/tasks/1/comments/1"　-d '{
  "created_by": user_001
}'

```

`created_by` は、コメントの作成者がコメントを削除する権限を持っているか確認するために使用する。

#### レスポンス

```json
{
  "message": "ok"
}
```

### USER-001

| API機能No. | USER-001 |
| --- | --- |
| API名 | ユーザー一覧取得 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | ユーザー一覧を取得するAPI |
| METHOD | GET |
| URI | `/api/v1/users` |

#### リクエスト

```bash
curl -X GET "http://localhost:8080/users"
```

#### レスポンス

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| user_id | string | 255 | ○ | × | ユーザーのUUID |
| name | string | 255 | ○ | × | ユーザーの名前 |
| icon | string | 255 | × | × | ユーザーのアイコン |

### USER-002

| API機能No. | USER-002 |
| --- | --- |
| API名 | ユーザーを取得する |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | user_id に紐づくユーザーを取得するAPI |
| METHOD | GET |
| URI | `/api/v1/users/{user_id}` |

#### リクエスト

```bash
curl -X GET "http://localhost:8080/users/user_001"
```

#### レスポンス

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| user_id | string | 255 | ○ | × | ユーザーのUUID |
| name | string | 255 | ○ | × | ユーザーの名前 |
| icon | string | 255 | × | × | ユーザーのアイコン |

レスポンス例

```json
{
  "user_id": "user_001",
  "name": "山田 太郎",
  "icon": "https://example.com/icons/user1.png"
}
```

### USER-003

| API機能No. | USER-003 |
| --- | --- |
| API名 | 新規ユーザーの作成 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | ユーザーを新規作成するAPI |
| METHOD | POST |
| URI | `/api/v1/users` |

#### リクエスト

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| name | string | 255 | ○ | × | ユーザーの名前 |
| icon | string | 255 | × | × | ユーザーのアイコン |

リクエスト例

```bash
curl -X POST "http://localhost:8080/users" -d '{
  "name": "山田 太郎",
  "icon": "https://example.com/icons/user1.png"
}'
```

#### レスポンス

```json
{
  "message": "ok"
}
```

### USER-004

| API機能No. | USER-004 |
| --- | --- |
| API名 | ユーザーの更新 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | ユーザーを更新するAPI |
| METHOD | PUT |
| URI | `/api/v1/users/{user_id}` |

#### リクエスト

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| user_id | string | 255 | ○ | × | ユーザーのUUID |
| name | string | 255 | ○ | × | ユーザーの名前 |
| icon | string | 255 | × | × | ユーザーのアイコン |

リクエスト例

```bash
curl -X PUT "http://localhost:8080/users/user_001" -d '{
  "name": "山田 太郎",
  "icon": "https://example.com/icons/user1.png"
}'
```

### USER-005

| API機能No. | USER-005 |
| --- | --- |
| API名 | ユーザーの削除 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | ユーザーを削除するAPI |
| METHOD | DELETE |
| URI | `/api/v1/users/{user_id}` |

#### リクエスト

```bash
curl -X DELETE "http://localhost:8080/users/user_001"
```

### USER-006

| API機能No. | USER-006 |
| --- | --- |
| API名 | ユーザーの権限変更 |
| 更新日/更新者 | 2025.02.11 / Amakura |
| 概要 | ユーザーの権限を変更するAPI |
| METHOD | PUT |  
| URI | `/api/v1/users/{user_id}/permission` |

#### リクエスト

| JSON Key | 型 | サイズ | 必須 | 繰返 | 値の説明 |
| --- | --- | --- | --- | --- | --- |
| permission | string | 255 | ○ | × | ユーザーの権限 |

リクエスト例

```bash
curl -X PUT "http://localhost:8080/users/user_001/permission" -d '{
  "permission": "admin"
}'
```

#### レスポンス

```json
{
  "message": "ok"
}
```


## 参考資料

[［API］ API仕様書の書き方](https://qiita.com/sunstripe2011/items/9230396febfab2eae2c2)
