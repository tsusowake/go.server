import { QueryClient } from "@tanstack/react-query";

// SSR + ブラウザで使い回す QueryClient を生成する。
// SSR では各リクエストごとに新しいインスタンスが望ましいため関数で提供する。
export const makeQueryClient = (): QueryClient =>
  new QueryClient({
    defaultOptions: {
      queries: {
        // SSR 中の即時 refetch を避ける。
        staleTime: 60 * 1000,
      },
    },
  });
