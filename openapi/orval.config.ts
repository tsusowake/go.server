import { defineConfig } from "orval";

// 単一スキーマ openapi/openapi.yaml から web / admin-console 双方の
// TypeScript クライアント (TanStack Query フック + 型) を生成する。
// orval はパスを「この設定ファイルのあるディレクトリ (openapi/)」基準で解決する。
// 将来 admin 専用 API を分けたい場合は openapi/admin.yaml を足し、
// admin ターゲットの input を差し替える。
const sharedOutput = {
  mode: "tags-split" as const,
  client: "react-query" as const,
  httpClient: "fetch" as const,
  clean: true,
  prettier: false,
};

export default defineConfig({
  web: {
    input: "./openapi.yaml",
    output: {
      ...sharedOutput,
      target: "../packages/web/app/api/endpoints",
      schemas: "../packages/web/app/api/model",
      override: {
        mutator: {
          path: "../packages/web/app/api/fetcher.ts",
          name: "customFetch",
        },
      },
    },
  },
  admin: {
    input: "./openapi.yaml",
    output: {
      ...sharedOutput,
      target: "../packages/admin-console/app/api/endpoints",
      schemas: "../packages/admin-console/app/api/model",
      override: {
        mutator: {
          path: "../packages/admin-console/app/api/fetcher.ts",
          name: "customFetch",
        },
      },
    },
  },
});
