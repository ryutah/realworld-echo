"use client";

import { useArticle } from "./store";

export default function Article() {
  const article = useArticle();

  if (!article) {
    return <></>;
  }
  return <>{article.title}</>;
}
