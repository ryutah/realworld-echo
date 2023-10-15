import { ArticleProps } from "@/app/domain";
import { articles } from "@/tests/testdata";
import { ReactNode } from "react";
import { ArticleProvider } from "./store";

export default async function ArticleLayout({
  children,
}: {
  children: ReactNode;
}) {
  const articles = await getArticles();

  return <ArticleProvider article={articles[0]}>{children}</ArticleProvider>;
}

async function getArticles(): Promise<ArticleProps[]> {
  return await new Promise((resolve) => {
    setTimeout(() => {
      resolve(articles);
    }, 500);
  });
}
