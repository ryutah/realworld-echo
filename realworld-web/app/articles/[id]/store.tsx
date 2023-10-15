"use client";

import { ArticleProps } from "@/app/domain";
import {
  Dispatch,
  ReactNode,
  createContext,
  useContext,
  useReducer,
} from "react";

type State = ArticleProps;

const Context = createContext<ArticleProps | null>(null);

const DispatchContext = createContext<Dispatch<Action> | undefined>(undefined);

type Action = {
  type: "article/store";
  payload: ArticleProps;
};

export const Actions = {
  storeArticle(article: ArticleProps): Action {
    return {
      type: "article/store",
      payload: article,
    };
  },
};

function reducers(state: State | null, action: Action) {
  switch (action.type) {
    case "article/store":
      return action.payload;
    default:
      return state;
  }
}

export const useArticle = () => {
  return useContext(Context);
};

export const useArticleDispatch = () => {
  return useContext(DispatchContext);
};

type Props = {
  children: ReactNode;
  article: ArticleProps;
};

export const ArticleProvider = ({ children, article }: Props) => {
  const [state, dispatch] = useReducer(reducers, article);

  return (
    <Context.Provider value={state}>
      <DispatchContext.Provider value={dispatch}>
        {children}
      </DispatchContext.Provider>
    </Context.Provider>
  );
};
