import { Profile } from "@/app/domain/profile";
import { NonFunctionProperties } from "@/app/lib/types";

export type ArticleProps = NonFunctionProperties<Article>;

export class Article {
  readonly slug: string;
  readonly title: string;
  readonly description: string;
  readonly body: string;
  readonly tagList: string[];
  readonly createdAt: Date;
  readonly updatedAt: Date;
  readonly favoritesCount: number;
  readonly favorited?: boolean;
  readonly author: Profile;

  constructor(article: ArticleProps) {
    this.slug = article.slug;
    this.title = article.title;
    this.description = article.description;
    this.body = article.body;
    this.tagList = article.tagList;
    this.createdAt = article.createdAt;
    this.updatedAt = article.updatedAt;
    this.favoritesCount = article.favoritesCount;
    this.author = article.author;
  }

  public formattedCreatedAt(): string {
    return this.createdAt.toDateString();
  }

  public formattedUpdatedAt(): string {
    return this.updatedAt.toDateString();
  }
}
