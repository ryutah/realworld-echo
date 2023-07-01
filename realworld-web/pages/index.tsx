import Link from "next/link";

const Banner = () => (
  <div className="banner">
    <div className="container">
      <h1 className="logo-font">conduit</h1>
      <p>A place to share your knowledge.</p>
    </div>
  </div>
);

const FeedToggle = () => (
  <div className="feed-toggle">
    <ul className="nav nav-pills outline-active">
      <li className="nav-item">
        <Link className="nav-link disabled" href="">
          Your Feed
        </Link>
      </li>
      <li className="nav-item">
        <Link className="nav-link active" href="">
          Global Feed
        </Link>
      </li>
    </ul>
  </div>
);

type Tag = {
  name: string;
};

const PopularTag = ({ tags }: { tags: Tag[] }) => (
  <div className="sidebar">
    <p>Popular Tags</p>

    <div className="tag-list">
      {tags.map((tag, idx) => (
        <a className="tag-pill tag-default" key={idx} href="">
          {tag.name}
        </a>
      ))}
    </div>
  </div>
);

type ArticleProp = {
  title: string;
  description: string;
  author: string;
  date: Date;
  likes: number;
};

const Article = ({ article }: { article: ArticleProp }) => (
  <div className="article-preview">
    <div className="article-meta">
      <Link href="/profile/eric/favorites">
        <img src="http://i.imgur.com/N4VcUeJ.jpg" />
      </Link>
      <div className="info">
        <Link href="" className="author">
          {article.author}
        </Link>
        <span className="date">{article.date.toDateString()}</span>
      </div>
      <button className="btn btn-outline-primary btn-sm pull-xs-right">
        <i className="ion-heart"></i> {article.likes}
      </button>
    </div>

    <Link href="/article/foobar" className="preview-link">
      <h1>{article.title}</h1>
      <p>{article.description}</p>
      <span>Read more...</span>
    </Link>
  </div>
);

const ArticleList = ({ articles }: { articles: ArticleProp[] }) => (
  <>
    <FeedToggle />
    {articles.map((article, idx) => (
      <Article key={idx} article={article} />
    ))}
  </>
);

export default function Home() {
  const articles: ArticleProp[] = [
    {
      author: "Eric Simons",
      date: new Date(),
      likes: 29,
      title: "How to build webapps that scale",
      description: "This is the description for the post.",
    },
    {
      author: "Albert Pai",
      date: new Date(),
      likes: 32,
      title:
        "The song you won't ever stop singing. No matter how hard you try.",
      description: "This is the description for the post.",
    },
  ];
  const tags: Tag[] = [
    { name: "programming" },
    { name: "javascript" },
    { name: "emberjs" },
    { name: "angularjs" },
    { name: "react" },
    { name: "mean" },
    { name: "node" },
    { name: "rails" },
  ];

  return (
    <div className="home-page">
      <Banner />

      <div className="container page">
        <div className="row">
          <div className="col-md-9">
            <ArticleList articles={articles} />
          </div>

          <div className="col-md-3">
            <PopularTag tags={tags} />
          </div>
        </div>
      </div>
    </div>
  );
}
