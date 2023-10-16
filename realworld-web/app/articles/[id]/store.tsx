"use client";

import { ArticleProps } from "@/app/domain";
import { ReactNode, createContext, useContext } from "react";

type State = ArticleProps;

const Context = createContext<State | null>(null);

export const useArticle = () => {
  return useContext(Context);
};

type Props = {
  children: ReactNode;
  article: ArticleProps;
};

export const ArticleProvider = ({ children, article }: Props) => {
  return <Context.Provider value={article}>{children}</Context.Provider>;
};
