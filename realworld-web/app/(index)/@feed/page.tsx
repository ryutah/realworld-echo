import { ArticleProps } from "@/app/domain";
import { articles } from "@/tests/testdata";
import FeedTab from "./FeedTab";
import { ArticlesProvider } from "./store";

export default async function Home() {
  const articles = await getArticles();

  return (
    <ArticlesProvider globalFeeds={articles}>
      <FeedTab initTab="global" />
    </ArticlesProvider>
  );
}

async function getArticles(): Promise<ArticleProps[]> {
  return await new Promise((resolve) => {
    setTimeout(() => {
      resolve(articles);
    }, 500);
  });
}
