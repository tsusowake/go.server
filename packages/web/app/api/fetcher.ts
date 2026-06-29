// orval が生成するクライアントが使う共通 fetch 実装 (mutator)。
// openapi/orval.config.ts の override.mutator から参照される。
//
// orval の httpClient: "fetch" は mutator が { status, data, headers } を
// 返すことを前提とする (生成された *Response 型がこの形)。
//
// SSR (サーバー側) では相対 URL を fetch できないため、API のベース URL を
// 解決して前置する。dev では vite proxy 経由で "/api" を Go へ転送する想定。
const getBaseUrl = (): string => {
  // サーバー実行時は環境変数を優先。
  if (typeof process !== "undefined" && process.env?.API_BASE_URL) {
    return process.env.API_BASE_URL;
  }
  // クライアント (ブラウザ) では同一オリジン + /api を使い vite proxy に委ねる。
  if (typeof window !== "undefined") {
    return `${window.location.origin}/api`;
  }
  // フォールバック (dev のローカル Go サーバー)。
  return "http://localhost:8080";
};

const getBody = async (response: Response): Promise<unknown> => {
  const contentType = response.headers.get("content-type");
  if (contentType?.includes("application/json")) {
    return response.json();
  }
  const text = await response.text();
  return text ? text : undefined;
};

export const customFetch = async <T>(
  url: string,
  options?: RequestInit,
): Promise<T> => {
  const requestUrl = `${getBaseUrl()}${url}`;
  const response = await fetch(requestUrl, options);
  const data = await getBody(response);

  return {
    status: response.status,
    data,
    headers: response.headers,
  } as T;
};
