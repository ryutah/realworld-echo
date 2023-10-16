"use client";

import { ArticleProps } from "@/app/domain";
import { ReactNode, createContext, useContext } from "react";

type State = {
  articles: {
    globalFeeds: ArticleProps[];
  };
};

const Context = createContext<State>({
  articles: {
    globalFeeds: [],
  },
});

type Props = {
  children: ReactNode;
  globalFeeds: ArticleProps[];
};

export function useArticles() {
  return useContext(Context);
}

export function ArticlesProvider({ children, globalFeeds }: Props) {
  return (
    <Context.Provider value={{ articles: { globalFeeds } }}>
      {children}
    </Context.Provider>
  );
}
