# TODOアプリ (Solid.js + DDD)

## 📌 プロジェクト概要
このプロジェクトは、Solid.js を使用して開発されるシンプルな TODO アプリです。ドメイン駆動開発 (DDD) の概念を取り入れ、拡張性とメンテナンス性の高いアーキテクチャを実現します。

## 🛠 技術スタック
- **フロントエンド**: Solid.js, TypeScript, Tailwind CSS
- **アーキテクチャ**: ドメイン駆動開発 (DDD)
- **状態管理**: Solid.js Signals / Stores
- **バックエンド (オプション)**: Firebase / Supabase

## 📂 ディレクトリ構成 (予定)
```
/src
 ├── application  # アプリケーション層 (ユースケース)
 ├── domain       # ドメイン層 (エンティティ・リポジトリ)
 ├── infrastructure # インフラ層 (API通信, DB操作)
 ├── ui           # UI コンポーネント (Solid.js コンポーネント)
 ├── main.tsx     # エントリポイント
```

## 🎯 機能 (予定)
- ✅ タスクの追加・削除・更新
- ✅ タスクの完了状態の変更
- ✅ ローカルストレージ / Firebase での永続化
- ✅ ユーザー認証 (オプション)

## 🚀 セットアップ方法
```bash
git clone https://github.com/your-repo/todo-solidjs.git
cd todo-solidjs
yarn install # または npm install
yarn dev # または npm run dev
```

## 📜 ライセンス
MIT License