import { defineConfig } from "orval";

// バンドル済みスキーマ openapi/openapi-bundled.yaml から web / admin-console
// 双方の TypeScript クライアント (TanStack Query フック + 型) を生成する。
// openapi-bundled.yaml は `bun run bundle` (redocly) が paths/ schemas/ を
// 解決して生成するため、orval 実行前にバンドルが必要。
// orval はパスを「この設定ファイルのあるディレクトリ (openapi/)」基準で解決する。
const sharedOutput = {
  mode: "tags-split" as const,
  client: "react-query" as const,
  httpClient: "fetch" as const,
  clean: true,
  prettier: false,
};

export default defineConfig({
  web: {
    input: "./openapi-bundled.yaml",
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
    input: "./openapi-bundled.yaml",
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
