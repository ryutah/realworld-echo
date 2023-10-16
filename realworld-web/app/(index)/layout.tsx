import { ArticleProps } from "@/app/domain";
import { articles } from "@/tests/testdata";
import { ReactNode } from "react";
import { ArticlesProvider } from "./store";

export default async function ArticleLayout({
  children,
}: {
  children: ReactNode;
}) {
  const articles = await getArticles();

  return <ArticlesProvider globalFeeds={articles}>{children}</ArticlesProvider>;
}

async function getArticles(): Promise<ArticleProps[]> {
  return await new Promise((resolve) => {
    setTimeout(() => {
      resolve(articles);
    }, 500);
  });
}
