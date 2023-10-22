import { ArticlesApi, BaseAPI, Configuration } from "./generated";

export function newClient<T extends BaseAPI>(type: {
  new (configuration?: Configuration): T;
}): T {
  return new type(
    new Configuration({
      basePath: "http://localhost:3000",
    })
  );
}

async function foo(): Promise<void> {
  const config = new Configuration({
    basePath: "http://localhost:3000",
  });
  const api = new ArticlesApi(config);
  const resp = await api.getArticle(
    {
      slug: "foo",
    },
    {
      headers: {},
    }
  );
}
